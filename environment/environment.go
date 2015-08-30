package environment

import (
	"github.com/ryo33/zenv/settings"
	"github.com/ryo33/zenv/storage"
	"github.com/ryo33/zenv/util"
	"github.com/ryo33/zenv/zenv"
	"path"
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
	ACTIVATED = "activated"
)

func GetGlobalEnv(name string) *Env {
	env := getGlobalEnv(name)
	if env == nil {
		util.PrintErrorMessage("not exists")
	}
	return env
}

// it is almost same as getLocalEnv but never return nil
func GetLocalEnv(name string) *Env {
	env := getLocalEnv(name)
	if env == nil {
		util.PrintErrorMessage("not exists")
	}
	return env
}

func getGlobalEnv(name string) *Env {
	if !ExistsGlobalEnv(name) {
		util.PrintErrorMessage("not exists global env")
	}
	return read(getGlobalPath(name))
}

func getLocalEnv(dir string) *Env {
	return read(getLocalPath(dir))
}

func read(name string) *Env {
	util.PrepareDir(path.Join(util.GetHomeDir(), zenv.ZENV, ENVS))
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

func Clean(current string) {
	settings.Initialize(getZenvPath())
	tmpPath := storage.GetTemporalPath()
	for _, dir := range util.GetAllDir(tmpPath) {
		if current != dir {
			activated := util.ReadFile(path.Join(tmpPath, dir, ACTIVATED))
			for _, act := range activated {
				env := getLocalEnv(act)
				if env != nil {
					env.ReadSettings()
					env.items.Clean(settings.NewInfo(getZenvPath(), env.dir))
				}
			}
			util.RemoveDir(path.Join(tmpPath, dir))
		}
	}
}

func Activate(pid string, dir string) {
	util.PrepareDir(getZenvPath())
	envs := getEnvs(dir)
	settings.Initialize(getZenvPath())
	for _, env := range envs {
		env.ReadSettings()
		env.Activate(pid)
	}
}

func Deactivate(pid string, dir string) {
	envs := getEnvs(dir)
	for _, env := range envs {
		env.ReadSettings()
		env.Deactivate(pid)
	}
}

func (env *Env) GetPath(sub string) string {
	return path.Join(env.dir, sub)
}

func getGlobalPath(name string) string {
	return path.Join(util.GetHomeDir(), zenv.ZENV, ENVS, name)
}

func getLocalPath(name string) string {
	return path.Join(name, ZENV_LOCAL)
}

func getZenvPath() string {
	return path.Join(util.GetHomeDir(), zenv.ZENV)
}

func GetActivated(pid string) []string {
	return storage.ReadTemporal(ACTIVATED, pid)
}

func isActivated(activated []string, name string) bool {
	for _, ac := range activated {
		if ac == name {
			return true
		}
	}
	return false
}

func (env *Env) Activate(pid string) {
	storage.ClearTemporal()
	activated := GetActivated(pid)
	if !isActivated(activated, env.name) {
		env.items.Activate(settings.NewInfo(getZenvPath(), env.dir))
		//TODO activate child envs
	}
	//Add to list
	storage.WriteTemporal(ACTIVATED, pid, append(activated, env.name))
}

func (env *Env) Deactivate(pid string) {
	//Remove from list
	activated := GetActivated(pid)
	for i, actName := range activated {
		if actName == env.name {
			activated = append(activated[:i], activated[i+1:]...)
			break
		}
	}
	storage.WriteTemporal(ACTIVATED, pid, activated)

	if !isActivated(activated, env.name) {
		env.items.Deactivate(settings.NewInfo(getZenvPath(), env.dir))
		//TODO deactivate child envs
	}
	if len(activated) == 0 {
		util.RemoveDir(storage.GetTemporalDir(pid))
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

func (env *Env) RemoveItems(lable string, remove func([]string, []string) bool, param [][]string) {
	its := env.GetItems(lable)
	if len(its) > 0 {
		result := [][]string{}
		for _, it := range its {
			rm := false
			for _, pa := range param {
				if remove(it, pa) {
					rm = true
				}
			}
			if !rm {
				result = append(result, it)
			}
		}
		env.items.ToMap()[lable] = result
	}
}
