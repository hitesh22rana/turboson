package internals

type TokenType string

const (
	BraceOpen  TokenType = "BraceOpen"
	BraceClose TokenType = "BraceClose"

	BracketOpen  TokenType = "BracketOpen"
	BracketClose TokenType = "BracketClose"

	Colon TokenType = "Colon"
	Comma TokenType = "Comma"

	String TokenType = "String"
	Number TokenType = "Number"

	True  TokenType = "True"
	False TokenType = "False"

	Null TokenType = "Null"
)

type Token struct {
	Type  TokenType
	Value string
}
