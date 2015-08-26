package commands

import (
	"github.com/codegangsta/cli"
	"github.com/ryo33/zenv/environment"
	"github.com/ryo33/zenv/util"
)

var local = cli.Command{
	Name:  "local",
	Usage: "initialize local environment",
	Description: `
	`,
	Flags: []cli.Flag{
		cli.BoolFlag{
			Name:  "force, f",
			Usage: "force to initialize",
		},
	},
	Action: doLocal,
}

func doLocal(c *cli.Context) {
	pwd := util.GetCurrentPath()
	if environment.ExistsLocalEnv(pwd) && !c.Bool("force") {
		util.PrintErrorMessage(`
		.zenv already exists
		--force flag to force to initialize
		`)
	} else {
		env := environment.NewEnv(false, pwd, c.Bool("recursive"), c.Bool("exclusive"))
		env.Write()
	}
}
