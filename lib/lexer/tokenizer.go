package lexer

import (
	"github.com/hitesh22rana/turboson/lib/internals"
)

func next(index *uint) {
	*index++
}

func Tokenize(input string) []internals.Token {
	if len(input) == 0 {
		panic("no input to tokenize")
	}

	var tokens []internals.Token

	var length uint = uint(len(input))
	var index uint = 0
	for index < length {
		var char byte = input[index]

		// open braces check
		if char == '{' {
			tokens = append(tokens, internals.Token{Type: internals.BraceOpen, Value: "{"})
			next(&index)
			continue
		}

		// close braces check
		if char == '}' {
			tokens = append(tokens, internals.Token{Type: internals.BraceClose, Value: "}"})
			next(&index)
			continue
		}

		// open brackets check
		if char == '[' {
			tokens = append(tokens, internals.Token{Type: internals.BracketOpen, Value: "["})
			next(&index)
			continue
		}

		// close brackets check
		if char == ']' {
			tokens = append(tokens, internals.Token{Type: internals.BracketClose, Value: "]"})
			next(&index)
			continue
		}

		// colon check
		if char == ':' {
			tokens = append(tokens, internals.Token{Type: internals.Colon, Value: ":"})
			next(&index)
			continue
		}

		// comma check
		if char == ',' {
			tokens = append(tokens, internals.Token{Type: internals.Comma, Value: ","})
			next(&index)
			continue
		}

		// string check
		if char == '"' {
			var value string = ""

			next(&index)
			for index < length && input[index] != '"' {
				value += string(input[index])
				next(&index)
			}

			next(&index)
			tokens = append(tokens, internals.Token{Type: internals.String, Value: value})
			continue
		}

		// number, boolean, null check
		if numBoolNullRegex.MatchString(string(char)) {
			var value string = ""
			for numBoolNullRegex.MatchString(string(input[index])) {
				value += string(input[index])
				next(&index)
			}

			if isNumber(value) {
				tokens = append(tokens, internals.Token{Type: internals.Number, Value: value})
			} else if isBooleanTrue(value) {
				tokens = append(tokens, internals.Token{Type: internals.True, Value: value})
			} else if isBooleanFalse(value) {
				tokens = append(tokens, internals.Token{Type: internals.False, Value: value})
			} else if isNull(value) {
				tokens = append(tokens, internals.Token{Type: internals.Null, Value: value})
			} else {
				panic("unexpected value: " + value)
			}
			continue
		}

		// skip whitespaces
		if spaceRegex.MatchString(string(char)) {
			next(&index)
			continue
		}

		panic("unexpected character: " + string(char))
	}

	return tokens
}
