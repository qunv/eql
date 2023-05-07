// Code generated from core/antlr/EqlLexer.g4 by ANTLR 4.12.0. DO NOT EDIT.

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
		"'MULTIPLY'", "'EQ'", "'CONCAT'", "'GT'",
	}
	staticData.symbolicNames = []string{
		"", "PLUS", "MINUS", "DIV", "MULT", "LPAREN", "RPAREN", "LCURLY", "RCURLY",
		"LBRACKET", "RBRACKET", "EQUAL", "SEMI", "COLON", "COMMA", "SUM", "ABS",
		"AVG", "ADD", "DIVIDE", "MULTIPLY", "EQ", "CONCAT", "GT", "DIGIT", "IDENTIFIER",
		"WS", "EOS",
	}
	staticData.ruleNames = []string{
		"PLUS", "MINUS", "DIV", "MULT", "LPAREN", "RPAREN", "LCURLY", "RCURLY",
		"LBRACKET", "RBRACKET", "EQUAL", "SEMI", "COLON", "COMMA", "SUM", "ABS",
		"AVG", "ADD", "DIVIDE", "MULTIPLY", "EQ", "CONCAT", "GT", "DIGIT", "IDENTIFIER",
		"WS", "EOS",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 27, 167, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7,
		20, 2, 21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25,
		2, 26, 7, 26, 1, 0, 1, 0, 1, 1, 1, 1, 1, 2, 1, 2, 1, 3, 1, 3, 1, 4, 1,
		4, 1, 5, 1, 5, 1, 6, 1, 6, 1, 7, 1, 7, 1, 8, 1, 8, 1, 9, 1, 9, 1, 10, 1,
		10, 1, 11, 1, 11, 1, 12, 1, 12, 1, 13, 1, 13, 1, 14, 1, 14, 1, 14, 1, 14,
		1, 15, 1, 15, 1, 15, 1, 15, 1, 16, 1, 16, 1, 16, 1, 16, 1, 17, 1, 17, 1,
		17, 1, 17, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 19, 1, 19,
		1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 20, 1, 20, 1, 20, 1,
		21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 22, 1, 22, 1, 22, 1, 23,
		4, 23, 130, 8, 23, 11, 23, 12, 23, 131, 1, 24, 4, 24, 135, 8, 24, 11, 24,
		12, 24, 136, 1, 25, 4, 25, 140, 8, 25, 11, 25, 12, 25, 141, 1, 25, 1, 25,
		1, 26, 4, 26, 147, 8, 26, 11, 26, 12, 26, 148, 1, 26, 1, 26, 1, 26, 1,
		26, 1, 26, 5, 26, 156, 8, 26, 10, 26, 12, 26, 159, 9, 26, 1, 26, 1, 26,
		1, 26, 3, 26, 164, 8, 26, 1, 26, 1, 26, 1, 157, 0, 27, 1, 1, 3, 2, 5, 3,
		7, 4, 9, 5, 11, 6, 13, 7, 15, 8, 17, 9, 19, 10, 21, 11, 23, 12, 25, 13,
		27, 14, 29, 15, 31, 16, 33, 17, 35, 18, 37, 19, 39, 20, 41, 21, 43, 22,
		45, 23, 47, 24, 49, 25, 51, 26, 53, 27, 1, 0, 4, 1, 0, 48, 57, 2, 0, 65,
		90, 97, 122, 3, 0, 9, 10, 13, 13, 32, 32, 2, 0, 10, 10, 13, 13, 174, 0,
		1, 1, 0, 0, 0, 0, 3, 1, 0, 0, 0, 0, 5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0, 0,
		9, 1, 0, 0, 0, 0, 11, 1, 0, 0, 0, 0, 13, 1, 0, 0, 0, 0, 15, 1, 0, 0, 0,
		0, 17, 1, 0, 0, 0, 0, 19, 1, 0, 0, 0, 0, 21, 1, 0, 0, 0, 0, 23, 1, 0, 0,
		0, 0, 25, 1, 0, 0, 0, 0, 27, 1, 0, 0, 0, 0, 29, 1, 0, 0, 0, 0, 31, 1, 0,
		0, 0, 0, 33, 1, 0, 0, 0, 0, 35, 1, 0, 0, 0, 0, 37, 1, 0, 0, 0, 0, 39, 1,
		0, 0, 0, 0, 41, 1, 0, 0, 0, 0, 43, 1, 0, 0, 0, 0, 45, 1, 0, 0, 0, 0, 47,
		1, 0, 0, 0, 0, 49, 1, 0, 0, 0, 0, 51, 1, 0, 0, 0, 0, 53, 1, 0, 0, 0, 1,
		55, 1, 0, 0, 0, 3, 57, 1, 0, 0, 0, 5, 59, 1, 0, 0, 0, 7, 61, 1, 0, 0, 0,
		9, 63, 1, 0, 0, 0, 11, 65, 1, 0, 0, 0, 13, 67, 1, 0, 0, 0, 15, 69, 1, 0,
		0, 0, 17, 71, 1, 0, 0, 0, 19, 73, 1, 0, 0, 0, 21, 75, 1, 0, 0, 0, 23, 77,
		1, 0, 0, 0, 25, 79, 1, 0, 0, 0, 27, 81, 1, 0, 0, 0, 29, 83, 1, 0, 0, 0,
		31, 87, 1, 0, 0, 0, 33, 91, 1, 0, 0, 0, 35, 95, 1, 0, 0, 0, 37, 99, 1,
		0, 0, 0, 39, 106, 1, 0, 0, 0, 41, 115, 1, 0, 0, 0, 43, 118, 1, 0, 0, 0,
		45, 125, 1, 0, 0, 0, 47, 129, 1, 0, 0, 0, 49, 134, 1, 0, 0, 0, 51, 139,
		1, 0, 0, 0, 53, 163, 1, 0, 0, 0, 55, 56, 5, 43, 0, 0, 56, 2, 1, 0, 0, 0,
		57, 58, 5, 45, 0, 0, 58, 4, 1, 0, 0, 0, 59, 60, 5, 47, 0, 0, 60, 6, 1,
		0, 0, 0, 61, 62, 5, 42, 0, 0, 62, 8, 1, 0, 0, 0, 63, 64, 5, 40, 0, 0, 64,
		10, 1, 0, 0, 0, 65, 66, 5, 41, 0, 0, 66, 12, 1, 0, 0, 0, 67, 68, 5, 123,
		0, 0, 68, 14, 1, 0, 0, 0, 69, 70, 5, 125, 0, 0, 70, 16, 1, 0, 0, 0, 71,
		72, 5, 91, 0, 0, 72, 18, 1, 0, 0, 0, 73, 74, 5, 93, 0, 0, 74, 20, 1, 0,
		0, 0, 75, 76, 5, 61, 0, 0, 76, 22, 1, 0, 0, 0, 77, 78, 5, 59, 0, 0, 78,
		24, 1, 0, 0, 0, 79, 80, 5, 58, 0, 0, 80, 26, 1, 0, 0, 0, 81, 82, 5, 44,
		0, 0, 82, 28, 1, 0, 0, 0, 83, 84, 5, 83, 0, 0, 84, 85, 5, 85, 0, 0, 85,
		86, 5, 77, 0, 0, 86, 30, 1, 0, 0, 0, 87, 88, 5, 65, 0, 0, 88, 89, 5, 66,
		0, 0, 89, 90, 5, 83, 0, 0, 90, 32, 1, 0, 0, 0, 91, 92, 5, 65, 0, 0, 92,
		93, 5, 86, 0, 0, 93, 94, 5, 71, 0, 0, 94, 34, 1, 0, 0, 0, 95, 96, 5, 65,
		0, 0, 96, 97, 5, 68, 0, 0, 97, 98, 5, 68, 0, 0, 98, 36, 1, 0, 0, 0, 99,
		100, 5, 68, 0, 0, 100, 101, 5, 73, 0, 0, 101, 102, 5, 86, 0, 0, 102, 103,
		5, 73, 0, 0, 103, 104, 5, 68, 0, 0, 104, 105, 5, 69, 0, 0, 105, 38, 1,
		0, 0, 0, 106, 107, 5, 77, 0, 0, 107, 108, 5, 85, 0, 0, 108, 109, 5, 76,
		0, 0, 109, 110, 5, 84, 0, 0, 110, 111, 5, 73, 0, 0, 111, 112, 5, 80, 0,
		0, 112, 113, 5, 76, 0, 0, 113, 114, 5, 89, 0, 0, 114, 40, 1, 0, 0, 0, 115,
		116, 5, 69, 0, 0, 116, 117, 5, 81, 0, 0, 117, 42, 1, 0, 0, 0, 118, 119,
		5, 67, 0, 0, 119, 120, 5, 79, 0, 0, 120, 121, 5, 78, 0, 0, 121, 122, 5,
		67, 0, 0, 122, 123, 5, 65, 0, 0, 123, 124, 5, 84, 0, 0, 124, 44, 1, 0,
		0, 0, 125, 126, 5, 71, 0, 0, 126, 127, 5, 84, 0, 0, 127, 46, 1, 0, 0, 0,
		128, 130, 7, 0, 0, 0, 129, 128, 1, 0, 0, 0, 130, 131, 1, 0, 0, 0, 131,
		129, 1, 0, 0, 0, 131, 132, 1, 0, 0, 0, 132, 48, 1, 0, 0, 0, 133, 135, 7,
		1, 0, 0, 134, 133, 1, 0, 0, 0, 135, 136, 1, 0, 0, 0, 136, 134, 1, 0, 0,
		0, 136, 137, 1, 0, 0, 0, 137, 50, 1, 0, 0, 0, 138, 140, 7, 2, 0, 0, 139,
		138, 1, 0, 0, 0, 140, 141, 1, 0, 0, 0, 141, 139, 1, 0, 0, 0, 141, 142,
		1, 0, 0, 0, 142, 143, 1, 0, 0, 0, 143, 144, 6, 25, 0, 0, 144, 52, 1, 0,
		0, 0, 145, 147, 7, 3, 0, 0, 146, 145, 1, 0, 0, 0, 147, 148, 1, 0, 0, 0,
		148, 146, 1, 0, 0, 0, 148, 149, 1, 0, 0, 0, 149, 164, 1, 0, 0, 0, 150,
		164, 5, 59, 0, 0, 151, 152, 5, 47, 0, 0, 152, 153, 5, 42, 0, 0, 153, 157,
		1, 0, 0, 0, 154, 156, 9, 0, 0, 0, 155, 154, 1, 0, 0, 0, 156, 159, 1, 0,
		0, 0, 157, 158, 1, 0, 0, 0, 157, 155, 1, 0, 0, 0, 158, 160, 1, 0, 0, 0,
		159, 157, 1, 0, 0, 0, 160, 161, 5, 42, 0, 0, 161, 164, 5, 47, 0, 0, 162,
		164, 5, 0, 0, 1, 163, 146, 1, 0, 0, 0, 163, 150, 1, 0, 0, 0, 163, 151,
		1, 0, 0, 0, 163, 162, 1, 0, 0, 0, 164, 165, 1, 0, 0, 0, 165, 166, 6, 26,
		1, 0, 166, 54, 1, 0, 0, 0, 7, 0, 131, 136, 141, 148, 157, 163, 2, 6, 0,
		0, 2, 0, 0,
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
	EqlLexerMULTIPLY   = 20
	EqlLexerEQ         = 21
	EqlLexerCONCAT     = 22
	EqlLexerGT         = 23
	EqlLexerDIGIT      = 24
	EqlLexerIDENTIFIER = 25
	EqlLexerWS         = 26
	EqlLexerEOS        = 27
)
