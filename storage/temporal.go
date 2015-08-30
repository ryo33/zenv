package storage

import (
	"github.com/ryo33/zenv/util"
	"github.com/ryo33/zenv/zenv"
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
	return path.Join(util.GetHomeDir(), zenv.ZENV, TMP, subkey, key)
}

func GetTemporalDir(subkey string) string {
	return path.Join(util.GetHomeDir(), zenv.ZENV, TMP, subkey)
}

func updateModified(pid string) {
	if util.Exists(GetTemporalDir(pid)) {
		util.WriteFile(getTemporalPath(MODIFIED, pid), []string{Now()})
	}
}

func ReadTemporal(key, subkey string) []string {
	fi := util.ReadFile(getTemporalPath(key, subkey))
	tmp := util.Remove(fi, "")
	return tmp
}

func WriteTemporal(key, subkey string, data []string) {
	util.PrepareDir(GetTemporalDir(subkey))
	util.WriteFile(getTemporalPath(key, subkey), data)
}
