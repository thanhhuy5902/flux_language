// Code generated from Variables.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parsing // Variables
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

type Variables struct {
	*antlr.BaseParser
}

var VariablesParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func variablesParserInit() {
	staticData := &VariablesParserStaticData
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
		"var_type", "var_name", "var_value", "default_var_declaration", "text_var_declaration",
		"number_var_declaration", "boolean_var_declaration", "ipv4_var_declaration",
		"var_declaration",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 18, 61, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 1, 0, 1, 0, 1, 1, 1,
		1, 1, 2, 1, 2, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 4, 1, 4, 1, 4, 1, 4, 1,
		4, 1, 4, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 6, 1, 6, 1, 6, 1, 6, 1,
		6, 1, 6, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 8, 1, 8, 1, 8, 1, 8, 1,
		8, 3, 8, 59, 8, 8, 1, 8, 0, 0, 9, 0, 2, 4, 6, 8, 10, 12, 14, 16, 0, 2,
		4, 0, 2, 2, 4, 4, 6, 6, 10, 10, 4, 0, 1, 1, 3, 3, 5, 5, 9, 9, 55, 0, 18,
		1, 0, 0, 0, 2, 20, 1, 0, 0, 0, 4, 22, 1, 0, 0, 0, 6, 24, 1, 0, 0, 0, 8,
		29, 1, 0, 0, 0, 10, 35, 1, 0, 0, 0, 12, 41, 1, 0, 0, 0, 14, 47, 1, 0, 0,
		0, 16, 58, 1, 0, 0, 0, 18, 19, 7, 0, 0, 0, 19, 1, 1, 0, 0, 0, 20, 21, 5,
		11, 0, 0, 21, 3, 1, 0, 0, 0, 22, 23, 7, 1, 0, 0, 23, 5, 1, 0, 0, 0, 24,
		25, 3, 0, 0, 0, 25, 26, 3, 2, 1, 0, 26, 27, 5, 15, 0, 0, 27, 28, 5, 16,
		0, 0, 28, 7, 1, 0, 0, 0, 29, 30, 5, 2, 0, 0, 30, 31, 3, 2, 1, 0, 31, 32,
		5, 15, 0, 0, 32, 33, 5, 1, 0, 0, 33, 34, 5, 16, 0, 0, 34, 9, 1, 0, 0, 0,
		35, 36, 5, 6, 0, 0, 36, 37, 3, 2, 1, 0, 37, 38, 5, 15, 0, 0, 38, 39, 5,
		5, 0, 0, 39, 40, 5, 16, 0, 0, 40, 11, 1, 0, 0, 0, 41, 42, 5, 4, 0, 0, 42,
		43, 3, 2, 1, 0, 43, 44, 5, 15, 0, 0, 44, 45, 5, 3, 0, 0, 45, 46, 5, 16,
		0, 0, 46, 13, 1, 0, 0, 0, 47, 48, 5, 10, 0, 0, 48, 49, 3, 2, 1, 0, 49,
		50, 5, 15, 0, 0, 50, 51, 5, 9, 0, 0, 51, 52, 5, 16, 0, 0, 52, 15, 1, 0,
		0, 0, 53, 59, 3, 6, 3, 0, 54, 59, 3, 8, 4, 0, 55, 59, 3, 10, 5, 0, 56,
		59, 3, 12, 6, 0, 57, 59, 3, 14, 7, 0, 58, 53, 1, 0, 0, 0, 58, 54, 1, 0,
		0, 0, 58, 55, 1, 0, 0, 0, 58, 56, 1, 0, 0, 0, 58, 57, 1, 0, 0, 0, 59, 17,
		1, 0, 0, 0, 1, 58,
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

// VariablesInit initializes any static state used to implement Variables. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewVariables(). You can call this function if you wish to initialize the static state ahead
// of time.
func VariablesInit() {
	staticData := &VariablesParserStaticData
	staticData.once.Do(variablesParserInit)
}

// NewVariables produces a new parser instance for the optional input antlr.TokenStream.
func NewVariables(input antlr.TokenStream) *Variables {
	VariablesInit()
	this := new(Variables)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &VariablesParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "Variables.g4"

	return this
}

// Variables tokens.
const (
	VariablesEOF                 = antlr.TokenEOF
	VariablesTEXT                = 1
	VariablesTEXT_TYPE           = 2
	VariablesBOOLEAN             = 3
	VariablesBOOLEAN_TYPE        = 4
	VariablesNUMBER              = 5
	VariablesNUMBER_TYPE         = 6
	VariablesDIGIT               = 7
	VariablesOCTET               = 8
	VariablesIPV4                = 9
	VariablesIPV4_TYPE           = 10
	VariablesVAR_IDENTIFIER      = 11
	VariablesCOMMON_IDENTIFIER   = 12
	VariablesARITHMETIC_OPERATOR = 13
	VariablesLOGICAL_OPERATOR    = 14
	VariablesL_BLOCK             = 15
	VariablesR_BLOCK             = 16
	VariablesAT                  = 17
	VariablesWS                  = 18
)

// Variables rules.
const (
	VariablesRULE_var_type                = 0
	VariablesRULE_var_name                = 1
	VariablesRULE_var_value               = 2
	VariablesRULE_default_var_declaration = 3
	VariablesRULE_text_var_declaration    = 4
	VariablesRULE_number_var_declaration  = 5
	VariablesRULE_boolean_var_declaration = 6
	VariablesRULE_ipv4_var_declaration    = 7
	VariablesRULE_var_declaration         = 8
)

// IVar_typeContext is an interface to support dynamic dispatch.
type IVar_typeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	TEXT_TYPE() antlr.TerminalNode
	NUMBER_TYPE() antlr.TerminalNode
	BOOLEAN_TYPE() antlr.TerminalNode
	IPV4_TYPE() antlr.TerminalNode

	// IsVar_typeContext differentiates from other interfaces.
	IsVar_typeContext()
}

type Var_typeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVar_typeContext() *Var_typeContext {
	var p = new(Var_typeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VariablesRULE_var_type
	return p
}

func InitEmptyVar_typeContext(p *Var_typeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VariablesRULE_var_type
}

func (*Var_typeContext) IsVar_typeContext() {}

func NewVar_typeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Var_typeContext {
	var p = new(Var_typeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VariablesRULE_var_type

	return p
}

func (s *Var_typeContext) GetParser() antlr.Parser { return s.parser }

func (s *Var_typeContext) TEXT_TYPE() antlr.TerminalNode {
	return s.GetToken(VariablesTEXT_TYPE, 0)
}

func (s *Var_typeContext) NUMBER_TYPE() antlr.TerminalNode {
	return s.GetToken(VariablesNUMBER_TYPE, 0)
}

func (s *Var_typeContext) BOOLEAN_TYPE() antlr.TerminalNode {
	return s.GetToken(VariablesBOOLEAN_TYPE, 0)
}

func (s *Var_typeContext) IPV4_TYPE() antlr.TerminalNode {
	return s.GetToken(VariablesIPV4_TYPE, 0)
}

func (s *Var_typeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Var_typeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Var_typeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VariablesListener); ok {
		listenerT.EnterVar_type(s)
	}
}

func (s *Var_typeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VariablesListener); ok {
		listenerT.ExitVar_type(s)
	}
}

func (p *Variables) Var_type() (localctx IVar_typeContext) {
	localctx = NewVar_typeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, VariablesRULE_var_type)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(18)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&1108) != 0) {
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

// IVar_nameContext is an interface to support dynamic dispatch.
type IVar_nameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	VAR_IDENTIFIER() antlr.TerminalNode

	// IsVar_nameContext differentiates from other interfaces.
	IsVar_nameContext()
}

type Var_nameContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVar_nameContext() *Var_nameContext {
	var p = new(Var_nameContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VariablesRULE_var_name
	return p
}

func InitEmptyVar_nameContext(p *Var_nameContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VariablesRULE_var_name
}

func (*Var_nameContext) IsVar_nameContext() {}

func NewVar_nameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Var_nameContext {
	var p = new(Var_nameContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VariablesRULE_var_name

	return p
}

func (s *Var_nameContext) GetParser() antlr.Parser { return s.parser }

func (s *Var_nameContext) VAR_IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(VariablesVAR_IDENTIFIER, 0)
}

func (s *Var_nameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Var_nameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Var_nameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VariablesListener); ok {
		listenerT.EnterVar_name(s)
	}
}

func (s *Var_nameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VariablesListener); ok {
		listenerT.ExitVar_name(s)
	}
}

func (p *Variables) Var_name() (localctx IVar_nameContext) {
	localctx = NewVar_nameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, VariablesRULE_var_name)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(20)
		p.Match(VariablesVAR_IDENTIFIER)
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

// IVar_valueContext is an interface to support dynamic dispatch.
type IVar_valueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IPV4() antlr.TerminalNode
	TEXT() antlr.TerminalNode
	NUMBER() antlr.TerminalNode
	BOOLEAN() antlr.TerminalNode

	// IsVar_valueContext differentiates from other interfaces.
	IsVar_valueContext()
}

type Var_valueContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVar_valueContext() *Var_valueContext {
	var p = new(Var_valueContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VariablesRULE_var_value
	return p
}

func InitEmptyVar_valueContext(p *Var_valueContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VariablesRULE_var_value
}

func (*Var_valueContext) IsVar_valueContext() {}

func NewVar_valueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Var_valueContext {
	var p = new(Var_valueContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VariablesRULE_var_value

	return p
}

func (s *Var_valueContext) GetParser() antlr.Parser { return s.parser }

func (s *Var_valueContext) IPV4() antlr.TerminalNode {
	return s.GetToken(VariablesIPV4, 0)
}

func (s *Var_valueContext) TEXT() antlr.TerminalNode {
	return s.GetToken(VariablesTEXT, 0)
}

func (s *Var_valueContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(VariablesNUMBER, 0)
}

func (s *Var_valueContext) BOOLEAN() antlr.TerminalNode {
	return s.GetToken(VariablesBOOLEAN, 0)
}

func (s *Var_valueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Var_valueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Var_valueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VariablesListener); ok {
		listenerT.EnterVar_value(s)
	}
}

func (s *Var_valueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VariablesListener); ok {
		listenerT.ExitVar_value(s)
	}
}

func (p *Variables) Var_value() (localctx IVar_valueContext) {
	localctx = NewVar_valueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, VariablesRULE_var_value)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(22)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&554) != 0) {
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

// IDefault_var_declarationContext is an interface to support dynamic dispatch.
type IDefault_var_declarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Var_type() IVar_typeContext
	Var_name() IVar_nameContext
	L_BLOCK() antlr.TerminalNode
	R_BLOCK() antlr.TerminalNode

	// IsDefault_var_declarationContext differentiates from other interfaces.
	IsDefault_var_declarationContext()
}

type Default_var_declarationContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDefault_var_declarationContext() *Default_var_declarationContext {
	var p = new(Default_var_declarationContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VariablesRULE_default_var_declaration
	return p
}

func InitEmptyDefault_var_declarationContext(p *Default_var_declarationContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VariablesRULE_default_var_declaration
}

func (*Default_var_declarationContext) IsDefault_var_declarationContext() {}

func NewDefault_var_declarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Default_var_declarationContext {
	var p = new(Default_var_declarationContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VariablesRULE_default_var_declaration

	return p
}

func (s *Default_var_declarationContext) GetParser() antlr.Parser { return s.parser }

func (s *Default_var_declarationContext) Var_type() IVar_typeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVar_typeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVar_typeContext)
}

func (s *Default_var_declarationContext) Var_name() IVar_nameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVar_nameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVar_nameContext)
}

func (s *Default_var_declarationContext) L_BLOCK() antlr.TerminalNode {
	return s.GetToken(VariablesL_BLOCK, 0)
}

func (s *Default_var_declarationContext) R_BLOCK() antlr.TerminalNode {
	return s.GetToken(VariablesR_BLOCK, 0)
}

func (s *Default_var_declarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Default_var_declarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Default_var_declarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VariablesListener); ok {
		listenerT.EnterDefault_var_declaration(s)
	}
}

func (s *Default_var_declarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VariablesListener); ok {
		listenerT.ExitDefault_var_declaration(s)
	}
}

func (p *Variables) Default_var_declaration() (localctx IDefault_var_declarationContext) {
	localctx = NewDefault_var_declarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, VariablesRULE_default_var_declaration)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(24)
		p.Var_type()
	}
	{
		p.SetState(25)
		p.Var_name()
	}
	{
		p.SetState(26)
		p.Match(VariablesL_BLOCK)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(27)
		p.Match(VariablesR_BLOCK)
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

// IText_var_declarationContext is an interface to support dynamic dispatch.
type IText_var_declarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	TEXT_TYPE() antlr.TerminalNode
	Var_name() IVar_nameContext
	L_BLOCK() antlr.TerminalNode
	TEXT() antlr.TerminalNode
	R_BLOCK() antlr.TerminalNode

	// IsText_var_declarationContext differentiates from other interfaces.
	IsText_var_declarationContext()
}

type Text_var_declarationContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyText_var_declarationContext() *Text_var_declarationContext {
	var p = new(Text_var_declarationContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VariablesRULE_text_var_declaration
	return p
}

func InitEmptyText_var_declarationContext(p *Text_var_declarationContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VariablesRULE_text_var_declaration
}

func (*Text_var_declarationContext) IsText_var_declarationContext() {}

func NewText_var_declarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Text_var_declarationContext {
	var p = new(Text_var_declarationContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VariablesRULE_text_var_declaration

	return p
}

func (s *Text_var_declarationContext) GetParser() antlr.Parser { return s.parser }

func (s *Text_var_declarationContext) TEXT_TYPE() antlr.TerminalNode {
	return s.GetToken(VariablesTEXT_TYPE, 0)
}

func (s *Text_var_declarationContext) Var_name() IVar_nameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVar_nameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVar_nameContext)
}

func (s *Text_var_declarationContext) L_BLOCK() antlr.TerminalNode {
	return s.GetToken(VariablesL_BLOCK, 0)
}

func (s *Text_var_declarationContext) TEXT() antlr.TerminalNode {
	return s.GetToken(VariablesTEXT, 0)
}

func (s *Text_var_declarationContext) R_BLOCK() antlr.TerminalNode {
	return s.GetToken(VariablesR_BLOCK, 0)
}

func (s *Text_var_declarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Text_var_declarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Text_var_declarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VariablesListener); ok {
		listenerT.EnterText_var_declaration(s)
	}
}

func (s *Text_var_declarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VariablesListener); ok {
		listenerT.ExitText_var_declaration(s)
	}
}

func (p *Variables) Text_var_declaration() (localctx IText_var_declarationContext) {
	localctx = NewText_var_declarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, VariablesRULE_text_var_declaration)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(29)
		p.Match(VariablesTEXT_TYPE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(30)
		p.Var_name()
	}
	{
		p.SetState(31)
		p.Match(VariablesL_BLOCK)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(32)
		p.Match(VariablesTEXT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(33)
		p.Match(VariablesR_BLOCK)
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

// INumber_var_declarationContext is an interface to support dynamic dispatch.
type INumber_var_declarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	NUMBER_TYPE() antlr.TerminalNode
	Var_name() IVar_nameContext
	L_BLOCK() antlr.TerminalNode
	NUMBER() antlr.TerminalNode
	R_BLOCK() antlr.TerminalNode

	// IsNumber_var_declarationContext differentiates from other interfaces.
	IsNumber_var_declarationContext()
}

type Number_var_declarationContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNumber_var_declarationContext() *Number_var_declarationContext {
	var p = new(Number_var_declarationContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VariablesRULE_number_var_declaration
	return p
}

func InitEmptyNumber_var_declarationContext(p *Number_var_declarationContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VariablesRULE_number_var_declaration
}

func (*Number_var_declarationContext) IsNumber_var_declarationContext() {}

func NewNumber_var_declarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Number_var_declarationContext {
	var p = new(Number_var_declarationContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VariablesRULE_number_var_declaration

	return p
}

func (s *Number_var_declarationContext) GetParser() antlr.Parser { return s.parser }

func (s *Number_var_declarationContext) NUMBER_TYPE() antlr.TerminalNode {
	return s.GetToken(VariablesNUMBER_TYPE, 0)
}

func (s *Number_var_declarationContext) Var_name() IVar_nameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVar_nameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVar_nameContext)
}

func (s *Number_var_declarationContext) L_BLOCK() antlr.TerminalNode {
	return s.GetToken(VariablesL_BLOCK, 0)
}

func (s *Number_var_declarationContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(VariablesNUMBER, 0)
}

func (s *Number_var_declarationContext) R_BLOCK() antlr.TerminalNode {
	return s.GetToken(VariablesR_BLOCK, 0)
}

func (s *Number_var_declarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Number_var_declarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Number_var_declarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VariablesListener); ok {
		listenerT.EnterNumber_var_declaration(s)
	}
}

func (s *Number_var_declarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VariablesListener); ok {
		listenerT.ExitNumber_var_declaration(s)
	}
}

func (p *Variables) Number_var_declaration() (localctx INumber_var_declarationContext) {
	localctx = NewNumber_var_declarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, VariablesRULE_number_var_declaration)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(35)
		p.Match(VariablesNUMBER_TYPE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(36)
		p.Var_name()
	}
	{
		p.SetState(37)
		p.Match(VariablesL_BLOCK)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(38)
		p.Match(VariablesNUMBER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(39)
		p.Match(VariablesR_BLOCK)
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

// IBoolean_var_declarationContext is an interface to support dynamic dispatch.
type IBoolean_var_declarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	BOOLEAN_TYPE() antlr.TerminalNode
	Var_name() IVar_nameContext
	L_BLOCK() antlr.TerminalNode
	BOOLEAN() antlr.TerminalNode
	R_BLOCK() antlr.TerminalNode

	// IsBoolean_var_declarationContext differentiates from other interfaces.
	IsBoolean_var_declarationContext()
}

type Boolean_var_declarationContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBoolean_var_declarationContext() *Boolean_var_declarationContext {
	var p = new(Boolean_var_declarationContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VariablesRULE_boolean_var_declaration
	return p
}

func InitEmptyBoolean_var_declarationContext(p *Boolean_var_declarationContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VariablesRULE_boolean_var_declaration
}

func (*Boolean_var_declarationContext) IsBoolean_var_declarationContext() {}

func NewBoolean_var_declarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Boolean_var_declarationContext {
	var p = new(Boolean_var_declarationContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VariablesRULE_boolean_var_declaration

	return p
}

func (s *Boolean_var_declarationContext) GetParser() antlr.Parser { return s.parser }

func (s *Boolean_var_declarationContext) BOOLEAN_TYPE() antlr.TerminalNode {
	return s.GetToken(VariablesBOOLEAN_TYPE, 0)
}

func (s *Boolean_var_declarationContext) Var_name() IVar_nameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVar_nameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVar_nameContext)
}

func (s *Boolean_var_declarationContext) L_BLOCK() antlr.TerminalNode {
	return s.GetToken(VariablesL_BLOCK, 0)
}

func (s *Boolean_var_declarationContext) BOOLEAN() antlr.TerminalNode {
	return s.GetToken(VariablesBOOLEAN, 0)
}

func (s *Boolean_var_declarationContext) R_BLOCK() antlr.TerminalNode {
	return s.GetToken(VariablesR_BLOCK, 0)
}

func (s *Boolean_var_declarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Boolean_var_declarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Boolean_var_declarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VariablesListener); ok {
		listenerT.EnterBoolean_var_declaration(s)
	}
}

func (s *Boolean_var_declarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VariablesListener); ok {
		listenerT.ExitBoolean_var_declaration(s)
	}
}

func (p *Variables) Boolean_var_declaration() (localctx IBoolean_var_declarationContext) {
	localctx = NewBoolean_var_declarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, VariablesRULE_boolean_var_declaration)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(41)
		p.Match(VariablesBOOLEAN_TYPE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(42)
		p.Var_name()
	}
	{
		p.SetState(43)
		p.Match(VariablesL_BLOCK)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(44)
		p.Match(VariablesBOOLEAN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(45)
		p.Match(VariablesR_BLOCK)
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

// IIpv4_var_declarationContext is an interface to support dynamic dispatch.
type IIpv4_var_declarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IPV4_TYPE() antlr.TerminalNode
	Var_name() IVar_nameContext
	L_BLOCK() antlr.TerminalNode
	IPV4() antlr.TerminalNode
	R_BLOCK() antlr.TerminalNode

	// IsIpv4_var_declarationContext differentiates from other interfaces.
	IsIpv4_var_declarationContext()
}

type Ipv4_var_declarationContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIpv4_var_declarationContext() *Ipv4_var_declarationContext {
	var p = new(Ipv4_var_declarationContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VariablesRULE_ipv4_var_declaration
	return p
}

func InitEmptyIpv4_var_declarationContext(p *Ipv4_var_declarationContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VariablesRULE_ipv4_var_declaration
}

func (*Ipv4_var_declarationContext) IsIpv4_var_declarationContext() {}

func NewIpv4_var_declarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Ipv4_var_declarationContext {
	var p = new(Ipv4_var_declarationContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VariablesRULE_ipv4_var_declaration

	return p
}

func (s *Ipv4_var_declarationContext) GetParser() antlr.Parser { return s.parser }

func (s *Ipv4_var_declarationContext) IPV4_TYPE() antlr.TerminalNode {
	return s.GetToken(VariablesIPV4_TYPE, 0)
}

func (s *Ipv4_var_declarationContext) Var_name() IVar_nameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVar_nameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVar_nameContext)
}

func (s *Ipv4_var_declarationContext) L_BLOCK() antlr.TerminalNode {
	return s.GetToken(VariablesL_BLOCK, 0)
}

func (s *Ipv4_var_declarationContext) IPV4() antlr.TerminalNode {
	return s.GetToken(VariablesIPV4, 0)
}

func (s *Ipv4_var_declarationContext) R_BLOCK() antlr.TerminalNode {
	return s.GetToken(VariablesR_BLOCK, 0)
}

func (s *Ipv4_var_declarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Ipv4_var_declarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Ipv4_var_declarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VariablesListener); ok {
		listenerT.EnterIpv4_var_declaration(s)
	}
}

func (s *Ipv4_var_declarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VariablesListener); ok {
		listenerT.ExitIpv4_var_declaration(s)
	}
}

func (p *Variables) Ipv4_var_declaration() (localctx IIpv4_var_declarationContext) {
	localctx = NewIpv4_var_declarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, VariablesRULE_ipv4_var_declaration)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(47)
		p.Match(VariablesIPV4_TYPE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(48)
		p.Var_name()
	}
	{
		p.SetState(49)
		p.Match(VariablesL_BLOCK)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(50)
		p.Match(VariablesIPV4)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(51)
		p.Match(VariablesR_BLOCK)
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

// IVar_declarationContext is an interface to support dynamic dispatch.
type IVar_declarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Default_var_declaration() IDefault_var_declarationContext
	Text_var_declaration() IText_var_declarationContext
	Number_var_declaration() INumber_var_declarationContext
	Boolean_var_declaration() IBoolean_var_declarationContext
	Ipv4_var_declaration() IIpv4_var_declarationContext

	// IsVar_declarationContext differentiates from other interfaces.
	IsVar_declarationContext()
}

type Var_declarationContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVar_declarationContext() *Var_declarationContext {
	var p = new(Var_declarationContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VariablesRULE_var_declaration
	return p
}

func InitEmptyVar_declarationContext(p *Var_declarationContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VariablesRULE_var_declaration
}

func (*Var_declarationContext) IsVar_declarationContext() {}

func NewVar_declarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Var_declarationContext {
	var p = new(Var_declarationContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VariablesRULE_var_declaration

	return p
}

func (s *Var_declarationContext) GetParser() antlr.Parser { return s.parser }

func (s *Var_declarationContext) Default_var_declaration() IDefault_var_declarationContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDefault_var_declarationContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDefault_var_declarationContext)
}

func (s *Var_declarationContext) Text_var_declaration() IText_var_declarationContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IText_var_declarationContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IText_var_declarationContext)
}

func (s *Var_declarationContext) Number_var_declaration() INumber_var_declarationContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INumber_var_declarationContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INumber_var_declarationContext)
}

func (s *Var_declarationContext) Boolean_var_declaration() IBoolean_var_declarationContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBoolean_var_declarationContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBoolean_var_declarationContext)
}

func (s *Var_declarationContext) Ipv4_var_declaration() IIpv4_var_declarationContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIpv4_var_declarationContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIpv4_var_declarationContext)
}

func (s *Var_declarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Var_declarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Var_declarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VariablesListener); ok {
		listenerT.EnterVar_declaration(s)
	}
}

func (s *Var_declarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VariablesListener); ok {
		listenerT.ExitVar_declaration(s)
	}
}

func (p *Variables) Var_declaration() (localctx IVar_declarationContext) {
	localctx = NewVar_declarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, VariablesRULE_var_declaration)
	p.SetState(58)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 0, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(53)
			p.Default_var_declaration()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(54)
			p.Text_var_declaration()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(55)
			p.Number_var_declaration()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(56)
			p.Boolean_var_declaration()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(57)
			p.Ipv4_var_declaration()
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
