package environment

import (
	"github.com/ryo33/zenv/util"
	"path"
	"sort"
	"strings"
)

func getEnvDirs() []string {
	return util.ReadFile(path.Join(util.GetHomeDir(), ZENV, DIRS))
}

func (env *Env) addEnvDir() {
	dirs := getEnvDirs()
	if !util.Contains(dirs, env.name) {
		util.WriteFile(path.Join(util.GetHomeDir(), ZENV, DIRS), append(dirs, env.name))
	}
}

func getEnvs(dir string) []*Env {
	tmp := getEnvDirs()
	dirs := []string{}
	for _, tmpDir := range tmp {
		if strings.HasPrefix(tmpDir, dir) {
			dirs = append(dirs, tmpDir)
		}
	}
	sort.Sort(ByLength(dirs))
	envs := []*Env{}
	for i := len(dirs) - 1; i >= 0; i-- {
		var env = GetLocalEnv(dirs[i])
		if env != nil {
			if dirs[i] == dir {
				envs = append(envs, env)
			}
			if env.exclusive {
				break
			}
			if env.recursive {
				envs = append(envs, env)
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
	return path.Join(util.GetHomeDir(), ZENV, ENVS, ENVS)
}

func GetGlovalEnvs() []string {
	envFile := GetEnvsPath()
	return util.ReadFile(envFile)
}

func (env *Env) addGlobalEnv() {
	envFile := GetEnvsPath()
	tmp := util.ReadFile(envFile)
	for _, en := range tmp {
		if en == env.name {
			util.PrintErrorMessage("") //TODO already exists  --force to overwrite
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
