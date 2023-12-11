package lexer

import (
	"regexp"
	"unicode"

	internals "github.com/hitesh22rana/turboson/lib/internals"
)

func advance(index *int) {
	*index++
}

func Tokenizer(input string) []internals.Token {
	var tokens []internals.Token

	var length int = len(input)
	var index int = 0

	for index < length {
		var char byte = input[index]

		// open braces check
		if char == '{' {
			tokens = append(tokens, internals.Token{Type: internals.BraceOpen, Value: "{"})
			advance(&index)

			continue
		}

		// close braces check
		if char == '}' {
			tokens = append(tokens, internals.Token{Type: internals.BraceClose, Value: "}"})
			advance(&index)

			continue
		}

		// open brackets check
		if char == '[' {
			tokens = append(tokens, internals.Token{Type: internals.BracketOpen, Value: "["})
			advance(&index)

			continue
		}

		// close brackets check
		if char == ']' {
			tokens = append(tokens, internals.Token{Type: internals.BracketClose, Value: "]"})
			advance(&index)

			continue
		}

		// colon check
		if char == ':' {
			tokens = append(tokens, internals.Token{Type: internals.Colon, Value: ":"})
			advance(&index)

			continue
		}

		// comma check
		if char == ',' {
			tokens = append(tokens, internals.Token{Type: internals.Comma, Value: ","})
			advance(&index)

			continue
		}

		// string check
		if char == '"' {
			var value string = ""

			advance(&index)
			for index < length && input[index] != '"' {
				value += string(input[index])
				advance(&index)
			}

			tokens = append(tokens, internals.Token{Type: internals.String, Value: value})
			advance(&index)

			continue
		}

		// number, boolean, null check
		if regexp.MustCompile(`[\d\w]`).MatchString(string(char)) {
			var value string = ""

			for regexp.MustCompile(`[\d\w]`).MatchString(string(char)) {
				value += string(char)
				char = input[index+1]
				advance(&index)
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
		if unicode.IsSpace(rune(char)) {
			advance(&index)

			continue
		}

		panic("unexpected character: " + string(char))
	}

	return tokens
}
