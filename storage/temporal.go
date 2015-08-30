package storage

import (
	"github.com/ryo33/zenv/util"
	"github.com/ryo33/zenv/zenv"
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

func GetTemporalPath() string {
	return path.Join(util.GetHomeDir(), zenv.ZENV, TMP)
}

func getTemporalFilePath(key, subkey string) string {
	return path.Join(util.GetHomeDir(), zenv.ZENV, TMP, subkey, key)
}

func GetTemporalDir(subkey string) string {
	return path.Join(util.GetHomeDir(), zenv.ZENV, TMP, subkey)
}

func updateLast(subkey string) {
	if util.Exists(GetTemporalDir(subkey)) {
		util.WriteFile(getTemporalFilePath(LAST, subkey), []string{Now()})
	}
}

func ReadTemporal(key, subkey string) []string {
	updateLast(subkey)
	fi := util.ReadFile(getTemporalFilePath(key, subkey))
	tmp := util.Remove(fi, "")
	return tmp
}

func WriteTemporal(key, subkey string, data []string) {
	util.PrepareDir(GetTemporalDir(subkey))
	util.WriteFile(getTemporalFilePath(key, subkey), data)
}

func ClearTemporal() {
	now := time.Now()
	for _, dir := range util.GetAllDir(GetTemporalPath()) {
		last := util.ReadFile(getTemporalFilePath(LAST, dir))
		if len(last) != 1 {
			updateLast(dir)
			continue
		}
		ti, err := time.Parse(time.ANSIC, last[0])
		if err == nil {
			if int(now.Sub(ti).Hours())/24 > remove {
				util.RemoveDir(GetTemporalDir(dir))
			}
		} else {
			updateLast(dir)
		}
	}
}
