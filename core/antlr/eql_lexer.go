// Code generated from core/antlr/EqlLexer.g4 by ANTLR 4.13.0. DO NOT EDIT.

package antlr

import (
	"fmt"
	"github.com/antlr4-go/antlr/v4"
	"sync"
	"unicode"
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

var EqlLexerLexerStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	ChannelNames           []string
	ModeNames              []string
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func eqllexerLexerInit() {
	staticData := &EqlLexerLexerStaticData
	staticData.ChannelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.ModeNames = []string{
		"DEFAULT_MODE",
	}
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
		"PLUS", "MINUS", "DIV", "MULT", "LPAREN", "RPAREN", "LCURLY", "RCURLY",
		"LBRACKET", "RBRACKET", "EQUAL", "SEMI", "COLON", "COMMA", "DOT", "GREATER_THAN",
		"LESS_THAN", "TRUE", "FALSE", "SUM", "ABS", "AVG", "ADD", "DIVIDE",
		"MULTIPLY", "EQ", "CONCAT", "GT", "GTE", "LAMBDA", "INT", "DECIMAL",
		"DIGIT", "STRING", "ESC_SEQ", "DIGITS", "IF", "IDENTIFIER", "WS", "EOS",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 38, 251, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7,
		20, 2, 21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25,
		2, 26, 7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2,
		31, 7, 31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 2, 35, 7, 35, 2, 36,
		7, 36, 2, 37, 7, 37, 2, 38, 7, 38, 2, 39, 7, 39, 1, 0, 1, 0, 1, 1, 1, 1,
		1, 2, 1, 2, 1, 3, 1, 3, 1, 4, 1, 4, 1, 5, 1, 5, 1, 6, 1, 6, 1, 7, 1, 7,
		1, 8, 1, 8, 1, 9, 1, 9, 1, 10, 1, 10, 1, 11, 1, 11, 1, 12, 1, 12, 1, 13,
		1, 13, 1, 14, 1, 14, 1, 15, 1, 15, 1, 16, 1, 16, 1, 17, 1, 17, 1, 17, 1,
		17, 1, 17, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 19, 1, 19, 1, 19,
		1, 19, 1, 20, 1, 20, 1, 20, 1, 20, 1, 21, 1, 21, 1, 21, 1, 21, 1, 22, 1,
		22, 1, 22, 1, 22, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 24,
		1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 25, 1, 25, 1,
		25, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 27, 1, 27, 1, 27,
		1, 28, 1, 28, 1, 28, 1, 28, 1, 29, 1, 29, 1, 29, 1, 29, 1, 29, 1, 29, 1,
		29, 1, 30, 3, 30, 184, 8, 30, 1, 30, 1, 30, 1, 31, 3, 31, 189, 8, 31, 1,
		31, 1, 31, 1, 31, 1, 31, 1, 32, 1, 32, 1, 33, 1, 33, 1, 33, 5, 33, 200,
		8, 33, 10, 33, 12, 33, 203, 9, 33, 1, 33, 1, 33, 1, 34, 1, 34, 1, 34, 1,
		35, 4, 35, 211, 8, 35, 11, 35, 12, 35, 212, 1, 36, 1, 36, 1, 36, 1, 37,
		4, 37, 219, 8, 37, 11, 37, 12, 37, 220, 1, 38, 4, 38, 224, 8, 38, 11, 38,
		12, 38, 225, 1, 38, 1, 38, 1, 39, 4, 39, 231, 8, 39, 11, 39, 12, 39, 232,
		1, 39, 1, 39, 1, 39, 1, 39, 1, 39, 5, 39, 240, 8, 39, 10, 39, 12, 39, 243,
		9, 39, 1, 39, 1, 39, 1, 39, 3, 39, 248, 8, 39, 1, 39, 1, 39, 1, 241, 0,
		40, 1, 1, 3, 2, 5, 3, 7, 4, 9, 5, 11, 6, 13, 7, 15, 8, 17, 9, 19, 10, 21,
		11, 23, 12, 25, 13, 27, 14, 29, 15, 31, 16, 33, 17, 35, 18, 37, 19, 39,
		20, 41, 21, 43, 22, 45, 23, 47, 24, 49, 25, 51, 26, 53, 27, 55, 28, 57,
		29, 59, 30, 61, 31, 63, 32, 65, 33, 67, 34, 69, 0, 71, 0, 73, 35, 75, 36,
		77, 37, 79, 38, 1, 0, 6, 1, 0, 48, 57, 2, 0, 34, 34, 92, 92, 8, 0, 34,
		34, 39, 39, 92, 92, 98, 98, 102, 102, 110, 110, 114, 114, 116, 116, 2,
		0, 65, 90, 97, 122, 3, 0, 9, 10, 13, 13, 32, 32, 2, 0, 10, 10, 13, 13,
		260, 0, 1, 1, 0, 0, 0, 0, 3, 1, 0, 0, 0, 0, 5, 1, 0, 0, 0, 0, 7, 1, 0,
		0, 0, 0, 9, 1, 0, 0, 0, 0, 11, 1, 0, 0, 0, 0, 13, 1, 0, 0, 0, 0, 15, 1,
		0, 0, 0, 0, 17, 1, 0, 0, 0, 0, 19, 1, 0, 0, 0, 0, 21, 1, 0, 0, 0, 0, 23,
		1, 0, 0, 0, 0, 25, 1, 0, 0, 0, 0, 27, 1, 0, 0, 0, 0, 29, 1, 0, 0, 0, 0,
		31, 1, 0, 0, 0, 0, 33, 1, 0, 0, 0, 0, 35, 1, 0, 0, 0, 0, 37, 1, 0, 0, 0,
		0, 39, 1, 0, 0, 0, 0, 41, 1, 0, 0, 0, 0, 43, 1, 0, 0, 0, 0, 45, 1, 0, 0,
		0, 0, 47, 1, 0, 0, 0, 0, 49, 1, 0, 0, 0, 0, 51, 1, 0, 0, 0, 0, 53, 1, 0,
		0, 0, 0, 55, 1, 0, 0, 0, 0, 57, 1, 0, 0, 0, 0, 59, 1, 0, 0, 0, 0, 61, 1,
		0, 0, 0, 0, 63, 1, 0, 0, 0, 0, 65, 1, 0, 0, 0, 0, 67, 1, 0, 0, 0, 0, 73,
		1, 0, 0, 0, 0, 75, 1, 0, 0, 0, 0, 77, 1, 0, 0, 0, 0, 79, 1, 0, 0, 0, 1,
		81, 1, 0, 0, 0, 3, 83, 1, 0, 0, 0, 5, 85, 1, 0, 0, 0, 7, 87, 1, 0, 0, 0,
		9, 89, 1, 0, 0, 0, 11, 91, 1, 0, 0, 0, 13, 93, 1, 0, 0, 0, 15, 95, 1, 0,
		0, 0, 17, 97, 1, 0, 0, 0, 19, 99, 1, 0, 0, 0, 21, 101, 1, 0, 0, 0, 23,
		103, 1, 0, 0, 0, 25, 105, 1, 0, 0, 0, 27, 107, 1, 0, 0, 0, 29, 109, 1,
		0, 0, 0, 31, 111, 1, 0, 0, 0, 33, 113, 1, 0, 0, 0, 35, 115, 1, 0, 0, 0,
		37, 120, 1, 0, 0, 0, 39, 126, 1, 0, 0, 0, 41, 130, 1, 0, 0, 0, 43, 134,
		1, 0, 0, 0, 45, 138, 1, 0, 0, 0, 47, 142, 1, 0, 0, 0, 49, 149, 1, 0, 0,
		0, 51, 158, 1, 0, 0, 0, 53, 161, 1, 0, 0, 0, 55, 168, 1, 0, 0, 0, 57, 171,
		1, 0, 0, 0, 59, 175, 1, 0, 0, 0, 61, 183, 1, 0, 0, 0, 63, 188, 1, 0, 0,
		0, 65, 194, 1, 0, 0, 0, 67, 196, 1, 0, 0, 0, 69, 206, 1, 0, 0, 0, 71, 210,
		1, 0, 0, 0, 73, 214, 1, 0, 0, 0, 75, 218, 1, 0, 0, 0, 77, 223, 1, 0, 0,
		0, 79, 247, 1, 0, 0, 0, 81, 82, 5, 43, 0, 0, 82, 2, 1, 0, 0, 0, 83, 84,
		5, 45, 0, 0, 84, 4, 1, 0, 0, 0, 85, 86, 5, 47, 0, 0, 86, 6, 1, 0, 0, 0,
		87, 88, 5, 42, 0, 0, 88, 8, 1, 0, 0, 0, 89, 90, 5, 40, 0, 0, 90, 10, 1,
		0, 0, 0, 91, 92, 5, 41, 0, 0, 92, 12, 1, 0, 0, 0, 93, 94, 5, 123, 0, 0,
		94, 14, 1, 0, 0, 0, 95, 96, 5, 125, 0, 0, 96, 16, 1, 0, 0, 0, 97, 98, 5,
		91, 0, 0, 98, 18, 1, 0, 0, 0, 99, 100, 5, 93, 0, 0, 100, 20, 1, 0, 0, 0,
		101, 102, 5, 61, 0, 0, 102, 22, 1, 0, 0, 0, 103, 104, 5, 59, 0, 0, 104,
		24, 1, 0, 0, 0, 105, 106, 5, 58, 0, 0, 106, 26, 1, 0, 0, 0, 107, 108, 5,
		44, 0, 0, 108, 28, 1, 0, 0, 0, 109, 110, 5, 46, 0, 0, 110, 30, 1, 0, 0,
		0, 111, 112, 5, 62, 0, 0, 112, 32, 1, 0, 0, 0, 113, 114, 5, 60, 0, 0, 114,
		34, 1, 0, 0, 0, 115, 116, 5, 84, 0, 0, 116, 117, 5, 82, 0, 0, 117, 118,
		5, 85, 0, 0, 118, 119, 5, 69, 0, 0, 119, 36, 1, 0, 0, 0, 120, 121, 5, 70,
		0, 0, 121, 122, 5, 65, 0, 0, 122, 123, 5, 76, 0, 0, 123, 124, 5, 83, 0,
		0, 124, 125, 5, 69, 0, 0, 125, 38, 1, 0, 0, 0, 126, 127, 5, 83, 0, 0, 127,
		128, 5, 85, 0, 0, 128, 129, 5, 77, 0, 0, 129, 40, 1, 0, 0, 0, 130, 131,
		5, 65, 0, 0, 131, 132, 5, 66, 0, 0, 132, 133, 5, 83, 0, 0, 133, 42, 1,
		0, 0, 0, 134, 135, 5, 65, 0, 0, 135, 136, 5, 86, 0, 0, 136, 137, 5, 71,
		0, 0, 137, 44, 1, 0, 0, 0, 138, 139, 5, 65, 0, 0, 139, 140, 5, 68, 0, 0,
		140, 141, 5, 68, 0, 0, 141, 46, 1, 0, 0, 0, 142, 143, 5, 68, 0, 0, 143,
		144, 5, 73, 0, 0, 144, 145, 5, 86, 0, 0, 145, 146, 5, 73, 0, 0, 146, 147,
		5, 68, 0, 0, 147, 148, 5, 69, 0, 0, 148, 48, 1, 0, 0, 0, 149, 150, 5, 77,
		0, 0, 150, 151, 5, 85, 0, 0, 151, 152, 5, 76, 0, 0, 152, 153, 5, 84, 0,
		0, 153, 154, 5, 73, 0, 0, 154, 155, 5, 80, 0, 0, 155, 156, 5, 76, 0, 0,
		156, 157, 5, 89, 0, 0, 157, 50, 1, 0, 0, 0, 158, 159, 5, 69, 0, 0, 159,
		160, 5, 81, 0, 0, 160, 52, 1, 0, 0, 0, 161, 162, 5, 67, 0, 0, 162, 163,
		5, 79, 0, 0, 163, 164, 5, 78, 0, 0, 164, 165, 5, 67, 0, 0, 165, 166, 5,
		65, 0, 0, 166, 167, 5, 84, 0, 0, 167, 54, 1, 0, 0, 0, 168, 169, 5, 71,
		0, 0, 169, 170, 5, 84, 0, 0, 170, 56, 1, 0, 0, 0, 171, 172, 5, 71, 0, 0,
		172, 173, 5, 84, 0, 0, 173, 174, 5, 69, 0, 0, 174, 58, 1, 0, 0, 0, 175,
		176, 5, 76, 0, 0, 176, 177, 5, 65, 0, 0, 177, 178, 5, 77, 0, 0, 178, 179,
		5, 66, 0, 0, 179, 180, 5, 68, 0, 0, 180, 181, 5, 65, 0, 0, 181, 60, 1,
		0, 0, 0, 182, 184, 3, 3, 1, 0, 183, 182, 1, 0, 0, 0, 183, 184, 1, 0, 0,
		0, 184, 185, 1, 0, 0, 0, 185, 186, 3, 71, 35, 0, 186, 62, 1, 0, 0, 0, 187,
		189, 3, 3, 1, 0, 188, 187, 1, 0, 0, 0, 188, 189, 1, 0, 0, 0, 189, 190,
		1, 0, 0, 0, 190, 191, 3, 71, 35, 0, 191, 192, 3, 29, 14, 0, 192, 193, 3,
		71, 35, 0, 193, 64, 1, 0, 0, 0, 194, 195, 7, 0, 0, 0, 195, 66, 1, 0, 0,
		0, 196, 201, 5, 34, 0, 0, 197, 200, 3, 69, 34, 0, 198, 200, 8, 1, 0, 0,
		199, 197, 1, 0, 0, 0, 199, 198, 1, 0, 0, 0, 200, 203, 1, 0, 0, 0, 201,
		199, 1, 0, 0, 0, 201, 202, 1, 0, 0, 0, 202, 204, 1, 0, 0, 0, 203, 201,
		1, 0, 0, 0, 204, 205, 5, 34, 0, 0, 205, 68, 1, 0, 0, 0, 206, 207, 5, 92,
		0, 0, 207, 208, 7, 2, 0, 0, 208, 70, 1, 0, 0, 0, 209, 211, 3, 65, 32, 0,
		210, 209, 1, 0, 0, 0, 211, 212, 1, 0, 0, 0, 212, 210, 1, 0, 0, 0, 212,
		213, 1, 0, 0, 0, 213, 72, 1, 0, 0, 0, 214, 215, 5, 73, 0, 0, 215, 216,
		5, 70, 0, 0, 216, 74, 1, 0, 0, 0, 217, 219, 7, 3, 0, 0, 218, 217, 1, 0,
		0, 0, 219, 220, 1, 0, 0, 0, 220, 218, 1, 0, 0, 0, 220, 221, 1, 0, 0, 0,
		221, 76, 1, 0, 0, 0, 222, 224, 7, 4, 0, 0, 223, 222, 1, 0, 0, 0, 224, 225,
		1, 0, 0, 0, 225, 223, 1, 0, 0, 0, 225, 226, 1, 0, 0, 0, 226, 227, 1, 0,
		0, 0, 227, 228, 6, 38, 0, 0, 228, 78, 1, 0, 0, 0, 229, 231, 7, 5, 0, 0,
		230, 229, 1, 0, 0, 0, 231, 232, 1, 0, 0, 0, 232, 230, 1, 0, 0, 0, 232,
		233, 1, 0, 0, 0, 233, 248, 1, 0, 0, 0, 234, 248, 5, 59, 0, 0, 235, 236,
		5, 47, 0, 0, 236, 237, 5, 42, 0, 0, 237, 241, 1, 0, 0, 0, 238, 240, 9,
		0, 0, 0, 239, 238, 1, 0, 0, 0, 240, 243, 1, 0, 0, 0, 241, 242, 1, 0, 0,
		0, 241, 239, 1, 0, 0, 0, 242, 244, 1, 0, 0, 0, 243, 241, 1, 0, 0, 0, 244,
		245, 5, 42, 0, 0, 245, 248, 5, 47, 0, 0, 246, 248, 5, 0, 0, 1, 247, 230,
		1, 0, 0, 0, 247, 234, 1, 0, 0, 0, 247, 235, 1, 0, 0, 0, 247, 246, 1, 0,
		0, 0, 248, 249, 1, 0, 0, 0, 249, 250, 6, 39, 1, 0, 250, 80, 1, 0, 0, 0,
		11, 0, 183, 188, 199, 201, 212, 220, 225, 232, 241, 247, 2, 6, 0, 0, 2,
		0, 0,
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
	staticData := &EqlLexerLexerStaticData
	staticData.once.Do(eqllexerLexerInit)
}

// NewEqlLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewEqlLexer(input antlr.CharStream) *EqlLexer {
	EqlLexerInit()
	l := new(EqlLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &EqlLexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	l.channelNames = staticData.ChannelNames
	l.modeNames = staticData.ModeNames
	l.RuleNames = staticData.RuleNames
	l.LiteralNames = staticData.LiteralNames
	l.SymbolicNames = staticData.SymbolicNames
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
	EqlLexerLAMBDA       = 30
	EqlLexerINT          = 31
	EqlLexerDECIMAL      = 32
	EqlLexerDIGIT        = 33
	EqlLexerSTRING       = 34
	EqlLexerIF           = 35
	EqlLexerIDENTIFIER   = 36
	EqlLexerWS           = 37
	EqlLexerEOS          = 38
)
