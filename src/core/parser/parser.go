package parser

import (
	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
	eqlantlr "github.com/qunv/eql/src/core/antlr"
)

type Parser struct {
	input [][]float64
}

func NewEqlParser(input [][]float64) *Parser {
	return &Parser{input: input}
}

func (p *Parser) Exec(q string) float64 {
	input := antlr.NewInputStream(q)
	lexer := eqlantlr.NewEqlLexer(input)
	eqlParser := eqlantlr.NewEqlParser(antlr.NewCommonTokenStream(lexer, 0))
	tree := eqlParser.Program()

	listener := NewEqlInterpreter(p.input)

	antlr.ParseTreeWalkerDefault.Walk(listener, tree)
	return listener.Result()
}
