// Code generated from src/core/antlr/EqlParser.g4 by ANTLR 4.12.0. DO NOT EDIT.

package antlr // EqlParser
import "github.com/antlr/antlr4/runtime/Go/antlr/v4"

// BaseEqlParserListener is a complete listener for a parse tree produced by EqlParser.
type BaseEqlParserListener struct{}

var _ EqlParserListener = &BaseEqlParserListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseEqlParserListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseEqlParserListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseEqlParserListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseEqlParserListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterProgram is called when production program is entered.
func (s *BaseEqlParserListener) EnterProgram(ctx *ProgramContext) {}

// ExitProgram is called when production program is exited.
func (s *BaseEqlParserListener) ExitProgram(ctx *ProgramContext) {}

// EnterStatement is called when production statement is entered.
func (s *BaseEqlParserListener) EnterStatement(ctx *StatementContext) {}

// ExitStatement is called when production statement is exited.
func (s *BaseEqlParserListener) ExitStatement(ctx *StatementContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseEqlParserListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseEqlParserListener) ExitExpression(ctx *ExpressionContext) {}

// EnterTerm is called when production term is entered.
func (s *BaseEqlParserListener) EnterTerm(ctx *TermContext) {}

// ExitTerm is called when production term is exited.
func (s *BaseEqlParserListener) ExitTerm(ctx *TermContext) {}

// EnterActSpec is called when production actSpec is entered.
func (s *BaseEqlParserListener) EnterActSpec(ctx *ActSpecContext) {}

// ExitActSpec is called when production actSpec is exited.
func (s *BaseEqlParserListener) ExitActSpec(ctx *ActSpecContext) {}

// EnterAct is called when production act is entered.
func (s *BaseEqlParserListener) EnterAct(ctx *ActContext) {}

// ExitAct is called when production act is exited.
func (s *BaseEqlParserListener) ExitAct(ctx *ActContext) {}

// EnterMath is called when production math is entered.
func (s *BaseEqlParserListener) EnterMath(ctx *MathContext) {}

// ExitMath is called when production math is exited.
func (s *BaseEqlParserListener) ExitMath(ctx *MathContext) {}

// EnterParam is called when production param is entered.
func (s *BaseEqlParserListener) EnterParam(ctx *ParamContext) {}

// ExitParam is called when production param is exited.
func (s *BaseEqlParserListener) ExitParam(ctx *ParamContext) {}

// EnterDef is called when production def is entered.
func (s *BaseEqlParserListener) EnterDef(ctx *DefContext) {}

// ExitDef is called when production def is exited.
func (s *BaseEqlParserListener) ExitDef(ctx *DefContext) {}

// EnterNumber is called when production number is entered.
func (s *BaseEqlParserListener) EnterNumber(ctx *NumberContext) {}

// ExitNumber is called when production number is exited.
func (s *BaseEqlParserListener) ExitNumber(ctx *NumberContext) {}

// EnterFactor is called when production factor is entered.
func (s *BaseEqlParserListener) EnterFactor(ctx *FactorContext) {}

// ExitFactor is called when production factor is exited.
func (s *BaseEqlParserListener) ExitFactor(ctx *FactorContext) {}