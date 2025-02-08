package utils

import (
	"fmt"
	"strconv"
	"strings"
)

func LogError(line int, col int, data string, err string) {
	lineAsString := strconv.Itoa(line)
	var lineDigits int = len(lineAsString)
	spaces := strings.Repeat(" ", 12+lineDigits+col)

	fmt.Printf("\n[ERROR]: %s at line %d:%d\n", err, line, col)
	fmt.Printf("   └───> %d | %s", line, data)
	fmt.Printf("%s^-- error\n", spaces)
}
