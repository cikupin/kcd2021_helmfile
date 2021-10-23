package cmd

import (
	"github.com/urfave/cli/v2"
)

var API = &cli.Command{
	Name:        "api",
	Description: "Run api for KCD 2021 demo app",
	Action: func(c *cli.Context) error {
		runAPI()
		return nil
	},
}

func runAPI() {
}
