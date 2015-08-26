package util

import (
	"fmt"
	"github.com/mitchellh/go-homedir"
	"os"
	"os/exec"
	"path/filepath"
)

var (
	Debug = false
	Quiet = false
	Eval  = false // zenv output will be evaluated
)

func format(str string) string {
	if Eval {
		return fmt.Sprintf("echo \"%s\"", str)
	} else {
		return str
	}
}

func Print(str string) {
	if !Quiet {
		fmt.Printf(str + "\n")
	}
}

func PrintDebug(str string) {
	if Debug {
		fmt.Printf(str + "\n")
	}
}

func PrintErrorMessage(err string) {
	fmt.Fprintf(os.Stderr, format(err)+"\n")
	os.Exit(1)
}

func PrintErrors(errs ...error) {
	exit := false
	for _, err := range errs {
		if err != nil {
			fmt.Fprintf(os.Stderr, format(err.Error())+"\n")
			exit = true
		}
	}
	if exit {
		os.Exit(1)
	}
}

func Exists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

func PrepareDir(dir string) {
	if !Exists(dir) {
		err := os.MkdirAll(dir, 0777)
		if err != nil {
			fmt.Println(err.Error())
		}
	}
}

func ExecCommand(command ...string) {
	exec.Command(command[0], command[1:]...).Run()
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

func Setenv(key, value string) {
	Print(fmt.Sprintf("export %s=%s", key, value))
}
