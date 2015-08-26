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
		var env = getLocalEnv(dirs[i])
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

func (env *Env) addGlobalEnv() {
	//envFile := path.Join(env.dir, ENVS)
	//TODO
}
