package migrations

import (
	"context"
	"fmt"

	"github.com/uptrace/bun"
	"github.com/verde-websites/fpldb/models"
)

func init() {
	Migrations.MustRegister(func(ctx context.Context, db *bun.DB) error {
		fmt.Print(" [up migration] ")
		_, err := db.NewCreateTable().Model((*models.Season)(nil)).Exec(ctx)
		if err != nil {
			return err
		}
		_, err = db.NewCreateTable().Model((*models.Team)(nil)).Exec(ctx)
		if err != nil {
			return err
		}
		_, err = db.NewCreateTable().Model((*models.Player)(nil)).Exec(ctx)
		if err != nil {
			return err
		}
		_, err = db.NewCreateTable().Model((*models.PlayerFplTracker)(nil)).Exec(ctx)
		if err != nil {
			return err
		}
		_, err = db.NewCreateTable().Model((*models.Gameweek)(nil)).Exec(ctx)
		if err != nil {
			return err
		}
		_, err = db.NewCreateTable().Model((*models.TeamFplTracker)(nil)).Exec(ctx)
		if err != nil {
			return err
		}
		_, err = db.NewCreateTable().Model((*models.Fixture)(nil)).Exec(ctx)
		if err != nil {
			return err
		}
		fmt.Println("done")
		return nil
	}, func(ctx context.Context, db *bun.DB) error {
		fmt.Print(" [down migration] ")
		_, err := db.NewDropTable().Model((*models.Season)(nil)).Exec(ctx)
		if err != nil {
			return err
		}
		_, err = db.NewDropTable().Model((*models.Team)(nil)).Exec(ctx)
		if err != nil {
			return err
		}
		_, err = db.NewDropTable().Model((*models.Player)(nil)).Exec(ctx)
		if err != nil {
			return err
		}
		_, err = db.NewDropTable().Model((*models.PlayerFplTracker)(nil)).Exec(ctx)
		if err != nil {
			return err
		}
		_, err = db.NewDropTable().Model((*models.Gameweek)(nil)).Exec(ctx)
		if err != nil {
			return err
		}
		_, err = db.NewDropTable().Model((*models.TeamFplTracker)(nil)).Exec(ctx)
		if err != nil {
			return err
		}
		_, err = db.NewDropTable().Model((*models.Fixture)(nil)).Exec(ctx)
		if err != nil {
			return err
		}
		fmt.Println("done")
		return nil
	})
}