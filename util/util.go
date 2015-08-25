package util

import (
	"fmt"
	"github.com/codegangsta/cli"
	"github.com/mitchellh/go-homedir"
	"os"
	"path/filepath"
)

func Print(c *cli.Context, str string) {
	if !c.Bool("quiet") {
		fmt.Printf(str + "\n")
	}
}

func PrintErrorMessage(err string) {
	fmt.Fprintf(os.Stderr, err+"\n")
	os.Exit(1)
}

func PrintErrors(errs ...error) {
	exit := false
	for _, err := range errs {
		if err != nil {
			fmt.Fprintf(os.Stderr, err.Error()+"\n")
			exit = true
		}
	}
	if exit {
		os.Exit(1)
	}
}

func PrintDebug(c *cli.Context, str string) {
	if c.Bool("debug") {
		fmt.Printf(str + "\n")
	}
}

func Exists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

func PrepareDir(dir string) {
	if !Exists(dir) {
		err := os.Mkdir(dir, 0777)
		if err != nil {
			fmt.Println(err.Error())
		}
	}
}

func GetCurrentPath() string {
	dir, err := filepath.Abs(".")
	PrintErrors(err)
	return dir
}

func GetHomeDir() string {
	dir, err := homedir.Dir()
	PrintErrors(err)
	return dir
}
