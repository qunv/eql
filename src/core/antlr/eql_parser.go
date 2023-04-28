// Code generated from src/antlr/EqlParser.g4 by ANTLR 4.12.0. DO NOT EDIT.

package antlr // EqlParser
import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type EqlParser struct {
	*antlr.BaseParser
}

var eqlparserParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func eqlparserParserInit() {
	staticData := &eqlparserParserStaticData
	staticData.literalNames = []string{
		"", "'+'", "'-'", "'/'", "'*'", "'('", "')'", "'{'", "'}'", "'['", "']'",
		"'='", "';'", "':'", "','", "'=SUM'", "'=AVG'",
	}
	staticData.symbolicNames = []string{
		"", "PLUS", "MINUS", "DIV", "MULT", "LPAREN", "RPAREN", "LCURLY", "RCURLY",
		"LBRACKET", "RBRACKET", "EQUAL", "SEMI", "COLON", "COMMA", "SUM", "AVG",
		"NUMBER", "IDENTIFIER", "WS", "EOS",
	}
	staticData.ruleNames = []string{
		"program", "statement", "expression", "term", "actSpec", "act", "param",
		"magic", "factor",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 20, 76, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 1, 0, 4, 0, 20, 8, 0,
		11, 0, 12, 0, 21, 1, 1, 1, 1, 1, 2, 1, 2, 1, 2, 5, 2, 29, 8, 2, 10, 2,
		12, 2, 32, 9, 2, 1, 3, 1, 3, 1, 3, 5, 3, 37, 8, 3, 10, 3, 12, 3, 40, 9,
		3, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 5, 4, 47, 8, 4, 10, 4, 12, 4, 50, 9, 4,
		1, 4, 1, 4, 1, 5, 1, 5, 1, 6, 1, 6, 1, 6, 3, 6, 59, 8, 6, 1, 6, 3, 6, 62,
		8, 6, 1, 7, 1, 7, 1, 7, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 3, 8,
		74, 8, 8, 1, 8, 0, 0, 9, 0, 2, 4, 6, 8, 10, 12, 14, 16, 0, 3, 1, 0, 1,
		2, 1, 0, 3, 4, 1, 0, 15, 16, 75, 0, 19, 1, 0, 0, 0, 2, 23, 1, 0, 0, 0,
		4, 25, 1, 0, 0, 0, 6, 33, 1, 0, 0, 0, 8, 41, 1, 0, 0, 0, 10, 53, 1, 0,
		0, 0, 12, 61, 1, 0, 0, 0, 14, 63, 1, 0, 0, 0, 16, 73, 1, 0, 0, 0, 18, 20,
		3, 2, 1, 0, 19, 18, 1, 0, 0, 0, 20, 21, 1, 0, 0, 0, 21, 19, 1, 0, 0, 0,
		21, 22, 1, 0, 0, 0, 22, 1, 1, 0, 0, 0, 23, 24, 3, 4, 2, 0, 24, 3, 1, 0,
		0, 0, 25, 30, 3, 6, 3, 0, 26, 27, 7, 0, 0, 0, 27, 29, 3, 6, 3, 0, 28, 26,
		1, 0, 0, 0, 29, 32, 1, 0, 0, 0, 30, 28, 1, 0, 0, 0, 30, 31, 1, 0, 0, 0,
		31, 5, 1, 0, 0, 0, 32, 30, 1, 0, 0, 0, 33, 38, 3, 16, 8, 0, 34, 35, 7,
		1, 0, 0, 35, 37, 3, 16, 8, 0, 36, 34, 1, 0, 0, 0, 37, 40, 1, 0, 0, 0, 38,
		36, 1, 0, 0, 0, 38, 39, 1, 0, 0, 0, 39, 7, 1, 0, 0, 0, 40, 38, 1, 0, 0,
		0, 41, 42, 3, 10, 5, 0, 42, 43, 5, 5, 0, 0, 43, 48, 3, 12, 6, 0, 44, 45,
		5, 12, 0, 0, 45, 47, 3, 12, 6, 0, 46, 44, 1, 0, 0, 0, 47, 50, 1, 0, 0,
		0, 48, 46, 1, 0, 0, 0, 48, 49, 1, 0, 0, 0, 49, 51, 1, 0, 0, 0, 50, 48,
		1, 0, 0, 0, 51, 52, 5, 6, 0, 0, 52, 9, 1, 0, 0, 0, 53, 54, 7, 2, 0, 0,
		54, 11, 1, 0, 0, 0, 55, 58, 3, 14, 7, 0, 56, 57, 5, 13, 0, 0, 57, 59, 3,
		14, 7, 0, 58, 56, 1, 0, 0, 0, 58, 59, 1, 0, 0, 0, 59, 62, 1, 0, 0, 0, 60,
		62, 3, 8, 4, 0, 61, 55, 1, 0, 0, 0, 61, 60, 1, 0, 0, 0, 62, 13, 1, 0, 0,
		0, 63, 64, 5, 18, 0, 0, 64, 65, 5, 17, 0, 0, 65, 15, 1, 0, 0, 0, 66, 74,
		5, 17, 0, 0, 67, 74, 3, 14, 7, 0, 68, 74, 3, 8, 4, 0, 69, 70, 5, 5, 0,
		0, 70, 71, 3, 4, 2, 0, 71, 72, 5, 6, 0, 0, 72, 74, 1, 0, 0, 0, 73, 66,
		1, 0, 0, 0, 73, 67, 1, 0, 0, 0, 73, 68, 1, 0, 0, 0, 73, 69, 1, 0, 0, 0,
		74, 17, 1, 0, 0, 0, 7, 21, 30, 38, 48, 58, 61, 73,
	}
	deserializer := antlr.NewATNDeserializer(nil)
	staticData.atn = deserializer.Deserialize(staticData.serializedATN)
	atn := staticData.atn
	staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
	decisionToDFA := staticData.decisionToDFA
	for index, state := range atn.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(state, index)
	}
}

// EqlParserInit initializes any static state used to implement EqlParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewEqlParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func EqlParserInit() {
	staticData := &eqlparserParserStaticData
	staticData.once.Do(eqlparserParserInit)
}

// NewEqlParser produces a new parser instance for the optional input antlr.TokenStream.
func NewEqlParser(input antlr.TokenStream) *EqlParser {
	EqlParserInit()
	this := new(EqlParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &eqlparserParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "EqlParser.g4"

	return this
}

// EqlParser tokens.
const (
	EqlParserEOF        = antlr.TokenEOF
	EqlParserPLUS       = 1
	EqlParserMINUS      = 2
	EqlParserDIV        = 3
	EqlParserMULT       = 4
	EqlParserLPAREN     = 5
	EqlParserRPAREN     = 6
	EqlParserLCURLY     = 7
	EqlParserRCURLY     = 8
	EqlParserLBRACKET   = 9
	EqlParserRBRACKET   = 10
	EqlParserEQUAL      = 11
	EqlParserSEMI       = 12
	EqlParserCOLON      = 13
	EqlParserCOMMA      = 14
	EqlParserSUM        = 15
	EqlParserAVG        = 16
	EqlParserNUMBER     = 17
	EqlParserIDENTIFIER = 18
	EqlParserWS         = 19
	EqlParserEOS        = 20
)

// EqlParser rules.
const (
	EqlParserRULE_program    = 0
	EqlParserRULE_statement  = 1
	EqlParserRULE_expression = 2
	EqlParserRULE_term       = 3
	EqlParserRULE_actSpec    = 4
	EqlParserRULE_act        = 5
	EqlParserRULE_param      = 6
	EqlParserRULE_magic      = 7
	EqlParserRULE_factor     = 8
)

// IProgramContext is an interface to support dynamic dispatch.
type IProgramContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllStatement() []IStatementContext
	Statement(i int) IStatementContext

	// IsProgramContext differentiates from other interfaces.
	IsProgramContext()
}

type ProgramContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProgramContext() *ProgramContext {
	var p = new(ProgramContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = EqlParserRULE_program
	return p
}

func (*ProgramContext) IsProgramContext() {}

func NewProgramContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProgramContext {
	var p = new(ProgramContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = EqlParserRULE_program

	return p
}

func (s *ProgramContext) GetParser() antlr.Parser { return s.parser }

func (s *ProgramContext) AllStatement() []IStatementContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStatementContext); ok {
			len++
		}
	}

	tst := make([]IStatementContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStatementContext); ok {
			tst[i] = t.(IStatementContext)
			i++
		}
	}

	return tst
}

func (s *ProgramContext) Statement(i int) IStatementContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatementContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *ProgramContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProgramContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProgramContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EqlParserListener); ok {
		listenerT.EnterProgram(s)
	}
}

func (s *ProgramContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EqlParserListener); ok {
		listenerT.ExitProgram(s)
	}
}

func (p *EqlParser) Program() (localctx IProgramContext) {
	this := p
	_ = this

	localctx = NewProgramContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, EqlParserRULE_program)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(19)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&491552) != 0) {
		{
			p.SetState(18)
			p.Statement()
		}

		p.SetState(21)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IStatementContext is an interface to support dynamic dispatch.
type IStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Expression() IExpressionContext

	// IsStatementContext differentiates from other interfaces.
	IsStatementContext()
}

type StatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStatementContext() *StatementContext {
	var p = new(StatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = EqlParserRULE_statement
	return p
}

func (*StatementContext) IsStatementContext() {}

func NewStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatementContext {
	var p = new(StatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = EqlParserRULE_statement

	return p
}

func (s *StatementContext) GetParser() antlr.Parser { return s.parser }

func (s *StatementContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *StatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EqlParserListener); ok {
		listenerT.EnterStatement(s)
	}
}

func (s *StatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EqlParserListener); ok {
		listenerT.ExitStatement(s)
	}
}

func (p *EqlParser) Statement() (localctx IStatementContext) {
	this := p
	_ = this

	localctx = NewStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, EqlParserRULE_statement)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(23)
		p.Expression()
	}

	return localctx
}

// IExpressionContext is an interface to support dynamic dispatch.
type IExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllTerm() []ITermContext
	Term(i int) ITermContext
	AllPLUS() []antlr.TerminalNode
	PLUS(i int) antlr.TerminalNode
	AllMINUS() []antlr.TerminalNode
	MINUS(i int) antlr.TerminalNode

	// IsExpressionContext differentiates from other interfaces.
	IsExpressionContext()
}

type ExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpressionContext() *ExpressionContext {
	var p = new(ExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = EqlParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = EqlParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) AllTerm() []ITermContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ITermContext); ok {
			len++
		}
	}

	tst := make([]ITermContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ITermContext); ok {
			tst[i] = t.(ITermContext)
			i++
		}
	}

	return tst
}

func (s *ExpressionContext) Term(i int) ITermContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITermContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITermContext)
}

func (s *ExpressionContext) AllPLUS() []antlr.TerminalNode {
	return s.GetTokens(EqlParserPLUS)
}

func (s *ExpressionContext) PLUS(i int) antlr.TerminalNode {
	return s.GetToken(EqlParserPLUS, i)
}

func (s *ExpressionContext) AllMINUS() []antlr.TerminalNode {
	return s.GetTokens(EqlParserMINUS)
}

func (s *ExpressionContext) MINUS(i int) antlr.TerminalNode {
	return s.GetToken(EqlParserMINUS, i)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EqlParserListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EqlParserListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *EqlParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, EqlParserRULE_expression)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(25)
		p.Term()
	}
	p.SetState(30)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == EqlParserPLUS || _la == EqlParserMINUS {
		{
			p.SetState(26)
			_la = p.GetTokenStream().LA(1)

			if !(_la == EqlParserPLUS || _la == EqlParserMINUS) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(27)
			p.Term()
		}

		p.SetState(32)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// ITermContext is an interface to support dynamic dispatch.
type ITermContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllFactor() []IFactorContext
	Factor(i int) IFactorContext
	AllMULT() []antlr.TerminalNode
	MULT(i int) antlr.TerminalNode
	AllDIV() []antlr.TerminalNode
	DIV(i int) antlr.TerminalNode

	// IsTermContext differentiates from other interfaces.
	IsTermContext()
}

type TermContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTermContext() *TermContext {
	var p = new(TermContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = EqlParserRULE_term
	return p
}

func (*TermContext) IsTermContext() {}

func NewTermContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TermContext {
	var p = new(TermContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = EqlParserRULE_term

	return p
}

func (s *TermContext) GetParser() antlr.Parser { return s.parser }

func (s *TermContext) AllFactor() []IFactorContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IFactorContext); ok {
			len++
		}
	}

	tst := make([]IFactorContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IFactorContext); ok {
			tst[i] = t.(IFactorContext)
			i++
		}
	}

	return tst
}

func (s *TermContext) Factor(i int) IFactorContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFactorContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFactorContext)
}

func (s *TermContext) AllMULT() []antlr.TerminalNode {
	return s.GetTokens(EqlParserMULT)
}

func (s *TermContext) MULT(i int) antlr.TerminalNode {
	return s.GetToken(EqlParserMULT, i)
}

func (s *TermContext) AllDIV() []antlr.TerminalNode {
	return s.GetTokens(EqlParserDIV)
}

func (s *TermContext) DIV(i int) antlr.TerminalNode {
	return s.GetToken(EqlParserDIV, i)
}

func (s *TermContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TermContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TermContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EqlParserListener); ok {
		listenerT.EnterTerm(s)
	}
}

func (s *TermContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EqlParserListener); ok {
		listenerT.ExitTerm(s)
	}
}

func (p *EqlParser) Term() (localctx ITermContext) {
	this := p
	_ = this

	localctx = NewTermContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, EqlParserRULE_term)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(33)
		p.Factor()
	}
	p.SetState(38)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == EqlParserDIV || _la == EqlParserMULT {
		{
			p.SetState(34)
			_la = p.GetTokenStream().LA(1)

			if !(_la == EqlParserDIV || _la == EqlParserMULT) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(35)
			p.Factor()
		}

		p.SetState(40)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IActSpecContext is an interface to support dynamic dispatch.
type IActSpecContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Act() IActContext
	LPAREN() antlr.TerminalNode
	RPAREN() antlr.TerminalNode
	AllParam() []IParamContext
	Param(i int) IParamContext
	AllSEMI() []antlr.TerminalNode
	SEMI(i int) antlr.TerminalNode

	// IsActSpecContext differentiates from other interfaces.
	IsActSpecContext()
}

type ActSpecContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyActSpecContext() *ActSpecContext {
	var p = new(ActSpecContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = EqlParserRULE_actSpec
	return p
}

func (*ActSpecContext) IsActSpecContext() {}

func NewActSpecContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ActSpecContext {
	var p = new(ActSpecContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = EqlParserRULE_actSpec

	return p
}

func (s *ActSpecContext) GetParser() antlr.Parser { return s.parser }

func (s *ActSpecContext) Act() IActContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IActContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IActContext)
}

func (s *ActSpecContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(EqlParserLPAREN, 0)
}

func (s *ActSpecContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(EqlParserRPAREN, 0)
}

func (s *ActSpecContext) AllParam() []IParamContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IParamContext); ok {
			len++
		}
	}

	tst := make([]IParamContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IParamContext); ok {
			tst[i] = t.(IParamContext)
			i++
		}
	}

	return tst
}

func (s *ActSpecContext) Param(i int) IParamContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IParamContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IParamContext)
}

func (s *ActSpecContext) AllSEMI() []antlr.TerminalNode {
	return s.GetTokens(EqlParserSEMI)
}

func (s *ActSpecContext) SEMI(i int) antlr.TerminalNode {
	return s.GetToken(EqlParserSEMI, i)
}

func (s *ActSpecContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ActSpecContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ActSpecContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EqlParserListener); ok {
		listenerT.EnterActSpec(s)
	}
}

func (s *ActSpecContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EqlParserListener); ok {
		listenerT.ExitActSpec(s)
	}
}

func (p *EqlParser) ActSpec() (localctx IActSpecContext) {
	this := p
	_ = this

	localctx = NewActSpecContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, EqlParserRULE_actSpec)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(41)
		p.Act()
	}
	{
		p.SetState(42)
		p.Match(EqlParserLPAREN)
	}

	{
		p.SetState(43)
		p.Param()
	}
	p.SetState(48)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == EqlParserSEMI {
		{
			p.SetState(44)
			p.Match(EqlParserSEMI)
		}
		{
			p.SetState(45)
			p.Param()
		}

		p.SetState(50)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	{
		p.SetState(51)
		p.Match(EqlParserRPAREN)
	}

	return localctx
}

// IActContext is an interface to support dynamic dispatch.
type IActContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	SUM() antlr.TerminalNode
	AVG() antlr.TerminalNode

	// IsActContext differentiates from other interfaces.
	IsActContext()
}

type ActContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyActContext() *ActContext {
	var p = new(ActContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = EqlParserRULE_act
	return p
}

func (*ActContext) IsActContext() {}

func NewActContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ActContext {
	var p = new(ActContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = EqlParserRULE_act

	return p
}

func (s *ActContext) GetParser() antlr.Parser { return s.parser }

func (s *ActContext) SUM() antlr.TerminalNode {
	return s.GetToken(EqlParserSUM, 0)
}

func (s *ActContext) AVG() antlr.TerminalNode {
	return s.GetToken(EqlParserAVG, 0)
}

func (s *ActContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ActContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ActContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EqlParserListener); ok {
		listenerT.EnterAct(s)
	}
}

func (s *ActContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EqlParserListener); ok {
		listenerT.ExitAct(s)
	}
}

func (p *EqlParser) Act() (localctx IActContext) {
	this := p
	_ = this

	localctx = NewActContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, EqlParserRULE_act)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(53)
		_la = p.GetTokenStream().LA(1)

		if !(_la == EqlParserSUM || _la == EqlParserAVG) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IParamContext is an interface to support dynamic dispatch.
type IParamContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllMagic() []IMagicContext
	Magic(i int) IMagicContext
	COLON() antlr.TerminalNode
	ActSpec() IActSpecContext

	// IsParamContext differentiates from other interfaces.
	IsParamContext()
}

type ParamContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyParamContext() *ParamContext {
	var p = new(ParamContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = EqlParserRULE_param
	return p
}

func (*ParamContext) IsParamContext() {}

func NewParamContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParamContext {
	var p = new(ParamContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = EqlParserRULE_param

	return p
}

func (s *ParamContext) GetParser() antlr.Parser { return s.parser }

func (s *ParamContext) AllMagic() []IMagicContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IMagicContext); ok {
			len++
		}
	}

	tst := make([]IMagicContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IMagicContext); ok {
			tst[i] = t.(IMagicContext)
			i++
		}
	}

	return tst
}

func (s *ParamContext) Magic(i int) IMagicContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMagicContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMagicContext)
}

func (s *ParamContext) COLON() antlr.TerminalNode {
	return s.GetToken(EqlParserCOLON, 0)
}

func (s *ParamContext) ActSpec() IActSpecContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IActSpecContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IActSpecContext)
}

func (s *ParamContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParamContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ParamContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EqlParserListener); ok {
		listenerT.EnterParam(s)
	}
}

func (s *ParamContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EqlParserListener); ok {
		listenerT.ExitParam(s)
	}
}

func (p *EqlParser) Param() (localctx IParamContext) {
	this := p
	_ = this

	localctx = NewParamContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, EqlParserRULE_param)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(61)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case EqlParserIDENTIFIER:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(55)
			p.Magic()
		}
		p.SetState(58)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == EqlParserCOLON {
			{
				p.SetState(56)
				p.Match(EqlParserCOLON)
			}
			{
				p.SetState(57)
				p.Magic()
			}

		}

	case EqlParserSUM, EqlParserAVG:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(60)
			p.ActSpec()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IMagicContext is an interface to support dynamic dispatch.
type IMagicContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IDENTIFIER() antlr.TerminalNode
	NUMBER() antlr.TerminalNode

	// IsMagicContext differentiates from other interfaces.
	IsMagicContext()
}

type MagicContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMagicContext() *MagicContext {
	var p = new(MagicContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = EqlParserRULE_magic
	return p
}

func (*MagicContext) IsMagicContext() {}

func NewMagicContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MagicContext {
	var p = new(MagicContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = EqlParserRULE_magic

	return p
}

func (s *MagicContext) GetParser() antlr.Parser { return s.parser }

func (s *MagicContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(EqlParserIDENTIFIER, 0)
}

func (s *MagicContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(EqlParserNUMBER, 0)
}

func (s *MagicContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MagicContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MagicContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EqlParserListener); ok {
		listenerT.EnterMagic(s)
	}
}

func (s *MagicContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EqlParserListener); ok {
		listenerT.ExitMagic(s)
	}
}

func (p *EqlParser) Magic() (localctx IMagicContext) {
	this := p
	_ = this

	localctx = NewMagicContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, EqlParserRULE_magic)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(63)
		p.Match(EqlParserIDENTIFIER)
	}
	{
		p.SetState(64)
		p.Match(EqlParserNUMBER)
	}

	return localctx
}

// IFactorContext is an interface to support dynamic dispatch.
type IFactorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	NUMBER() antlr.TerminalNode
	Magic() IMagicContext
	ActSpec() IActSpecContext
	LPAREN() antlr.TerminalNode
	Expression() IExpressionContext
	RPAREN() antlr.TerminalNode

	// IsFactorContext differentiates from other interfaces.
	IsFactorContext()
}

type FactorContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFactorContext() *FactorContext {
	var p = new(FactorContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = EqlParserRULE_factor
	return p
}

func (*FactorContext) IsFactorContext() {}

func NewFactorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FactorContext {
	var p = new(FactorContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = EqlParserRULE_factor

	return p
}

func (s *FactorContext) GetParser() antlr.Parser { return s.parser }

func (s *FactorContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(EqlParserNUMBER, 0)
}

func (s *FactorContext) Magic() IMagicContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMagicContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMagicContext)
}

func (s *FactorContext) ActSpec() IActSpecContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IActSpecContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IActSpecContext)
}

func (s *FactorContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(EqlParserLPAREN, 0)
}

func (s *FactorContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *FactorContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(EqlParserRPAREN, 0)
}

func (s *FactorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FactorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FactorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EqlParserListener); ok {
		listenerT.EnterFactor(s)
	}
}

func (s *FactorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EqlParserListener); ok {
		listenerT.ExitFactor(s)
	}
}

func (p *EqlParser) Factor() (localctx IFactorContext) {
	this := p
	_ = this

	localctx = NewFactorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, EqlParserRULE_factor)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(73)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case EqlParserNUMBER:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(66)
			p.Match(EqlParserNUMBER)
		}

	case EqlParserIDENTIFIER:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(67)
			p.Magic()
		}

	case EqlParserSUM, EqlParserAVG:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(68)
			p.ActSpec()
		}

	case EqlParserLPAREN:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(69)
			p.Match(EqlParserLPAREN)
		}
		{
			p.SetState(70)
			p.Expression()
		}
		{
			p.SetState(71)
			p.Match(EqlParserRPAREN)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}
