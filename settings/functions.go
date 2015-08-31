package settings

import (
	"path"
)

func initialize(dir string)    {}
func equal(a, b []string) bool { return false }

func getGitDir(dir string) string {
	return "--git-dir=" + path.Join(dir, ".git")
}
