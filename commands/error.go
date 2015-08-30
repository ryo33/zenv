package commands

import (
	"fmt"
	"github.com/ryo33/zenv/util"
)

func printArgumentError(num int) {
	if num == 1 {
		util.PrintErrorMessage(fmt.Sprintf("needs %d arg", num))
	} else {
		util.PrintErrorMessage(fmt.Sprintf("needs %d args", num))
	}
}

func printArgumentMoreError(num int) {
	util.PrintErrorMessage(fmt.Sprintf("needs %d or more args", num))
}
