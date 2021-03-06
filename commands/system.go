package commands

import (
	"github.com/codegangsta/cli"
	"github.com/ryo33/zenv/environment"
	"github.com/ryo33/zenv/util"
)

var system = cli.Command{
	Name:  "system",
	Usage: "system command",
	Description: `
	`,
	Flags: []cli.Flag{
		cli.BoolFlag{
			Name:  "clean",
			Usage: "",
		},
		cli.BoolFlag{
			Name:  "cd-before",
			Usage: "",
		},
		cli.BoolFlag{
			Name:  "cd-after",
			Usage: "",
		},
		cli.BoolFlag{
			Name:  "cd",
			Usage: "",
		},
	},
	Action: doSystem,
}

func doSystem(c *cli.Context) {
	args := c.Args()
	if c.Bool("cd-before") {
		if len(args) == 2 {
			environment.Deactivate(args[0], args[1])
		} else {
			util.PrintArgumentError(2)
		}
	} else if c.Bool("cd-after") {
		if len(args) == 2 {
			environment.Activate(args[0], args[1])
		} else {
			util.PrintArgumentError(2)
		}
	} else if c.Bool("cd") {
		if len(args) == 3 {
			environment.Deactivate(args[0], args[1])
			environment.Activate(args[0], args[2])
		} else {
			util.PrintArgumentError(3)
		}
	} else if c.Bool("clean") {
		if len(args) == 1 {
			environment.Clean(args[0])
		} else {
			util.PrintArgumentError(1)
		}
	}
}
