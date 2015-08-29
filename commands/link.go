package commands

import (
	"github.com/codegangsta/cli"
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
		env := GetEnv()
		env.ReadSettings()
		env.RemoveItems("link", removeLink, args)
		env.Write()
	} else if len(args) == 0 {
		env := GetEnv()
		env.ReadSettings()
		for _, link := range env.GetItems("link") {
			util.Print(link[0] + " " + link[1])
		}
	} else if len(args) != 2 {
		printArgumentError(2)
	} else {
		env := GetEnv()
		env.ReadSettings()
		env.AddItems("link", c.Bool("force"), []string{args[0], args[1]})
		env.Write()
	}
}

func removeLink(it []string, param []string) bool {
	if it[0] == param[0] {
		return true
	}
	return false
}
