package internals

type ASTNodeType string

const (
	ObjectNode  ASTNodeType = "Object"
	ArrayNode   ASTNodeType = "Array"
	StringNode  ASTNodeType = "String"
	NumberNode  ASTNodeType = "Number"
	BooleanNode ASTNodeType = "Boolean"
	NullNode    ASTNodeType = "Null"
)

type ASTNode struct {
	Type  ASTNodeType
	Value any
}
