package environment

import (
	"fmt"
	"github.com/ryo33/zenv/util"
	"path"
	"strings"
)

const (
	ZENV       = ".zenv"  // Dir
	ZENV_LOCAL = ".zenvl" // Dir
	DIRS       = "dirs"   // File
	ENVS       = "envs"   // Both
	LOADS      = "loads"  // File
	INFO       = "info"   // File
)

//Read env
func readEnv(pa string) *Env {
	data := util.ReadFile(path.Join(pa, INFO))
	var name, dir string
	var global, recursive, exclusive bool
	for _, d := range data {
		re := strings.Split(d, "=")
		if len(re) == 2 {
			switch re[0] {
			case "name":
				name = re[1]
			case "dir":
				dir = re[1]
			case "global":
				global = false
				if re[1] == "true" {
					global = true
				}
			case "recursive":
				recursive = false
				if re[1] == "true" {
					recursive = true
				}
			case "exclusive":
				exclusive = false
				if re[1] == "true" {
					exclusive = true
				}
			}
		}
	}
	if len(name) == 0 || len(dir) == 0 {
		util.PrintErrorMessage(fmt.Sprintf("Can't read environment info in %s", pa))
	}
	env := NewEnv(global, name, recursive, exclusive)
	//TODO read others
	return env
}

//writeEnv
func (env *Env) writeEnv() {
	info := append([]string{}, "name="+env.name, "dir="+env.dir)
	if env.global {
		info = append(info, "global=true")
	} else {
		info = append(info, "global=false")
	}
	if env.recursive {
		info = append(info, "recursive=true")
	} else {
		info = append(info, "recursive=false")
	}
	if env.exclusive {
		info = append(info, "exclusive=true")
	} else {
		info = append(info, "exclusive=false")
	}
	util.WriteFile(path.Join(env.dir, INFO), info)
	//TODO write others
}

//Read all global envs
func readEnvs() []string {
	return util.ReadFile(path.Join(util.GetHomeDir(), ZENV, ENVS))
}

//Exists global environment
func ExistsGlobalEnv(name string) bool {
	envs := readEnvs()
	for _, env := range envs {
		if env == name {
			return true
		}
	}
	return false
}

//Exists local environment
func ExistsLocalEnv(name string) bool {
	p := path.Join(name, ZENV)
	if util.Exists(p) {
		return true
	}
	return false
}
