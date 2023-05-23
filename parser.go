package eql

import (
	"github.com/antlr4-go/antlr/v4"
	"github.com/qunv/eql/core"
	eqlantlr "github.com/qunv/eql/core/antlr"
	"github.com/qunv/eql/core/val"
)

type Parser struct {
	input [][]string
}

func NewEqlParser(input [][]string) *Parser {
	return &Parser{input: input}
}

func (p *Parser) Exec(q string) (val.EqlValue, error) {
	input := antlr.NewInputStream(q)
	lexer := eqlantlr.NewEqlLexer(input)
	eqlParser := eqlantlr.NewEqlParser(antlr.NewCommonTokenStream(lexer, 0))
	tree := eqlParser.Program()

	listener := core.NewEqlInterpreter(p.input)

	antlr.ParseTreeWalkerDefault.Walk(listener, tree)
	return listener.Result()
}
