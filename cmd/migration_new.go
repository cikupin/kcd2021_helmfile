package cmd

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/urfave/cli/v2"
)

// MigrationDown will execute command for creating new migration file
var MigrationNew = &cli.Command{
	Name:        "migration-new",
	Description: "Create new migration file",
	Action: func(c *cli.Context) error {
		return runMigrationNew(c.Args().Get(0))
	},
}

func runMigrationNew(migrationName string) error {
	var migrationContent = `-- +migrate Up
-- SQL in section 'Up' is executed when this migration is applied

-- [your SQL script here]

-- +migrate Down
-- SQL section 'Down' is executed when this migration is rolled back

-- [your SQL script here]
`
	filename := fmt.Sprintf("%d_%s.sql", time.Now().Unix(), migrationName)
	filepath := fmt.Sprintf("%s%s", "migrations/", filename)

	f, err := os.Create(filepath)
	if err != nil {
		log.Printf("[ERROR] Error create migration file %s: %s\n", filepath, err.Error())
		return err
	}
	defer f.Close()

	f.WriteString(migrationContent)
	f.Sync()

	log.Printf("[INFO] New migration file has been created: %s\n", filepath)
	return nil
}
