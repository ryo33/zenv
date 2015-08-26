package commands

import (
	"github.com/codegangsta/cli"
	"github.com/ryo33/zenv/environment"
	"github.com/ryo33/zenv/util"
)

var activate = cli.Command{
	Name:  "activate",
	Usage: "activate the environment",
	Description: `
	`,
	Action: doActivate,
}

func doActivate(c *cli.Context) {
	args := c.Args()
	if len(args) == 1 {
		environment.GetGlobalEnv(args[0]).Activate()
	} else {
		util.PrintErrorMessage("needs 1 arg")
	}
}
