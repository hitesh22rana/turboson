package parser

import (
	"strconv"

	"github.com/hitesh22rana/turboson/lib/internals"
)

func parse(d *Data) internals.ASTNode {
	switch (d.tokens)[d.index].Type {
	case internals.String:
		return internals.ASTNode{
			Type:  internals.StringNode,
			Value: d.tokens[d.index].Value,
		}
	case internals.Number:
		num, err := strconv.ParseFloat(d.tokens[d.index].Value, 64)
		if err != nil {
			panic(err.Error())
		}

		return internals.ASTNode{
			Type:  internals.NumberNode,
			Value: num,
		}
	case internals.True:
		return internals.ASTNode{
			Type:  internals.BooleanNode,
			Value: true,
		}
	case internals.False:
		return internals.ASTNode{
			Type:  internals.BooleanNode,
			Value: false,
		}
	case internals.Null:
		return internals.ASTNode{
			Type:  internals.NullNode,
			Value: nil,
		}
	case internals.BraceOpen:
		return parseObject(d)
	case internals.BracketOpen:
		return parseArray(d)
	default:
		panic("unexpected token type: " + d.tokens[d.index].Type)
	}
}

func parseObject(d *Data) internals.ASTNode {
	var node internals.ASTNode = internals.ASTNode{
		Type:  internals.ObjectNode,
		Value: map[string]internals.ASTNode{},
	}

	d.next()
	for d.tokens[d.index].Type != internals.BraceClose {
		if d.tokens[d.index].Type != internals.String {
			panic("expected string key in object, received: " + d.tokens[d.index].Type)
		}

		key := d.tokens[d.index].Value

		d.next()
		if d.tokens[d.index].Type != internals.Colon {
			panic("expected colon in key-value pair")
		}

		d.next()
		node.Value.(map[string]internals.ASTNode)[key] = parse(d)

		d.next()
		if d.tokens[d.index].Type == internals.Comma {
			d.next()
		}
	}

	return node
}

func parseArray(d *Data) internals.ASTNode {
	var node internals.ASTNode = internals.ASTNode{
		Type:  internals.ArrayNode,
		Value: []internals.ASTNode{},
	}

	d.next()
	for d.tokens[d.index].Type != internals.BracketClose {
		node.Value = append(node.Value.([]internals.ASTNode), parse(d))

		d.next()
		if d.tokens[d.index].Type == internals.Comma {
			d.next()
		}
	}

	return node
}
