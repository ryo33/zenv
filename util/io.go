package util

import (
	"io/ioutil"
	"os"
	"strings"
)

func ReadFile(name string) []string {
	str := ""
	bytes, err := ioutil.ReadFile(name)
	if err == nil {
		str = string(bytes)
	} else if !os.IsNotExist(err) {
		PrintErrors(err)
	}
	return strings.Split(str, "\n")
}

func WriteFile(name string, data []string) {
	data = Remove(data, "")
	err := ioutil.WriteFile(name, []byte(strings.Join(data, "\n")), 0777)
	PrintErrors(err)
}
