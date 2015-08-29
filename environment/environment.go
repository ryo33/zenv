package environment

import (
	"github.com/ryo33/zenv/settings"
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
	items     settings.Items
}

const (
	SEPARATOR = "="
)

func GetGlobalEnv(name string) *Env {
	if !ExistsGlobalEnv(name) {
		util.PrintErrorMessage("not exists global env")
	}
	return read(getGlobalPath(name))
}

func GetLocalEnv(dir string) *Env {
	return read(getLocalPath(dir))
}

func read(name string) *Env {
	util.PrepareDir(path.Join(util.GetHomeDir(), ZENV, ENVS))
	env := readInfo(name)
	return env
}

func (env *Env) ReadSettings() {
	env.items = settings.Read(env.dir)
}

func (env *Env) Write() {
	util.RemoveDir(env.dir)
	util.PrepareDir(env.dir)
	util.PrepareDir(getZenvPath())
	util.PrepareDir(getEnvsPath())
	env.writeInfo()
	env.items.Write(env.dir)
	if !env.global {
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
		items:     make(settings.Items),
	}
	return env
}

func GetCurrentEnv() *Env {
	return read(getLocalPath(util.GetCurrentPath()))
}

func Activate(dir string) {
	util.PrepareDir(getZenvPath())
	envs := getEnvs(dir)
	settings.Initialize(getZenvPath())
	for _, env := range envs {
		env.ReadSettings()
		env.Activate()
	}
}

func Deactivate(dir string) {
	envs := getEnvs(dir)
	for _, env := range envs {
		env.ReadSettings()
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

func getZenvPath() string {
	return path.Join(util.GetHomeDir(), ZENV)
}

func (env *Env) Activate() {
	isActivated := IsActivated(env.name)
	if !isActivated {
		env.items.Activate(settings.NewInfo(getZenvPath(), env.dir))
		//TODO activate child envs
	}
	//Add to list
	util.Setenv(ZENV_ACTIVATED, strings.Join(append(GetActivated(), env.name), VAR_SEPARATOR))
}

func (env *Env) Deactivate() {
	//Remove from list
	activated := GetActivated()
	for i, actName := range activated {
		if actName == env.name {
			activated = append(activated[:i], activated[i+1:]...)
			break
		}
	}
	util.Setenv(ZENV_ACTIVATED, strings.Join(activated, VAR_SEPARATOR))

	if !IsActivated(env.name) {
		env.items.Deactivate(settings.NewInfo(getZenvPath(), env.dir))
		//TODO deactivate child envs
	}
}

func (env *Env) GetItems(lable string) [][]string {
	its, ok := env.items.ToMap()[lable]
	if ok {
		return its
	} else {
		return [][]string{}
	}
}

func (env *Env) AddItems(lable string, force bool, nits ...[]string) {
	env.items.AddItems(lable, force, nits)
}

func (env *Env) RemoveItems(lable string, remove func([]string, []string) bool, param []string) {
	its := env.GetItems(lable)
	if len(its) > 0 {
		result := [][]string{}
		for _, it := range its {
			if !remove(it, param) {
				result = append(result, it)
			}
		}
		env.items.ToMap()[lable] = result
	}
}
