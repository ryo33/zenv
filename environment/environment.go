package environment

import (
	"github.com/ryo33/zenv/util"
	"os"
	"path"
	"strings"
)

type Env struct {
	name      string
	dir       string
	global    bool
	recursive bool
	exclusive bool
}

func getGlobalEnv(name string) *Env {
	if !ExistsGlobalEnv(name) {
		//TODO error
	}
	return readEnv(getGlobalPath(name))
}

func getLocalEnv(dir string) *Env {
	return readEnv(getLocalPath(dir))
}

func NewEnv(global bool, name string, recursive, exclusive bool) *Env {
	var dir string
	if global {
		dir = getGlobalPath(name)
	} else {
		dir = getLocalPath(name)
	}
	util.PrepareDir(dir)
	env := &Env{
		global:    global,
		name:      name,
		dir:       dir,
		recursive: recursive,
		exclusive: exclusive,
	}
	return env
}

func Remove(name string) {
	//TODO remove global environment
}

func Activate(dir string) {
	envs := getEnvs(dir)
	for _, env := range envs {
		env.activate()
	}
}

func Deactivate(dir string) {
	envs := getEnvs(dir)
	for _, env := range envs {
		env.deactivate()
	}
}

func (env *Env) GetPath(sub string) string {
	return path.Join(env.dir, sub)
}

func getGlobalPath(name string) string {
	return path.Join(util.GetHomeDir(), ZENV, ENVS, name)
}

func getLocalPath(name string) string {
	return path.Join(name, ZENV)
}

func (env *Env) Write() {
	util.PrepareDir(path.Join(util.GetHomeDir(), ZENV))
	if env.global {
		env.writeEnv()
		env.addGlobalEnv()
	} else {
		env.writeEnv()
		env.addEnvDir()
	}
}

func (env *Env) activate() {
	activated := IsActivated(env.name)
	if !activated {
		//Add to path
		os.Setenv("PATH", env.GetLinksPath()+":"+os.Getenv("PATH"))
	}
	//Add to list
	os.Setenv(ZENV_ACTIVATED, os.Getenv(ZENV_ACTIVATED)+VAR_SEPARATOR+env.GetLinksPath())

	if !activated {
		//TODO activate child envs
	}
}

func (env *Env) deactivate() {
	//Remove from list
	activated := GetActivated()
	for i, actName := range activated {
		if actName == env.name {
			activated = append(activated[:i], activated[i+1:]...)
			break
		}
	}
	os.Setenv(ZENV_ACTIVATED, strings.Join(activated, VAR_SEPARATOR))

	newPath := []string{}
	if !IsActivated(env.name) {
		//Remove from path
		for _, p := range GetPath() {
			if p != env.GetLinksPath() {
				newPath = append(newPath, p)
			}
		}
	}
	os.Setenv("PATH", strings.Join(newPath, ":"))

	if !IsActivated(env.name) {
		//TODO deactivate child envs
	}
}
