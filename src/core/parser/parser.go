package parser

import (
	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
	"github.com/qunv/eql/src/core/action"
	eqlantlr "github.com/qunv/eql/src/core/antlr"
)

type Parser struct {
	input [][]interface{}
}

func NewEqlParser(input [][]interface{}) *Parser {
	return &Parser{input: input}
}

func (p *Parser) Exec(q string) (action.EqlValue, error) {
	input := antlr.NewInputStream(q)
	lexer := eqlantlr.NewEqlLexer(input)
	eqlParser := eqlantlr.NewEqlParser(antlr.NewCommonTokenStream(lexer, 0))
	tree := eqlParser.Program()

	listener := NewEqlInterpreter(p.input)

	antlr.ParseTreeWalkerDefault.Walk(listener, tree)
	return listener.Result()
}
