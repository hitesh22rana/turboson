package lexer

import "strconv"

func isNumber(value string) bool {
	_, err := strconv.ParseInt(value, 10, 64)
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
