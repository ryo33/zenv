package commands

import (
	"github.com/codegangsta/cli"
	"github.com/ryo33/zenv/environment"
	"github.com/ryo33/zenv/util"
)

var global = cli.Command{
	Name:  "global",
	Usage: "make global environment",
	Description: `
	`,
	Flags: []cli.Flag{
		cli.BoolFlag{
			Name:  "remove",
			Usage: "",
		},
	},
	Action: doGlobal,
}

func doGlobal(c *cli.Context) {
	args := c.Args()
	if c.Bool("remove") {
		if len(args) != 1 {
			util.PrintErrorMessage("needs only one arg")
		}
		environment.Remove(args[0])
	} else {
		//create new global environment
		name := args[0]
		env := environment.NewEnv(true, name, c.Bool("recursive"), c.Bool("exclusive"))
		env.Write()
	}
}
