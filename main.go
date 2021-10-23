package main

import (
	"log"
	"os"

	"github.com/cikupin/kcd2021_helmfile/cmd"
	"github.com/urfave/cli/v2"
)

func main() {
	app := cli.NewApp()
	app.Name = "kcd2021-helmfile"
	app.Description = "Run demo app for KCD Indonesia 2021"
	app.Usage = "Run demo app for KCD Indonesia 2021"
	app.UsageText = "kcd2021-helmfile [command]"

	app.Commands = []*cli.Command{
		cmd.API,
		cmd.MigrationNew,
		cmd.MigrationUp,
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
