package settings

import (
	"fmt"
	"github.com/ryo33/zenv/storage"
	"github.com/ryo33/zenv/util"
	"os"
	"strings"
)

const (
	GIT_CHECKOUT = "git-checkout"
)

var gitCheckout = setting{
	read:       readGitCheckout,
	write:      writeGitCheckout,
	activate:   activateGitCheckout,
	deactivate: deactivateGitCheckout,
	initialize: initialize,
	equal:      equalGitCheckout,
}

func readGitCheckout(str string) []string {
	re := strings.Split(str, " ")
	if len(re) != 2 {
		return []string{}
	}
	return re
}

func writeGitCheckout(args []string) string {
	if len(args) != 2 {
		return ""
	}
	return args[0] + " " + args[1]
}

func escapeDir(dir string) string {
	return strings.Join(strings.Split(dir, "/"), "-")
}

func activateGitCheckout(args []string, info *Info) bool {
	if len(args) != 2 {
		return false
	}
	defer os.Chdir(util.GetCurrentPath())
	os.Chdir(args[0])
	branch, err := util.ExecCommand("git", "rev-parse", "--abbrev-ref", "HEAD")
	if err != nil || len(branch) == 0 {
		util.PrintErrorMessageContinue("can't find branch message in " + args[0])
		return false
	}
	dir := escapeDir(args[0])
	branches := storage.ReadStorage(SETTINGS, GIT_CHECKOUT, dir)
	storage.WriteStorage(SETTINGS, GIT_CHECKOUT, dir, append(branches, info.envdir+"="+branch))
	_, err = util.ExecCommand("git", "checkout", args[1])
	if err != nil {
		util.PrintErrorMessageContinue(fmt.Sprintf("can't checkout to %s in %s", args[1], args[0]))
		return false
	}
	return true
}

func deactivateGitCheckout(args []string, info *Info) bool {
	if len(args) != 2 {
		return false
	}
	dir := escapeDir(args[0])
	branches := storage.ReadStorage(SETTINGS, GIT_CHECKOUT, dir)
	prefix := info.envdir + "="
	for i, line := range branches {
		if strings.HasPrefix(line, prefix) {
			defer os.Chdir(util.GetCurrentPath())
			os.Chdir(args[0])
			util.ExecCommand("git", "checkout", strings.TrimPrefix(line, prefix))
			storage.WriteStorage(SETTINGS, GIT_CHECKOUT, dir, append(branches[:i], branches[i+1:]...))
			break
		}
	}
	return true
}

func equalGitCheckout(a []string, b []string) bool {
	if len(a) >= 1 && len(b) >= 1 && a[0] == b[0] {
		return true
	} else {
		return false
	}
}
