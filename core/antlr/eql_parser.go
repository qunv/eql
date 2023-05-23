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
		"'='", "';'", "':'", "','", "'.'", "'>'", "'<'", "'$'", "'TRUE'", "'FALSE'",
		"'print'", "'SUM'", "'ABS'", "'AVG'", "'ADD'", "'DIVIDE'", "'MULTIPLY'",
		"'EQ'", "'CONCAT'", "'GT'", "'GTE'", "'LAMBDA'", "", "", "", "", "'IF'",
	}
	staticData.SymbolicNames = []string{
		"", "PLUS", "MINUS", "DIV", "MULT", "LPAREN", "RPAREN", "LCURLY", "RCURLY",
		"LBRACKET", "RBRACKET", "EQUAL", "SEMI", "COLON", "COMMA", "DOT", "GREATER_THAN",
		"LESS_THAN", "DOLAR", "TRUE", "FALSE", "PRINT", "SUM", "ABS", "AVG",
		"ADD", "DIVIDE", "MULTIPLY", "EQ", "CONCAT", "GT", "GTE", "LAMBDA",
		"INT", "DECIMAL", "DIGIT", "STRING", "IF", "ALPHABET", "IDENT", "WS",
		"EOS",
	}
	staticData.RuleNames = []string{
		"program", "statement", "declarement", "function", "expression", "compair",
		"term", "actionSpec", "param", "type", "math", "operator", "logical",
		"inputRange", "cell", "number", "factor",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 41, 137, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 1, 0, 4, 0, 36, 8, 0, 11, 0, 12, 0, 37, 1, 1, 1, 1, 1, 1,
		3, 1, 43, 8, 1, 1, 2, 1, 2, 3, 2, 47, 8, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1,
		3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 4, 1, 4, 1, 4, 5, 4, 61, 8, 4, 10, 4, 12,
		4, 64, 9, 4, 1, 4, 1, 4, 1, 4, 1, 4, 3, 4, 70, 8, 4, 3, 4, 72, 8, 4, 1,
		5, 3, 5, 75, 8, 5, 1, 5, 3, 5, 78, 8, 5, 1, 6, 1, 6, 1, 6, 5, 6, 83, 8,
		6, 10, 6, 12, 6, 86, 9, 6, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 5, 7, 93, 8, 7,
		10, 7, 12, 7, 96, 9, 7, 1, 7, 1, 7, 1, 8, 1, 8, 3, 8, 102, 8, 8, 1, 9,
		1, 9, 1, 9, 3, 9, 107, 8, 9, 1, 10, 1, 10, 1, 11, 1, 11, 1, 12, 1, 12,
		1, 13, 1, 13, 1, 13, 1, 13, 1, 14, 1, 14, 1, 14, 1, 15, 1, 15, 1, 16, 1,
		16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 3, 16,
		135, 8, 16, 1, 16, 0, 0, 17, 0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22,
		24, 26, 28, 30, 32, 0, 7, 1, 0, 1, 2, 1, 0, 16, 17, 1, 0, 3, 4, 1, 0, 22,
		23, 1, 0, 24, 31, 2, 0, 32, 32, 37, 37, 1, 0, 33, 34, 140, 0, 35, 1, 0,
		0, 0, 2, 42, 1, 0, 0, 0, 4, 46, 1, 0, 0, 0, 6, 52, 1, 0, 0, 0, 8, 71, 1,
		0, 0, 0, 10, 74, 1, 0, 0, 0, 12, 79, 1, 0, 0, 0, 14, 87, 1, 0, 0, 0, 16,
		101, 1, 0, 0, 0, 18, 106, 1, 0, 0, 0, 20, 108, 1, 0, 0, 0, 22, 110, 1,
		0, 0, 0, 24, 112, 1, 0, 0, 0, 26, 114, 1, 0, 0, 0, 28, 118, 1, 0, 0, 0,
		30, 121, 1, 0, 0, 0, 32, 134, 1, 0, 0, 0, 34, 36, 3, 2, 1, 0, 35, 34, 1,
		0, 0, 0, 36, 37, 1, 0, 0, 0, 37, 35, 1, 0, 0, 0, 37, 38, 1, 0, 0, 0, 38,
		1, 1, 0, 0, 0, 39, 43, 3, 8, 4, 0, 40, 43, 3, 4, 2, 0, 41, 43, 3, 6, 3,
		0, 42, 39, 1, 0, 0, 0, 42, 40, 1, 0, 0, 0, 42, 41, 1, 0, 0, 0, 43, 3, 1,
		0, 0, 0, 44, 47, 3, 28, 14, 0, 45, 47, 5, 39, 0, 0, 46, 44, 1, 0, 0, 0,
		46, 45, 1, 0, 0, 0, 47, 48, 1, 0, 0, 0, 48, 49, 5, 13, 0, 0, 49, 50, 5,
		11, 0, 0, 50, 51, 3, 8, 4, 0, 51, 5, 1, 0, 0, 0, 52, 53, 5, 21, 0, 0, 53,
		54, 5, 5, 0, 0, 54, 55, 3, 8, 4, 0, 55, 56, 5, 6, 0, 0, 56, 7, 1, 0, 0,
		0, 57, 62, 3, 12, 6, 0, 58, 59, 7, 0, 0, 0, 59, 61, 3, 12, 6, 0, 60, 58,
		1, 0, 0, 0, 61, 64, 1, 0, 0, 0, 62, 60, 1, 0, 0, 0, 62, 63, 1, 0, 0, 0,
		63, 72, 1, 0, 0, 0, 64, 62, 1, 0, 0, 0, 65, 69, 3, 12, 6, 0, 66, 67, 3,
		10, 5, 0, 67, 68, 3, 12, 6, 0, 68, 70, 1, 0, 0, 0, 69, 66, 1, 0, 0, 0,
		69, 70, 1, 0, 0, 0, 70, 72, 1, 0, 0, 0, 71, 57, 1, 0, 0, 0, 71, 65, 1,
		0, 0, 0, 72, 9, 1, 0, 0, 0, 73, 75, 7, 1, 0, 0, 74, 73, 1, 0, 0, 0, 74,
		75, 1, 0, 0, 0, 75, 77, 1, 0, 0, 0, 76, 78, 5, 11, 0, 0, 77, 76, 1, 0,
		0, 0, 77, 78, 1, 0, 0, 0, 78, 11, 1, 0, 0, 0, 79, 84, 3, 32, 16, 0, 80,
		81, 7, 2, 0, 0, 81, 83, 3, 32, 16, 0, 82, 80, 1, 0, 0, 0, 83, 86, 1, 0,
		0, 0, 84, 82, 1, 0, 0, 0, 84, 85, 1, 0, 0, 0, 85, 13, 1, 0, 0, 0, 86, 84,
		1, 0, 0, 0, 87, 88, 3, 18, 9, 0, 88, 89, 5, 5, 0, 0, 89, 94, 3, 16, 8,
		0, 90, 91, 5, 14, 0, 0, 91, 93, 3, 16, 8, 0, 92, 90, 1, 0, 0, 0, 93, 96,
		1, 0, 0, 0, 94, 92, 1, 0, 0, 0, 94, 95, 1, 0, 0, 0, 95, 97, 1, 0, 0, 0,
		96, 94, 1, 0, 0, 0, 97, 98, 5, 6, 0, 0, 98, 15, 1, 0, 0, 0, 99, 102, 3,
		26, 13, 0, 100, 102, 3, 8, 4, 0, 101, 99, 1, 0, 0, 0, 101, 100, 1, 0, 0,
		0, 102, 17, 1, 0, 0, 0, 103, 107, 3, 20, 10, 0, 104, 107, 3, 22, 11, 0,
		105, 107, 3, 24, 12, 0, 106, 103, 1, 0, 0, 0, 106, 104, 1, 0, 0, 0, 106,
		105, 1, 0, 0, 0, 107, 19, 1, 0, 0, 0, 108, 109, 7, 3, 0, 0, 109, 21, 1,
		0, 0, 0, 110, 111, 7, 4, 0, 0, 111, 23, 1, 0, 0, 0, 112, 113, 7, 5, 0,
		0, 113, 25, 1, 0, 0, 0, 114, 115, 3, 28, 14, 0, 115, 116, 5, 13, 0, 0,
		116, 117, 3, 28, 14, 0, 117, 27, 1, 0, 0, 0, 118, 119, 5, 38, 0, 0, 119,
		120, 5, 33, 0, 0, 120, 29, 1, 0, 0, 0, 121, 122, 7, 6, 0, 0, 122, 31, 1,
		0, 0, 0, 123, 135, 3, 30, 15, 0, 124, 135, 3, 28, 14, 0, 125, 135, 3, 14,
		7, 0, 126, 135, 5, 19, 0, 0, 127, 135, 5, 20, 0, 0, 128, 135, 5, 39, 0,
		0, 129, 135, 5, 36, 0, 0, 130, 131, 5, 5, 0, 0, 131, 132, 3, 8, 4, 0, 132,
		133, 5, 6, 0, 0, 133, 135, 1, 0, 0, 0, 134, 123, 1, 0, 0, 0, 134, 124,
		1, 0, 0, 0, 134, 125, 1, 0, 0, 0, 134, 126, 1, 0, 0, 0, 134, 127, 1, 0,
		0, 0, 134, 128, 1, 0, 0, 0, 134, 129, 1, 0, 0, 0, 134, 130, 1, 0, 0, 0,
		135, 33, 1, 0, 0, 0, 13, 37, 42, 46, 62, 69, 71, 74, 77, 84, 94, 101, 106,
		134,
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
	EqlParserDOLAR        = 18
	EqlParserTRUE         = 19
	EqlParserFALSE        = 20
	EqlParserPRINT        = 21
	EqlParserSUM          = 22
	EqlParserABS          = 23
	EqlParserAVG          = 24
	EqlParserADD          = 25
	EqlParserDIVIDE       = 26
	EqlParserMULTIPLY     = 27
	EqlParserEQ           = 28
	EqlParserCONCAT       = 29
	EqlParserGT           = 30
	EqlParserGTE          = 31
	EqlParserLAMBDA       = 32
	EqlParserINT          = 33
	EqlParserDECIMAL      = 34
	EqlParserDIGIT        = 35
	EqlParserSTRING       = 36
	EqlParserIF           = 37
	EqlParserALPHABET     = 38
	EqlParserIDENT        = 39
	EqlParserWS           = 40
	EqlParserEOS          = 41
)

// EqlParser rules.
const (
	EqlParserRULE_program     = 0
	EqlParserRULE_statement   = 1
	EqlParserRULE_declarement = 2
	EqlParserRULE_function    = 3
	EqlParserRULE_expression  = 4
	EqlParserRULE_compair     = 5
	EqlParserRULE_term        = 6
	EqlParserRULE_actionSpec  = 7
	EqlParserRULE_param       = 8
	EqlParserRULE_type        = 9
	EqlParserRULE_math        = 10
	EqlParserRULE_operator    = 11
	EqlParserRULE_logical     = 12
	EqlParserRULE_inputRange  = 13
	EqlParserRULE_cell        = 14
	EqlParserRULE_number      = 15
	EqlParserRULE_factor      = 16
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
	p.SetState(35)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&1065151365152) != 0) {
		{
			p.SetState(34)
			p.Statement()
		}

		p.SetState(37)
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
	Declarement() IDeclarementContext
	Function() IFunctionContext

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

func (s *StatementContext) Declarement() IDeclarementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDeclarementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDeclarementContext)
}

func (s *StatementContext) Function() IFunctionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFunctionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFunctionContext)
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
	p.SetState(42)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 1, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(39)
			p.Expression()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(40)
			p.Declarement()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(41)
			p.Function()
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

// IDeclarementContext is an interface to support dynamic dispatch.
type IDeclarementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	COLON() antlr.TerminalNode
	EQUAL() antlr.TerminalNode
	Expression() IExpressionContext
	Cell() ICellContext
	IDENT() antlr.TerminalNode

	// IsDeclarementContext differentiates from other interfaces.
	IsDeclarementContext()
}

type DeclarementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDeclarementContext() *DeclarementContext {
	var p = new(DeclarementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = EqlParserRULE_declarement
	return p
}

func InitEmptyDeclarementContext(p *DeclarementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = EqlParserRULE_declarement
}

func (*DeclarementContext) IsDeclarementContext() {}

func NewDeclarementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DeclarementContext {
	var p = new(DeclarementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = EqlParserRULE_declarement

	return p
}

func (s *DeclarementContext) GetParser() antlr.Parser { return s.parser }

func (s *DeclarementContext) COLON() antlr.TerminalNode {
	return s.GetToken(EqlParserCOLON, 0)
}

func (s *DeclarementContext) EQUAL() antlr.TerminalNode {
	return s.GetToken(EqlParserEQUAL, 0)
}

func (s *DeclarementContext) Expression() IExpressionContext {
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

func (s *DeclarementContext) Cell() ICellContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICellContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICellContext)
}

func (s *DeclarementContext) IDENT() antlr.TerminalNode {
	return s.GetToken(EqlParserIDENT, 0)
}

func (s *DeclarementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DeclarementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DeclarementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EqlParserListener); ok {
		listenerT.EnterDeclarement(s)
	}
}

func (s *DeclarementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EqlParserListener); ok {
		listenerT.ExitDeclarement(s)
	}
}

func (p *EqlParser) Declarement() (localctx IDeclarementContext) {
	localctx = NewDeclarementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, EqlParserRULE_declarement)
	p.EnterOuterAlt(localctx, 1)
	p.SetState(46)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case EqlParserALPHABET:
		{
			p.SetState(44)
			p.Cell()
		}

	case EqlParserIDENT:
		{
			p.SetState(45)
			p.Match(EqlParserIDENT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}
	{
		p.SetState(48)
		p.Match(EqlParserCOLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(49)
		p.Match(EqlParserEQUAL)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(50)
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

// IFunctionContext is an interface to support dynamic dispatch.
type IFunctionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	PRINT() antlr.TerminalNode
	LPAREN() antlr.TerminalNode
	Expression() IExpressionContext
	RPAREN() antlr.TerminalNode

	// IsFunctionContext differentiates from other interfaces.
	IsFunctionContext()
}

type FunctionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunctionContext() *FunctionContext {
	var p = new(FunctionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = EqlParserRULE_function
	return p
}

func InitEmptyFunctionContext(p *FunctionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = EqlParserRULE_function
}

func (*FunctionContext) IsFunctionContext() {}

func NewFunctionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunctionContext {
	var p = new(FunctionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = EqlParserRULE_function

	return p
}

func (s *FunctionContext) GetParser() antlr.Parser { return s.parser }

func (s *FunctionContext) PRINT() antlr.TerminalNode {
	return s.GetToken(EqlParserPRINT, 0)
}

func (s *FunctionContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(EqlParserLPAREN, 0)
}

func (s *FunctionContext) Expression() IExpressionContext {
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

func (s *FunctionContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(EqlParserRPAREN, 0)
}

func (s *FunctionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FunctionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EqlParserListener); ok {
		listenerT.EnterFunction(s)
	}
}

func (s *FunctionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EqlParserListener); ok {
		listenerT.ExitFunction(s)
	}
}

func (p *EqlParser) Function() (localctx IFunctionContext) {
	localctx = NewFunctionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, EqlParserRULE_function)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(52)
		p.Match(EqlParserPRINT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(53)
		p.Match(EqlParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(54)
		p.Expression()
	}
	{
		p.SetState(55)
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
	p.EnterRule(localctx, 8, EqlParserRULE_expression)
	var _la int

	p.SetState(71)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 5, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(57)
			p.Term()
		}
		p.SetState(62)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == EqlParserPLUS || _la == EqlParserMINUS {
			{
				p.SetState(58)
				_la = p.GetTokenStream().LA(1)

				if !(_la == EqlParserPLUS || _la == EqlParserMINUS) {
					p.GetErrorHandler().RecoverInline(p)
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
			}
			{
				p.SetState(59)
				p.Term()
			}

			p.SetState(64)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(65)
			p.Term()
		}
		p.SetState(69)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 4, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(66)
				p.Compair()
			}
			{
				p.SetState(67)
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
	p.EnterRule(localctx, 10, EqlParserRULE_compair)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(74)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == EqlParserGREATER_THAN || _la == EqlParserLESS_THAN {
		{
			p.SetState(73)
			_la = p.GetTokenStream().LA(1)

			if !(_la == EqlParserGREATER_THAN || _la == EqlParserLESS_THAN) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

	}
	p.SetState(77)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == EqlParserEQUAL {
		{
			p.SetState(76)
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
	p.EnterRule(localctx, 12, EqlParserRULE_term)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(79)
		p.Factor()
	}
	p.SetState(84)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == EqlParserDIV || _la == EqlParserMULT {
		{
			p.SetState(80)
			_la = p.GetTokenStream().LA(1)

			if !(_la == EqlParserDIV || _la == EqlParserMULT) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(81)
			p.Factor()
		}

		p.SetState(86)
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
	p.EnterRule(localctx, 14, EqlParserRULE_actionSpec)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(87)
		p.Type_()
	}
	{
		p.SetState(88)
		p.Match(EqlParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

	{
		p.SetState(89)
		p.Param()
	}
	p.SetState(94)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == EqlParserCOMMA {
		{
			p.SetState(90)
			p.Match(EqlParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(91)
			p.Param()
		}

		p.SetState(96)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

	{
		p.SetState(97)
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
	p.EnterRule(localctx, 16, EqlParserRULE_param)
	p.SetState(101)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 10, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(99)
			p.InputRange()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(100)
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
	p.EnterRule(localctx, 18, EqlParserRULE_type)
	p.SetState(106)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case EqlParserSUM, EqlParserABS:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(103)
			p.Math()
		}

	case EqlParserAVG, EqlParserADD, EqlParserDIVIDE, EqlParserMULTIPLY, EqlParserEQ, EqlParserCONCAT, EqlParserGT, EqlParserGTE:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(104)
			p.Operator()
		}

	case EqlParserLAMBDA, EqlParserIF:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(105)
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
	p.EnterRule(localctx, 20, EqlParserRULE_math)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(108)
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
	p.EnterRule(localctx, 22, EqlParserRULE_operator)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(110)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&4278190080) != 0) {
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
	p.EnterRule(localctx, 24, EqlParserRULE_logical)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(112)
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
	AllCell() []ICellContext
	Cell(i int) ICellContext
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

func (s *InputRangeContext) AllCell() []ICellContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ICellContext); ok {
			len++
		}
	}

	tst := make([]ICellContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ICellContext); ok {
			tst[i] = t.(ICellContext)
			i++
		}
	}

	return tst
}

func (s *InputRangeContext) Cell(i int) ICellContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICellContext); ok {
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

	return t.(ICellContext)
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
	p.EnterRule(localctx, 26, EqlParserRULE_inputRange)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(114)
		p.Cell()
	}
	{
		p.SetState(115)
		p.Match(EqlParserCOLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(116)
		p.Cell()
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

// ICellContext is an interface to support dynamic dispatch.
type ICellContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ALPHABET() antlr.TerminalNode
	INT() antlr.TerminalNode

	// IsCellContext differentiates from other interfaces.
	IsCellContext()
}

type CellContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCellContext() *CellContext {
	var p = new(CellContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = EqlParserRULE_cell
	return p
}

func InitEmptyCellContext(p *CellContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = EqlParserRULE_cell
}

func (*CellContext) IsCellContext() {}

func NewCellContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CellContext {
	var p = new(CellContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = EqlParserRULE_cell

	return p
}

func (s *CellContext) GetParser() antlr.Parser { return s.parser }

func (s *CellContext) ALPHABET() antlr.TerminalNode {
	return s.GetToken(EqlParserALPHABET, 0)
}

func (s *CellContext) INT() antlr.TerminalNode {
	return s.GetToken(EqlParserINT, 0)
}

func (s *CellContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CellContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CellContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EqlParserListener); ok {
		listenerT.EnterCell(s)
	}
}

func (s *CellContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EqlParserListener); ok {
		listenerT.ExitCell(s)
	}
}

func (p *EqlParser) Cell() (localctx ICellContext) {
	localctx = NewCellContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, EqlParserRULE_cell)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(118)
		p.Match(EqlParserALPHABET)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(119)
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
	p.EnterRule(localctx, 30, EqlParserRULE_number)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(121)
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
	Cell() ICellContext
	ActionSpec() IActionSpecContext
	TRUE() antlr.TerminalNode
	FALSE() antlr.TerminalNode
	IDENT() antlr.TerminalNode
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

func (s *FactorContext) Cell() ICellContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICellContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICellContext)
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

func (s *FactorContext) IDENT() antlr.TerminalNode {
	return s.GetToken(EqlParserIDENT, 0)
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
	p.EnterRule(localctx, 32, EqlParserRULE_factor)
	p.SetState(134)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case EqlParserINT, EqlParserDECIMAL:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(123)
			p.Number()
		}

	case EqlParserALPHABET:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(124)
			p.Cell()
		}

	case EqlParserSUM, EqlParserABS, EqlParserAVG, EqlParserADD, EqlParserDIVIDE, EqlParserMULTIPLY, EqlParserEQ, EqlParserCONCAT, EqlParserGT, EqlParserGTE, EqlParserLAMBDA, EqlParserIF:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(125)
			p.ActionSpec()
		}

	case EqlParserTRUE:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(126)
			p.Match(EqlParserTRUE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case EqlParserFALSE:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(127)
			p.Match(EqlParserFALSE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case EqlParserIDENT:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(128)
			p.Match(EqlParserIDENT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case EqlParserSTRING:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(129)
			p.Match(EqlParserSTRING)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case EqlParserLPAREN:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(130)
			p.Match(EqlParserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(131)
			p.Expression()
		}
		{
			p.SetState(132)
			p.Match(EqlParserRPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
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
