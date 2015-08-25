package util

import (
	"io/ioutil"
	"os"
	"strings"
)

func ReadFile(name string) []string {
	str := ""
	bytes, err := ioutil.ReadFile(name)
	if err == nil || os.IsNotExist(err) {
		PrintErrors(err)
		str = string(bytes)
	}
	return strings.Split(str, "\n")
}

func WriteFile(name string, data []string) {
	err := ioutil.WriteFile(name, []byte(strings.Join(data, "\n")), 0777)
	PrintErrors(err)
}
