package util

import (
	"fmt"
)

func PrintArgumentError(num int) {
	if num == 1 {
		PrintErrorMessage(fmt.Sprintf("needs %d arg", num))
	} else {
		PrintErrorMessage(fmt.Sprintf("needs %d args", num))
	}
}

func PrintArgumentMoreError(num int) {
	PrintErrorMessage(fmt.Sprintf("needs %d or more args", num))
}
