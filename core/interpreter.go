package core

import (
	"github.com/qunv/eql/core/action"
	"github.com/qunv/eql/core/antlr"
	"github.com/qunv/eql/core/mem"
	"github.com/qunv/eql/core/val"
)

type EqlInterpreter struct {
	*antlr.BaseEqlParserListener
	input  action.EqlInput
	result val.EqlValue
	store  map[string]val.EqlValue
	err    error
}

func NewEqlInterpreter(input action.EqlInput) *EqlInterpreter {
	return &EqlInterpreter{
		input: input,
	}
}

func (e *EqlInterpreter) Result() (val.EqlValue, error) {
	return e.result, e.err
}

func (e *EqlInterpreter) ExitProgram(ctx *antlr.ProgramContext) {
	stmtCtxs := ctx.AllStatement()
	store := mem.NewMemory(nil)
	for _, stmtCtx := range stmtCtxs {
		stmt := action.NewStatement(stmtCtx, e.input, store)
		value, err := stmt.Evaluate()
		if value == nil && err == nil {
			continue
		}
		e.result, e.err = value, err
	}
}
