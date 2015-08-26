package environment

import (
	"github.com/ryo33/zenv/util"
	"os"
	"strings"
)

const (
	BASE           = "ZENV_"
	ZENV_ACTIVATED = "ZENV_ACTIVATED"
	VAR_SEPARATOR  = ":"
)

func GetActivated() []string {
	return util.Remove(strings.Split(os.Getenv(ZENV_ACTIVATED), VAR_SEPARATOR), "")
}

func GetPath() []string {
	return util.Remove(strings.Split(os.Getenv("PATH"), ":"), "")
}

func IsActivated(name string) bool {
	for _, actName := range GetActivated() {
		if actName == name {
			return true
		}
	}
	return false
}

func Getenv(name string) string {
	return os.Getenv(BASE + name)
}

func Setenv(name, value string) {
	os.Setenv(BASE+name, value)
}
