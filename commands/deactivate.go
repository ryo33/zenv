package commands

import (
	"github.com/codegangsta/cli"
	"github.com/ryo33/zenv/environment"
	"github.com/ryo33/zenv/util"
)

var deactivate = cli.Command{
	Name:  "deactivate",
	Usage: "deactivate the environment",
	Description: `
	`,
	Action: doDeactivate,
}

func doDeactivate(c *cli.Context) {
	util.Eval = true
	args := c.Args()
	if len(args) == 1 {
		environment.GetGlobalEnv(args[0]).Deactivate()
	} else {
		util.PrintErrorMessage("needs 1 arg")
	}
}
