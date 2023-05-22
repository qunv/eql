// Code generated from core/antlr/EqlParser.g4 by ANTLR 4.13.0. DO NOT EDIT.

package antlr // EqlParser
import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr4-go/antlr/v4"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type EqlParser struct {
	*antlr.BaseParser
}

var EqlParserParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func eqlparserParserInit() {
	staticData := &EqlParserParserStaticData
	staticData.LiteralNames = []string{
		"", "'+'", "'-'", "'/'", "'*'", "'('", "')'", "'{'", "'}'", "'['", "']'",
		"'='", "';'", "':'", "','", "'.'", "'>'", "'<'", "'TRUE'", "'FALSE'",
		"'SUM'", "'ABS'", "'AVG'", "'ADD'", "'DIVIDE'", "'MULTIPLY'", "'EQ'",
		"'CONCAT'", "'GT'", "'GTE'", "'LAMBDA'", "", "", "", "", "'IF'",
	}
	staticData.SymbolicNames = []string{
		"", "PLUS", "MINUS", "DIV", "MULT", "LPAREN", "RPAREN", "LCURLY", "RCURLY",
		"LBRACKET", "RBRACKET", "EQUAL", "SEMI", "COLON", "COMMA", "DOT", "GREATER_THAN",
		"LESS_THAN", "TRUE", "FALSE", "SUM", "ABS", "AVG", "ADD", "DIVIDE",
		"MULTIPLY", "EQ", "CONCAT", "GT", "GTE", "LAMBDA", "INT", "DECIMAL",
		"DIGIT", "STRING", "IF", "IDENTIFIER", "WS", "EOS",
	}
	staticData.RuleNames = []string{
		"program", "statement", "expression", "compair", "term", "actionSpec",
		"param", "type", "math", "operator", "logical", "inputRange", "def",
		"number", "factor",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 38, 117, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 1, 0, 4, 0,
		32, 8, 0, 11, 0, 12, 0, 33, 1, 1, 1, 1, 1, 2, 1, 2, 1, 2, 5, 2, 41, 8,
		2, 10, 2, 12, 2, 44, 9, 2, 1, 2, 1, 2, 1, 2, 1, 2, 3, 2, 50, 8, 2, 3, 2,
		52, 8, 2, 1, 3, 3, 3, 55, 8, 3, 1, 3, 3, 3, 58, 8, 3, 1, 4, 1, 4, 1, 4,
		5, 4, 63, 8, 4, 10, 4, 12, 4, 66, 9, 4, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 5,
		5, 73, 8, 5, 10, 5, 12, 5, 76, 9, 5, 1, 5, 1, 5, 1, 6, 1, 6, 3, 6, 82,
		8, 6, 1, 7, 1, 7, 1, 7, 3, 7, 87, 8, 7, 1, 8, 1, 8, 1, 9, 1, 9, 1, 10,
		1, 10, 1, 11, 1, 11, 1, 11, 1, 11, 1, 12, 1, 12, 1, 12, 1, 13, 1, 13, 1,
		14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14,
		3, 14, 115, 8, 14, 1, 14, 0, 0, 15, 0, 2, 4, 6, 8, 10, 12, 14, 16, 18,
		20, 22, 24, 26, 28, 0, 7, 1, 0, 1, 2, 1, 0, 16, 17, 1, 0, 3, 4, 1, 0, 20,
		21, 1, 0, 22, 29, 2, 0, 30, 30, 35, 35, 1, 0, 31, 32, 119, 0, 31, 1, 0,
		0, 0, 2, 35, 1, 0, 0, 0, 4, 51, 1, 0, 0, 0, 6, 54, 1, 0, 0, 0, 8, 59, 1,
		0, 0, 0, 10, 67, 1, 0, 0, 0, 12, 81, 1, 0, 0, 0, 14, 86, 1, 0, 0, 0, 16,
		88, 1, 0, 0, 0, 18, 90, 1, 0, 0, 0, 20, 92, 1, 0, 0, 0, 22, 94, 1, 0, 0,
		0, 24, 98, 1, 0, 0, 0, 26, 101, 1, 0, 0, 0, 28, 114, 1, 0, 0, 0, 30, 32,
		3, 2, 1, 0, 31, 30, 1, 0, 0, 0, 32, 33, 1, 0, 0, 0, 33, 31, 1, 0, 0, 0,
		33, 34, 1, 0, 0, 0, 34, 1, 1, 0, 0, 0, 35, 36, 3, 4, 2, 0, 36, 3, 1, 0,
		0, 0, 37, 42, 3, 8, 4, 0, 38, 39, 7, 0, 0, 0, 39, 41, 3, 8, 4, 0, 40, 38,
		1, 0, 0, 0, 41, 44, 1, 0, 0, 0, 42, 40, 1, 0, 0, 0, 42, 43, 1, 0, 0, 0,
		43, 52, 1, 0, 0, 0, 44, 42, 1, 0, 0, 0, 45, 49, 3, 8, 4, 0, 46, 47, 3,
		6, 3, 0, 47, 48, 3, 8, 4, 0, 48, 50, 1, 0, 0, 0, 49, 46, 1, 0, 0, 0, 49,
		50, 1, 0, 0, 0, 50, 52, 1, 0, 0, 0, 51, 37, 1, 0, 0, 0, 51, 45, 1, 0, 0,
		0, 52, 5, 1, 0, 0, 0, 53, 55, 7, 1, 0, 0, 54, 53, 1, 0, 0, 0, 54, 55, 1,
		0, 0, 0, 55, 57, 1, 0, 0, 0, 56, 58, 5, 11, 0, 0, 57, 56, 1, 0, 0, 0, 57,
		58, 1, 0, 0, 0, 58, 7, 1, 0, 0, 0, 59, 64, 3, 28, 14, 0, 60, 61, 7, 2,
		0, 0, 61, 63, 3, 28, 14, 0, 62, 60, 1, 0, 0, 0, 63, 66, 1, 0, 0, 0, 64,
		62, 1, 0, 0, 0, 64, 65, 1, 0, 0, 0, 65, 9, 1, 0, 0, 0, 66, 64, 1, 0, 0,
		0, 67, 68, 3, 14, 7, 0, 68, 69, 5, 5, 0, 0, 69, 74, 3, 12, 6, 0, 70, 71,
		5, 14, 0, 0, 71, 73, 3, 12, 6, 0, 72, 70, 1, 0, 0, 0, 73, 76, 1, 0, 0,
		0, 74, 72, 1, 0, 0, 0, 74, 75, 1, 0, 0, 0, 75, 77, 1, 0, 0, 0, 76, 74,
		1, 0, 0, 0, 77, 78, 5, 6, 0, 0, 78, 11, 1, 0, 0, 0, 79, 82, 3, 22, 11,
		0, 80, 82, 3, 4, 2, 0, 81, 79, 1, 0, 0, 0, 81, 80, 1, 0, 0, 0, 82, 13,
		1, 0, 0, 0, 83, 87, 3, 16, 8, 0, 84, 87, 3, 18, 9, 0, 85, 87, 3, 20, 10,
		0, 86, 83, 1, 0, 0, 0, 86, 84, 1, 0, 0, 0, 86, 85, 1, 0, 0, 0, 87, 15,
		1, 0, 0, 0, 88, 89, 7, 3, 0, 0, 89, 17, 1, 0, 0, 0, 90, 91, 7, 4, 0, 0,
		91, 19, 1, 0, 0, 0, 92, 93, 7, 5, 0, 0, 93, 21, 1, 0, 0, 0, 94, 95, 3,
		24, 12, 0, 95, 96, 5, 13, 0, 0, 96, 97, 3, 24, 12, 0, 97, 23, 1, 0, 0,
		0, 98, 99, 5, 36, 0, 0, 99, 100, 5, 31, 0, 0, 100, 25, 1, 0, 0, 0, 101,
		102, 7, 6, 0, 0, 102, 27, 1, 0, 0, 0, 103, 115, 3, 26, 13, 0, 104, 115,
		3, 24, 12, 0, 105, 115, 3, 10, 5, 0, 106, 115, 5, 18, 0, 0, 107, 115, 5,
		19, 0, 0, 108, 115, 5, 36, 0, 0, 109, 115, 5, 34, 0, 0, 110, 111, 5, 5,
		0, 0, 111, 112, 3, 4, 2, 0, 112, 113, 5, 6, 0, 0, 113, 115, 1, 0, 0, 0,
		114, 103, 1, 0, 0, 0, 114, 104, 1, 0, 0, 0, 114, 105, 1, 0, 0, 0, 114,
		106, 1, 0, 0, 0, 114, 107, 1, 0, 0, 0, 114, 108, 1, 0, 0, 0, 114, 109,
		1, 0, 0, 0, 114, 110, 1, 0, 0, 0, 115, 29, 1, 0, 0, 0, 11, 33, 42, 49,
		51, 54, 57, 64, 74, 81, 86, 114,
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
	staticData := &EqlParserParserStaticData
	staticData.once.Do(eqlparserParserInit)
}

// NewEqlParser produces a new parser instance for the optional input antlr.TokenStream.
func NewEqlParser(input antlr.TokenStream) *EqlParser {
	EqlParserInit()
	this := new(EqlParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &EqlParserParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "EqlParser.g4"

	return this
}

// EqlParser tokens.
const (
	EqlParserEOF          = antlr.TokenEOF
	EqlParserPLUS         = 1
	EqlParserMINUS        = 2
	EqlParserDIV          = 3
	EqlParserMULT         = 4
	EqlParserLPAREN       = 5
	EqlParserRPAREN       = 6
	EqlParserLCURLY       = 7
	EqlParserRCURLY       = 8
	EqlParserLBRACKET     = 9
	EqlParserRBRACKET     = 10
	EqlParserEQUAL        = 11
	EqlParserSEMI         = 12
	EqlParserCOLON        = 13
	EqlParserCOMMA        = 14
	EqlParserDOT          = 15
	EqlParserGREATER_THAN = 16
	EqlParserLESS_THAN    = 17
	EqlParserTRUE         = 18
	EqlParserFALSE        = 19
	EqlParserSUM          = 20
	EqlParserABS          = 21
	EqlParserAVG          = 22
	EqlParserADD          = 23
	EqlParserDIVIDE       = 24
	EqlParserMULTIPLY     = 25
	EqlParserEQ           = 26
	EqlParserCONCAT       = 27
	EqlParserGT           = 28
	EqlParserGTE          = 29
	EqlParserLAMBDA       = 30
	EqlParserINT          = 31
	EqlParserDECIMAL      = 32
	EqlParserDIGIT        = 33
	EqlParserSTRING       = 34
	EqlParserIF           = 35
	EqlParserIDENTIFIER   = 36
	EqlParserWS           = 37
	EqlParserEOS          = 38
)

// EqlParser rules.
const (
	EqlParserRULE_program    = 0
	EqlParserRULE_statement  = 1
	EqlParserRULE_expression = 2
	EqlParserRULE_compair    = 3
	EqlParserRULE_term       = 4
	EqlParserRULE_actionSpec = 5
	EqlParserRULE_param      = 6
	EqlParserRULE_type       = 7
	EqlParserRULE_math       = 8
	EqlParserRULE_operator   = 9
	EqlParserRULE_logical    = 10
	EqlParserRULE_inputRange = 11
	EqlParserRULE_def        = 12
	EqlParserRULE_number     = 13
	EqlParserRULE_factor     = 14
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
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProgramContext() *ProgramContext {
	var p = new(ProgramContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = EqlParserRULE_program
	return p
}

func InitEmptyProgramContext(p *ProgramContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = EqlParserRULE_program
}

func (*ProgramContext) IsProgramContext() {}

func NewProgramContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProgramContext {
	var p = new(ProgramContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

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
	localctx = NewProgramContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, EqlParserRULE_program)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(31)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&128848756768) != 0) {
		{
			p.SetState(30)
			p.Statement()
		}

		p.SetState(33)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
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
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStatementContext() *StatementContext {
	var p = new(StatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = EqlParserRULE_statement
	return p
}

func InitEmptyStatementContext(p *StatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = EqlParserRULE_statement
}

func (*StatementContext) IsStatementContext() {}

func NewStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatementContext {
	var p = new(StatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

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
	localctx = NewStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, EqlParserRULE_statement)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(35)
		p.Expression()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
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
	Compair() ICompairContext

	// IsExpressionContext differentiates from other interfaces.
	IsExpressionContext()
}

type ExpressionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpressionContext() *ExpressionContext {
	var p = new(ExpressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = EqlParserRULE_expression
	return p
}

func InitEmptyExpressionContext(p *ExpressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = EqlParserRULE_expression
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

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

func (s *ExpressionContext) Compair() ICompairContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICompairContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICompairContext)
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
	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, EqlParserRULE_expression)
	var _la int

	p.SetState(51)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 3, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(37)
			p.Term()
		}
		p.SetState(42)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == EqlParserPLUS || _la == EqlParserMINUS {
			{
				p.SetState(38)
				_la = p.GetTokenStream().LA(1)

				if !(_la == EqlParserPLUS || _la == EqlParserMINUS) {
					p.GetErrorHandler().RecoverInline(p)
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
			}
			{
				p.SetState(39)
				p.Term()
			}

			p.SetState(44)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(45)
			p.Term()
		}
		p.SetState(49)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 2, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(46)
				p.Compair()
			}
			{
				p.SetState(47)
				p.Term()
			}

		} else if p.HasError() { // JIM
			goto errorExit
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ICompairContext is an interface to support dynamic dispatch.
type ICompairContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	EQUAL() antlr.TerminalNode
	GREATER_THAN() antlr.TerminalNode
	LESS_THAN() antlr.TerminalNode

	// IsCompairContext differentiates from other interfaces.
	IsCompairContext()
}

type CompairContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCompairContext() *CompairContext {
	var p = new(CompairContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = EqlParserRULE_compair
	return p
}

func InitEmptyCompairContext(p *CompairContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = EqlParserRULE_compair
}

func (*CompairContext) IsCompairContext() {}

func NewCompairContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CompairContext {
	var p = new(CompairContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = EqlParserRULE_compair

	return p
}

func (s *CompairContext) GetParser() antlr.Parser { return s.parser }

func (s *CompairContext) EQUAL() antlr.TerminalNode {
	return s.GetToken(EqlParserEQUAL, 0)
}

func (s *CompairContext) GREATER_THAN() antlr.TerminalNode {
	return s.GetToken(EqlParserGREATER_THAN, 0)
}

func (s *CompairContext) LESS_THAN() antlr.TerminalNode {
	return s.GetToken(EqlParserLESS_THAN, 0)
}

func (s *CompairContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CompairContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CompairContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EqlParserListener); ok {
		listenerT.EnterCompair(s)
	}
}

func (s *CompairContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EqlParserListener); ok {
		listenerT.ExitCompair(s)
	}
}

func (p *EqlParser) Compair() (localctx ICompairContext) {
	localctx = NewCompairContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, EqlParserRULE_compair)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(54)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == EqlParserGREATER_THAN || _la == EqlParserLESS_THAN {
		{
			p.SetState(53)
			_la = p.GetTokenStream().LA(1)

			if !(_la == EqlParserGREATER_THAN || _la == EqlParserLESS_THAN) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

	}
	p.SetState(57)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == EqlParserEQUAL {
		{
			p.SetState(56)
			p.Match(EqlParserEQUAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
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
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTermContext() *TermContext {
	var p = new(TermContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = EqlParserRULE_term
	return p
}

func InitEmptyTermContext(p *TermContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = EqlParserRULE_term
}

func (*TermContext) IsTermContext() {}

func NewTermContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TermContext {
	var p = new(TermContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

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
	localctx = NewTermContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, EqlParserRULE_term)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(59)
		p.Factor()
	}
	p.SetState(64)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == EqlParserDIV || _la == EqlParserMULT {
		{
			p.SetState(60)
			_la = p.GetTokenStream().LA(1)

			if !(_la == EqlParserDIV || _la == EqlParserMULT) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(61)
			p.Factor()
		}

		p.SetState(66)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IActionSpecContext is an interface to support dynamic dispatch.
type IActionSpecContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Type_() ITypeContext
	LPAREN() antlr.TerminalNode
	RPAREN() antlr.TerminalNode
	AllParam() []IParamContext
	Param(i int) IParamContext
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsActionSpecContext differentiates from other interfaces.
	IsActionSpecContext()
}

type ActionSpecContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyActionSpecContext() *ActionSpecContext {
	var p = new(ActionSpecContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = EqlParserRULE_actionSpec
	return p
}

func InitEmptyActionSpecContext(p *ActionSpecContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = EqlParserRULE_actionSpec
}

func (*ActionSpecContext) IsActionSpecContext() {}

func NewActionSpecContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ActionSpecContext {
	var p = new(ActionSpecContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = EqlParserRULE_actionSpec

	return p
}

func (s *ActionSpecContext) GetParser() antlr.Parser { return s.parser }

func (s *ActionSpecContext) Type_() ITypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeContext)
}

func (s *ActionSpecContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(EqlParserLPAREN, 0)
}

func (s *ActionSpecContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(EqlParserRPAREN, 0)
}

func (s *ActionSpecContext) AllParam() []IParamContext {
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

func (s *ActionSpecContext) Param(i int) IParamContext {
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

func (s *ActionSpecContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(EqlParserCOMMA)
}

func (s *ActionSpecContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(EqlParserCOMMA, i)
}

func (s *ActionSpecContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ActionSpecContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ActionSpecContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EqlParserListener); ok {
		listenerT.EnterActionSpec(s)
	}
}

func (s *ActionSpecContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EqlParserListener); ok {
		listenerT.ExitActionSpec(s)
	}
}

func (p *EqlParser) ActionSpec() (localctx IActionSpecContext) {
	localctx = NewActionSpecContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, EqlParserRULE_actionSpec)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(67)
		p.Type_()
	}
	{
		p.SetState(68)
		p.Match(EqlParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

	{
		p.SetState(69)
		p.Param()
	}
	p.SetState(74)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == EqlParserCOMMA {
		{
			p.SetState(70)
			p.Match(EqlParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(71)
			p.Param()
		}

		p.SetState(76)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

	{
		p.SetState(77)
		p.Match(EqlParserRPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IParamContext is an interface to support dynamic dispatch.
type IParamContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	InputRange() IInputRangeContext
	Expression() IExpressionContext

	// IsParamContext differentiates from other interfaces.
	IsParamContext()
}

type ParamContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyParamContext() *ParamContext {
	var p = new(ParamContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = EqlParserRULE_param
	return p
}

func InitEmptyParamContext(p *ParamContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = EqlParserRULE_param
}

func (*ParamContext) IsParamContext() {}

func NewParamContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParamContext {
	var p = new(ParamContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = EqlParserRULE_param

	return p
}

func (s *ParamContext) GetParser() antlr.Parser { return s.parser }

func (s *ParamContext) InputRange() IInputRangeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IInputRangeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IInputRangeContext)
}

func (s *ParamContext) Expression() IExpressionContext {
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
	localctx = NewParamContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, EqlParserRULE_param)
	p.SetState(81)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 8, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(79)
			p.InputRange()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(80)
			p.Expression()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ITypeContext is an interface to support dynamic dispatch.
type ITypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Math() IMathContext
	Operator() IOperatorContext
	Logical() ILogicalContext

	// IsTypeContext differentiates from other interfaces.
	IsTypeContext()
}

type TypeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypeContext() *TypeContext {
	var p = new(TypeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = EqlParserRULE_type
	return p
}

func InitEmptyTypeContext(p *TypeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = EqlParserRULE_type
}

func (*TypeContext) IsTypeContext() {}

func NewTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeContext {
	var p = new(TypeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = EqlParserRULE_type

	return p
}

func (s *TypeContext) GetParser() antlr.Parser { return s.parser }

func (s *TypeContext) Math() IMathContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMathContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMathContext)
}

func (s *TypeContext) Operator() IOperatorContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOperatorContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOperatorContext)
}

func (s *TypeContext) Logical() ILogicalContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILogicalContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILogicalContext)
}

func (s *TypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EqlParserListener); ok {
		listenerT.EnterType(s)
	}
}

func (s *TypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EqlParserListener); ok {
		listenerT.ExitType(s)
	}
}

func (p *EqlParser) Type_() (localctx ITypeContext) {
	localctx = NewTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, EqlParserRULE_type)
	p.SetState(86)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case EqlParserSUM, EqlParserABS:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(83)
			p.Math()
		}

	case EqlParserAVG, EqlParserADD, EqlParserDIVIDE, EqlParserMULTIPLY, EqlParserEQ, EqlParserCONCAT, EqlParserGT, EqlParserGTE:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(84)
			p.Operator()
		}

	case EqlParserLAMBDA, EqlParserIF:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(85)
			p.Logical()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IMathContext is an interface to support dynamic dispatch.
type IMathContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	SUM() antlr.TerminalNode
	ABS() antlr.TerminalNode

	// IsMathContext differentiates from other interfaces.
	IsMathContext()
}

type MathContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMathContext() *MathContext {
	var p = new(MathContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = EqlParserRULE_math
	return p
}

func InitEmptyMathContext(p *MathContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = EqlParserRULE_math
}

func (*MathContext) IsMathContext() {}

func NewMathContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MathContext {
	var p = new(MathContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = EqlParserRULE_math

	return p
}

func (s *MathContext) GetParser() antlr.Parser { return s.parser }

func (s *MathContext) SUM() antlr.TerminalNode {
	return s.GetToken(EqlParserSUM, 0)
}

func (s *MathContext) ABS() antlr.TerminalNode {
	return s.GetToken(EqlParserABS, 0)
}

func (s *MathContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MathContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MathContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EqlParserListener); ok {
		listenerT.EnterMath(s)
	}
}

func (s *MathContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EqlParserListener); ok {
		listenerT.ExitMath(s)
	}
}

func (p *EqlParser) Math() (localctx IMathContext) {
	localctx = NewMathContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, EqlParserRULE_math)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(88)
		_la = p.GetTokenStream().LA(1)

		if !(_la == EqlParserSUM || _la == EqlParserABS) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IOperatorContext is an interface to support dynamic dispatch.
type IOperatorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AVG() antlr.TerminalNode
	ADD() antlr.TerminalNode
	EQ() antlr.TerminalNode
	CONCAT() antlr.TerminalNode
	DIVIDE() antlr.TerminalNode
	MULTIPLY() antlr.TerminalNode
	GT() antlr.TerminalNode
	GTE() antlr.TerminalNode

	// IsOperatorContext differentiates from other interfaces.
	IsOperatorContext()
}

type OperatorContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOperatorContext() *OperatorContext {
	var p = new(OperatorContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = EqlParserRULE_operator
	return p
}

func InitEmptyOperatorContext(p *OperatorContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = EqlParserRULE_operator
}

func (*OperatorContext) IsOperatorContext() {}

func NewOperatorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OperatorContext {
	var p = new(OperatorContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = EqlParserRULE_operator

	return p
}

func (s *OperatorContext) GetParser() antlr.Parser { return s.parser }

func (s *OperatorContext) AVG() antlr.TerminalNode {
	return s.GetToken(EqlParserAVG, 0)
}

func (s *OperatorContext) ADD() antlr.TerminalNode {
	return s.GetToken(EqlParserADD, 0)
}

func (s *OperatorContext) EQ() antlr.TerminalNode {
	return s.GetToken(EqlParserEQ, 0)
}

func (s *OperatorContext) CONCAT() antlr.TerminalNode {
	return s.GetToken(EqlParserCONCAT, 0)
}

func (s *OperatorContext) DIVIDE() antlr.TerminalNode {
	return s.GetToken(EqlParserDIVIDE, 0)
}

func (s *OperatorContext) MULTIPLY() antlr.TerminalNode {
	return s.GetToken(EqlParserMULTIPLY, 0)
}

func (s *OperatorContext) GT() antlr.TerminalNode {
	return s.GetToken(EqlParserGT, 0)
}

func (s *OperatorContext) GTE() antlr.TerminalNode {
	return s.GetToken(EqlParserGTE, 0)
}

func (s *OperatorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OperatorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OperatorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EqlParserListener); ok {
		listenerT.EnterOperator(s)
	}
}

func (s *OperatorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EqlParserListener); ok {
		listenerT.ExitOperator(s)
	}
}

func (p *EqlParser) Operator() (localctx IOperatorContext) {
	localctx = NewOperatorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, EqlParserRULE_operator)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(90)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&1069547520) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ILogicalContext is an interface to support dynamic dispatch.
type ILogicalContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IF() antlr.TerminalNode
	LAMBDA() antlr.TerminalNode

	// IsLogicalContext differentiates from other interfaces.
	IsLogicalContext()
}

type LogicalContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLogicalContext() *LogicalContext {
	var p = new(LogicalContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = EqlParserRULE_logical
	return p
}

func InitEmptyLogicalContext(p *LogicalContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = EqlParserRULE_logical
}

func (*LogicalContext) IsLogicalContext() {}

func NewLogicalContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LogicalContext {
	var p = new(LogicalContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = EqlParserRULE_logical

	return p
}

func (s *LogicalContext) GetParser() antlr.Parser { return s.parser }

func (s *LogicalContext) IF() antlr.TerminalNode {
	return s.GetToken(EqlParserIF, 0)
}

func (s *LogicalContext) LAMBDA() antlr.TerminalNode {
	return s.GetToken(EqlParserLAMBDA, 0)
}

func (s *LogicalContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LogicalContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LogicalContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EqlParserListener); ok {
		listenerT.EnterLogical(s)
	}
}

func (s *LogicalContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EqlParserListener); ok {
		listenerT.ExitLogical(s)
	}
}

func (p *EqlParser) Logical() (localctx ILogicalContext) {
	localctx = NewLogicalContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, EqlParserRULE_logical)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(92)
		_la = p.GetTokenStream().LA(1)

		if !(_la == EqlParserLAMBDA || _la == EqlParserIF) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IInputRangeContext is an interface to support dynamic dispatch.
type IInputRangeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllDef() []IDefContext
	Def(i int) IDefContext
	COLON() antlr.TerminalNode

	// IsInputRangeContext differentiates from other interfaces.
	IsInputRangeContext()
}

type InputRangeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInputRangeContext() *InputRangeContext {
	var p = new(InputRangeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = EqlParserRULE_inputRange
	return p
}

func InitEmptyInputRangeContext(p *InputRangeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = EqlParserRULE_inputRange
}

func (*InputRangeContext) IsInputRangeContext() {}

func NewInputRangeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InputRangeContext {
	var p = new(InputRangeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = EqlParserRULE_inputRange

	return p
}

func (s *InputRangeContext) GetParser() antlr.Parser { return s.parser }

func (s *InputRangeContext) AllDef() []IDefContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IDefContext); ok {
			len++
		}
	}

	tst := make([]IDefContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IDefContext); ok {
			tst[i] = t.(IDefContext)
			i++
		}
	}

	return tst
}

func (s *InputRangeContext) Def(i int) IDefContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDefContext); ok {
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

	return t.(IDefContext)
}

func (s *InputRangeContext) COLON() antlr.TerminalNode {
	return s.GetToken(EqlParserCOLON, 0)
}

func (s *InputRangeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InputRangeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *InputRangeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EqlParserListener); ok {
		listenerT.EnterInputRange(s)
	}
}

func (s *InputRangeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EqlParserListener); ok {
		listenerT.ExitInputRange(s)
	}
}

func (p *EqlParser) InputRange() (localctx IInputRangeContext) {
	localctx = NewInputRangeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, EqlParserRULE_inputRange)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(94)
		p.Def()
	}
	{
		p.SetState(95)
		p.Match(EqlParserCOLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(96)
		p.Def()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IDefContext is an interface to support dynamic dispatch.
type IDefContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IDENTIFIER() antlr.TerminalNode
	INT() antlr.TerminalNode

	// IsDefContext differentiates from other interfaces.
	IsDefContext()
}

type DefContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDefContext() *DefContext {
	var p = new(DefContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = EqlParserRULE_def
	return p
}

func InitEmptyDefContext(p *DefContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = EqlParserRULE_def
}

func (*DefContext) IsDefContext() {}

func NewDefContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DefContext {
	var p = new(DefContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = EqlParserRULE_def

	return p
}

func (s *DefContext) GetParser() antlr.Parser { return s.parser }

func (s *DefContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(EqlParserIDENTIFIER, 0)
}

func (s *DefContext) INT() antlr.TerminalNode {
	return s.GetToken(EqlParserINT, 0)
}

func (s *DefContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DefContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DefContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EqlParserListener); ok {
		listenerT.EnterDef(s)
	}
}

func (s *DefContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EqlParserListener); ok {
		listenerT.ExitDef(s)
	}
}

func (p *EqlParser) Def() (localctx IDefContext) {
	localctx = NewDefContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, EqlParserRULE_def)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(98)
		p.Match(EqlParserIDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(99)
		p.Match(EqlParserINT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// INumberContext is an interface to support dynamic dispatch.
type INumberContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	INT() antlr.TerminalNode
	DECIMAL() antlr.TerminalNode

	// IsNumberContext differentiates from other interfaces.
	IsNumberContext()
}

type NumberContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNumberContext() *NumberContext {
	var p = new(NumberContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = EqlParserRULE_number
	return p
}

func InitEmptyNumberContext(p *NumberContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = EqlParserRULE_number
}

func (*NumberContext) IsNumberContext() {}

func NewNumberContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NumberContext {
	var p = new(NumberContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = EqlParserRULE_number

	return p
}

func (s *NumberContext) GetParser() antlr.Parser { return s.parser }

func (s *NumberContext) INT() antlr.TerminalNode {
	return s.GetToken(EqlParserINT, 0)
}

func (s *NumberContext) DECIMAL() antlr.TerminalNode {
	return s.GetToken(EqlParserDECIMAL, 0)
}

func (s *NumberContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NumberContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NumberContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EqlParserListener); ok {
		listenerT.EnterNumber(s)
	}
}

func (s *NumberContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EqlParserListener); ok {
		listenerT.ExitNumber(s)
	}
}

func (p *EqlParser) Number() (localctx INumberContext) {
	localctx = NewNumberContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, EqlParserRULE_number)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(101)
		_la = p.GetTokenStream().LA(1)

		if !(_la == EqlParserINT || _la == EqlParserDECIMAL) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFactorContext is an interface to support dynamic dispatch.
type IFactorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Number() INumberContext
	Def() IDefContext
	ActionSpec() IActionSpecContext
	TRUE() antlr.TerminalNode
	FALSE() antlr.TerminalNode
	IDENTIFIER() antlr.TerminalNode
	STRING() antlr.TerminalNode
	LPAREN() antlr.TerminalNode
	Expression() IExpressionContext
	RPAREN() antlr.TerminalNode

	// IsFactorContext differentiates from other interfaces.
	IsFactorContext()
}

type FactorContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFactorContext() *FactorContext {
	var p = new(FactorContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = EqlParserRULE_factor
	return p
}

func InitEmptyFactorContext(p *FactorContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = EqlParserRULE_factor
}

func (*FactorContext) IsFactorContext() {}

func NewFactorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FactorContext {
	var p = new(FactorContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = EqlParserRULE_factor

	return p
}

func (s *FactorContext) GetParser() antlr.Parser { return s.parser }

func (s *FactorContext) Number() INumberContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INumberContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INumberContext)
}

func (s *FactorContext) Def() IDefContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDefContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDefContext)
}

func (s *FactorContext) ActionSpec() IActionSpecContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IActionSpecContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IActionSpecContext)
}

func (s *FactorContext) TRUE() antlr.TerminalNode {
	return s.GetToken(EqlParserTRUE, 0)
}

func (s *FactorContext) FALSE() antlr.TerminalNode {
	return s.GetToken(EqlParserFALSE, 0)
}

func (s *FactorContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(EqlParserIDENTIFIER, 0)
}

func (s *FactorContext) STRING() antlr.TerminalNode {
	return s.GetToken(EqlParserSTRING, 0)
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
	localctx = NewFactorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, EqlParserRULE_factor)
	p.SetState(114)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 10, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(103)
			p.Number()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(104)
			p.Def()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(105)
			p.ActionSpec()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(106)
			p.Match(EqlParserTRUE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(107)
			p.Match(EqlParserFALSE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(108)
			p.Match(EqlParserIDENTIFIER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(109)
			p.Match(EqlParserSTRING)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 8:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(110)
			p.Match(EqlParserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(111)
			p.Expression()
		}
		{
			p.SetState(112)
			p.Match(EqlParserRPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}
