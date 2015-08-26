package environment

import (
	"github.com/ryo33/zenv/util"
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

func GetGlobalEnv(name string) *Env {
	if !ExistsGlobalEnv(name) {
		//TODO error
	}
	return read(getGlobalPath(name))
}

func GetLocalEnv(dir string) *Env {
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
	return read(getLocalPath(util.GetCurrentPath()))
}

func Activate(dir string) {
	envs := getEnvs(dir)
	for _, env := range envs {
		env.Activate()
	}
}

func Deactivate(dir string) {
	envs := getEnvs(dir)
	for _, env := range envs {
		env.Deactivate()
	}
}

func (env *Env) GetPath(sub string) string {
	return path.Join(env.dir, sub)
}

func getGlobalPath(name string) string {
	return path.Join(util.GetHomeDir(), ZENV, ENVS, name)
}

func getLocalPath(name string) string {
	return path.Join(name, ZENV_LOCAL)
}

func (env *Env) Activate() {
	util.PrintDebug("[activate] " + env.name)
	isActivated := IsActivated(env.name)
	if !isActivated {
		//Add to path
		path := GetPath()
		path = append([]string{env.GetLinksPath()}, path...)
		util.Setenv("PATH", strings.Join(path, ":"))
	}
	//Add to list
	util.Setenv(ZENV_ACTIVATED, strings.Join(append(GetActivated(), env.GetLinksPath()), VAR_SEPARATOR))
	if !isActivated {
		//TODO activate child envs
	}
}

func (env *Env) Deactivate() {
	util.PrintDebug("[deactivate] " + env.name)
	//Remove from list
	activated := GetActivated()
	for i, actName := range activated {
		if actName == env.name {
			activated = append(activated[:i], activated[i+1:]...)
			break
		}
	}
	util.Setenv(ZENV_ACTIVATED, strings.Join(activated, VAR_SEPARATOR))

	newPath := []string{}
	if !IsActivated(env.name) {
		//Remove from path
		for _, p := range GetPath() {
			if p != env.GetLinksPath() {
				newPath = append(newPath, p)
			}
		}
	}
	util.Setenv("PATH", strings.Join(newPath, ":"))

	if !IsActivated(env.name) {
		//TODO deactivate child envs
	}
}
