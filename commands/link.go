package commands

import (
	"github.com/codegangsta/cli"
	"github.com/ryo33/zenv/environment"
	"github.com/ryo33/zenv/util"
)

var link = cli.Command{
	Name: "link",
	Description: `
	add to path
	`,
	Flags: []cli.Flag{
		cli.BoolFlag{
			Name:  "force, f",
			Usage: "overwrite links",
		},
		cli.BoolFlag{
			Name:  "remove, r",
			Usage: "remove links",
		},
	},
	Action: doLink,
}

func doLink(c *cli.Context) {
	args := c.Args()
	if c.Bool("remove") {
		env := environment.GetCurrentEnv()
		env.RemoveItems("link", removeLink, args)
		env.Write()
	} else if len(args) == 0 {
		env := environment.GetCurrentEnv()
		env.ReadSettings()
		for _, link := range env.GetItems("link") {
			util.Print(link[0] + " " + link[1])
		}
	} else if len(args) != 2 {
		util.PrintErrorMessage("needs 2 args")
	} else {
		env := environment.GetCurrentEnv()
		env.AddItems("link", []string{args[0], args[1]})
		env.Write()
	}
}

func removeLink(it []string, param []string) bool {
	if it[0] == param[0] {
		return true
	}
	return false
}
