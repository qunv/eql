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
	if s.ctx.Expression() != nil {
		return s.evaluateExpression(s.ctx.Expression())
	}
	if s.ctx.Declarement() != nil {
		return nil, s.evaluateDeclarement(s.ctx.Declarement())
	}
	if s.ctx.Function() != nil {
		return nil, s.evaluateFunction(s.ctx.Function())
	}
	if s.ctx.Loop() != nil {
		return nil, s.evaluateLoop(s.ctx.Loop())
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

func (s *Statement) evaluateLoop(ctx antlr.ILoopContext) error {
	store := mem.NewMemory(s.mem)
	alias := ctx.IDENT().GetText()

	magic1 := ctx.InputRange().Cell(0)
	row1, col1, err := utils.GetRowAndColum(magic1)
	if err != nil {
		return err
	}
	magic2 := ctx.InputRange().Cell(1)
	row2, col2, err := utils.GetRowAndColum(magic2)
	if err != nil {
		return err
	}

	stmtCtxs := ctx.AllStatement()
	for i := row1; i <= row2; i++ {
		if col1 != col2 {
			for j := col1; j <= col2; j++ {
				err = s.handleLoopBody(stmtCtxs, alias, store, i, j)
				if err != nil {
					return err
				}
			}
		} else {
			err = s.handleLoopBody(stmtCtxs, alias, store, i, col1)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func (s *Statement) handleLoopBody(
	stmtCtxs []antlr.IStatementContext,
	alias string,
	store mem.Mem,
	row, col int,
) error {
	store.Set(alias, val.NewEqlValue(s.input.Get(row, col)))
	for _, stmtCtx := range stmtCtxs {
		if stmtCtx.Declarement() != nil &&
			stmtCtx.Declarement().IDENT().GetText() == alias {
			expr := newExpression(stmtCtx.Declarement().Expression(), s.input, store)
			result, err := expr.Evaluate()
			if err != nil {
				return err
			}
			s.input.Set(row, col, result.String())
			store.Set(alias, result)
			continue
		}
		stmt := NewStatement(stmtCtx, s.input, store)
		_, err := stmt.Evaluate()
		if err != nil {
			return err
		}
	}
	return nil
}
