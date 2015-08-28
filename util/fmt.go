package util

import (
	"fmt"
)

func FormatErrorAtLine(file string, line int, message string) string {
	return fmt.Sprintf("%s:%d: %s", file, line, message)
}
