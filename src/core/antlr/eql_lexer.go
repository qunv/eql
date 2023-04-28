// Code generated from src/antlr/EqlLexer.g4 by ANTLR 4.12.0. DO NOT EDIT.

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
		"'='", "';'", "':'", "','", "'=SUM'", "'=AVG'",
	}
	staticData.symbolicNames = []string{
		"", "PLUS", "MINUS", "DIV", "MULT", "LPAREN", "RPAREN", "LCURLY", "RCURLY",
		"LBRACKET", "RBRACKET", "EQUAL", "SEMI", "COLON", "COMMA", "SUM", "AVG",
		"NUMBER", "IDENTIFIER", "WS", "EOS",
	}
	staticData.ruleNames = []string{
		"PLUS", "MINUS", "DIV", "MULT", "LPAREN", "RPAREN", "LCURLY", "RCURLY",
		"LBRACKET", "RBRACKET", "EQUAL", "SEMI", "COLON", "COMMA", "SUM", "AVG",
		"NUMBER", "IDENTIFIER", "WS", "EOS",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 20, 118, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 1, 0, 1,
		0, 1, 1, 1, 1, 1, 2, 1, 2, 1, 3, 1, 3, 1, 4, 1, 4, 1, 5, 1, 5, 1, 6, 1,
		6, 1, 7, 1, 7, 1, 8, 1, 8, 1, 9, 1, 9, 1, 10, 1, 10, 1, 11, 1, 11, 1, 12,
		1, 12, 1, 13, 1, 13, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 15, 1, 15, 1,
		15, 1, 15, 1, 15, 1, 16, 4, 16, 81, 8, 16, 11, 16, 12, 16, 82, 1, 17, 4,
		17, 86, 8, 17, 11, 17, 12, 17, 87, 1, 18, 4, 18, 91, 8, 18, 11, 18, 12,
		18, 92, 1, 18, 1, 18, 1, 19, 4, 19, 98, 8, 19, 11, 19, 12, 19, 99, 1, 19,
		1, 19, 1, 19, 1, 19, 1, 19, 5, 19, 107, 8, 19, 10, 19, 12, 19, 110, 9,
		19, 1, 19, 1, 19, 1, 19, 3, 19, 115, 8, 19, 1, 19, 1, 19, 1, 108, 0, 20,
		1, 1, 3, 2, 5, 3, 7, 4, 9, 5, 11, 6, 13, 7, 15, 8, 17, 9, 19, 10, 21, 11,
		23, 12, 25, 13, 27, 14, 29, 15, 31, 16, 33, 17, 35, 18, 37, 19, 39, 20,
		1, 0, 4, 1, 0, 48, 57, 2, 0, 65, 90, 97, 122, 3, 0, 9, 10, 13, 13, 32,
		32, 2, 0, 10, 10, 13, 13, 125, 0, 1, 1, 0, 0, 0, 0, 3, 1, 0, 0, 0, 0, 5,
		1, 0, 0, 0, 0, 7, 1, 0, 0, 0, 0, 9, 1, 0, 0, 0, 0, 11, 1, 0, 0, 0, 0, 13,
		1, 0, 0, 0, 0, 15, 1, 0, 0, 0, 0, 17, 1, 0, 0, 0, 0, 19, 1, 0, 0, 0, 0,
		21, 1, 0, 0, 0, 0, 23, 1, 0, 0, 0, 0, 25, 1, 0, 0, 0, 0, 27, 1, 0, 0, 0,
		0, 29, 1, 0, 0, 0, 0, 31, 1, 0, 0, 0, 0, 33, 1, 0, 0, 0, 0, 35, 1, 0, 0,
		0, 0, 37, 1, 0, 0, 0, 0, 39, 1, 0, 0, 0, 1, 41, 1, 0, 0, 0, 3, 43, 1, 0,
		0, 0, 5, 45, 1, 0, 0, 0, 7, 47, 1, 0, 0, 0, 9, 49, 1, 0, 0, 0, 11, 51,
		1, 0, 0, 0, 13, 53, 1, 0, 0, 0, 15, 55, 1, 0, 0, 0, 17, 57, 1, 0, 0, 0,
		19, 59, 1, 0, 0, 0, 21, 61, 1, 0, 0, 0, 23, 63, 1, 0, 0, 0, 25, 65, 1,
		0, 0, 0, 27, 67, 1, 0, 0, 0, 29, 69, 1, 0, 0, 0, 31, 74, 1, 0, 0, 0, 33,
		80, 1, 0, 0, 0, 35, 85, 1, 0, 0, 0, 37, 90, 1, 0, 0, 0, 39, 114, 1, 0,
		0, 0, 41, 42, 5, 43, 0, 0, 42, 2, 1, 0, 0, 0, 43, 44, 5, 45, 0, 0, 44,
		4, 1, 0, 0, 0, 45, 46, 5, 47, 0, 0, 46, 6, 1, 0, 0, 0, 47, 48, 5, 42, 0,
		0, 48, 8, 1, 0, 0, 0, 49, 50, 5, 40, 0, 0, 50, 10, 1, 0, 0, 0, 51, 52,
		5, 41, 0, 0, 52, 12, 1, 0, 0, 0, 53, 54, 5, 123, 0, 0, 54, 14, 1, 0, 0,
		0, 55, 56, 5, 125, 0, 0, 56, 16, 1, 0, 0, 0, 57, 58, 5, 91, 0, 0, 58, 18,
		1, 0, 0, 0, 59, 60, 5, 93, 0, 0, 60, 20, 1, 0, 0, 0, 61, 62, 5, 61, 0,
		0, 62, 22, 1, 0, 0, 0, 63, 64, 5, 59, 0, 0, 64, 24, 1, 0, 0, 0, 65, 66,
		5, 58, 0, 0, 66, 26, 1, 0, 0, 0, 67, 68, 5, 44, 0, 0, 68, 28, 1, 0, 0,
		0, 69, 70, 5, 61, 0, 0, 70, 71, 5, 83, 0, 0, 71, 72, 5, 85, 0, 0, 72, 73,
		5, 77, 0, 0, 73, 30, 1, 0, 0, 0, 74, 75, 5, 61, 0, 0, 75, 76, 5, 65, 0,
		0, 76, 77, 5, 86, 0, 0, 77, 78, 5, 71, 0, 0, 78, 32, 1, 0, 0, 0, 79, 81,
		7, 0, 0, 0, 80, 79, 1, 0, 0, 0, 81, 82, 1, 0, 0, 0, 82, 80, 1, 0, 0, 0,
		82, 83, 1, 0, 0, 0, 83, 34, 1, 0, 0, 0, 84, 86, 7, 1, 0, 0, 85, 84, 1,
		0, 0, 0, 86, 87, 1, 0, 0, 0, 87, 85, 1, 0, 0, 0, 87, 88, 1, 0, 0, 0, 88,
		36, 1, 0, 0, 0, 89, 91, 7, 2, 0, 0, 90, 89, 1, 0, 0, 0, 91, 92, 1, 0, 0,
		0, 92, 90, 1, 0, 0, 0, 92, 93, 1, 0, 0, 0, 93, 94, 1, 0, 0, 0, 94, 95,
		6, 18, 0, 0, 95, 38, 1, 0, 0, 0, 96, 98, 7, 3, 0, 0, 97, 96, 1, 0, 0, 0,
		98, 99, 1, 0, 0, 0, 99, 97, 1, 0, 0, 0, 99, 100, 1, 0, 0, 0, 100, 115,
		1, 0, 0, 0, 101, 115, 5, 59, 0, 0, 102, 103, 5, 47, 0, 0, 103, 104, 5,
		42, 0, 0, 104, 108, 1, 0, 0, 0, 105, 107, 9, 0, 0, 0, 106, 105, 1, 0, 0,
		0, 107, 110, 1, 0, 0, 0, 108, 109, 1, 0, 0, 0, 108, 106, 1, 0, 0, 0, 109,
		111, 1, 0, 0, 0, 110, 108, 1, 0, 0, 0, 111, 112, 5, 42, 0, 0, 112, 115,
		5, 47, 0, 0, 113, 115, 5, 0, 0, 1, 114, 97, 1, 0, 0, 0, 114, 101, 1, 0,
		0, 0, 114, 102, 1, 0, 0, 0, 114, 113, 1, 0, 0, 0, 115, 116, 1, 0, 0, 0,
		116, 117, 6, 19, 1, 0, 117, 40, 1, 0, 0, 0, 7, 0, 82, 87, 92, 99, 108,
		114, 2, 6, 0, 0, 2, 0, 0,
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
	EqlLexerAVG        = 16
	EqlLexerNUMBER     = 17
	EqlLexerIDENTIFIER = 18
	EqlLexerWS         = 19
	EqlLexerEOS        = 20
)
