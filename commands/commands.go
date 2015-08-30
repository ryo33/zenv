package commands

import (
	"github.com/codegangsta/cli"
	"github.com/ryo33/zenv/environment"
)

var Commands = []cli.Command{
	global,
	local,
	activate,
	deactivate,
	system,
	link,
	git,
}

var Env *environment.Env

func GetEnv() *environment.Env {
	if Env != nil {
		return Env
	}
	return environment.GetCurrentEnv()
}
