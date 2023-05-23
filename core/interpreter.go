package core

import (
	"errors"
	"fmt"
	"github.com/qunv/eql/core/action"
	"github.com/qunv/eql/core/antlr"
	"github.com/qunv/eql/core/mem"
	"github.com/qunv/eql/core/utils"
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

func (e *EqlInterpreter) ExitStatement(ctx *antlr.StatementContext) {
	if ctx.Expression() != nil {
		e.result, e.err = e.evaluateExpression(ctx.Expression())
		return
	}
	if ctx.Declarement() != nil {
		e.evaluateDeclarement(ctx.Declarement())
		return
	}
	if ctx.Function() != nil {
		e.evaluateFunction(ctx.Function())
		return
	}
}

func (e *EqlInterpreter) evaluateFunction(ctx antlr.IFunctionContext) {
	value, err := e.evaluateExpression(ctx.Expression())
	if err != nil {
		e.err = err
		return
	}
	if ctx.PRINT() != nil {
		fmt.Println(value.String())
		return
	}
}

func (e *EqlInterpreter) evaluateDeclarement(ctx antlr.IDeclarementContext) {
	if ctx.Cell() != nil {
		row, col, err := utils.GetRowAndColum(ctx.Cell())
		if err != nil {
			e.err = err
			return
		}
		value, err := e.evaluateExpression(ctx.Expression())
		if err != nil {
			e.err = err
			return
		}
		e.input.Set(row, col, value.String())
		return
	}
	if ctx.IDENT() != nil {
		id := ctx.IDENT().GetText()
		value, err := e.evaluateExpression(ctx.Expression())
		if err != nil {
			e.err = err
			return
		}
		mem.Set(id, value)
		return
	}
	e.err = errors.New("declare syntax error")
}

func (e *EqlInterpreter) evaluateExpression(ctx antlr.IExpressionContext) (val.EqlValue, error) {
	expr := action.NewExpression(ctx, e.input)
	result, err := expr.Evaluate()
	if err != nil {
		return nil, err
	}
	return result, nil
}
