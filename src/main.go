package main

import (
	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
	gogoantlr "github.com/qunv/gogo/src/antlr"
)

func main() {
	// Create a CharStream from the file
	input, err := antlr.NewFileStream("src/main.gogo")
	if err != nil {
		panic(err)
	}
	lexer := gogoantlr.NewEqlLexer(input)
	p := gogoantlr.NewEqlParser(antlr.NewCommonTokenStream(lexer, 0))
	tree := p.Program()

	listener := gogoantlr.NewGoGoInterpreter()

	antlr.ParseTreeWalkerDefault.Walk(listener, tree)

	//visitor := gogoantlr.NewSimpleMathVisitor()
	//result := visitor.Visit(tree)
	//
	//fmt.Println(result) // Output: 17
}
