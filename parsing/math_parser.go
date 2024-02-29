// Code generated from Math.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parsing // Math
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

type Math struct {
	*antlr.BaseParser
}

var MathParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func mathParserInit() {
	staticData := &MathParserStaticData
	staticData.LiteralNames = []string{
		"", "", "'text'", "", "'boolean'", "", "", "", "", "'num'", "", "",
		"'ipv4'", "", "", "", "", "'{'", "'}'", "'@'",
	}
	staticData.SymbolicNames = []string{
		"", "TEXT", "TEXT_TYPE", "BOOLEAN", "BOOLEAN_TYPE", "DIGIT", "NUMBER",
		"UNIT", "SI_POSTFIX", "NUMBER_TYPE", "OCTET", "IPV4", "IPV4_TYPE", "VAR_IDENTIFIER",
		"COMMON_IDENTIFIER", "ARITHMETIC_OPERATOR", "LOGICAL_OPERATOR", "L_BLOCK",
		"R_BLOCK", "AT", "WS", "OP_MULTIPLY", "OP_DIVIDE", "OP_MOD", "OP_POWER",
		"OP_PLUS", "OP_MINUS", "OP_AND", "OP_OR", "OP_XOR", "OP_EQUAL", "OP_NOT_EQUAL",
		"OP_LESS", "OP_LESS_EQUAL", "OP_GREATER", "OP_GREATER_EQUAL", "OP_NOT",
		"L_PAREN", "R_PAREN", "L_SQUARE", "R_SQUARE", "DOT", "COMMA",
	}
	staticData.RuleNames = []string{
		"op_level1", "op_level2", "op_level3", "op_level4", "op_level5", "numberic_expression",
		"logical_expression", "comparative_expression", "get_variable", "math_expression",
		"get_array_element", "get_child", "function_call",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 42, 158, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 1, 0, 1, 0, 1, 1, 1, 1, 1, 2, 1, 2, 1,
		3, 1, 3, 1, 4, 1, 4, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1,
		5, 1, 5, 3, 5, 47, 8, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5,
		5, 5, 57, 8, 5, 10, 5, 12, 5, 60, 9, 5, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1,
		6, 1, 6, 1, 6, 1, 6, 1, 6, 3, 6, 72, 8, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6,
		1, 6, 1, 6, 1, 6, 5, 6, 82, 8, 6, 10, 6, 12, 6, 85, 9, 6, 1, 7, 1, 7, 1,
		7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 3,
		7, 101, 8, 7, 1, 7, 1, 7, 1, 7, 1, 7, 5, 7, 107, 8, 7, 10, 7, 12, 7, 110,
		9, 7, 1, 8, 1, 8, 1, 8, 3, 8, 115, 8, 8, 1, 9, 1, 9, 1, 9, 3, 9, 120, 8,
		9, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11,
		1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 3,
		11, 142, 8, 11, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 5, 12, 149, 8, 12, 10,
		12, 12, 12, 152, 9, 12, 3, 12, 154, 8, 12, 1, 12, 1, 12, 1, 12, 0, 3, 10,
		12, 14, 13, 0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 0, 4, 1, 0,
		21, 24, 1, 0, 25, 26, 1, 0, 27, 29, 1, 0, 30, 35, 170, 0, 26, 1, 0, 0,
		0, 2, 28, 1, 0, 0, 0, 4, 30, 1, 0, 0, 0, 6, 32, 1, 0, 0, 0, 8, 34, 1, 0,
		0, 0, 10, 46, 1, 0, 0, 0, 12, 71, 1, 0, 0, 0, 14, 100, 1, 0, 0, 0, 16,
		114, 1, 0, 0, 0, 18, 119, 1, 0, 0, 0, 20, 121, 1, 0, 0, 0, 22, 141, 1,
		0, 0, 0, 24, 143, 1, 0, 0, 0, 26, 27, 7, 0, 0, 0, 27, 1, 1, 0, 0, 0, 28,
		29, 7, 1, 0, 0, 29, 3, 1, 0, 0, 0, 30, 31, 7, 2, 0, 0, 31, 5, 1, 0, 0,
		0, 32, 33, 7, 3, 0, 0, 33, 7, 1, 0, 0, 0, 34, 35, 5, 36, 0, 0, 35, 9, 1,
		0, 0, 0, 36, 37, 6, 5, -1, 0, 37, 38, 5, 37, 0, 0, 38, 39, 3, 10, 5, 0,
		39, 40, 5, 38, 0, 0, 40, 47, 1, 0, 0, 0, 41, 47, 5, 6, 0, 0, 42, 47, 3,
		16, 8, 0, 43, 47, 3, 24, 12, 0, 44, 45, 5, 26, 0, 0, 45, 47, 3, 10, 5,
		1, 46, 36, 1, 0, 0, 0, 46, 41, 1, 0, 0, 0, 46, 42, 1, 0, 0, 0, 46, 43,
		1, 0, 0, 0, 46, 44, 1, 0, 0, 0, 47, 58, 1, 0, 0, 0, 48, 49, 10, 6, 0, 0,
		49, 50, 3, 0, 0, 0, 50, 51, 3, 10, 5, 7, 51, 57, 1, 0, 0, 0, 52, 53, 10,
		5, 0, 0, 53, 54, 3, 2, 1, 0, 54, 55, 3, 10, 5, 6, 55, 57, 1, 0, 0, 0, 56,
		48, 1, 0, 0, 0, 56, 52, 1, 0, 0, 0, 57, 60, 1, 0, 0, 0, 58, 56, 1, 0, 0,
		0, 58, 59, 1, 0, 0, 0, 59, 11, 1, 0, 0, 0, 60, 58, 1, 0, 0, 0, 61, 62,
		6, 6, -1, 0, 62, 63, 5, 37, 0, 0, 63, 64, 3, 12, 6, 0, 64, 65, 5, 38, 0,
		0, 65, 72, 1, 0, 0, 0, 66, 67, 5, 36, 0, 0, 67, 72, 3, 12, 6, 4, 68, 72,
		5, 3, 0, 0, 69, 72, 3, 16, 8, 0, 70, 72, 3, 24, 12, 0, 71, 61, 1, 0, 0,
		0, 71, 66, 1, 0, 0, 0, 71, 68, 1, 0, 0, 0, 71, 69, 1, 0, 0, 0, 71, 70,
		1, 0, 0, 0, 72, 83, 1, 0, 0, 0, 73, 74, 10, 6, 0, 0, 74, 75, 3, 4, 2, 0,
		75, 76, 3, 12, 6, 7, 76, 82, 1, 0, 0, 0, 77, 78, 10, 5, 0, 0, 78, 79, 3,
		6, 3, 0, 79, 80, 3, 12, 6, 6, 80, 82, 1, 0, 0, 0, 81, 73, 1, 0, 0, 0, 81,
		77, 1, 0, 0, 0, 82, 85, 1, 0, 0, 0, 83, 81, 1, 0, 0, 0, 83, 84, 1, 0, 0,
		0, 84, 13, 1, 0, 0, 0, 85, 83, 1, 0, 0, 0, 86, 87, 6, 7, -1, 0, 87, 88,
		3, 10, 5, 0, 88, 89, 3, 6, 3, 0, 89, 90, 3, 10, 5, 0, 90, 101, 1, 0, 0,
		0, 91, 92, 3, 12, 6, 0, 92, 93, 3, 4, 2, 0, 93, 94, 3, 12, 6, 0, 94, 101,
		1, 0, 0, 0, 95, 96, 3, 8, 4, 0, 96, 97, 5, 37, 0, 0, 97, 98, 3, 14, 7,
		0, 98, 99, 5, 38, 0, 0, 99, 101, 1, 0, 0, 0, 100, 86, 1, 0, 0, 0, 100,
		91, 1, 0, 0, 0, 100, 95, 1, 0, 0, 0, 101, 108, 1, 0, 0, 0, 102, 103, 10,
		2, 0, 0, 103, 104, 3, 4, 2, 0, 104, 105, 3, 12, 6, 0, 105, 107, 1, 0, 0,
		0, 106, 102, 1, 0, 0, 0, 107, 110, 1, 0, 0, 0, 108, 106, 1, 0, 0, 0, 108,
		109, 1, 0, 0, 0, 109, 15, 1, 0, 0, 0, 110, 108, 1, 0, 0, 0, 111, 115, 5,
		13, 0, 0, 112, 115, 3, 20, 10, 0, 113, 115, 3, 22, 11, 0, 114, 111, 1,
		0, 0, 0, 114, 112, 1, 0, 0, 0, 114, 113, 1, 0, 0, 0, 115, 17, 1, 0, 0,
		0, 116, 120, 3, 16, 8, 0, 117, 120, 3, 10, 5, 0, 118, 120, 3, 12, 6, 0,
		119, 116, 1, 0, 0, 0, 119, 117, 1, 0, 0, 0, 119, 118, 1, 0, 0, 0, 120,
		19, 1, 0, 0, 0, 121, 122, 5, 13, 0, 0, 122, 123, 5, 39, 0, 0, 123, 124,
		3, 10, 5, 0, 124, 125, 5, 40, 0, 0, 125, 21, 1, 0, 0, 0, 126, 127, 5, 13,
		0, 0, 127, 128, 5, 41, 0, 0, 128, 142, 5, 13, 0, 0, 129, 130, 5, 13, 0,
		0, 130, 131, 5, 41, 0, 0, 131, 142, 3, 20, 10, 0, 132, 133, 5, 13, 0, 0,
		133, 134, 5, 41, 0, 0, 134, 142, 3, 22, 11, 0, 135, 136, 3, 20, 10, 0,
		136, 137, 5, 41, 0, 0, 137, 138, 3, 22, 11, 0, 138, 142, 1, 0, 0, 0, 139,
		142, 3, 20, 10, 0, 140, 142, 5, 13, 0, 0, 141, 126, 1, 0, 0, 0, 141, 129,
		1, 0, 0, 0, 141, 132, 1, 0, 0, 0, 141, 135, 1, 0, 0, 0, 141, 139, 1, 0,
		0, 0, 141, 140, 1, 0, 0, 0, 142, 23, 1, 0, 0, 0, 143, 144, 5, 13, 0, 0,
		144, 153, 5, 37, 0, 0, 145, 150, 3, 18, 9, 0, 146, 147, 5, 42, 0, 0, 147,
		149, 3, 18, 9, 0, 148, 146, 1, 0, 0, 0, 149, 152, 1, 0, 0, 0, 150, 148,
		1, 0, 0, 0, 150, 151, 1, 0, 0, 0, 151, 154, 1, 0, 0, 0, 152, 150, 1, 0,
		0, 0, 153, 145, 1, 0, 0, 0, 153, 154, 1, 0, 0, 0, 154, 155, 1, 0, 0, 0,
		155, 156, 5, 38, 0, 0, 156, 25, 1, 0, 0, 0, 13, 46, 56, 58, 71, 81, 83,
		100, 108, 114, 119, 141, 150, 153,
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

// MathInit initializes any static state used to implement Math. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewMath(). You can call this function if you wish to initialize the static state ahead
// of time.
func MathInit() {
	staticData := &MathParserStaticData
	staticData.once.Do(mathParserInit)
}

// NewMath produces a new parser instance for the optional input antlr.TokenStream.
func NewMath(input antlr.TokenStream) *Math {
	MathInit()
	this := new(Math)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &MathParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "Math.g4"

	return this
}

// Math tokens.
const (
	MathEOF                 = antlr.TokenEOF
	MathTEXT                = 1
	MathTEXT_TYPE           = 2
	MathBOOLEAN             = 3
	MathBOOLEAN_TYPE        = 4
	MathDIGIT               = 5
	MathNUMBER              = 6
	MathUNIT                = 7
	MathSI_POSTFIX          = 8
	MathNUMBER_TYPE         = 9
	MathOCTET               = 10
	MathIPV4                = 11
	MathIPV4_TYPE           = 12
	MathVAR_IDENTIFIER      = 13
	MathCOMMON_IDENTIFIER   = 14
	MathARITHMETIC_OPERATOR = 15
	MathLOGICAL_OPERATOR    = 16
	MathL_BLOCK             = 17
	MathR_BLOCK             = 18
	MathAT                  = 19
	MathWS                  = 20
	MathOP_MULTIPLY         = 21
	MathOP_DIVIDE           = 22
	MathOP_MOD              = 23
	MathOP_POWER            = 24
	MathOP_PLUS             = 25
	MathOP_MINUS            = 26
	MathOP_AND              = 27
	MathOP_OR               = 28
	MathOP_XOR              = 29
	MathOP_EQUAL            = 30
	MathOP_NOT_EQUAL        = 31
	MathOP_LESS             = 32
	MathOP_LESS_EQUAL       = 33
	MathOP_GREATER          = 34
	MathOP_GREATER_EQUAL    = 35
	MathOP_NOT              = 36
	MathL_PAREN             = 37
	MathR_PAREN             = 38
	MathL_SQUARE            = 39
	MathR_SQUARE            = 40
	MathDOT                 = 41
	MathCOMMA               = 42
)

// Math rules.
const (
	MathRULE_op_level1              = 0
	MathRULE_op_level2              = 1
	MathRULE_op_level3              = 2
	MathRULE_op_level4              = 3
	MathRULE_op_level5              = 4
	MathRULE_numberic_expression    = 5
	MathRULE_logical_expression     = 6
	MathRULE_comparative_expression = 7
	MathRULE_get_variable           = 8
	MathRULE_math_expression        = 9
	MathRULE_get_array_element      = 10
	MathRULE_get_child              = 11
	MathRULE_function_call          = 12
)

// IOp_level1Context is an interface to support dynamic dispatch.
type IOp_level1Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	OP_MULTIPLY() antlr.TerminalNode
	OP_DIVIDE() antlr.TerminalNode
	OP_MOD() antlr.TerminalNode
	OP_POWER() antlr.TerminalNode

	// IsOp_level1Context differentiates from other interfaces.
	IsOp_level1Context()
}

type Op_level1Context struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOp_level1Context() *Op_level1Context {
	var p = new(Op_level1Context)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MathRULE_op_level1
	return p
}

func InitEmptyOp_level1Context(p *Op_level1Context) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MathRULE_op_level1
}

func (*Op_level1Context) IsOp_level1Context() {}

func NewOp_level1Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Op_level1Context {
	var p = new(Op_level1Context)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MathRULE_op_level1

	return p
}

func (s *Op_level1Context) GetParser() antlr.Parser { return s.parser }

func (s *Op_level1Context) OP_MULTIPLY() antlr.TerminalNode {
	return s.GetToken(MathOP_MULTIPLY, 0)
}

func (s *Op_level1Context) OP_DIVIDE() antlr.TerminalNode {
	return s.GetToken(MathOP_DIVIDE, 0)
}

func (s *Op_level1Context) OP_MOD() antlr.TerminalNode {
	return s.GetToken(MathOP_MOD, 0)
}

func (s *Op_level1Context) OP_POWER() antlr.TerminalNode {
	return s.GetToken(MathOP_POWER, 0)
}

func (s *Op_level1Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Op_level1Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Op_level1Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MathListener); ok {
		listenerT.EnterOp_level1(s)
	}
}

func (s *Op_level1Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MathListener); ok {
		listenerT.ExitOp_level1(s)
	}
}

func (p *Math) Op_level1() (localctx IOp_level1Context) {
	localctx = NewOp_level1Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, MathRULE_op_level1)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(26)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&31457280) != 0) {
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

// IOp_level2Context is an interface to support dynamic dispatch.
type IOp_level2Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	OP_PLUS() antlr.TerminalNode
	OP_MINUS() antlr.TerminalNode

	// IsOp_level2Context differentiates from other interfaces.
	IsOp_level2Context()
}

type Op_level2Context struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOp_level2Context() *Op_level2Context {
	var p = new(Op_level2Context)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MathRULE_op_level2
	return p
}

func InitEmptyOp_level2Context(p *Op_level2Context) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MathRULE_op_level2
}

func (*Op_level2Context) IsOp_level2Context() {}

func NewOp_level2Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Op_level2Context {
	var p = new(Op_level2Context)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MathRULE_op_level2

	return p
}

func (s *Op_level2Context) GetParser() antlr.Parser { return s.parser }

func (s *Op_level2Context) OP_PLUS() antlr.TerminalNode {
	return s.GetToken(MathOP_PLUS, 0)
}

func (s *Op_level2Context) OP_MINUS() antlr.TerminalNode {
	return s.GetToken(MathOP_MINUS, 0)
}

func (s *Op_level2Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Op_level2Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Op_level2Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MathListener); ok {
		listenerT.EnterOp_level2(s)
	}
}

func (s *Op_level2Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MathListener); ok {
		listenerT.ExitOp_level2(s)
	}
}

func (p *Math) Op_level2() (localctx IOp_level2Context) {
	localctx = NewOp_level2Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, MathRULE_op_level2)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(28)
		_la = p.GetTokenStream().LA(1)

		if !(_la == MathOP_PLUS || _la == MathOP_MINUS) {
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

// IOp_level3Context is an interface to support dynamic dispatch.
type IOp_level3Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	OP_AND() antlr.TerminalNode
	OP_OR() antlr.TerminalNode
	OP_XOR() antlr.TerminalNode

	// IsOp_level3Context differentiates from other interfaces.
	IsOp_level3Context()
}

type Op_level3Context struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOp_level3Context() *Op_level3Context {
	var p = new(Op_level3Context)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MathRULE_op_level3
	return p
}

func InitEmptyOp_level3Context(p *Op_level3Context) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MathRULE_op_level3
}

func (*Op_level3Context) IsOp_level3Context() {}

func NewOp_level3Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Op_level3Context {
	var p = new(Op_level3Context)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MathRULE_op_level3

	return p
}

func (s *Op_level3Context) GetParser() antlr.Parser { return s.parser }

func (s *Op_level3Context) OP_AND() antlr.TerminalNode {
	return s.GetToken(MathOP_AND, 0)
}

func (s *Op_level3Context) OP_OR() antlr.TerminalNode {
	return s.GetToken(MathOP_OR, 0)
}

func (s *Op_level3Context) OP_XOR() antlr.TerminalNode {
	return s.GetToken(MathOP_XOR, 0)
}

func (s *Op_level3Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Op_level3Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Op_level3Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MathListener); ok {
		listenerT.EnterOp_level3(s)
	}
}

func (s *Op_level3Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MathListener); ok {
		listenerT.ExitOp_level3(s)
	}
}

func (p *Math) Op_level3() (localctx IOp_level3Context) {
	localctx = NewOp_level3Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, MathRULE_op_level3)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(30)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&939524096) != 0) {
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

// IOp_level4Context is an interface to support dynamic dispatch.
type IOp_level4Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	OP_EQUAL() antlr.TerminalNode
	OP_NOT_EQUAL() antlr.TerminalNode
	OP_LESS() antlr.TerminalNode
	OP_LESS_EQUAL() antlr.TerminalNode
	OP_GREATER() antlr.TerminalNode
	OP_GREATER_EQUAL() antlr.TerminalNode

	// IsOp_level4Context differentiates from other interfaces.
	IsOp_level4Context()
}

type Op_level4Context struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOp_level4Context() *Op_level4Context {
	var p = new(Op_level4Context)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MathRULE_op_level4
	return p
}

func InitEmptyOp_level4Context(p *Op_level4Context) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MathRULE_op_level4
}

func (*Op_level4Context) IsOp_level4Context() {}

func NewOp_level4Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Op_level4Context {
	var p = new(Op_level4Context)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MathRULE_op_level4

	return p
}

func (s *Op_level4Context) GetParser() antlr.Parser { return s.parser }

func (s *Op_level4Context) OP_EQUAL() antlr.TerminalNode {
	return s.GetToken(MathOP_EQUAL, 0)
}

func (s *Op_level4Context) OP_NOT_EQUAL() antlr.TerminalNode {
	return s.GetToken(MathOP_NOT_EQUAL, 0)
}

func (s *Op_level4Context) OP_LESS() antlr.TerminalNode {
	return s.GetToken(MathOP_LESS, 0)
}

func (s *Op_level4Context) OP_LESS_EQUAL() antlr.TerminalNode {
	return s.GetToken(MathOP_LESS_EQUAL, 0)
}

func (s *Op_level4Context) OP_GREATER() antlr.TerminalNode {
	return s.GetToken(MathOP_GREATER, 0)
}

func (s *Op_level4Context) OP_GREATER_EQUAL() antlr.TerminalNode {
	return s.GetToken(MathOP_GREATER_EQUAL, 0)
}

func (s *Op_level4Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Op_level4Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Op_level4Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MathListener); ok {
		listenerT.EnterOp_level4(s)
	}
}

func (s *Op_level4Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MathListener); ok {
		listenerT.ExitOp_level4(s)
	}
}

func (p *Math) Op_level4() (localctx IOp_level4Context) {
	localctx = NewOp_level4Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, MathRULE_op_level4)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(32)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&67645734912) != 0) {
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

// IOp_level5Context is an interface to support dynamic dispatch.
type IOp_level5Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	OP_NOT() antlr.TerminalNode

	// IsOp_level5Context differentiates from other interfaces.
	IsOp_level5Context()
}

type Op_level5Context struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOp_level5Context() *Op_level5Context {
	var p = new(Op_level5Context)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MathRULE_op_level5
	return p
}

func InitEmptyOp_level5Context(p *Op_level5Context) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MathRULE_op_level5
}

func (*Op_level5Context) IsOp_level5Context() {}

func NewOp_level5Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Op_level5Context {
	var p = new(Op_level5Context)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MathRULE_op_level5

	return p
}

func (s *Op_level5Context) GetParser() antlr.Parser { return s.parser }

func (s *Op_level5Context) OP_NOT() antlr.TerminalNode {
	return s.GetToken(MathOP_NOT, 0)
}

func (s *Op_level5Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Op_level5Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Op_level5Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MathListener); ok {
		listenerT.EnterOp_level5(s)
	}
}

func (s *Op_level5Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MathListener); ok {
		listenerT.ExitOp_level5(s)
	}
}

func (p *Math) Op_level5() (localctx IOp_level5Context) {
	localctx = NewOp_level5Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, MathRULE_op_level5)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(34)
		p.Match(MathOP_NOT)
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

// INumberic_expressionContext is an interface to support dynamic dispatch.
type INumberic_expressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	L_PAREN() antlr.TerminalNode
	AllNumberic_expression() []INumberic_expressionContext
	Numberic_expression(i int) INumberic_expressionContext
	R_PAREN() antlr.TerminalNode
	NUMBER() antlr.TerminalNode
	Get_variable() IGet_variableContext
	Function_call() IFunction_callContext
	OP_MINUS() antlr.TerminalNode
	Op_level1() IOp_level1Context
	Op_level2() IOp_level2Context

	// IsNumberic_expressionContext differentiates from other interfaces.
	IsNumberic_expressionContext()
}

type Numberic_expressionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNumberic_expressionContext() *Numberic_expressionContext {
	var p = new(Numberic_expressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MathRULE_numberic_expression
	return p
}

func InitEmptyNumberic_expressionContext(p *Numberic_expressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MathRULE_numberic_expression
}

func (*Numberic_expressionContext) IsNumberic_expressionContext() {}

func NewNumberic_expressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Numberic_expressionContext {
	var p = new(Numberic_expressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MathRULE_numberic_expression

	return p
}

func (s *Numberic_expressionContext) GetParser() antlr.Parser { return s.parser }

func (s *Numberic_expressionContext) L_PAREN() antlr.TerminalNode {
	return s.GetToken(MathL_PAREN, 0)
}

func (s *Numberic_expressionContext) AllNumberic_expression() []INumberic_expressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(INumberic_expressionContext); ok {
			len++
		}
	}

	tst := make([]INumberic_expressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(INumberic_expressionContext); ok {
			tst[i] = t.(INumberic_expressionContext)
			i++
		}
	}

	return tst
}

func (s *Numberic_expressionContext) Numberic_expression(i int) INumberic_expressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INumberic_expressionContext); ok {
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

	return t.(INumberic_expressionContext)
}

func (s *Numberic_expressionContext) R_PAREN() antlr.TerminalNode {
	return s.GetToken(MathR_PAREN, 0)
}

func (s *Numberic_expressionContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(MathNUMBER, 0)
}

func (s *Numberic_expressionContext) Get_variable() IGet_variableContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IGet_variableContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IGet_variableContext)
}

func (s *Numberic_expressionContext) Function_call() IFunction_callContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFunction_callContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFunction_callContext)
}

func (s *Numberic_expressionContext) OP_MINUS() antlr.TerminalNode {
	return s.GetToken(MathOP_MINUS, 0)
}

func (s *Numberic_expressionContext) Op_level1() IOp_level1Context {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOp_level1Context); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOp_level1Context)
}

func (s *Numberic_expressionContext) Op_level2() IOp_level2Context {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOp_level2Context); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOp_level2Context)
}

func (s *Numberic_expressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Numberic_expressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Numberic_expressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MathListener); ok {
		listenerT.EnterNumberic_expression(s)
	}
}

func (s *Numberic_expressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MathListener); ok {
		listenerT.ExitNumberic_expression(s)
	}
}

func (p *Math) Numberic_expression() (localctx INumberic_expressionContext) {
	return p.numberic_expression(0)
}

func (p *Math) numberic_expression(_p int) (localctx INumberic_expressionContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewNumberic_expressionContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx INumberic_expressionContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 10
	p.EnterRecursionRule(localctx, 10, MathRULE_numberic_expression, _p)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(46)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 0, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(37)
			p.Match(MathL_PAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(38)
			p.numberic_expression(0)
		}
		{
			p.SetState(39)
			p.Match(MathR_PAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		{
			p.SetState(41)
			p.Match(MathNUMBER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 3:
		{
			p.SetState(42)
			p.Get_variable()
		}

	case 4:
		{
			p.SetState(43)
			p.Function_call()
		}

	case 5:
		{
			p.SetState(44)
			p.Match(MathOP_MINUS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(45)
			p.numberic_expression(1)
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(58)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 2, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(56)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 1, p.GetParserRuleContext()) {
			case 1:
				localctx = NewNumberic_expressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, MathRULE_numberic_expression)
				p.SetState(48)

				if !(p.Precpred(p.GetParserRuleContext(), 6)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 6)", ""))
					goto errorExit
				}
				{
					p.SetState(49)
					p.Op_level1()
				}
				{
					p.SetState(50)
					p.numberic_expression(7)
				}

			case 2:
				localctx = NewNumberic_expressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, MathRULE_numberic_expression)
				p.SetState(52)

				if !(p.Precpred(p.GetParserRuleContext(), 5)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 5)", ""))
					goto errorExit
				}
				{
					p.SetState(53)
					p.Op_level2()
				}
				{
					p.SetState(54)
					p.numberic_expression(6)
				}

			case antlr.ATNInvalidAltNumber:
				goto errorExit
			}

		}
		p.SetState(60)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 2, p.GetParserRuleContext())
		if p.HasError() {
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
	p.UnrollRecursionContexts(_parentctx)
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ILogical_expressionContext is an interface to support dynamic dispatch.
type ILogical_expressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	L_PAREN() antlr.TerminalNode
	AllLogical_expression() []ILogical_expressionContext
	Logical_expression(i int) ILogical_expressionContext
	R_PAREN() antlr.TerminalNode
	OP_NOT() antlr.TerminalNode
	BOOLEAN() antlr.TerminalNode
	Get_variable() IGet_variableContext
	Function_call() IFunction_callContext
	Op_level3() IOp_level3Context
	Op_level4() IOp_level4Context

	// IsLogical_expressionContext differentiates from other interfaces.
	IsLogical_expressionContext()
}

type Logical_expressionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLogical_expressionContext() *Logical_expressionContext {
	var p = new(Logical_expressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MathRULE_logical_expression
	return p
}

func InitEmptyLogical_expressionContext(p *Logical_expressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MathRULE_logical_expression
}

func (*Logical_expressionContext) IsLogical_expressionContext() {}

func NewLogical_expressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Logical_expressionContext {
	var p = new(Logical_expressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MathRULE_logical_expression

	return p
}

func (s *Logical_expressionContext) GetParser() antlr.Parser { return s.parser }

func (s *Logical_expressionContext) L_PAREN() antlr.TerminalNode {
	return s.GetToken(MathL_PAREN, 0)
}

func (s *Logical_expressionContext) AllLogical_expression() []ILogical_expressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ILogical_expressionContext); ok {
			len++
		}
	}

	tst := make([]ILogical_expressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ILogical_expressionContext); ok {
			tst[i] = t.(ILogical_expressionContext)
			i++
		}
	}

	return tst
}

func (s *Logical_expressionContext) Logical_expression(i int) ILogical_expressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILogical_expressionContext); ok {
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

	return t.(ILogical_expressionContext)
}

func (s *Logical_expressionContext) R_PAREN() antlr.TerminalNode {
	return s.GetToken(MathR_PAREN, 0)
}

func (s *Logical_expressionContext) OP_NOT() antlr.TerminalNode {
	return s.GetToken(MathOP_NOT, 0)
}

func (s *Logical_expressionContext) BOOLEAN() antlr.TerminalNode {
	return s.GetToken(MathBOOLEAN, 0)
}

func (s *Logical_expressionContext) Get_variable() IGet_variableContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IGet_variableContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IGet_variableContext)
}

func (s *Logical_expressionContext) Function_call() IFunction_callContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFunction_callContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFunction_callContext)
}

func (s *Logical_expressionContext) Op_level3() IOp_level3Context {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOp_level3Context); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOp_level3Context)
}

func (s *Logical_expressionContext) Op_level4() IOp_level4Context {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOp_level4Context); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOp_level4Context)
}

func (s *Logical_expressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Logical_expressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Logical_expressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MathListener); ok {
		listenerT.EnterLogical_expression(s)
	}
}

func (s *Logical_expressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MathListener); ok {
		listenerT.ExitLogical_expression(s)
	}
}

func (p *Math) Logical_expression() (localctx ILogical_expressionContext) {
	return p.logical_expression(0)
}

func (p *Math) logical_expression(_p int) (localctx ILogical_expressionContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewLogical_expressionContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx ILogical_expressionContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 12
	p.EnterRecursionRule(localctx, 12, MathRULE_logical_expression, _p)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(71)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 3, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(62)
			p.Match(MathL_PAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(63)
			p.logical_expression(0)
		}
		{
			p.SetState(64)
			p.Match(MathR_PAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		{
			p.SetState(66)
			p.Match(MathOP_NOT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(67)
			p.logical_expression(4)
		}

	case 3:
		{
			p.SetState(68)
			p.Match(MathBOOLEAN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 4:
		{
			p.SetState(69)
			p.Get_variable()
		}

	case 5:
		{
			p.SetState(70)
			p.Function_call()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(83)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 5, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(81)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 4, p.GetParserRuleContext()) {
			case 1:
				localctx = NewLogical_expressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, MathRULE_logical_expression)
				p.SetState(73)

				if !(p.Precpred(p.GetParserRuleContext(), 6)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 6)", ""))
					goto errorExit
				}
				{
					p.SetState(74)
					p.Op_level3()
				}
				{
					p.SetState(75)
					p.logical_expression(7)
				}

			case 2:
				localctx = NewLogical_expressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, MathRULE_logical_expression)
				p.SetState(77)

				if !(p.Precpred(p.GetParserRuleContext(), 5)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 5)", ""))
					goto errorExit
				}
				{
					p.SetState(78)
					p.Op_level4()
				}
				{
					p.SetState(79)
					p.logical_expression(6)
				}

			case antlr.ATNInvalidAltNumber:
				goto errorExit
			}

		}
		p.SetState(85)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 5, p.GetParserRuleContext())
		if p.HasError() {
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
	p.UnrollRecursionContexts(_parentctx)
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IComparative_expressionContext is an interface to support dynamic dispatch.
type IComparative_expressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllNumberic_expression() []INumberic_expressionContext
	Numberic_expression(i int) INumberic_expressionContext
	Op_level4() IOp_level4Context
	AllLogical_expression() []ILogical_expressionContext
	Logical_expression(i int) ILogical_expressionContext
	Op_level3() IOp_level3Context
	Op_level5() IOp_level5Context
	L_PAREN() antlr.TerminalNode
	Comparative_expression() IComparative_expressionContext
	R_PAREN() antlr.TerminalNode

	// IsComparative_expressionContext differentiates from other interfaces.
	IsComparative_expressionContext()
}

type Comparative_expressionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyComparative_expressionContext() *Comparative_expressionContext {
	var p = new(Comparative_expressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MathRULE_comparative_expression
	return p
}

func InitEmptyComparative_expressionContext(p *Comparative_expressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MathRULE_comparative_expression
}

func (*Comparative_expressionContext) IsComparative_expressionContext() {}

func NewComparative_expressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Comparative_expressionContext {
	var p = new(Comparative_expressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MathRULE_comparative_expression

	return p
}

func (s *Comparative_expressionContext) GetParser() antlr.Parser { return s.parser }

func (s *Comparative_expressionContext) AllNumberic_expression() []INumberic_expressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(INumberic_expressionContext); ok {
			len++
		}
	}

	tst := make([]INumberic_expressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(INumberic_expressionContext); ok {
			tst[i] = t.(INumberic_expressionContext)
			i++
		}
	}

	return tst
}

func (s *Comparative_expressionContext) Numberic_expression(i int) INumberic_expressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INumberic_expressionContext); ok {
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

	return t.(INumberic_expressionContext)
}

func (s *Comparative_expressionContext) Op_level4() IOp_level4Context {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOp_level4Context); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOp_level4Context)
}

func (s *Comparative_expressionContext) AllLogical_expression() []ILogical_expressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ILogical_expressionContext); ok {
			len++
		}
	}

	tst := make([]ILogical_expressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ILogical_expressionContext); ok {
			tst[i] = t.(ILogical_expressionContext)
			i++
		}
	}

	return tst
}

func (s *Comparative_expressionContext) Logical_expression(i int) ILogical_expressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILogical_expressionContext); ok {
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

	return t.(ILogical_expressionContext)
}

func (s *Comparative_expressionContext) Op_level3() IOp_level3Context {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOp_level3Context); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOp_level3Context)
}

func (s *Comparative_expressionContext) Op_level5() IOp_level5Context {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOp_level5Context); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOp_level5Context)
}

func (s *Comparative_expressionContext) L_PAREN() antlr.TerminalNode {
	return s.GetToken(MathL_PAREN, 0)
}

func (s *Comparative_expressionContext) Comparative_expression() IComparative_expressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IComparative_expressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IComparative_expressionContext)
}

func (s *Comparative_expressionContext) R_PAREN() antlr.TerminalNode {
	return s.GetToken(MathR_PAREN, 0)
}

func (s *Comparative_expressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Comparative_expressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Comparative_expressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MathListener); ok {
		listenerT.EnterComparative_expression(s)
	}
}

func (s *Comparative_expressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MathListener); ok {
		listenerT.ExitComparative_expression(s)
	}
}

func (p *Math) Comparative_expression() (localctx IComparative_expressionContext) {
	return p.comparative_expression(0)
}

func (p *Math) comparative_expression(_p int) (localctx IComparative_expressionContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewComparative_expressionContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IComparative_expressionContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 14
	p.EnterRecursionRule(localctx, 14, MathRULE_comparative_expression, _p)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(100)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 6, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(87)
			p.numberic_expression(0)
		}
		{
			p.SetState(88)
			p.Op_level4()
		}
		{
			p.SetState(89)
			p.numberic_expression(0)
		}

	case 2:
		{
			p.SetState(91)
			p.logical_expression(0)
		}
		{
			p.SetState(92)
			p.Op_level3()
		}
		{
			p.SetState(93)
			p.logical_expression(0)
		}

	case 3:
		{
			p.SetState(95)
			p.Op_level5()
		}
		{
			p.SetState(96)
			p.Match(MathL_PAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(97)
			p.comparative_expression(0)
		}
		{
			p.SetState(98)
			p.Match(MathR_PAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(108)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 7, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewComparative_expressionContext(p, _parentctx, _parentState)
			p.PushNewRecursionContext(localctx, _startState, MathRULE_comparative_expression)
			p.SetState(102)

			if !(p.Precpred(p.GetParserRuleContext(), 2)) {
				p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
				goto errorExit
			}
			{
				p.SetState(103)
				p.Op_level3()
			}
			{
				p.SetState(104)
				p.logical_expression(0)
			}

		}
		p.SetState(110)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 7, p.GetParserRuleContext())
		if p.HasError() {
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
	p.UnrollRecursionContexts(_parentctx)
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IGet_variableContext is an interface to support dynamic dispatch.
type IGet_variableContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	VAR_IDENTIFIER() antlr.TerminalNode
	Get_array_element() IGet_array_elementContext
	Get_child() IGet_childContext

	// IsGet_variableContext differentiates from other interfaces.
	IsGet_variableContext()
}

type Get_variableContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyGet_variableContext() *Get_variableContext {
	var p = new(Get_variableContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MathRULE_get_variable
	return p
}

func InitEmptyGet_variableContext(p *Get_variableContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MathRULE_get_variable
}

func (*Get_variableContext) IsGet_variableContext() {}

func NewGet_variableContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Get_variableContext {
	var p = new(Get_variableContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MathRULE_get_variable

	return p
}

func (s *Get_variableContext) GetParser() antlr.Parser { return s.parser }

func (s *Get_variableContext) VAR_IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(MathVAR_IDENTIFIER, 0)
}

func (s *Get_variableContext) Get_array_element() IGet_array_elementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IGet_array_elementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IGet_array_elementContext)
}

func (s *Get_variableContext) Get_child() IGet_childContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IGet_childContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IGet_childContext)
}

func (s *Get_variableContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Get_variableContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Get_variableContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MathListener); ok {
		listenerT.EnterGet_variable(s)
	}
}

func (s *Get_variableContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MathListener); ok {
		listenerT.ExitGet_variable(s)
	}
}

func (p *Math) Get_variable() (localctx IGet_variableContext) {
	localctx = NewGet_variableContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, MathRULE_get_variable)
	p.SetState(114)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 8, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(111)
			p.Match(MathVAR_IDENTIFIER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(112)
			p.Get_array_element()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(113)
			p.Get_child()
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

// IMath_expressionContext is an interface to support dynamic dispatch.
type IMath_expressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Get_variable() IGet_variableContext
	Numberic_expression() INumberic_expressionContext
	Logical_expression() ILogical_expressionContext

	// IsMath_expressionContext differentiates from other interfaces.
	IsMath_expressionContext()
}

type Math_expressionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMath_expressionContext() *Math_expressionContext {
	var p = new(Math_expressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MathRULE_math_expression
	return p
}

func InitEmptyMath_expressionContext(p *Math_expressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MathRULE_math_expression
}

func (*Math_expressionContext) IsMath_expressionContext() {}

func NewMath_expressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Math_expressionContext {
	var p = new(Math_expressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MathRULE_math_expression

	return p
}

func (s *Math_expressionContext) GetParser() antlr.Parser { return s.parser }

func (s *Math_expressionContext) Get_variable() IGet_variableContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IGet_variableContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IGet_variableContext)
}

func (s *Math_expressionContext) Numberic_expression() INumberic_expressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INumberic_expressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INumberic_expressionContext)
}

func (s *Math_expressionContext) Logical_expression() ILogical_expressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILogical_expressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILogical_expressionContext)
}

func (s *Math_expressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Math_expressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Math_expressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MathListener); ok {
		listenerT.EnterMath_expression(s)
	}
}

func (s *Math_expressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MathListener); ok {
		listenerT.ExitMath_expression(s)
	}
}

func (p *Math) Math_expression() (localctx IMath_expressionContext) {
	localctx = NewMath_expressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, MathRULE_math_expression)
	p.SetState(119)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 9, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(116)
			p.Get_variable()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(117)
			p.numberic_expression(0)
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(118)
			p.logical_expression(0)
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

// IGet_array_elementContext is an interface to support dynamic dispatch.
type IGet_array_elementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	VAR_IDENTIFIER() antlr.TerminalNode
	L_SQUARE() antlr.TerminalNode
	Numberic_expression() INumberic_expressionContext
	R_SQUARE() antlr.TerminalNode

	// IsGet_array_elementContext differentiates from other interfaces.
	IsGet_array_elementContext()
}

type Get_array_elementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyGet_array_elementContext() *Get_array_elementContext {
	var p = new(Get_array_elementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MathRULE_get_array_element
	return p
}

func InitEmptyGet_array_elementContext(p *Get_array_elementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MathRULE_get_array_element
}

func (*Get_array_elementContext) IsGet_array_elementContext() {}

func NewGet_array_elementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Get_array_elementContext {
	var p = new(Get_array_elementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MathRULE_get_array_element

	return p
}

func (s *Get_array_elementContext) GetParser() antlr.Parser { return s.parser }

func (s *Get_array_elementContext) VAR_IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(MathVAR_IDENTIFIER, 0)
}

func (s *Get_array_elementContext) L_SQUARE() antlr.TerminalNode {
	return s.GetToken(MathL_SQUARE, 0)
}

func (s *Get_array_elementContext) Numberic_expression() INumberic_expressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INumberic_expressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INumberic_expressionContext)
}

func (s *Get_array_elementContext) R_SQUARE() antlr.TerminalNode {
	return s.GetToken(MathR_SQUARE, 0)
}

func (s *Get_array_elementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Get_array_elementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Get_array_elementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MathListener); ok {
		listenerT.EnterGet_array_element(s)
	}
}

func (s *Get_array_elementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MathListener); ok {
		listenerT.ExitGet_array_element(s)
	}
}

func (p *Math) Get_array_element() (localctx IGet_array_elementContext) {
	localctx = NewGet_array_elementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, MathRULE_get_array_element)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(121)
		p.Match(MathVAR_IDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(122)
		p.Match(MathL_SQUARE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(123)
		p.numberic_expression(0)
	}
	{
		p.SetState(124)
		p.Match(MathR_SQUARE)
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

// IGet_childContext is an interface to support dynamic dispatch.
type IGet_childContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllVAR_IDENTIFIER() []antlr.TerminalNode
	VAR_IDENTIFIER(i int) antlr.TerminalNode
	DOT() antlr.TerminalNode
	Get_array_element() IGet_array_elementContext
	Get_child() IGet_childContext

	// IsGet_childContext differentiates from other interfaces.
	IsGet_childContext()
}

type Get_childContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyGet_childContext() *Get_childContext {
	var p = new(Get_childContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MathRULE_get_child
	return p
}

func InitEmptyGet_childContext(p *Get_childContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MathRULE_get_child
}

func (*Get_childContext) IsGet_childContext() {}

func NewGet_childContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Get_childContext {
	var p = new(Get_childContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MathRULE_get_child

	return p
}

func (s *Get_childContext) GetParser() antlr.Parser { return s.parser }

func (s *Get_childContext) AllVAR_IDENTIFIER() []antlr.TerminalNode {
	return s.GetTokens(MathVAR_IDENTIFIER)
}

func (s *Get_childContext) VAR_IDENTIFIER(i int) antlr.TerminalNode {
	return s.GetToken(MathVAR_IDENTIFIER, i)
}

func (s *Get_childContext) DOT() antlr.TerminalNode {
	return s.GetToken(MathDOT, 0)
}

func (s *Get_childContext) Get_array_element() IGet_array_elementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IGet_array_elementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IGet_array_elementContext)
}

func (s *Get_childContext) Get_child() IGet_childContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IGet_childContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IGet_childContext)
}

func (s *Get_childContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Get_childContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Get_childContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MathListener); ok {
		listenerT.EnterGet_child(s)
	}
}

func (s *Get_childContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MathListener); ok {
		listenerT.ExitGet_child(s)
	}
}

func (p *Math) Get_child() (localctx IGet_childContext) {
	localctx = NewGet_childContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, MathRULE_get_child)
	p.SetState(141)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 10, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(126)
			p.Match(MathVAR_IDENTIFIER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(127)
			p.Match(MathDOT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(128)
			p.Match(MathVAR_IDENTIFIER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(129)
			p.Match(MathVAR_IDENTIFIER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(130)
			p.Match(MathDOT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(131)
			p.Get_array_element()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(132)
			p.Match(MathVAR_IDENTIFIER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(133)
			p.Match(MathDOT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(134)
			p.Get_child()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(135)
			p.Get_array_element()
		}
		{
			p.SetState(136)
			p.Match(MathDOT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(137)
			p.Get_child()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(139)
			p.Get_array_element()
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(140)
			p.Match(MathVAR_IDENTIFIER)
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

// IFunction_callContext is an interface to support dynamic dispatch.
type IFunction_callContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	VAR_IDENTIFIER() antlr.TerminalNode
	L_PAREN() antlr.TerminalNode
	R_PAREN() antlr.TerminalNode
	AllMath_expression() []IMath_expressionContext
	Math_expression(i int) IMath_expressionContext
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsFunction_callContext differentiates from other interfaces.
	IsFunction_callContext()
}

type Function_callContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunction_callContext() *Function_callContext {
	var p = new(Function_callContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MathRULE_function_call
	return p
}

func InitEmptyFunction_callContext(p *Function_callContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MathRULE_function_call
}

func (*Function_callContext) IsFunction_callContext() {}

func NewFunction_callContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Function_callContext {
	var p = new(Function_callContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MathRULE_function_call

	return p
}

func (s *Function_callContext) GetParser() antlr.Parser { return s.parser }

func (s *Function_callContext) VAR_IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(MathVAR_IDENTIFIER, 0)
}

func (s *Function_callContext) L_PAREN() antlr.TerminalNode {
	return s.GetToken(MathL_PAREN, 0)
}

func (s *Function_callContext) R_PAREN() antlr.TerminalNode {
	return s.GetToken(MathR_PAREN, 0)
}

func (s *Function_callContext) AllMath_expression() []IMath_expressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IMath_expressionContext); ok {
			len++
		}
	}

	tst := make([]IMath_expressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IMath_expressionContext); ok {
			tst[i] = t.(IMath_expressionContext)
			i++
		}
	}

	return tst
}

func (s *Function_callContext) Math_expression(i int) IMath_expressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMath_expressionContext); ok {
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

	return t.(IMath_expressionContext)
}

func (s *Function_callContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(MathCOMMA)
}

func (s *Function_callContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(MathCOMMA, i)
}

func (s *Function_callContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Function_callContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Function_callContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MathListener); ok {
		listenerT.EnterFunction_call(s)
	}
}

func (s *Function_callContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MathListener); ok {
		listenerT.ExitFunction_call(s)
	}
}

func (p *Math) Function_call() (localctx IFunction_callContext) {
	localctx = NewFunction_callContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, MathRULE_function_call)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(143)
		p.Match(MathVAR_IDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(144)
		p.Match(MathL_PAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(153)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&206225547336) != 0 {
		{
			p.SetState(145)
			p.Math_expression()
		}
		p.SetState(150)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == MathCOMMA {
			{
				p.SetState(146)
				p.Match(MathCOMMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(147)
				p.Math_expression()
			}

			p.SetState(152)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(155)
		p.Match(MathR_PAREN)
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

func (p *Math) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 5:
		var t *Numberic_expressionContext = nil
		if localctx != nil {
			t = localctx.(*Numberic_expressionContext)
		}
		return p.Numberic_expression_Sempred(t, predIndex)

	case 6:
		var t *Logical_expressionContext = nil
		if localctx != nil {
			t = localctx.(*Logical_expressionContext)
		}
		return p.Logical_expression_Sempred(t, predIndex)

	case 7:
		var t *Comparative_expressionContext = nil
		if localctx != nil {
			t = localctx.(*Comparative_expressionContext)
		}
		return p.Comparative_expression_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *Math) Numberic_expression_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 6)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 5)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *Math) Logical_expression_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 2:
		return p.Precpred(p.GetParserRuleContext(), 6)

	case 3:
		return p.Precpred(p.GetParserRuleContext(), 5)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *Math) Comparative_expression_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 4:
		return p.Precpred(p.GetParserRuleContext(), 2)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
