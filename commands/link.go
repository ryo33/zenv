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
		env.RemoveLinks(args)
		env.Write()
	} else if len(args) == 0 {
		env := environment.GetCurrentEnv()
		for _, link := range env.Links() {
			util.Print(link.Name() + " " + link.Source())
		}
	} else if len(args) != 2 {
		util.PrintErrorMessage("needs 2 args")
	} else {
		env := environment.GetCurrentEnv()
		env.AddLink(environment.NewLink(args[0], args[1]), c.Bool("force"))
		env.Write()
	}
}
