// Code generated from core/antlr/EqlParser.g4 by ANTLR 4.13.0. DO NOT EDIT.

package antlr // EqlParser
import "github.com/antlr4-go/antlr/v4"

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

// EnterDeclarement is called when production declarement is entered.
func (s *BaseEqlParserListener) EnterDeclarement(ctx *DeclarementContext) {}

// ExitDeclarement is called when production declarement is exited.
func (s *BaseEqlParserListener) ExitDeclarement(ctx *DeclarementContext) {}

// EnterFunction is called when production function is entered.
func (s *BaseEqlParserListener) EnterFunction(ctx *FunctionContext) {}

// ExitFunction is called when production function is exited.
func (s *BaseEqlParserListener) ExitFunction(ctx *FunctionContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseEqlParserListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseEqlParserListener) ExitExpression(ctx *ExpressionContext) {}

// EnterCompair is called when production compair is entered.
func (s *BaseEqlParserListener) EnterCompair(ctx *CompairContext) {}

// ExitCompair is called when production compair is exited.
func (s *BaseEqlParserListener) ExitCompair(ctx *CompairContext) {}

// EnterTerm is called when production term is entered.
func (s *BaseEqlParserListener) EnterTerm(ctx *TermContext) {}

// ExitTerm is called when production term is exited.
func (s *BaseEqlParserListener) ExitTerm(ctx *TermContext) {}

// EnterActionSpec is called when production actionSpec is entered.
func (s *BaseEqlParserListener) EnterActionSpec(ctx *ActionSpecContext) {}

// ExitActionSpec is called when production actionSpec is exited.
func (s *BaseEqlParserListener) ExitActionSpec(ctx *ActionSpecContext) {}

// EnterParam is called when production param is entered.
func (s *BaseEqlParserListener) EnterParam(ctx *ParamContext) {}

// ExitParam is called when production param is exited.
func (s *BaseEqlParserListener) ExitParam(ctx *ParamContext) {}

// EnterType is called when production type is entered.
func (s *BaseEqlParserListener) EnterType(ctx *TypeContext) {}

// ExitType is called when production type is exited.
func (s *BaseEqlParserListener) ExitType(ctx *TypeContext) {}

// EnterMath is called when production math is entered.
func (s *BaseEqlParserListener) EnterMath(ctx *MathContext) {}

// ExitMath is called when production math is exited.
func (s *BaseEqlParserListener) ExitMath(ctx *MathContext) {}

// EnterOperator is called when production operator is entered.
func (s *BaseEqlParserListener) EnterOperator(ctx *OperatorContext) {}

// ExitOperator is called when production operator is exited.
func (s *BaseEqlParserListener) ExitOperator(ctx *OperatorContext) {}

// EnterLogical is called when production logical is entered.
func (s *BaseEqlParserListener) EnterLogical(ctx *LogicalContext) {}

// ExitLogical is called when production logical is exited.
func (s *BaseEqlParserListener) ExitLogical(ctx *LogicalContext) {}

// EnterInputRange is called when production inputRange is entered.
func (s *BaseEqlParserListener) EnterInputRange(ctx *InputRangeContext) {}

// ExitInputRange is called when production inputRange is exited.
func (s *BaseEqlParserListener) ExitInputRange(ctx *InputRangeContext) {}

// EnterCell is called when production cell is entered.
func (s *BaseEqlParserListener) EnterCell(ctx *CellContext) {}

// ExitCell is called when production cell is exited.
func (s *BaseEqlParserListener) ExitCell(ctx *CellContext) {}

// EnterNumber is called when production number is entered.
func (s *BaseEqlParserListener) EnterNumber(ctx *NumberContext) {}

// ExitNumber is called when production number is exited.
func (s *BaseEqlParserListener) ExitNumber(ctx *NumberContext) {}

// EnterFactor is called when production factor is entered.
func (s *BaseEqlParserListener) EnterFactor(ctx *FactorContext) {}

// ExitFactor is called when production factor is exited.
func (s *BaseEqlParserListener) ExitFactor(ctx *FactorContext) {}
