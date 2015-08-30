package storage

import (
	"github.com/ryo33/zenv/util"
	"github.com/ryo33/zenv/zenv"
	"path"
)

const (
	STORAGE = "storage"
)

func GetStoragePath(name string) string {
	if len(name) == 0 {
		name = STORAGE
	}
	return path.Join(util.GetHomeDir(), zenv.ZENV, name)
}

func getStorageFilePath(name, key, subkey string) string {
	if len(name) == 0 {
		name = STORAGE
	}
	return path.Join(util.GetHomeDir(), zenv.ZENV, name, key, subkey)
}

func GetStorageDir(name, key string) string {
	if len(name) == 0 {
		name = STORAGE
	}
	return path.Join(util.GetHomeDir(), zenv.ZENV, name, key)
}

func ReadStorage(name, key, subkey string) []string {
	if len(name) == 0 {
		name = STORAGE
	}
	fi := util.ReadFile(getStorageFilePath(name, key, subkey))
	return util.Remove(fi, "")
}

func WriteStorage(name, key, subkey string, data []string) {
	if len(name) == 0 {
		name = STORAGE
	}
	util.PrepareDir(GetStorageDir(name, key))
	util.WriteFile(getStorageFilePath(name, key, subkey), data)
}
