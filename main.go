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
	app.Before = commands.ZenvSetup
	app.Flags = commands.ZenvFlag
	app.Action = commands.DoZenv
	app.Run(os.Args)
}
