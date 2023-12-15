package main

import (
	"fmt"
	"time"

	"github.com/hitesh22rana/turboson/lib/internals"
	"github.com/hitesh22rana/turboson/lib/lexer"
	"github.com/hitesh22rana/turboson/lib/parser"
)

func main() {
	start := time.Now()
	tokens := lexer.Tokenize(`{
		"id": "647ceaf3657eade56f8224eb",
		"index": 0,
		"anArray": [],
		"boolean": true,
		"nullValue": null
		}`)

	parsed := parser.Parse(tokens)

	PrintJson(parsed, 0)

	fmt.Println("total time taken", time.Since(start))
}

func PrintJson(json internals.ASTNode, depth uint) {
	fmt.Printf("{\n\ttype: %q\n\tvalue: \n", json.Type)
	for key, value := range json.Value.(map[string]internals.ASTNode) {
		fmt.Printf("\t%v: %q, value: %v\n", key, value.Type, value.Value)
	}
	fmt.Println("}")
}
