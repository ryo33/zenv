package commands

import (
	"github.com/codegangsta/cli"
	"github.com/ryo33/zenv/environment"
	"github.com/ryo33/zenv/util"
)

var edit = cli.Command{
	Name:  "edit",
	Usage: "edit the environment",
	Description: `
	`,
	Flags: []cli.Flag{
		cli.BoolFlag{
			Name:  "finish",
			Usage: "",
		},
	},
	Action: doEdit,
}

func doEdit(c *cli.Context) {
	args := c.Args()
	if c.Bool("finish") {
		p := environment.Getenv("EDITING")
		if len(p) > 0 {
			//TODO cd to p
		} else {
			util.PrintErrorMessage("You're not editing.")
		}
	}
	if len(args) == 1 {
		//TODO util.Setenv("EDITING", `pwd`)
		//TODO cd to dir
		util.Print(c, "") //TODO zenv edit --finish
	} else if len(args) == 0 {
		//TODO list environments
	}
}
