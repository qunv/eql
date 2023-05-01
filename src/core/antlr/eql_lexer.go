// Code generated from src/core/antlr/EqlLexer.g4 by ANTLR 4.12.0. DO NOT EDIT.

package antlr

import (
	"fmt"
	"sync"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = sync.Once{}
var _ = unicode.IsLetter

type EqlLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var eqllexerLexerStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	channelNames           []string
	modeNames              []string
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func eqllexerLexerInit() {
	staticData := &eqllexerLexerStaticData
	staticData.channelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.modeNames = []string{
		"DEFAULT_MODE",
	}
	staticData.literalNames = []string{
		"", "'+'", "'-'", "'/'", "'*'", "'('", "')'", "'{'", "'}'", "'['", "']'",
		"'='", "';'", "':'", "','", "'SUM'", "'ABS'", "'AVG'", "'ADD'", "'DIVIDE'",
		"'EQ'", "'CONCAT'",
	}
	staticData.symbolicNames = []string{
		"", "PLUS", "MINUS", "DIV", "MULT", "LPAREN", "RPAREN", "LCURLY", "RCURLY",
		"LBRACKET", "RBRACKET", "EQUAL", "SEMI", "COLON", "COMMA", "SUM", "ABS",
		"AVG", "ADD", "DIVIDE", "EQ", "CONCAT", "DIGIT", "IDENTIFIER", "WS",
		"EOS",
	}
	staticData.ruleNames = []string{
		"PLUS", "MINUS", "DIV", "MULT", "LPAREN", "RPAREN", "LCURLY", "RCURLY",
		"LBRACKET", "RBRACKET", "EQUAL", "SEMI", "COLON", "COMMA", "SUM", "ABS",
		"AVG", "ADD", "DIVIDE", "EQ", "CONCAT", "DIGIT", "IDENTIFIER", "WS",
		"EOS",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 25, 151, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7,
		20, 2, 21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 1, 0, 1, 0,
		1, 1, 1, 1, 1, 2, 1, 2, 1, 3, 1, 3, 1, 4, 1, 4, 1, 5, 1, 5, 1, 6, 1, 6,
		1, 7, 1, 7, 1, 8, 1, 8, 1, 9, 1, 9, 1, 10, 1, 10, 1, 11, 1, 11, 1, 12,
		1, 12, 1, 13, 1, 13, 1, 14, 1, 14, 1, 14, 1, 14, 1, 15, 1, 15, 1, 15, 1,
		15, 1, 16, 1, 16, 1, 16, 1, 16, 1, 17, 1, 17, 1, 17, 1, 17, 1, 18, 1, 18,
		1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 19, 1, 19, 1, 19, 1, 20, 1, 20, 1,
		20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 21, 4, 21, 114, 8, 21, 11, 21, 12, 21,
		115, 1, 22, 4, 22, 119, 8, 22, 11, 22, 12, 22, 120, 1, 23, 4, 23, 124,
		8, 23, 11, 23, 12, 23, 125, 1, 23, 1, 23, 1, 24, 4, 24, 131, 8, 24, 11,
		24, 12, 24, 132, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 5, 24, 140, 8, 24,
		10, 24, 12, 24, 143, 9, 24, 1, 24, 1, 24, 1, 24, 3, 24, 148, 8, 24, 1,
		24, 1, 24, 1, 141, 0, 25, 1, 1, 3, 2, 5, 3, 7, 4, 9, 5, 11, 6, 13, 7, 15,
		8, 17, 9, 19, 10, 21, 11, 23, 12, 25, 13, 27, 14, 29, 15, 31, 16, 33, 17,
		35, 18, 37, 19, 39, 20, 41, 21, 43, 22, 45, 23, 47, 24, 49, 25, 1, 0, 4,
		1, 0, 48, 57, 2, 0, 65, 90, 97, 122, 3, 0, 9, 10, 13, 13, 32, 32, 2, 0,
		10, 10, 13, 13, 158, 0, 1, 1, 0, 0, 0, 0, 3, 1, 0, 0, 0, 0, 5, 1, 0, 0,
		0, 0, 7, 1, 0, 0, 0, 0, 9, 1, 0, 0, 0, 0, 11, 1, 0, 0, 0, 0, 13, 1, 0,
		0, 0, 0, 15, 1, 0, 0, 0, 0, 17, 1, 0, 0, 0, 0, 19, 1, 0, 0, 0, 0, 21, 1,
		0, 0, 0, 0, 23, 1, 0, 0, 0, 0, 25, 1, 0, 0, 0, 0, 27, 1, 0, 0, 0, 0, 29,
		1, 0, 0, 0, 0, 31, 1, 0, 0, 0, 0, 33, 1, 0, 0, 0, 0, 35, 1, 0, 0, 0, 0,
		37, 1, 0, 0, 0, 0, 39, 1, 0, 0, 0, 0, 41, 1, 0, 0, 0, 0, 43, 1, 0, 0, 0,
		0, 45, 1, 0, 0, 0, 0, 47, 1, 0, 0, 0, 0, 49, 1, 0, 0, 0, 1, 51, 1, 0, 0,
		0, 3, 53, 1, 0, 0, 0, 5, 55, 1, 0, 0, 0, 7, 57, 1, 0, 0, 0, 9, 59, 1, 0,
		0, 0, 11, 61, 1, 0, 0, 0, 13, 63, 1, 0, 0, 0, 15, 65, 1, 0, 0, 0, 17, 67,
		1, 0, 0, 0, 19, 69, 1, 0, 0, 0, 21, 71, 1, 0, 0, 0, 23, 73, 1, 0, 0, 0,
		25, 75, 1, 0, 0, 0, 27, 77, 1, 0, 0, 0, 29, 79, 1, 0, 0, 0, 31, 83, 1,
		0, 0, 0, 33, 87, 1, 0, 0, 0, 35, 91, 1, 0, 0, 0, 37, 95, 1, 0, 0, 0, 39,
		102, 1, 0, 0, 0, 41, 105, 1, 0, 0, 0, 43, 113, 1, 0, 0, 0, 45, 118, 1,
		0, 0, 0, 47, 123, 1, 0, 0, 0, 49, 147, 1, 0, 0, 0, 51, 52, 5, 43, 0, 0,
		52, 2, 1, 0, 0, 0, 53, 54, 5, 45, 0, 0, 54, 4, 1, 0, 0, 0, 55, 56, 5, 47,
		0, 0, 56, 6, 1, 0, 0, 0, 57, 58, 5, 42, 0, 0, 58, 8, 1, 0, 0, 0, 59, 60,
		5, 40, 0, 0, 60, 10, 1, 0, 0, 0, 61, 62, 5, 41, 0, 0, 62, 12, 1, 0, 0,
		0, 63, 64, 5, 123, 0, 0, 64, 14, 1, 0, 0, 0, 65, 66, 5, 125, 0, 0, 66,
		16, 1, 0, 0, 0, 67, 68, 5, 91, 0, 0, 68, 18, 1, 0, 0, 0, 69, 70, 5, 93,
		0, 0, 70, 20, 1, 0, 0, 0, 71, 72, 5, 61, 0, 0, 72, 22, 1, 0, 0, 0, 73,
		74, 5, 59, 0, 0, 74, 24, 1, 0, 0, 0, 75, 76, 5, 58, 0, 0, 76, 26, 1, 0,
		0, 0, 77, 78, 5, 44, 0, 0, 78, 28, 1, 0, 0, 0, 79, 80, 5, 83, 0, 0, 80,
		81, 5, 85, 0, 0, 81, 82, 5, 77, 0, 0, 82, 30, 1, 0, 0, 0, 83, 84, 5, 65,
		0, 0, 84, 85, 5, 66, 0, 0, 85, 86, 5, 83, 0, 0, 86, 32, 1, 0, 0, 0, 87,
		88, 5, 65, 0, 0, 88, 89, 5, 86, 0, 0, 89, 90, 5, 71, 0, 0, 90, 34, 1, 0,
		0, 0, 91, 92, 5, 65, 0, 0, 92, 93, 5, 68, 0, 0, 93, 94, 5, 68, 0, 0, 94,
		36, 1, 0, 0, 0, 95, 96, 5, 68, 0, 0, 96, 97, 5, 73, 0, 0, 97, 98, 5, 86,
		0, 0, 98, 99, 5, 73, 0, 0, 99, 100, 5, 68, 0, 0, 100, 101, 5, 69, 0, 0,
		101, 38, 1, 0, 0, 0, 102, 103, 5, 69, 0, 0, 103, 104, 5, 81, 0, 0, 104,
		40, 1, 0, 0, 0, 105, 106, 5, 67, 0, 0, 106, 107, 5, 79, 0, 0, 107, 108,
		5, 78, 0, 0, 108, 109, 5, 67, 0, 0, 109, 110, 5, 65, 0, 0, 110, 111, 5,
		84, 0, 0, 111, 42, 1, 0, 0, 0, 112, 114, 7, 0, 0, 0, 113, 112, 1, 0, 0,
		0, 114, 115, 1, 0, 0, 0, 115, 113, 1, 0, 0, 0, 115, 116, 1, 0, 0, 0, 116,
		44, 1, 0, 0, 0, 117, 119, 7, 1, 0, 0, 118, 117, 1, 0, 0, 0, 119, 120, 1,
		0, 0, 0, 120, 118, 1, 0, 0, 0, 120, 121, 1, 0, 0, 0, 121, 46, 1, 0, 0,
		0, 122, 124, 7, 2, 0, 0, 123, 122, 1, 0, 0, 0, 124, 125, 1, 0, 0, 0, 125,
		123, 1, 0, 0, 0, 125, 126, 1, 0, 0, 0, 126, 127, 1, 0, 0, 0, 127, 128,
		6, 23, 0, 0, 128, 48, 1, 0, 0, 0, 129, 131, 7, 3, 0, 0, 130, 129, 1, 0,
		0, 0, 131, 132, 1, 0, 0, 0, 132, 130, 1, 0, 0, 0, 132, 133, 1, 0, 0, 0,
		133, 148, 1, 0, 0, 0, 134, 148, 5, 59, 0, 0, 135, 136, 5, 47, 0, 0, 136,
		137, 5, 42, 0, 0, 137, 141, 1, 0, 0, 0, 138, 140, 9, 0, 0, 0, 139, 138,
		1, 0, 0, 0, 140, 143, 1, 0, 0, 0, 141, 142, 1, 0, 0, 0, 141, 139, 1, 0,
		0, 0, 142, 144, 1, 0, 0, 0, 143, 141, 1, 0, 0, 0, 144, 145, 5, 42, 0, 0,
		145, 148, 5, 47, 0, 0, 146, 148, 5, 0, 0, 1, 147, 130, 1, 0, 0, 0, 147,
		134, 1, 0, 0, 0, 147, 135, 1, 0, 0, 0, 147, 146, 1, 0, 0, 0, 148, 149,
		1, 0, 0, 0, 149, 150, 6, 24, 1, 0, 150, 50, 1, 0, 0, 0, 7, 0, 115, 120,
		125, 132, 141, 147, 2, 6, 0, 0, 2, 0, 0,
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

// EqlLexerInit initializes any static state used to implement EqlLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewEqlLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func EqlLexerInit() {
	staticData := &eqllexerLexerStaticData
	staticData.once.Do(eqllexerLexerInit)
}

// NewEqlLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewEqlLexer(input antlr.CharStream) *EqlLexer {
	EqlLexerInit()
	l := new(EqlLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &eqllexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	l.channelNames = staticData.channelNames
	l.modeNames = staticData.modeNames
	l.RuleNames = staticData.ruleNames
	l.LiteralNames = staticData.literalNames
	l.SymbolicNames = staticData.symbolicNames
	l.GrammarFileName = "EqlLexer.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// EqlLexer tokens.
const (
	EqlLexerPLUS       = 1
	EqlLexerMINUS      = 2
	EqlLexerDIV        = 3
	EqlLexerMULT       = 4
	EqlLexerLPAREN     = 5
	EqlLexerRPAREN     = 6
	EqlLexerLCURLY     = 7
	EqlLexerRCURLY     = 8
	EqlLexerLBRACKET   = 9
	EqlLexerRBRACKET   = 10
	EqlLexerEQUAL      = 11
	EqlLexerSEMI       = 12
	EqlLexerCOLON      = 13
	EqlLexerCOMMA      = 14
	EqlLexerSUM        = 15
	EqlLexerABS        = 16
	EqlLexerAVG        = 17
	EqlLexerADD        = 18
	EqlLexerDIVIDE     = 19
	EqlLexerEQ         = 20
	EqlLexerCONCAT     = 21
	EqlLexerDIGIT      = 22
	EqlLexerIDENTIFIER = 23
	EqlLexerWS         = 24
	EqlLexerEOS        = 25
)
