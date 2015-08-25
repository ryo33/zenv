package main

import (
	"github.com/codegangsta/cli"
	"github.com/ryo33/zenv/commands"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Name = "zenv"
	app.Usage = "powerful environments management"
	app.Commands = commands.Commands
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
	app.Action = func(c *cli.Context) {
	}
	app.Run(os.Args)
}
