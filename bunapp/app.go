package bunapp

import (
	"context"
	"database/sql"
	"errors"
	"math/rand"
	"os"
	"os/signal"
	"sync"
	"sync/atomic"
	"syscall"
	"time"

	"github.com/benbjohnson/clock"
	"github.com/uptrace/bun"

	//	"github.com/uptrace/bun/dialect/pgdialect"
	//"github.com/uptrace/bun/driver/pgdriver"
	_ "github.com/go-sql-driver/mysql"
	"github.com/uptrace/bun/dialect/mysqldialect"
	"github.com/uptrace/bun/extra/bundebug"
	"github.com/urfave/cli/v2"
)

type appCtxKey struct{}

func AppFromContext(ctx context.Context) *App {
	return ctx.Value(appCtxKey{}).(*App)
}

func ContextWithApp(ctx context.Context, app *App) context.Context {
	ctx = context.WithValue(ctx, appCtxKey{}, app)
	return ctx
}

type App struct {
	ctx context.Context
	cfg *AppConfig

	stopping uint32
	stopCh   chan struct{}

	onStop      appHooks
	onAfterStop appHooks

	clock clock.Clock

	// lazy init
	dbOnce sync.Once
	db     *bun.DB
}

func New(ctx context.Context, cfg *AppConfig) *App {
	app := &App{
		cfg:    cfg,
		stopCh: make(chan struct{}),
		clock:  clock.New(),
	}
	app.ctx = ContextWithApp(ctx, app)
	return app
}

func StartCLI(c *cli.Context) (context.Context, *App, error) {
	return Start(c.Context, c.Command.Name, c.String("env"), c.String("dsn"))
}

func Start(ctx context.Context, service, envName string, dsn string) (context.Context, *App, error) {
	cfg, err := ReadConfig(FS(), service, envName)
	if envName == "prod" {
		if dsn == "" {
			return nil, nil, errors.New("the postgres dsn must be set via the --dsn flag while running in prod")
		}
		cfg.PGX.DSN = dsn
	}
	if dsn != "" {
		cfg.PGX.DSN = dsn
	}
	if err != nil {
		return nil, nil, err
	}
	return StartConfig(ctx, cfg)
}

func StartConfig(ctx context.Context, cfg *AppConfig) (context.Context, *App, error) {
	rand.Seed(time.Now().UnixNano())

	app := New(ctx, cfg)
	if err := onStart.Run(ctx, app); err != nil {
		return nil, nil, err
	}
	return app.Context(), app, nil
}

func (app *App) Stop() {
	_ = app.onStop.Run(app.ctx, app)
	_ = app.onAfterStop.Run(app.ctx, app)
}

func (app *App) OnStop(name string, fn HookFunc) {
	app.onStop.Add(newHook(name, fn))
}

func (app *App) OnAfterStop(name string, fn HookFunc) {
	app.onAfterStop.Add(newHook(name, fn))
}

func (app *App) Context() context.Context {
	return app.ctx
}

func (app *App) Config() *AppConfig {
	return app.cfg
}

func (app *App) Running() bool {
	return !app.Stopping()
}

func (app *App) Stopping() bool {
	return atomic.LoadUint32(&app.stopping) == 1
}

func (app *App) IsDebug() bool {
	return app.cfg.Debug
}

func (app *App) Clock() clock.Clock {
	return app.clock
}

func (app *App) SetClock(clock clock.Clock) {
	app.clock = clock
}

func (app *App) DB() *bun.DB {
	app.dbOnce.Do(func() {
		mysqldb, err := sql.Open("mysql", app.cfg.MYSQL.DSN)
		if err != nil {
			panic(err)
		}
		//sqldb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(app.cfg.PGX.DSN)))

		db := bun.NewDB(
			mysqldb,
			mysqldialect.New(),
		)
		//		db := bun.NewDB(sqldb, pgdialect.New())
		db.AddQueryHook(bundebug.NewQueryHook(
			bundebug.WithEnabled(app.IsDebug()),
			bundebug.FromEnv(""),
		))

		app.db = db
	})
	return app.db
}

func WaitExitSignal() os.Signal {
	ch := make(chan os.Signal, 3)
	signal.Notify(
		ch,
		syscall.SIGINT,
		syscall.SIGQUIT,
		syscall.SIGTERM,
	)
	return <-ch
}
