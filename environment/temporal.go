package environment

import (
	"github.com/ryo33/zenv/util"
	"path"
	"time"
)

const (
	remove   = 30    // 30 days
	TMP      = "tmp" // file
	MODIFIED = ".modified"
)

func Now() string {
	return time.Now().Format(time.ANSIC)
}

func getTemporalPath(key, subkey string) string {
	return path.Join(util.GetHomeDir(), ZENV, TMP, subkey, key)
}

func getTemporalDir(subkey string) string {
	return path.Join(util.GetHomeDir(), ZENV, TMP, subkey)
}

func updateModified(pid string) {
	if util.Exists(getTemporalDir(pid)) {
		util.WriteFile(getTemporalPath(MODIFIED, pid), []string{Now()})
	}
}

func readTemporal(key, subkey string) []string {
	fi := util.ReadFile(getTemporalPath(key, subkey))
	tmp := util.Remove(fi, "")
	return tmp
}

func writeTemporal(key, subkey string, data []string) {
	util.PrepareDir(getTemporalDir(subkey))
	util.WriteFile(getTemporalPath(key, subkey), data)
}
