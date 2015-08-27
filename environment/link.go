package environment

import (
	"github.com/ryo33/zenv/util"
	"path"
)

const (
	VBIN = "vbin" // Dir
)

func GetVBinDir() string {
	return path.Join(util.GetHomeDir(), VBIN)
}

/* Activate
func (env *Env) writeLinks() {
	for _, li := range env.links {
		//TODO is there smart way?
		util.ExecCommand("ln", li.source, path.Join(linksPath, li.name))
	}
}
*/
