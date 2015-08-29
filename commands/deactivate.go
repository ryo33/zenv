package commands

import (
	"github.com/codegangsta/cli"
	"github.com/ryo33/zenv/environment"
)

var deactivate = cli.Command{
	Name:  "deactivate",
	Usage: "deactivate the environment",
	Description: `
	`,
	Action: doDeactivate,
}

func doDeactivate(c *cli.Context) {
	args := c.Args()
	if len(args) == 2 {
		environment.GetGlobalEnv(args[1]).Deactivate(args[0])
	} else {
		printArgumentError(2)
	}
}
