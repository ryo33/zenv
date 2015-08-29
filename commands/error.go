package commands

import (
	"fmt"
	"github.com/ryo33/zenv/util"
)

func printArgumentError(num int) {
	util.PrintErrorMessage(fmt.Sprintf("needs %d arg", num))
}
