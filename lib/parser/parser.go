package parser

import (
	"github.com/hitesh22rana/turboson/lib/internals"
)

type Data struct {
	tokens []internals.Token
	index  uint
}

func (d *Data) next() {
	d.index++
}

func Parse(tokens []internals.Token) internals.ASTNode {
	if len(tokens) == 0 {
		panic("no tokens to parse")
	}

	return parse(&Data{
		tokens: tokens,
		index:  0,
	})
}
