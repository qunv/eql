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
		"'='", "';'", "':'", "','", "'.'", "'>'", "'<'", "'TRUE'", "'FALSE'",
		"'SUM'", "'ABS'", "'AVG'", "'ADD'", "'DIVIDE'", "'MULTIPLY'", "'EQ'",
		"'CONCAT'", "'GT'", "'GTE'", "", "", "", "", "'IF'",
	}
	staticData.symbolicNames = []string{
		"", "PLUS", "MINUS", "DIV", "MULT", "LPAREN", "RPAREN", "LCURLY", "RCURLY",
		"LBRACKET", "RBRACKET", "EQUAL", "SEMI", "COLON", "COMMA", "DOT", "GREATER_THAN",
		"LESS_THAN", "TRUE", "FALSE", "SUM", "ABS", "AVG", "ADD", "DIVIDE",
		"MULTIPLY", "EQ", "CONCAT", "GT", "GTE", "INT", "DECIMAL", "DIGIT",
		"STRING", "IF", "IDENTIFIER", "WS", "EOS",
	}
	staticData.ruleNames = []string{
		"PLUS", "MINUS", "DIV", "MULT", "LPAREN", "RPAREN", "LCURLY", "RCURLY",
		"LBRACKET", "RBRACKET", "EQUAL", "SEMI", "COLON", "COMMA", "DOT", "GREATER_THAN",
		"LESS_THAN", "TRUE", "FALSE", "SUM", "ABS", "AVG", "ADD", "DIVIDE",
		"MULTIPLY", "EQ", "CONCAT", "GT", "GTE", "INT", "DECIMAL", "DIGIT",
		"STRING", "ESC_SEQ", "DIGITS", "IF", "IDENTIFIER", "WS", "EOS",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 37, 242, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7,
		20, 2, 21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25,
		2, 26, 7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2,
		31, 7, 31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 2, 35, 7, 35, 2, 36,
		7, 36, 2, 37, 7, 37, 2, 38, 7, 38, 1, 0, 1, 0, 1, 1, 1, 1, 1, 2, 1, 2,
		1, 3, 1, 3, 1, 4, 1, 4, 1, 5, 1, 5, 1, 6, 1, 6, 1, 7, 1, 7, 1, 8, 1, 8,
		1, 9, 1, 9, 1, 10, 1, 10, 1, 11, 1, 11, 1, 12, 1, 12, 1, 13, 1, 13, 1,
		14, 1, 14, 1, 15, 1, 15, 1, 16, 1, 16, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17,
		1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 19, 1, 19, 1, 19, 1, 19, 1,
		20, 1, 20, 1, 20, 1, 20, 1, 21, 1, 21, 1, 21, 1, 21, 1, 22, 1, 22, 1, 22,
		1, 22, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 24, 1, 24, 1,
		24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 25, 1, 25, 1, 25, 1, 26,
		1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 27, 1, 27, 1, 27, 1, 28, 1,
		28, 1, 28, 1, 28, 1, 29, 3, 29, 175, 8, 29, 1, 29, 1, 29, 1, 30, 3, 30,
		180, 8, 30, 1, 30, 1, 30, 1, 30, 1, 30, 1, 31, 1, 31, 1, 32, 1, 32, 1,
		32, 5, 32, 191, 8, 32, 10, 32, 12, 32, 194, 9, 32, 1, 32, 1, 32, 1, 33,
		1, 33, 1, 33, 1, 34, 4, 34, 202, 8, 34, 11, 34, 12, 34, 203, 1, 35, 1,
		35, 1, 35, 1, 36, 4, 36, 210, 8, 36, 11, 36, 12, 36, 211, 1, 37, 4, 37,
		215, 8, 37, 11, 37, 12, 37, 216, 1, 37, 1, 37, 1, 38, 4, 38, 222, 8, 38,
		11, 38, 12, 38, 223, 1, 38, 1, 38, 1, 38, 1, 38, 1, 38, 5, 38, 231, 8,
		38, 10, 38, 12, 38, 234, 9, 38, 1, 38, 1, 38, 1, 38, 3, 38, 239, 8, 38,
		1, 38, 1, 38, 1, 232, 0, 39, 1, 1, 3, 2, 5, 3, 7, 4, 9, 5, 11, 6, 13, 7,
		15, 8, 17, 9, 19, 10, 21, 11, 23, 12, 25, 13, 27, 14, 29, 15, 31, 16, 33,
		17, 35, 18, 37, 19, 39, 20, 41, 21, 43, 22, 45, 23, 47, 24, 49, 25, 51,
		26, 53, 27, 55, 28, 57, 29, 59, 30, 61, 31, 63, 32, 65, 33, 67, 0, 69,
		0, 71, 34, 73, 35, 75, 36, 77, 37, 1, 0, 6, 1, 0, 48, 57, 2, 0, 34, 34,
		92, 92, 8, 0, 34, 34, 39, 39, 92, 92, 98, 98, 102, 102, 110, 110, 114,
		114, 116, 116, 2, 0, 65, 90, 97, 122, 3, 0, 9, 10, 13, 13, 32, 32, 2, 0,
		10, 10, 13, 13, 251, 0, 1, 1, 0, 0, 0, 0, 3, 1, 0, 0, 0, 0, 5, 1, 0, 0,
		0, 0, 7, 1, 0, 0, 0, 0, 9, 1, 0, 0, 0, 0, 11, 1, 0, 0, 0, 0, 13, 1, 0,
		0, 0, 0, 15, 1, 0, 0, 0, 0, 17, 1, 0, 0, 0, 0, 19, 1, 0, 0, 0, 0, 21, 1,
		0, 0, 0, 0, 23, 1, 0, 0, 0, 0, 25, 1, 0, 0, 0, 0, 27, 1, 0, 0, 0, 0, 29,
		1, 0, 0, 0, 0, 31, 1, 0, 0, 0, 0, 33, 1, 0, 0, 0, 0, 35, 1, 0, 0, 0, 0,
		37, 1, 0, 0, 0, 0, 39, 1, 0, 0, 0, 0, 41, 1, 0, 0, 0, 0, 43, 1, 0, 0, 0,
		0, 45, 1, 0, 0, 0, 0, 47, 1, 0, 0, 0, 0, 49, 1, 0, 0, 0, 0, 51, 1, 0, 0,
		0, 0, 53, 1, 0, 0, 0, 0, 55, 1, 0, 0, 0, 0, 57, 1, 0, 0, 0, 0, 59, 1, 0,
		0, 0, 0, 61, 1, 0, 0, 0, 0, 63, 1, 0, 0, 0, 0, 65, 1, 0, 0, 0, 0, 71, 1,
		0, 0, 0, 0, 73, 1, 0, 0, 0, 0, 75, 1, 0, 0, 0, 0, 77, 1, 0, 0, 0, 1, 79,
		1, 0, 0, 0, 3, 81, 1, 0, 0, 0, 5, 83, 1, 0, 0, 0, 7, 85, 1, 0, 0, 0, 9,
		87, 1, 0, 0, 0, 11, 89, 1, 0, 0, 0, 13, 91, 1, 0, 0, 0, 15, 93, 1, 0, 0,
		0, 17, 95, 1, 0, 0, 0, 19, 97, 1, 0, 0, 0, 21, 99, 1, 0, 0, 0, 23, 101,
		1, 0, 0, 0, 25, 103, 1, 0, 0, 0, 27, 105, 1, 0, 0, 0, 29, 107, 1, 0, 0,
		0, 31, 109, 1, 0, 0, 0, 33, 111, 1, 0, 0, 0, 35, 113, 1, 0, 0, 0, 37, 118,
		1, 0, 0, 0, 39, 124, 1, 0, 0, 0, 41, 128, 1, 0, 0, 0, 43, 132, 1, 0, 0,
		0, 45, 136, 1, 0, 0, 0, 47, 140, 1, 0, 0, 0, 49, 147, 1, 0, 0, 0, 51, 156,
		1, 0, 0, 0, 53, 159, 1, 0, 0, 0, 55, 166, 1, 0, 0, 0, 57, 169, 1, 0, 0,
		0, 59, 174, 1, 0, 0, 0, 61, 179, 1, 0, 0, 0, 63, 185, 1, 0, 0, 0, 65, 187,
		1, 0, 0, 0, 67, 197, 1, 0, 0, 0, 69, 201, 1, 0, 0, 0, 71, 205, 1, 0, 0,
		0, 73, 209, 1, 0, 0, 0, 75, 214, 1, 0, 0, 0, 77, 238, 1, 0, 0, 0, 79, 80,
		5, 43, 0, 0, 80, 2, 1, 0, 0, 0, 81, 82, 5, 45, 0, 0, 82, 4, 1, 0, 0, 0,
		83, 84, 5, 47, 0, 0, 84, 6, 1, 0, 0, 0, 85, 86, 5, 42, 0, 0, 86, 8, 1,
		0, 0, 0, 87, 88, 5, 40, 0, 0, 88, 10, 1, 0, 0, 0, 89, 90, 5, 41, 0, 0,
		90, 12, 1, 0, 0, 0, 91, 92, 5, 123, 0, 0, 92, 14, 1, 0, 0, 0, 93, 94, 5,
		125, 0, 0, 94, 16, 1, 0, 0, 0, 95, 96, 5, 91, 0, 0, 96, 18, 1, 0, 0, 0,
		97, 98, 5, 93, 0, 0, 98, 20, 1, 0, 0, 0, 99, 100, 5, 61, 0, 0, 100, 22,
		1, 0, 0, 0, 101, 102, 5, 59, 0, 0, 102, 24, 1, 0, 0, 0, 103, 104, 5, 58,
		0, 0, 104, 26, 1, 0, 0, 0, 105, 106, 5, 44, 0, 0, 106, 28, 1, 0, 0, 0,
		107, 108, 5, 46, 0, 0, 108, 30, 1, 0, 0, 0, 109, 110, 5, 62, 0, 0, 110,
		32, 1, 0, 0, 0, 111, 112, 5, 60, 0, 0, 112, 34, 1, 0, 0, 0, 113, 114, 5,
		84, 0, 0, 114, 115, 5, 82, 0, 0, 115, 116, 5, 85, 0, 0, 116, 117, 5, 69,
		0, 0, 117, 36, 1, 0, 0, 0, 118, 119, 5, 70, 0, 0, 119, 120, 5, 65, 0, 0,
		120, 121, 5, 76, 0, 0, 121, 122, 5, 83, 0, 0, 122, 123, 5, 69, 0, 0, 123,
		38, 1, 0, 0, 0, 124, 125, 5, 83, 0, 0, 125, 126, 5, 85, 0, 0, 126, 127,
		5, 77, 0, 0, 127, 40, 1, 0, 0, 0, 128, 129, 5, 65, 0, 0, 129, 130, 5, 66,
		0, 0, 130, 131, 5, 83, 0, 0, 131, 42, 1, 0, 0, 0, 132, 133, 5, 65, 0, 0,
		133, 134, 5, 86, 0, 0, 134, 135, 5, 71, 0, 0, 135, 44, 1, 0, 0, 0, 136,
		137, 5, 65, 0, 0, 137, 138, 5, 68, 0, 0, 138, 139, 5, 68, 0, 0, 139, 46,
		1, 0, 0, 0, 140, 141, 5, 68, 0, 0, 141, 142, 5, 73, 0, 0, 142, 143, 5,
		86, 0, 0, 143, 144, 5, 73, 0, 0, 144, 145, 5, 68, 0, 0, 145, 146, 5, 69,
		0, 0, 146, 48, 1, 0, 0, 0, 147, 148, 5, 77, 0, 0, 148, 149, 5, 85, 0, 0,
		149, 150, 5, 76, 0, 0, 150, 151, 5, 84, 0, 0, 151, 152, 5, 73, 0, 0, 152,
		153, 5, 80, 0, 0, 153, 154, 5, 76, 0, 0, 154, 155, 5, 89, 0, 0, 155, 50,
		1, 0, 0, 0, 156, 157, 5, 69, 0, 0, 157, 158, 5, 81, 0, 0, 158, 52, 1, 0,
		0, 0, 159, 160, 5, 67, 0, 0, 160, 161, 5, 79, 0, 0, 161, 162, 5, 78, 0,
		0, 162, 163, 5, 67, 0, 0, 163, 164, 5, 65, 0, 0, 164, 165, 5, 84, 0, 0,
		165, 54, 1, 0, 0, 0, 166, 167, 5, 71, 0, 0, 167, 168, 5, 84, 0, 0, 168,
		56, 1, 0, 0, 0, 169, 170, 5, 71, 0, 0, 170, 171, 5, 84, 0, 0, 171, 172,
		5, 69, 0, 0, 172, 58, 1, 0, 0, 0, 173, 175, 3, 3, 1, 0, 174, 173, 1, 0,
		0, 0, 174, 175, 1, 0, 0, 0, 175, 176, 1, 0, 0, 0, 176, 177, 3, 69, 34,
		0, 177, 60, 1, 0, 0, 0, 178, 180, 3, 3, 1, 0, 179, 178, 1, 0, 0, 0, 179,
		180, 1, 0, 0, 0, 180, 181, 1, 0, 0, 0, 181, 182, 3, 69, 34, 0, 182, 183,
		3, 29, 14, 0, 183, 184, 3, 69, 34, 0, 184, 62, 1, 0, 0, 0, 185, 186, 7,
		0, 0, 0, 186, 64, 1, 0, 0, 0, 187, 192, 5, 34, 0, 0, 188, 191, 3, 67, 33,
		0, 189, 191, 8, 1, 0, 0, 190, 188, 1, 0, 0, 0, 190, 189, 1, 0, 0, 0, 191,
		194, 1, 0, 0, 0, 192, 190, 1, 0, 0, 0, 192, 193, 1, 0, 0, 0, 193, 195,
		1, 0, 0, 0, 194, 192, 1, 0, 0, 0, 195, 196, 5, 34, 0, 0, 196, 66, 1, 0,
		0, 0, 197, 198, 5, 92, 0, 0, 198, 199, 7, 2, 0, 0, 199, 68, 1, 0, 0, 0,
		200, 202, 3, 63, 31, 0, 201, 200, 1, 0, 0, 0, 202, 203, 1, 0, 0, 0, 203,
		201, 1, 0, 0, 0, 203, 204, 1, 0, 0, 0, 204, 70, 1, 0, 0, 0, 205, 206, 5,
		73, 0, 0, 206, 207, 5, 70, 0, 0, 207, 72, 1, 0, 0, 0, 208, 210, 7, 3, 0,
		0, 209, 208, 1, 0, 0, 0, 210, 211, 1, 0, 0, 0, 211, 209, 1, 0, 0, 0, 211,
		212, 1, 0, 0, 0, 212, 74, 1, 0, 0, 0, 213, 215, 7, 4, 0, 0, 214, 213, 1,
		0, 0, 0, 215, 216, 1, 0, 0, 0, 216, 214, 1, 0, 0, 0, 216, 217, 1, 0, 0,
		0, 217, 218, 1, 0, 0, 0, 218, 219, 6, 37, 0, 0, 219, 76, 1, 0, 0, 0, 220,
		222, 7, 5, 0, 0, 221, 220, 1, 0, 0, 0, 222, 223, 1, 0, 0, 0, 223, 221,
		1, 0, 0, 0, 223, 224, 1, 0, 0, 0, 224, 239, 1, 0, 0, 0, 225, 239, 5, 59,
		0, 0, 226, 227, 5, 47, 0, 0, 227, 228, 5, 42, 0, 0, 228, 232, 1, 0, 0,
		0, 229, 231, 9, 0, 0, 0, 230, 229, 1, 0, 0, 0, 231, 234, 1, 0, 0, 0, 232,
		233, 1, 0, 0, 0, 232, 230, 1, 0, 0, 0, 233, 235, 1, 0, 0, 0, 234, 232,
		1, 0, 0, 0, 235, 236, 5, 42, 0, 0, 236, 239, 5, 47, 0, 0, 237, 239, 5,
		0, 0, 1, 238, 221, 1, 0, 0, 0, 238, 225, 1, 0, 0, 0, 238, 226, 1, 0, 0,
		0, 238, 237, 1, 0, 0, 0, 239, 240, 1, 0, 0, 0, 240, 241, 6, 38, 1, 0, 241,
		78, 1, 0, 0, 0, 11, 0, 174, 179, 190, 192, 203, 211, 216, 223, 232, 238,
		2, 6, 0, 0, 2, 0, 0,
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
	EqlLexerPLUS         = 1
	EqlLexerMINUS        = 2
	EqlLexerDIV          = 3
	EqlLexerMULT         = 4
	EqlLexerLPAREN       = 5
	EqlLexerRPAREN       = 6
	EqlLexerLCURLY       = 7
	EqlLexerRCURLY       = 8
	EqlLexerLBRACKET     = 9
	EqlLexerRBRACKET     = 10
	EqlLexerEQUAL        = 11
	EqlLexerSEMI         = 12
	EqlLexerCOLON        = 13
	EqlLexerCOMMA        = 14
	EqlLexerDOT          = 15
	EqlLexerGREATER_THAN = 16
	EqlLexerLESS_THAN    = 17
	EqlLexerTRUE         = 18
	EqlLexerFALSE        = 19
	EqlLexerSUM          = 20
	EqlLexerABS          = 21
	EqlLexerAVG          = 22
	EqlLexerADD          = 23
	EqlLexerDIVIDE       = 24
	EqlLexerMULTIPLY     = 25
	EqlLexerEQ           = 26
	EqlLexerCONCAT       = 27
	EqlLexerGT           = 28
	EqlLexerGTE          = 29
	EqlLexerINT          = 30
	EqlLexerDECIMAL      = 31
	EqlLexerDIGIT        = 32
	EqlLexerSTRING       = 33
	EqlLexerIF           = 34
	EqlLexerIDENTIFIER   = 35
	EqlLexerWS           = 36
	EqlLexerEOS          = 37
)
