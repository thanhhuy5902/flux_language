// Code generated from Primitives.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parsing

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

type Primitives struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var PrimitivesLexerStaticData struct {
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

func primitivesLexerInit() {
	staticData := &PrimitivesLexerStaticData
	staticData.ChannelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.ModeNames = []string{
		"DEFAULT_MODE",
	}
	staticData.LiteralNames = []string{
		"", "", "'text'", "", "'boolean'", "", "'num'", "", "", "", "'ipv4'",
		"", "", "", "", "'{'", "'}'", "'@'",
	}
	staticData.SymbolicNames = []string{
		"", "TEXT", "TEXT_TYPE", "BOOLEAN", "BOOLEAN_TYPE", "NUMBER", "NUMBER_TYPE",
		"DIGIT", "OCTET", "IPV4", "IPV4_TYPE", "VAR_IDENTIFIER", "COMMON_IDENTIFIER",
		"ARITHMETIC_OPERATOR", "LOGICAL_OPERATOR", "L_BLOCK", "R_BLOCK", "AT",
		"WS",
	}
	staticData.RuleNames = []string{
		"TEXT", "TEXT_TYPE", "BOOLEAN", "BOOLEAN_TYPE", "NUMBER", "NUMBER_TYPE",
		"DIGIT", "OCTET", "IPV4", "IPV4_TYPE", "VAR_IDENTIFIER", "COMMON_IDENTIFIER",
		"ARITHMETIC_OPERATOR", "LOGICAL_OPERATOR", "L_BLOCK", "R_BLOCK", "AT",
		"WS",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 18, 183, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 1, 0, 1, 0, 5, 0, 40, 8, 0, 10, 0, 12,
		0, 43, 9, 0, 1, 0, 1, 0, 1, 0, 5, 0, 48, 8, 0, 10, 0, 12, 0, 51, 9, 0,
		1, 0, 3, 0, 54, 8, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 1, 2, 1, 2, 1,
		2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 3, 2, 70, 8, 2, 1, 3, 1, 3, 1, 3, 1, 3,
		1, 3, 1, 3, 1, 3, 1, 3, 1, 4, 3, 4, 81, 8, 4, 1, 4, 4, 4, 84, 8, 4, 11,
		4, 12, 4, 85, 1, 4, 1, 4, 5, 4, 90, 8, 4, 10, 4, 12, 4, 93, 9, 4, 3, 4,
		95, 8, 4, 1, 5, 1, 5, 1, 5, 1, 5, 1, 6, 1, 6, 1, 7, 1, 7, 1, 7, 3, 7, 106,
		8, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7,
		3, 7, 119, 8, 7, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 9,
		1, 9, 1, 9, 1, 9, 1, 9, 1, 10, 1, 10, 5, 10, 136, 8, 10, 10, 10, 12, 10,
		139, 9, 10, 1, 11, 1, 11, 5, 11, 143, 8, 11, 10, 11, 12, 11, 146, 9, 11,
		1, 12, 1, 12, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1,
		13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13,
		3, 13, 169, 8, 13, 1, 14, 1, 14, 1, 15, 1, 15, 1, 16, 1, 16, 1, 17, 4,
		17, 178, 8, 17, 11, 17, 12, 17, 179, 1, 17, 1, 17, 1, 49, 0, 18, 1, 1,
		3, 2, 5, 3, 7, 4, 9, 5, 11, 6, 13, 7, 15, 8, 17, 9, 19, 10, 21, 11, 23,
		12, 25, 13, 27, 14, 29, 15, 31, 16, 33, 17, 35, 18, 1, 0, 11, 1, 0, 34,
		34, 1, 0, 48, 57, 1, 0, 49, 57, 1, 0, 48, 52, 1, 0, 48, 53, 1, 0, 97, 122,
		3, 0, 48, 57, 65, 90, 97, 122, 2, 0, 65, 90, 97, 122, 5, 0, 37, 37, 42,
		43, 45, 45, 47, 47, 94, 94, 2, 0, 60, 60, 62, 62, 3, 0, 9, 10, 13, 13,
		32, 32, 206, 0, 1, 1, 0, 0, 0, 0, 3, 1, 0, 0, 0, 0, 5, 1, 0, 0, 0, 0, 7,
		1, 0, 0, 0, 0, 9, 1, 0, 0, 0, 0, 11, 1, 0, 0, 0, 0, 13, 1, 0, 0, 0, 0,
		15, 1, 0, 0, 0, 0, 17, 1, 0, 0, 0, 0, 19, 1, 0, 0, 0, 0, 21, 1, 0, 0, 0,
		0, 23, 1, 0, 0, 0, 0, 25, 1, 0, 0, 0, 0, 27, 1, 0, 0, 0, 0, 29, 1, 0, 0,
		0, 0, 31, 1, 0, 0, 0, 0, 33, 1, 0, 0, 0, 0, 35, 1, 0, 0, 0, 1, 53, 1, 0,
		0, 0, 3, 55, 1, 0, 0, 0, 5, 69, 1, 0, 0, 0, 7, 71, 1, 0, 0, 0, 9, 80, 1,
		0, 0, 0, 11, 96, 1, 0, 0, 0, 13, 100, 1, 0, 0, 0, 15, 118, 1, 0, 0, 0,
		17, 120, 1, 0, 0, 0, 19, 128, 1, 0, 0, 0, 21, 133, 1, 0, 0, 0, 23, 140,
		1, 0, 0, 0, 25, 147, 1, 0, 0, 0, 27, 168, 1, 0, 0, 0, 29, 170, 1, 0, 0,
		0, 31, 172, 1, 0, 0, 0, 33, 174, 1, 0, 0, 0, 35, 177, 1, 0, 0, 0, 37, 41,
		5, 34, 0, 0, 38, 40, 8, 0, 0, 0, 39, 38, 1, 0, 0, 0, 40, 43, 1, 0, 0, 0,
		41, 39, 1, 0, 0, 0, 41, 42, 1, 0, 0, 0, 42, 44, 1, 0, 0, 0, 43, 41, 1,
		0, 0, 0, 44, 54, 5, 34, 0, 0, 45, 49, 5, 39, 0, 0, 46, 48, 9, 0, 0, 0,
		47, 46, 1, 0, 0, 0, 48, 51, 1, 0, 0, 0, 49, 50, 1, 0, 0, 0, 49, 47, 1,
		0, 0, 0, 50, 52, 1, 0, 0, 0, 51, 49, 1, 0, 0, 0, 52, 54, 5, 39, 0, 0, 53,
		37, 1, 0, 0, 0, 53, 45, 1, 0, 0, 0, 54, 2, 1, 0, 0, 0, 55, 56, 5, 116,
		0, 0, 56, 57, 5, 101, 0, 0, 57, 58, 5, 120, 0, 0, 58, 59, 5, 116, 0, 0,
		59, 4, 1, 0, 0, 0, 60, 61, 5, 116, 0, 0, 61, 62, 5, 114, 0, 0, 62, 63,
		5, 117, 0, 0, 63, 70, 5, 101, 0, 0, 64, 65, 5, 102, 0, 0, 65, 66, 5, 97,
		0, 0, 66, 67, 5, 108, 0, 0, 67, 68, 5, 115, 0, 0, 68, 70, 5, 101, 0, 0,
		69, 60, 1, 0, 0, 0, 69, 64, 1, 0, 0, 0, 70, 6, 1, 0, 0, 0, 71, 72, 5, 98,
		0, 0, 72, 73, 5, 111, 0, 0, 73, 74, 5, 111, 0, 0, 74, 75, 5, 108, 0, 0,
		75, 76, 5, 101, 0, 0, 76, 77, 5, 97, 0, 0, 77, 78, 5, 110, 0, 0, 78, 8,
		1, 0, 0, 0, 79, 81, 5, 45, 0, 0, 80, 79, 1, 0, 0, 0, 80, 81, 1, 0, 0, 0,
		81, 83, 1, 0, 0, 0, 82, 84, 7, 1, 0, 0, 83, 82, 1, 0, 0, 0, 84, 85, 1,
		0, 0, 0, 85, 83, 1, 0, 0, 0, 85, 86, 1, 0, 0, 0, 86, 94, 1, 0, 0, 0, 87,
		91, 5, 46, 0, 0, 88, 90, 7, 1, 0, 0, 89, 88, 1, 0, 0, 0, 90, 93, 1, 0,
		0, 0, 91, 89, 1, 0, 0, 0, 91, 92, 1, 0, 0, 0, 92, 95, 1, 0, 0, 0, 93, 91,
		1, 0, 0, 0, 94, 87, 1, 0, 0, 0, 94, 95, 1, 0, 0, 0, 95, 10, 1, 0, 0, 0,
		96, 97, 5, 110, 0, 0, 97, 98, 5, 117, 0, 0, 98, 99, 5, 109, 0, 0, 99, 12,
		1, 0, 0, 0, 100, 101, 7, 1, 0, 0, 101, 14, 1, 0, 0, 0, 102, 119, 5, 48,
		0, 0, 103, 105, 7, 2, 0, 0, 104, 106, 3, 13, 6, 0, 105, 104, 1, 0, 0, 0,
		105, 106, 1, 0, 0, 0, 106, 119, 1, 0, 0, 0, 107, 108, 5, 49, 0, 0, 108,
		109, 3, 13, 6, 0, 109, 110, 3, 13, 6, 0, 110, 119, 1, 0, 0, 0, 111, 112,
		5, 50, 0, 0, 112, 113, 7, 3, 0, 0, 113, 119, 3, 13, 6, 0, 114, 115, 5,
		50, 0, 0, 115, 116, 5, 53, 0, 0, 116, 117, 1, 0, 0, 0, 117, 119, 7, 4,
		0, 0, 118, 102, 1, 0, 0, 0, 118, 103, 1, 0, 0, 0, 118, 107, 1, 0, 0, 0,
		118, 111, 1, 0, 0, 0, 118, 114, 1, 0, 0, 0, 119, 16, 1, 0, 0, 0, 120, 121,
		3, 15, 7, 0, 121, 122, 5, 46, 0, 0, 122, 123, 3, 15, 7, 0, 123, 124, 5,
		46, 0, 0, 124, 125, 3, 15, 7, 0, 125, 126, 5, 46, 0, 0, 126, 127, 3, 15,
		7, 0, 127, 18, 1, 0, 0, 0, 128, 129, 5, 105, 0, 0, 129, 130, 5, 112, 0,
		0, 130, 131, 5, 118, 0, 0, 131, 132, 5, 52, 0, 0, 132, 20, 1, 0, 0, 0,
		133, 137, 7, 5, 0, 0, 134, 136, 7, 6, 0, 0, 135, 134, 1, 0, 0, 0, 136,
		139, 1, 0, 0, 0, 137, 135, 1, 0, 0, 0, 137, 138, 1, 0, 0, 0, 138, 22, 1,
		0, 0, 0, 139, 137, 1, 0, 0, 0, 140, 144, 7, 7, 0, 0, 141, 143, 7, 6, 0,
		0, 142, 141, 1, 0, 0, 0, 143, 146, 1, 0, 0, 0, 144, 142, 1, 0, 0, 0, 144,
		145, 1, 0, 0, 0, 145, 24, 1, 0, 0, 0, 146, 144, 1, 0, 0, 0, 147, 148, 7,
		8, 0, 0, 148, 26, 1, 0, 0, 0, 149, 150, 5, 120, 0, 0, 150, 151, 5, 111,
		0, 0, 151, 169, 5, 114, 0, 0, 152, 153, 5, 97, 0, 0, 153, 154, 5, 110,
		0, 0, 154, 169, 5, 100, 0, 0, 155, 156, 5, 111, 0, 0, 156, 169, 5, 114,
		0, 0, 157, 158, 5, 110, 0, 0, 158, 159, 5, 111, 0, 0, 159, 169, 5, 116,
		0, 0, 160, 169, 5, 61, 0, 0, 161, 162, 5, 33, 0, 0, 162, 169, 5, 61, 0,
		0, 163, 169, 7, 9, 0, 0, 164, 165, 5, 60, 0, 0, 165, 169, 5, 61, 0, 0,
		166, 167, 5, 62, 0, 0, 167, 169, 5, 61, 0, 0, 168, 149, 1, 0, 0, 0, 168,
		152, 1, 0, 0, 0, 168, 155, 1, 0, 0, 0, 168, 157, 1, 0, 0, 0, 168, 160,
		1, 0, 0, 0, 168, 161, 1, 0, 0, 0, 168, 163, 1, 0, 0, 0, 168, 164, 1, 0,
		0, 0, 168, 166, 1, 0, 0, 0, 169, 28, 1, 0, 0, 0, 170, 171, 5, 123, 0, 0,
		171, 30, 1, 0, 0, 0, 172, 173, 5, 125, 0, 0, 173, 32, 1, 0, 0, 0, 174,
		175, 5, 64, 0, 0, 175, 34, 1, 0, 0, 0, 176, 178, 7, 10, 0, 0, 177, 176,
		1, 0, 0, 0, 178, 179, 1, 0, 0, 0, 179, 177, 1, 0, 0, 0, 179, 180, 1, 0,
		0, 0, 180, 181, 1, 0, 0, 0, 181, 182, 6, 17, 0, 0, 182, 36, 1, 0, 0, 0,
		15, 0, 41, 49, 53, 69, 80, 85, 91, 94, 105, 118, 137, 144, 168, 179, 1,
		6, 0, 0,
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

// PrimitivesInit initializes any static state used to implement Primitives. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewPrimitives(). You can call this function if you wish to initialize the static state ahead
// of time.
func PrimitivesInit() {
	staticData := &PrimitivesLexerStaticData
	staticData.once.Do(primitivesLexerInit)
}

// NewPrimitives produces a new lexer instance for the optional input antlr.CharStream.
func NewPrimitives(input antlr.CharStream) *Primitives {
	PrimitivesInit()
	l := new(Primitives)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &PrimitivesLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	l.channelNames = staticData.ChannelNames
	l.modeNames = staticData.ModeNames
	l.RuleNames = staticData.RuleNames
	l.LiteralNames = staticData.LiteralNames
	l.SymbolicNames = staticData.SymbolicNames
	l.GrammarFileName = "Primitives.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// Primitives tokens.
const (
	PrimitivesTEXT                = 1
	PrimitivesTEXT_TYPE           = 2
	PrimitivesBOOLEAN             = 3
	PrimitivesBOOLEAN_TYPE        = 4
	PrimitivesNUMBER              = 5
	PrimitivesNUMBER_TYPE         = 6
	PrimitivesDIGIT               = 7
	PrimitivesOCTET               = 8
	PrimitivesIPV4                = 9
	PrimitivesIPV4_TYPE           = 10
	PrimitivesVAR_IDENTIFIER      = 11
	PrimitivesCOMMON_IDENTIFIER   = 12
	PrimitivesARITHMETIC_OPERATOR = 13
	PrimitivesLOGICAL_OPERATOR    = 14
	PrimitivesL_BLOCK             = 15
	PrimitivesR_BLOCK             = 16
	PrimitivesAT                  = 17
	PrimitivesWS                  = 18
)
