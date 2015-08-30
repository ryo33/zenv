package util

import (
	"fmt"
	"github.com/mitchellh/go-homedir"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"strings"
)

var (
	Debug = false
	Quiet = false
)

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

func PrintErrorMessageContinue(err string) {
	fmt.Fprintf(os.Stderr, err+"\n")
}

func PrintErrorMessage(err string) {
	PrintErrorMessageContinue(err)
	os.Exit(1)
}

func PrintErrorsContinue(errs ...error) bool {
	exit := false
	for _, err := range errs {
		if err != nil {
			fmt.Fprintf(os.Stderr, err.Error()+"\n")
			exit = true
		}
	}
	return exit
}

func PrintErrors(errs ...error) {
	if PrintErrorsContinue(errs...) {
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

func ExecCommand(command ...string) (string, error) {
	output, err := exec.Command(command[0], command[1:]...).Output()
	return string(output), err
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
	os.Setenv(key, value)
	Print(fmt.Sprintf("export %s=%s", key, value))
}

func ExpandPath(pa string) string {
	var err error
	if strings.HasPrefix(pa, "./") {
		pa = path.Join(GetCurrentPath(), strings.TrimPrefix(pa, "./"))
	} else if strings.HasPrefix(pa, "~/") {
		pa, err = homedir.Expand(pa)
		PrintErrors(err)
	}
	return path.Clean(pa)
}
