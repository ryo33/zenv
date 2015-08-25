package environment

import (
	"github.com/ryo33/zenv/util"
	"os"
)

type Link struct {
	source string
	dest   string
}

const (
	LINKS = "links" // Dir
)

func (env *Env) GetLinksPath() string {
	return env.GetPath(LINKS)
}

func validate(env *Env, dest string, source string) {
	if IsExistLink(dest) {
		//TODO error
	}
	if !util.Exists(source) {
		//TODO error
	}
	info, err := os.Stat(source)
	util.PrintErrors(err)
	if info.Mode()&0111 == 0 {
		//TODO error
	}
}

func IsExistLink(dest string) bool {
	//TODO
	return false
}
