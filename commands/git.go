package commands

import (
	"github.com/codegangsta/cli"
	"github.com/ryo33/zenv/util"
	"strings"
)

const (
	preGit = "git-"
)

var git = cli.Command{
	Name: "git",
	Description: `
	add to path
	`,
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:  "directory, d",
			Usage: "specify the git directory",
		},
		cli.BoolFlag{
			Name:  "force, f",
			Usage: "overwrite gits",
		},
		cli.BoolFlag{
			Name:  "remove, r",
			Usage: "remove gits",
		},
	},
	Action: doGit,
}

var subCommands = []string{
	"checkout",
	"config",
}

func doGit(c *cli.Context) {
	args := c.Args()
	argc := len(args)
	var dir, sub string
	if dir = c.String("directory"); len(dir) > 0 {
		dir = util.ExpandPath(dir)
	} else {
		dir = util.GetCurrentPath()
	}
	if argc == 0 {
		printArgumentMoreError(1)
	}
	sub = args[0]
	args = args[1:]
	argc--
	env := GetEnv()
	env.ReadSettings()
	if !util.Contains(subCommands, sub) {
		util.PrintErrorMessage(sub + " is undefined")
	}
	// list settings
	if argc == 0 {
		switch sub {
		default:
			for _, i := range env.GetItems(preGit + sub) {
				util.Print(strings.Join(i, " "))
			}
		}
	} else if c.Bool("remove") {
		switch sub {
		default:
			env.RemoveItems(preGit+sub, removeGit1, util.Wrap(args))
		}
		env.Write()
	} else {
		switch sub {
		// 2 value
		case "config":
			if argc != 2 {
				printArgumentError(2)
			}
			env.AddItems(preGit+sub, c.Bool("force"), []string{dir, args[0], args[1]})
		// 1 value
		default:
			if argc != 1 {
				printArgumentError(1)
			}
			env.AddItems(preGit+sub, c.Bool("force"), []string{dir, args[0]})
		}
		env.Write()
	}
}

// match 0 value
func removeGit0(it []string, param []string) bool {
	// [0] is dir
	if it[0] == param[0] {
		return true
	}
	return false
}

// match 1 value
func removeGit1(it []string, param []string) bool {
	// [0] is dir
	if it[0] == param[0] && it[1] == param[1] {
		return true
	}
	return false
}
