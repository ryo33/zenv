package commands

import (
	"github.com/codegangsta/cli"
	"github.com/ryo33/zenv/commands/git"
	"github.com/ryo33/zenv/util"
)

var git_ = cli.Command{
	Name: "git",
	Description: `
	add to path
	`,
	Before: beforeGit,
	Subcommands: []cli.Command{
		git.Checkout,
		git.Config,
	},
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:  "directory, d",
			Usage: "specify the git directory",
		},
	},
	Action: doGit,
}

func doGit(c *cli.Context) {
}

func beforeGit(c *cli.Context) error {
	if dir := c.String("directory"); len(dir) > 0 {
		git.Dir = util.ExpandPath(dir)
	} else {
		git.Dir = util.GetCurrentPath()
	}
	git.Env = GetEnv()
	return nil
}
