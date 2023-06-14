package util

import (
	"fmt"
)

const DEBUG = false

func Dprintf(format string, a ...any) {
	if DEBUG {
		fmt.Printf(format, a...)
	}
}
