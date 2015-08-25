package commands

import (
	"github.com/codegangsta/cli"
)

var activate = cli.Command{
	Name:  "activate",
	Usage: "activate the environment",
	Description: `
	`,
	Action: doActivate,
}

func doActivate(c *cli.Context) {
}
