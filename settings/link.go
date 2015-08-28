package settings

import (
	"github.com/ryo33/zenv/util"
	"os"
	"path"
	"strings"
)

const (
	VBIN = "vbin" // Dir
)

var link = Setting{
	read:       readLink,
	write:      writeLink,
	activate:   activateLink,
	deactivate: deactivateLink,
	initialize: initializeLink,
}

func initializeLink(dir string) {
	pa := GetVBinPath(dir)
	util.RemoveDir(pa)
	util.PrepareDir(pa)
}

func readLink(str string) []string {
	re := strings.Split(str, "=")
	if len(re) != 2 {
		return []string{}
	}
	return re
}

func writeLink(link []string) string {
	if len(link) != 2 {
		return ""
	}
	return link[0] + "=" + link[1]
}

func activateLink(link []string, info *Info) bool {
	if len(link) != 2 {
		return false
	}
	util.ExecCommand("ln", link[1], info.GetLinkPath(link[0]))
	return true
}

func deactivateLink(link []string, info *Info) bool {
	if len(link) != 2 {
		return false
	}
	os.Remove(info.GetLinkPath(link[0]))
	return true
}

func GetVBinPath(dir string) string {
	return path.Join(dir, VBIN)
}

func (info *Info) GetLinkPath(link string) string {
	return path.Join(info.zenv, VBIN, link)
}
