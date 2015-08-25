package commands

import (
	"github.com/codegangsta/cli"
)

var Commands = []cli.Command{
	global,
	local,
	activate,
	deactivate,
	system,
	edit,
}
