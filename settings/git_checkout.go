package settings

import (
	"strings"
)

var gitCheckout = Setting{
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

func writeGitCheckout(link []string) string {
	if len(link) != 2 {
		return ""
	}
	return link[0] + " " + link[1]
}

func activateGitCheckout(link []string, info *Info) bool {
	if len(link) != 2 {
		return false
	}
	// TODO
	return true
}

func deactivateGitCheckout(link []string, info *Info) bool {
	if len(link) != 2 {
		return false
	}
	// TODO
	return true
}

func equalGitCheckout(a []string, b []string) bool {
	if len(a) >= 1 && len(b) >= 1 && a[0] == b[0] {
		return true
	} else {
		return false
	}
}
