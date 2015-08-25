package commands

import (
	"github.com/codegangsta/cli"
)

var deactivate = cli.Command{
	Name:  "deactivate",
	Usage: "deactivate the environment",
	Description: `
	`,
	Action: doDeactivate,
}

func doDeactivate(c *cli.Context) {
}
