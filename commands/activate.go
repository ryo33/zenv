package commands

import (
	"github.com/codegangsta/cli"
	"github.com/ryo33/zenv/environment"
)

var activate = cli.Command{
	Name:  "activate",
	Usage: "activate the environment",
	Description: `
	`,
	Action: doActivate,
}

func doActivate(c *cli.Context) {
	util.Eval = true
	args := c.Args()
	if len(args) == 2 {
		environment.GetGlobalEnv(args[1]).Activate(args[0])
	} else {
		printArgumentError(2)
	}
}
