package environment

import (
	"fmt"
	"github.com/ryo33/zenv/util"
	"github.com/ryo33/zenv/zenv"
	"path"
	"sort"
	"strings"
)

func getEnvDirs() []string {
	return util.ReadFile(path.Join(util.GetHomeDir(), zenv.ZENV, DIRS))
}

func (env *Env) addEnvDir() {
	dirs := getEnvDirs()
	if !util.Contains(dirs, env.name) {
		util.WriteFile(path.Join(util.GetHomeDir(), zenv.ZENV, DIRS), append(dirs, env.name))
	}
}

func removeEnvDir(name string) {
	dirs := getEnvDirs()
	for i, dir := range dirs {
		if name == dir {
			util.WriteFile(path.Join(util.GetHomeDir(), zenv.ZENV, DIRS), append(dirs[:i], dirs[i+1:]...))
			return
		}
	}
}

func getEnvs(dir string) []*Env {
	tmp := getEnvDirs()
	dirs := []string{}
	for _, tmpDir := range tmp {
		if strings.HasPrefix(dir, tmpDir) {
			dirs = append(dirs, tmpDir)
		}
	}
	sort.Sort(ByLength(dirs))
	envs := []*Env{}
	for i := len(dirs) - 1; i >= 0; i-- {
		name := dirs[i]
		var env = getLocalEnv(name)
		if env != nil {
			if dirs[i] == dir {
				envs = append(envs, env)
			} else if env.recursive {
				envs = append(envs, env)
			}
			if env.exclusive {
				break
			}
		} else {
			util.PrintErrorMessageContinue(fmt.Sprintf("%s not exists", name))
			if util.YNPrompt("remove?", false) {
				removeEnvDir(name)
			}
		}
	}
	return envs
}

type ByLength []string

func (b ByLength) Len() int {
	return len(b)
}

func (b ByLength) Swap(i, j int) {
	tmp := b[i]
	b[i] = b[j]
	b[j] = tmp
}

func (b ByLength) Less(i, j int) bool {
	return len(b[i]) < len(b[j])
}

func GetEnvsPath() string {
	return path.Join(util.GetHomeDir(), zenv.ZENV, ENVS, ENVS)
}

func GetGlovalEnvs() []string {
	envFile := GetEnvsPath()
	return util.ReadFile(envFile)
}

func (env *Env) AddGlobalEnv(force bool) {
	envFile := GetEnvsPath()
	tmp := util.ReadFile(envFile)
	for _, en := range tmp {
		if en == env.name {
			if !force {
				util.PrintErrorMessage("already exists\n--force to overwrite") //TODO already exists  --force to overwrite
			}
			break
		}
	}
	util.WriteFile(envFile, append(tmp, env.name))
}

func RemoveGlobalEnv(name string) {
	envFile := GetEnvsPath()
	tmp := util.ReadFile(envFile)
	result := []string{}
	for _, en := range tmp {
		if en != name {
			result = append(result, en)
		}
	}
	util.WriteFile(envFile, result)
}
