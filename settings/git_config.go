package settings

import (
	"fmt"
	"github.com/ryo33/zenv/storage"
	"github.com/ryo33/zenv/util"
	"strings"
)

const (
	GIT_CONFIG = "git-config"
)

var gitConfig = setting{
	read:       readGitConfig,
	write:      writeGitConfig,
	activate:   activateGitConfig,
	deactivate: deactivateGitConfig,
	initialize: initialize,
	equal:      equalGitConfig,
}

func readGitConfig(str string) []string {
	re := strings.Split(str, " ")
	if len(re) == 2 || len(re) == 3 { // global
		return re
	} else {
		return []string{}
	}
	return re
}

func writeGitConfig(args []string) string {
	if len(args) == 2 { // global
		return args[0] + " " + args[1]
	} else if len(args) == 3 { // local
		return args[0] + " " + args[1] + " " + args[2]
	} else {
		return ""
	}
}

func activateGitConfig(args []string, info *Info) bool {
	var global bool
	var dir, gitDir, name, value string
	if len(args) == 2 {
		global = true
		dir = util.GetCurrentPath()
		name = args[0]
		value = args[1]
	} else if len(args) == 3 {
		dir = args[0]
		name = args[1]
		value = args[2]
	} else {
		return false
	}
	gitDir = getGitDir(dir)

	var preValue string
	var err error
	if global {
		preValue, err = util.ExecCommand("git", "config", "--global", name)
	} else {
		preValue, err = util.ExecCommand("git", gitDir, "config", name)
	}
	if err != nil {
		util.PrintErrorMessageContinue(fmt.Sprintf("can't change config %s to %s in %s", name, value, dir))
		return false
	}
	escapedDir := escapeDir(dir)
	values := storage.ReadStorage(SETTINGS, GIT_CONFIG, escapedDir)
	storage.WriteStorage(SETTINGS, GIT_CONFIG, escapedDir, append(values, info.envdir+"="+preValue))
	if global {
		_, err = util.ExecCommand("git", "config", "--global", name, value)
	} else {
		_, err = util.ExecCommand("git", gitDir, "config", name, value)
	}
	if err != nil {
		util.PrintErrorMessageContinue(fmt.Sprintf("can't change config %s to %s in %s", name, value, dir))
		return false
	}
	return true
}

func deactivateGitConfig(args []string, info *Info) bool {
	var global bool
	var dir, gitDir, name string
	if len(args) == 2 {
		global = true
		name = args[0]
	} else if len(args) == 3 {
		dir = args[0]
		gitDir = getGitDir(dir)
		name = args[1]
	} else {
		return false
	}
	escapedDir := escapeDir(dir)
	values := storage.ReadStorage(SETTINGS, GIT_CONFIG, escapedDir)
	prefix := info.envdir + "="
	for i, line := range values {
		if strings.HasPrefix(line, prefix) {
			var err error
			if global {
				_, err = util.ExecCommand("git", "config", "--global", name, strings.TrimPrefix(line, prefix))
			} else {
				_, err = util.ExecCommand("git", gitDir, "config", name, strings.TrimPrefix(line, prefix))
			}
			if err != nil {
				util.PrintErrorsContinue(err)
			}
			storage.WriteStorage(SETTINGS, GIT_CONFIG, escapedDir, append(values[:i], values[i+1:]...))
			break
		}
	}
	return true
}

func equalGitConfig(a []string, b []string) bool {
	if len(a) == 2 && len(b) == 2 && a[0] == b[0] {
		return true
	}
	if len(a) == 3 && len(b) == 3 && a[1] == b[1] {
		return true
	}
	return false
}
