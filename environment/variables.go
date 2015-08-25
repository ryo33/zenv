package environment

import (
	"os"
	"strings"
)

const (
	BASE           = "ZENV_"
	ZENV_ACTIVATED = "ZENV_ACTIVATED"
	SEPARATOR      = ":"
)

func GetActivated() []string {
	return strings.Split(os.Getenv(ZENV_ACTIVATED), SEPARATOR)
}

func GetPath() []string {
	return strings.Split(os.Getenv("PATH"), ":")
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
