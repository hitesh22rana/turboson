package lexer

import (
	"regexp"
	"strconv"
)

var (
	numBoolNullRegex = regexp.MustCompile(`[\d\w]`)
	spaceRegex       = regexp.MustCompile(`\s`)
)

func isNumber(value string) bool {
	_, err := strconv.ParseFloat(value, 64)
	return err == nil
}

func isBooleanTrue(value string) bool {
	return value == "true"
}

func isBooleanFalse(value string) bool {
	return value == "false"
}

func isNull(value string) bool {
	return value == "null"
}
