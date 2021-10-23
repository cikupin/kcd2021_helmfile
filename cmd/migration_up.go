package cmd

import (
	"log"

	"github.com/cikupin/kcd2021_helmfile/internal/bootstrap"
	migrate "github.com/rubenv/sql-migrate"
	"github.com/urfave/cli/v2"
)

var MigrationUp = &cli.Command{
	Name:        "migration-up",
	Description: "Run database migration for KCD 2021 demo app",
	Action: func(c *cli.Context) error {
		return migrateDatabase()
	},
}

func migrateDatabase() error {
	config := bootstrap.LoadConfig()
	db, err := bootstrap.NewMysqlDatabase(config.DBOptions)
	if err != nil {
		log.Printf("[ERROR] Fail initiating database: %s\n", err.Error())
		return err
	}

	source := migrate.FileMigrationSource{
		Dir: "migrations/",
	}

	total, err := migrate.Exec(db.Db, "mysql", source, migrate.Up)
	if err != nil {
		log.Printf("[ERROR] Fail migration: %s\n", err.Error())
		return err
	}

	if total == 0 {
		log.Println("[INFO] Database is up to date")
	} else {
		log.Printf("[INFO] Migrate Up success, total migrated: %d\n", total)
	}
	return nil
}
