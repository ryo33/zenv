package git

import (
	"github.com/codegangsta/cli"
	"github.com/ryo33/zenv/util"
	"strings"
)

const (
	GIT_CONFIG = "git-config"
)

var Config = cli.Command{
	Name: "config",
	Description: `
	git config
	`,
	Flags: []cli.Flag{
		cli.BoolFlag{
			Name:  "global, g",
			Usage: "global config",
		},
		cli.BoolFlag{
			Name:  "force, f",
			Usage: "overwrite",
		},
		cli.BoolFlag{
			Name:  "remove, r",
			Usage: "remove",
		},
	},
	Action: doConfig,
}

func doConfig(c *cli.Context) {
	args := c.Args()
	argc := len(args)
	Env.ReadSettings()
	// list settings
	if argc == 0 {
		for _, i := range Env.GetItems(GIT_CONFIG) {
			util.Print(strings.Join(i, " "))
		}
	} else if c.Bool("remove") {
		p := [][]string{}
		var global bool
		if c.Bool("global") {
			global = true
		}
		for _, a := range args {
			if global {
				p = append(p, []string{a})
			} else {
				p = append(p, []string{Dir, a})
			}
		}
		Env.RemoveItems(GIT_CONFIG, removeConfig, p)
		Env.Write()
	} else {
		if argc != 2 {
			util.PrintArgumentError(2)
		}
		if c.Bool("global") {
			Env.AddItems(GIT_CONFIG, c.Bool("force"), []string{args[0], args[1]})
		} else {
			Env.AddItems(GIT_CONFIG, c.Bool("force"), []string{Dir, args[0], args[1]})
		}
		Env.Write()
	}
}

func removeConfig(it []string, param []string) bool {
	if len(param) == 2 { // local
		if len(it) == 3 && it[0] == param[0] && it[1] == param[1] {
			return true
		}
	} else if len(param) == 1 { // global
		if len(it) == 2 && it[0] == param[0] {
			return true
		}
	}
	return false
}
