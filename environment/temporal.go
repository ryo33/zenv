package environment

import (
	"github.com/ryo33/zenv/util"
	"path"
	"time"
)

const (
	remove = 30    // 30 days
	TMP    = "tmp" // file
	LAST   = ".last"
)

func Now() string {
	return time.Now().Format(time.ANSIC)
}

func getTemporalPath() string {
	return path.Join(util.GetHomeDir(), ZENV, TMP)
}

func getTemporalFilePath(key, subkey string) string {
	return path.Join(util.GetHomeDir(), ZENV, TMP, subkey, key)
}

func getTemporalDir(subkey string) string {
	return path.Join(util.GetHomeDir(), ZENV, TMP, subkey)
}

func updateLast(subkey string) {
	if util.Exists(getTemporalDir(subkey)) {
		util.WriteFile(getTemporalFilePath(LAST, subkey), []string{Now()})
	}
}

func readTemporal(key, subkey string) []string {
	updateLast(subkey)
	fi := util.ReadFile(getTemporalFilePath(key, subkey))
	tmp := util.Remove(fi, "")
	return tmp
}

func writeTemporal(key, subkey string, data []string) {
	util.PrepareDir(getTemporalDir(subkey))
	util.WriteFile(getTemporalFilePath(key, subkey), data)
}

func clearTemporal() {
	now := time.Now()
	for _, dir := range util.GetAllDir(getTemporalPath()) {
		last := util.ReadFile(getTemporalFilePath(LAST, dir))
		if len(last) != 1 {
			updateLast(dir)
			continue
		}
		ti, err := time.Parse(time.ANSIC, last[0])
		if err == nil {
			if int(now.Sub(ti).Hours())/24 > remove {
				util.RemoveDir(getTemporalDir(dir))
			}
		} else {
			updateLast(dir)
		}
	}
}
