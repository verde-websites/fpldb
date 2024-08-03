package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/uptrace/bun/dbfixture"
	"github.com/uptrace/bun/migrate"
	"github.com/urfave/cli/v2"
	"github.com/verde-websites/fpldb/bunapp"
	"github.com/verde-websites/fpldb/cmd/bun/migrations"
	"github.com/verde-websites/fpldb/generate"
	"github.com/verde-websites/fpldb/models"
)

func main() {
	app := &cli.App{
		Name: "bun",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "env",
				Value: "dev",
				Usage: "environment",
			},
			&cli.StringFlag{
				Name:  "dsn",
				Usage: "postgres dsn",
			},
		},
		Commands: []*cli.Command{
			newfplFixtureGeneratorCommand,
			newDBCommand(migrations.Migrations),
		},
	}
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

var newfplFixtureGeneratorCommand = &cli.Command{
	Name:  "generate",
	Usage: "generate all database fixtures",
	Action: func(c *cli.Context) error {
		_, app, err := bunapp.StartCLI(c)
		if err != nil {
			return err
		}
		defer app.Stop()
		err = generate.GenerateDatabaseFixtures()
		if err != nil {
			return err
		}
		return nil
	},
}

func newDBCommand(migrations *migrate.Migrations) *cli.Command {
	return &cli.Command{
		Name:  "db",
		Usage: "manage database migrations",
		Subcommands: []*cli.Command{
			{
				Name:  "init",
				Usage: "create migration tables",
				Action: func(c *cli.Context) error {
					ctx, app, err := bunapp.StartCLI(c)
					if err != nil {
						return err
					}
					defer app.Stop()

					migrator := migrate.NewMigrator(app.DB(), migrations)
					return migrator.Init(ctx)
				},
			},
			{
				Name:  "migrate",
				Usage: "migrate database",
				Action: func(c *cli.Context) error {
					ctx, app, err := bunapp.StartCLI(c)
					if err != nil {
						return err
					}
					defer app.Stop()

					migrator := migrate.NewMigrator(app.DB(), migrations)

					group, err := migrator.Migrate(ctx)
					if err != nil {
						return err
					}

					if group.ID == 0 {
						fmt.Printf("there are no new migrations to run\n")
						return nil
					}

					fmt.Printf("migrated to %s\n", group)
					return nil
				},
			},
			{
				Name:  "rollback",
				Usage: "rollback the last migration group",
				Action: func(c *cli.Context) error {
					ctx, app, err := bunapp.StartCLI(c)
					if err != nil {
						return err
					}
					defer app.Stop()

					migrator := migrate.NewMigrator(app.DB(), migrations)

					group, err := migrator.Rollback(ctx)
					if err != nil {
						return err
					}

					if group.ID == 0 {
						fmt.Printf("there are no groups to roll back\n")
						return nil
					}

					fmt.Printf("rolled back %s\n", group)
					return nil
				},
			},
			{
				Name:  "lock",
				Usage: "lock migrations",
				Action: func(c *cli.Context) error {
					ctx, app, err := bunapp.StartCLI(c)
					if err != nil {
						return err
					}
					defer app.Stop()

					migrator := migrate.NewMigrator(app.DB(), migrations)
					return migrator.Lock(ctx)
				},
			},
			{
				Name:  "unlock",
				Usage: "unlock migrations",
				Action: func(c *cli.Context) error {
					ctx, app, err := bunapp.StartCLI(c)
					if err != nil {
						return err
					}
					defer app.Stop()

					migrator := migrate.NewMigrator(app.DB(), migrations)
					return migrator.Unlock(ctx)
				},
			},
			{
				Name:  "create_go",
				Usage: "create Go migration",
				Action: func(c *cli.Context) error {
					ctx, app, err := bunapp.StartCLI(c)
					if err != nil {
						return err
					}
					defer app.Stop()

					migrator := migrate.NewMigrator(app.DB(), migrations)

					name := strings.Join(c.Args().Slice(), "_")
					mf, err := migrator.CreateGoMigration(ctx, name)
					if err != nil {
						return err
					}
					fmt.Printf("created migration %s (%s)\n", mf.Name, mf.Path)

					return nil
				},
			},
			{
				Name:  "create_sql",
				Usage: "create up and down SQL migrations",
				Action: func(c *cli.Context) error {
					ctx, app, err := bunapp.StartCLI(c)
					if err != nil {
						return err
					}
					defer app.Stop()

					migrator := migrate.NewMigrator(app.DB(), migrations)

					name := strings.Join(c.Args().Slice(), "_")
					files, err := migrator.CreateSQLMigrations(ctx, name)
					if err != nil {
						return err
					}

					for _, mf := range files {
						fmt.Printf("created migration %s (%s)\n", mf.Name, mf.Path)
					}

					return nil
				},
			},
			{
				Name:  "status",
				Usage: "print migrations status",
				Action: func(c *cli.Context) error {
					ctx, app, err := bunapp.StartCLI(c)
					if err != nil {
						return err
					}
					defer app.Stop()

					migrator := migrate.NewMigrator(app.DB(), migrations)

					ms, err := migrator.MigrationsWithStatus(ctx)
					if err != nil {
						return err
					}
					fmt.Printf("migrations: %s\n", ms)
					fmt.Printf("unapplied migrations: %s\n", ms.Unapplied())
					fmt.Printf("last migration group: %s\n", ms.LastGroup())

					return nil
				},
			},
			{
				Name:  "mark_applied",
				Usage: "mark migrations as applied without actually running them",
				Action: func(c *cli.Context) error {
					ctx, app, err := bunapp.StartCLI(c)
					if err != nil {
						return err
					}
					defer app.Stop()

					migrator := migrate.NewMigrator(app.DB(), migrations)

					group, err := migrator.Migrate(ctx, migrate.WithNopMigration())
					if err != nil {
						return err
					}

					if group.ID == 0 {
						fmt.Printf("there are no new migrations to mark as applied\n")
						return nil
					}

					fmt.Printf("marked as applied %s\n", group)
					return nil
				},
			},
			{
				Name:  "seed",
				Usage: "seed the database with database fixtures",
				Action: func(c *cli.Context) error {
					_, app, err := bunapp.StartCLI(c)
					if err != nil {
						return err
					}
					defer app.Stop()
					app.DB().RegisterModel((*models.Season)(nil), (*models.Team)(nil), (*models.Player)(nil), (*models.PlayerFplTracker)(nil), (*models.Gameweek)(nil), (*models.TeamFplTracker)(nil), (*models.Fixture)(nil))
					fixture := dbfixture.New(app.DB(), dbfixture.WithTruncateTables())

					// seed the season fixtures
					err = fixture.Load(context.Background(), bunapp.FS(), "fixture/season.yml")
					if err != nil {
						return err
					}
					// seed the team fixtures
					err = fixture.Load(context.Background(), bunapp.FS(), "fixture/team.yml")
					if err != nil {
						return err
					}
					// seed the player fixtures
					err = fixture.Load(context.Background(), bunapp.FS(), "fixture/player.yml")
					if err != nil {
						return err
					}
					// seed the playerFplTracker fixtures
					err = fixture.Load(context.Background(), bunapp.FS(), "fixture/playerFplTracker.yml")
					if err != nil {
						return err
					}
					// seed the gameweek fixtures
					err = fixture.Load(context.Background(), bunapp.FS(), "fixture/gameweek.yml")
					if err != nil {
						return err
					}
					// seed the teamFplTracker fixtures
					err = fixture.Load(context.Background(), bunapp.FS(), "fixture/teamFplTracker.yml")
					if err != nil {
						return err
					}
					// seed the fixture fixtures
					err = fixture.Load(context.Background(), bunapp.FS(), "fixture/fixture.yml")
					if err != nil {
						return err
					}
					return nil
				},
			},
		},
	}
}
