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
	Flags:  []cli.Flag{},
	Action: doLocal,
}

func doLocal(c *cli.Context) {
	env := environment.NewEnv(false, util.GetCurrentPath(), c.Bool("recursive"), c.Bool("exclusive"))
	env.Write()
}
