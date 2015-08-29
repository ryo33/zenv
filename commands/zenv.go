package commands

import (
	"github.com/codegangsta/cli"
	"github.com/ryo33/zenv/environment"
	"github.com/ryo33/zenv/util"
)

var ZenvFlag = []cli.Flag{
	cli.StringFlag{
		Name:  "global, g",
		Usage: "edit global environment",
	},
	cli.StringFlag{
		Name:  "local, l",
		Usage: "edit local environment",
	},
	cli.BoolFlag{
		Name:  "quiet, q",
		Usage: "do not print normal output",
	},
	cli.BoolFlag{
		Name:  "debug, d",
		Usage: "print debugging information",
	},
}

func DoZenv(c *cli.Context) {
	args := c.Args()
	if len(args) == 1 {
		for _, env := range environment.GetActivated(args[0]) {
			util.Print(env)
		}
	} else {
		util.PrintErrorMessage("needs 1 arg")
	}
}

func ZenvSetup(c *cli.Context) error {
	if c.Bool("quiet") {
		util.Quiet = true
	}
	if c.Bool("debug") {
		util.Debug = true
	}
	if g := c.String("global"); len(g) != 0 {
		Env = environment.GetGlobalEnv(g)
	}
	if l := c.String("local"); len(l) != 0 {
		Env = environment.GetLocalEnv(l)
	}
	return nil
}
