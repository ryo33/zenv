package storage

import (
	"github.com/ryo33/zenv/util"
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

func updateLast(key string) {
	if util.Exists(GetStorageDir(TMP, key)) {
		util.WriteFile(getStorageFilePath(TMP, key, LAST), []string{Now()})
	}
}

func ReadTemporal(key, subkey string) []string {
	return ReadStorage(TMP, key, subkey)
}

func WriteTemporal(key, subkey string, data []string) {
	WriteStorage(TMP, key, subkey, data)
}

func ClearTemporal() {
	now := time.Now()
	for _, dir := range util.GetAllDir(GetStoragePath(TMP)) {
		last := util.ReadFile(getStorageFilePath(TMP, dir, LAST))
		if len(last) != 1 {
			updateLast(dir)
			continue
		}
		ti, err := time.Parse(time.ANSIC, last[0])
		if err == nil {
			if int(now.Sub(ti).Hours())/24 > remove {
				util.RemoveDir(GetStorageDir(TMP, dir))
			}
		} else {
			updateLast(dir)
		}
	}
}
