package util

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func YNPrompt(prompt string, yes bool) bool {
	yn := "[y/N]"
	if yes {
		yn = "[Y/n]"
	}
	for {
		fmt.Printf(prompt + " " + yn)
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		switch strings.ToLower(scanner.Text()) {
		case "y", "yes":
			return true
		case "n", "no":
			return false
		case "":
			return yes
		}
	}
}
