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
	links     []*Link
}

const (
	SEPARATOR = "="
)

func getGlobalEnv(name string) *Env {
	if !ExistsGlobalEnv(name) {
		//TODO error
	}
	return read(getGlobalPath(name))
}

func getLocalEnv(dir string) *Env {
	return read(getLocalPath(dir))
}

func read(name string) *Env {
	env := readInfo(name)
	env.readLinks()
	return env
}

func (env *Env) Write() {
	pa := path.Join(util.GetHomeDir(), ZENV)
	util.PrepareDir(pa)
	util.RemoveDir(env.dir)
	util.PrepareDir(env.dir)
	env.writeInfo()
	env.writeLinks()
	if env.global {
		env.addGlobalEnv()
	} else {
		env.addEnvDir()
	}
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
		links:     []*Link{},
	}
	return env
}

func GetCurrentEnv() *Env {
	return read(path.Join(util.GetCurrentPath(), ZENV))
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
