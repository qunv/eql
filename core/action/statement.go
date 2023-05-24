package action

import (
	"errors"
	"fmt"
	"github.com/qunv/eql/core/antlr"
	"github.com/qunv/eql/core/mem"
	"github.com/qunv/eql/core/utils"
	"github.com/qunv/eql/core/val"
)

type Statement struct {
	ctx   antlr.IStatementContext
	input EqlInput
	mem   mem.Mem
}

func NewStatement(ctx antlr.IStatementContext, input EqlInput, mem mem.Mem) *Statement {
	return &Statement{
		ctx:   ctx,
		input: input,
		mem:   mem,
	}
}

func void() (val.EqlValue, error) {
	return nil, nil
}

func (s *Statement) Evaluate() (val.EqlValue, error) {
	s.mem.Set("test-store", val.NewEqlValue("test-store"))
	if s.ctx.Expression() != nil {
		return s.evaluateExpression(s.ctx.Expression())
	}
	if s.ctx.Declarement() != nil {
		return nil, s.evaluateDeclarement(s.ctx.Declarement())
	}
	if s.ctx.Function() != nil {
		return nil, s.evaluateFunction(s.ctx.Function())
	}
	return void()
}

func (s *Statement) evaluateFunction(ctx antlr.IFunctionContext) error {
	value, err := s.evaluateExpression(ctx.Expression())
	if err != nil {
		return err
	}
	if ctx.PRINT() != nil {
		fmt.Println(value.String())
		return err
	}
	return nil
}

func (s *Statement) evaluateDeclarement(ctx antlr.IDeclarementContext) error {
	if ctx.Cell() != nil {
		row, col, err := utils.GetRowAndColum(ctx.Cell())
		if err != nil {
			return err
		}
		value, err := s.evaluateExpression(ctx.Expression())
		if err != nil {
			return err
		}
		s.input.Set(row, col, value.String())
		return nil
	}
	if ctx.IDENT() != nil {
		id := ctx.IDENT().GetText()
		value, err := s.evaluateExpression(ctx.Expression())
		if err != nil {
			return err
		}
		s.mem.Set(id, value)
		return nil
	}
	return errors.New("declare syntax error")
}

func (s *Statement) evaluateExpression(ctx antlr.IExpressionContext) (val.EqlValue, error) {
	expr := newExpression(ctx, s.input, s.mem)
	result, err := expr.Evaluate()
	if err != nil {
		return nil, err
	}
	return result, nil
}
