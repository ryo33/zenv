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
	return Remove(strings.Split(str, "\n"), "")
}

func WriteFile(name string, data []string) {
	data = Remove(data, "")
	err := ioutil.WriteFile(name, []byte(strings.Join(data, "\n")), 0777)
	PrintErrors(err)
}

func RemoveDir(name string) {
	PrintErrors(os.RemoveAll(name))
}

func GetAllDir(dir string) []string {
	entries, err := ioutil.ReadDir(dir)
	PrintErrors(err)
	result := []string{}
	for _, entry := range entries {
		if entry.IsDir() {
			result = append(result, entry.Name())
		}
	}
	return result
}
