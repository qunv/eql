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
		"'='", "';'", "':'", "','", "'SUM'", "'AVG'", "'ABS'", "'ADD'",
	}
	staticData.symbolicNames = []string{
		"", "PLUS", "MINUS", "DIV", "MULT", "LPAREN", "RPAREN", "LCURLY", "RCURLY",
		"LBRACKET", "RBRACKET", "EQUAL", "SEMI", "COLON", "COMMA", "SUM", "AVG",
		"ABS", "ADD", "DIGIT", "IDENTIFIER", "WS", "EOS",
	}
	staticData.ruleNames = []string{
		"PLUS", "MINUS", "DIV", "MULT", "LPAREN", "RPAREN", "LCURLY", "RCURLY",
		"LBRACKET", "RBRACKET", "EQUAL", "SEMI", "COLON", "COMMA", "SUM", "AVG",
		"ABS", "ADD", "DIGIT", "IDENTIFIER", "WS", "EOS",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 22, 128, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7,
		20, 2, 21, 7, 21, 1, 0, 1, 0, 1, 1, 1, 1, 1, 2, 1, 2, 1, 3, 1, 3, 1, 4,
		1, 4, 1, 5, 1, 5, 1, 6, 1, 6, 1, 7, 1, 7, 1, 8, 1, 8, 1, 9, 1, 9, 1, 10,
		1, 10, 1, 11, 1, 11, 1, 12, 1, 12, 1, 13, 1, 13, 1, 14, 1, 14, 1, 14, 1,
		14, 1, 15, 1, 15, 1, 15, 1, 15, 1, 16, 1, 16, 1, 16, 1, 16, 1, 17, 1, 17,
		1, 17, 1, 17, 1, 18, 4, 18, 91, 8, 18, 11, 18, 12, 18, 92, 1, 19, 4, 19,
		96, 8, 19, 11, 19, 12, 19, 97, 1, 20, 4, 20, 101, 8, 20, 11, 20, 12, 20,
		102, 1, 20, 1, 20, 1, 21, 4, 21, 108, 8, 21, 11, 21, 12, 21, 109, 1, 21,
		1, 21, 1, 21, 1, 21, 1, 21, 5, 21, 117, 8, 21, 10, 21, 12, 21, 120, 9,
		21, 1, 21, 1, 21, 1, 21, 3, 21, 125, 8, 21, 1, 21, 1, 21, 1, 118, 0, 22,
		1, 1, 3, 2, 5, 3, 7, 4, 9, 5, 11, 6, 13, 7, 15, 8, 17, 9, 19, 10, 21, 11,
		23, 12, 25, 13, 27, 14, 29, 15, 31, 16, 33, 17, 35, 18, 37, 19, 39, 20,
		41, 21, 43, 22, 1, 0, 4, 1, 0, 48, 57, 2, 0, 65, 90, 97, 122, 3, 0, 9,
		10, 13, 13, 32, 32, 2, 0, 10, 10, 13, 13, 135, 0, 1, 1, 0, 0, 0, 0, 3,
		1, 0, 0, 0, 0, 5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0, 0, 9, 1, 0, 0, 0, 0, 11,
		1, 0, 0, 0, 0, 13, 1, 0, 0, 0, 0, 15, 1, 0, 0, 0, 0, 17, 1, 0, 0, 0, 0,
		19, 1, 0, 0, 0, 0, 21, 1, 0, 0, 0, 0, 23, 1, 0, 0, 0, 0, 25, 1, 0, 0, 0,
		0, 27, 1, 0, 0, 0, 0, 29, 1, 0, 0, 0, 0, 31, 1, 0, 0, 0, 0, 33, 1, 0, 0,
		0, 0, 35, 1, 0, 0, 0, 0, 37, 1, 0, 0, 0, 0, 39, 1, 0, 0, 0, 0, 41, 1, 0,
		0, 0, 0, 43, 1, 0, 0, 0, 1, 45, 1, 0, 0, 0, 3, 47, 1, 0, 0, 0, 5, 49, 1,
		0, 0, 0, 7, 51, 1, 0, 0, 0, 9, 53, 1, 0, 0, 0, 11, 55, 1, 0, 0, 0, 13,
		57, 1, 0, 0, 0, 15, 59, 1, 0, 0, 0, 17, 61, 1, 0, 0, 0, 19, 63, 1, 0, 0,
		0, 21, 65, 1, 0, 0, 0, 23, 67, 1, 0, 0, 0, 25, 69, 1, 0, 0, 0, 27, 71,
		1, 0, 0, 0, 29, 73, 1, 0, 0, 0, 31, 77, 1, 0, 0, 0, 33, 81, 1, 0, 0, 0,
		35, 85, 1, 0, 0, 0, 37, 90, 1, 0, 0, 0, 39, 95, 1, 0, 0, 0, 41, 100, 1,
		0, 0, 0, 43, 124, 1, 0, 0, 0, 45, 46, 5, 43, 0, 0, 46, 2, 1, 0, 0, 0, 47,
		48, 5, 45, 0, 0, 48, 4, 1, 0, 0, 0, 49, 50, 5, 47, 0, 0, 50, 6, 1, 0, 0,
		0, 51, 52, 5, 42, 0, 0, 52, 8, 1, 0, 0, 0, 53, 54, 5, 40, 0, 0, 54, 10,
		1, 0, 0, 0, 55, 56, 5, 41, 0, 0, 56, 12, 1, 0, 0, 0, 57, 58, 5, 123, 0,
		0, 58, 14, 1, 0, 0, 0, 59, 60, 5, 125, 0, 0, 60, 16, 1, 0, 0, 0, 61, 62,
		5, 91, 0, 0, 62, 18, 1, 0, 0, 0, 63, 64, 5, 93, 0, 0, 64, 20, 1, 0, 0,
		0, 65, 66, 5, 61, 0, 0, 66, 22, 1, 0, 0, 0, 67, 68, 5, 59, 0, 0, 68, 24,
		1, 0, 0, 0, 69, 70, 5, 58, 0, 0, 70, 26, 1, 0, 0, 0, 71, 72, 5, 44, 0,
		0, 72, 28, 1, 0, 0, 0, 73, 74, 5, 83, 0, 0, 74, 75, 5, 85, 0, 0, 75, 76,
		5, 77, 0, 0, 76, 30, 1, 0, 0, 0, 77, 78, 5, 65, 0, 0, 78, 79, 5, 86, 0,
		0, 79, 80, 5, 71, 0, 0, 80, 32, 1, 0, 0, 0, 81, 82, 5, 65, 0, 0, 82, 83,
		5, 66, 0, 0, 83, 84, 5, 83, 0, 0, 84, 34, 1, 0, 0, 0, 85, 86, 5, 65, 0,
		0, 86, 87, 5, 68, 0, 0, 87, 88, 5, 68, 0, 0, 88, 36, 1, 0, 0, 0, 89, 91,
		7, 0, 0, 0, 90, 89, 1, 0, 0, 0, 91, 92, 1, 0, 0, 0, 92, 90, 1, 0, 0, 0,
		92, 93, 1, 0, 0, 0, 93, 38, 1, 0, 0, 0, 94, 96, 7, 1, 0, 0, 95, 94, 1,
		0, 0, 0, 96, 97, 1, 0, 0, 0, 97, 95, 1, 0, 0, 0, 97, 98, 1, 0, 0, 0, 98,
		40, 1, 0, 0, 0, 99, 101, 7, 2, 0, 0, 100, 99, 1, 0, 0, 0, 101, 102, 1,
		0, 0, 0, 102, 100, 1, 0, 0, 0, 102, 103, 1, 0, 0, 0, 103, 104, 1, 0, 0,
		0, 104, 105, 6, 20, 0, 0, 105, 42, 1, 0, 0, 0, 106, 108, 7, 3, 0, 0, 107,
		106, 1, 0, 0, 0, 108, 109, 1, 0, 0, 0, 109, 107, 1, 0, 0, 0, 109, 110,
		1, 0, 0, 0, 110, 125, 1, 0, 0, 0, 111, 125, 5, 59, 0, 0, 112, 113, 5, 47,
		0, 0, 113, 114, 5, 42, 0, 0, 114, 118, 1, 0, 0, 0, 115, 117, 9, 0, 0, 0,
		116, 115, 1, 0, 0, 0, 117, 120, 1, 0, 0, 0, 118, 119, 1, 0, 0, 0, 118,
		116, 1, 0, 0, 0, 119, 121, 1, 0, 0, 0, 120, 118, 1, 0, 0, 0, 121, 122,
		5, 42, 0, 0, 122, 125, 5, 47, 0, 0, 123, 125, 5, 0, 0, 1, 124, 107, 1,
		0, 0, 0, 124, 111, 1, 0, 0, 0, 124, 112, 1, 0, 0, 0, 124, 123, 1, 0, 0,
		0, 125, 126, 1, 0, 0, 0, 126, 127, 6, 21, 1, 0, 127, 44, 1, 0, 0, 0, 7,
		0, 92, 97, 102, 109, 118, 124, 2, 6, 0, 0, 2, 0, 0,
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
	EqlLexerABS        = 17
	EqlLexerADD        = 18
	EqlLexerDIGIT      = 19
	EqlLexerIDENTIFIER = 20
	EqlLexerWS         = 21
	EqlLexerEOS        = 22
)
