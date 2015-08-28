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
		cli.BoolFlag{
			Name:  "force, f",
			Usage: "force to initialize",
		},
	},
	Action: doGlobal,
}

func doGlobal(c *cli.Context) {
	args := c.Args()
	if len(args) == 0 {
		for _, arg := range environment.GetGlovalEnvs() {
			util.Print(arg)
		}
	} else if c.Bool("remove") {
		for _, arg := range args {
			environment.RemoveGlobalEnv(arg)
		}
	} else {
		for _, arg := range args {
			//create new global environment
			env := environment.NewEnv(true, arg, c.Bool("recursive"), c.Bool("exclusive"))
			env.Write()
			env.AddGlobalEnv(c.Bool("force"))
		}
	}
}
