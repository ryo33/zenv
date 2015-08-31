package git

import (
	"github.com/codegangsta/cli"
	"github.com/ryo33/zenv/util"
	"strings"
)

const (
	GIT_CHECKOUT = "git-checkout"
)

var Checkout = cli.Command{
	Name: "checkout",
	Description: `
	git checkout
	`,
	Flags: []cli.Flag{
		cli.BoolFlag{
			Name:  "force, f",
			Usage: "overwrite",
		},
		cli.BoolFlag{
			Name:  "remove, r",
			Usage: "remove",
		},
	},
	Action: doCheckout,
}

func doCheckout(c *cli.Context) {
	args := c.Args()
	argc := len(args)
	Env.ReadSettings()
	// list settings
	if argc == 0 {
		for _, i := range Env.GetItems(GIT_CHECKOUT) {
			util.Print(strings.Join(i, " "))
		}
	} else if c.Bool("remove") {
		p := [][]string{}
		for _, a := range args {
			p = append(p, []string{util.ExpandPath(a)})
		}
		Env.RemoveItems(GIT_CHECKOUT, remove0, p)
		Env.Write()
	} else {
		if argc != 1 {
			util.PrintArgumentError(1)
		}
		Env.AddItems(GIT_CHECKOUT, c.Bool("force"), []string{Dir, args[0]})
		Env.Write()
	}
}
