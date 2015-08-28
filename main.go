package main

import (
	"github.com/codegangsta/cli"
	"github.com/ryo33/zenv/commands"
	"github.com/ryo33/zenv/environment"
	"github.com/ryo33/zenv/util"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Name = "zenv"
	app.Usage = "powerful environments management"
	app.Commands = commands.Commands
	app.Before = setup
	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:  "quiet, q",
			Usage: "do not print normal output",
		},
		cli.BoolFlag{
			Name:  "debug, d",
			Usage: "print debugging information",
		},
	}
	app.Action = doZenv
	app.Run(os.Args)
}

func setup(c *cli.Context) error {
	if c.Bool("quiet") {
		util.Quiet = true
	}
	if c.Bool("debug") {
		util.Debug = true
	}
	return nil
}

func doZenv(c *cli.Context) {
	for _, env := range environment.GetActivated() {
		util.Print(env)
	}
}
