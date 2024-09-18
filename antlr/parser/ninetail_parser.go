// Code generated from /home/ankan/Documents/git/me/9tail/antlr/NinetailParser.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // NinetailParser

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

type NinetailParser struct {
	*antlr.BaseParser
}

var NinetailParserParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func ninetailparserParserInit() {
	staticData := &NinetailParserParserStaticData
	staticData.LiteralNames = []string{
		"", "'@input'", "'@output'", "'@meta'", "", "", "'='",
	}
	staticData.SymbolicNames = []string{
		"", "INPUT", "OUTPUT", "META", "TYPE", "VAR", "ASSIGN", "STRING", "WS",
	}
	staticData.RuleNames = []string{
		"program", "declaration", "metaDeclaration",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 8, 23, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 1, 0, 1, 0, 5, 0, 9, 8,
		0, 10, 0, 12, 0, 12, 9, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 1, 2, 1, 2, 1,
		2, 1, 2, 1, 2, 0, 0, 3, 0, 2, 4, 0, 1, 1, 0, 1, 2, 21, 0, 10, 1, 0, 0,
		0, 2, 13, 1, 0, 0, 0, 4, 17, 1, 0, 0, 0, 6, 9, 3, 2, 1, 0, 7, 9, 3, 4,
		2, 0, 8, 6, 1, 0, 0, 0, 8, 7, 1, 0, 0, 0, 9, 12, 1, 0, 0, 0, 10, 8, 1,
		0, 0, 0, 10, 11, 1, 0, 0, 0, 11, 1, 1, 0, 0, 0, 12, 10, 1, 0, 0, 0, 13,
		14, 7, 0, 0, 0, 14, 15, 5, 5, 0, 0, 15, 16, 5, 4, 0, 0, 16, 3, 1, 0, 0,
		0, 17, 18, 5, 3, 0, 0, 18, 19, 5, 5, 0, 0, 19, 20, 5, 6, 0, 0, 20, 21,
		5, 7, 0, 0, 21, 5, 1, 0, 0, 0, 2, 8, 10,
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

// NinetailParserInit initializes any static state used to implement NinetailParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewNinetailParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func NinetailParserInit() {
	staticData := &NinetailParserParserStaticData
	staticData.once.Do(ninetailparserParserInit)
}

// NewNinetailParser produces a new parser instance for the optional input antlr.TokenStream.
func NewNinetailParser(input antlr.TokenStream) *NinetailParser {
	NinetailParserInit()
	this := new(NinetailParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &NinetailParserParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "NinetailParser.g4"

	return this
}

// NinetailParser tokens.
const (
	NinetailParserEOF    = antlr.TokenEOF
	NinetailParserINPUT  = 1
	NinetailParserOUTPUT = 2
	NinetailParserMETA   = 3
	NinetailParserTYPE   = 4
	NinetailParserVAR    = 5
	NinetailParserASSIGN = 6
	NinetailParserSTRING = 7
	NinetailParserWS     = 8
)

// NinetailParser rules.
const (
	NinetailParserRULE_program         = 0
	NinetailParserRULE_declaration     = 1
	NinetailParserRULE_metaDeclaration = 2
)

// IProgramContext is an interface to support dynamic dispatch.
type IProgramContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllDeclaration() []IDeclarationContext
	Declaration(i int) IDeclarationContext
	AllMetaDeclaration() []IMetaDeclarationContext
	MetaDeclaration(i int) IMetaDeclarationContext

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
	p.RuleIndex = NinetailParserRULE_program
	return p
}

func InitEmptyProgramContext(p *ProgramContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = NinetailParserRULE_program
}

func (*ProgramContext) IsProgramContext() {}

func NewProgramContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProgramContext {
	var p = new(ProgramContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = NinetailParserRULE_program

	return p
}

func (s *ProgramContext) GetParser() antlr.Parser { return s.parser }

func (s *ProgramContext) AllDeclaration() []IDeclarationContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IDeclarationContext); ok {
			len++
		}
	}

	tst := make([]IDeclarationContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IDeclarationContext); ok {
			tst[i] = t.(IDeclarationContext)
			i++
		}
	}

	return tst
}

func (s *ProgramContext) Declaration(i int) IDeclarationContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDeclarationContext); ok {
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

	return t.(IDeclarationContext)
}

func (s *ProgramContext) AllMetaDeclaration() []IMetaDeclarationContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IMetaDeclarationContext); ok {
			len++
		}
	}

	tst := make([]IMetaDeclarationContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IMetaDeclarationContext); ok {
			tst[i] = t.(IMetaDeclarationContext)
			i++
		}
	}

	return tst
}

func (s *ProgramContext) MetaDeclaration(i int) IMetaDeclarationContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMetaDeclarationContext); ok {
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

	return t.(IMetaDeclarationContext)
}

func (s *ProgramContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProgramContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProgramContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NinetailParserListener); ok {
		listenerT.EnterProgram(s)
	}
}

func (s *ProgramContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NinetailParserListener); ok {
		listenerT.ExitProgram(s)
	}
}

func (s *ProgramContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case NinetailParserVisitor:
		return t.VisitProgram(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *NinetailParser) Program() (localctx IProgramContext) {
	localctx = NewProgramContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, NinetailParserRULE_program)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(10)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&14) != 0 {
		p.SetState(8)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}

		switch p.GetTokenStream().LA(1) {
		case NinetailParserINPUT, NinetailParserOUTPUT:
			{
				p.SetState(6)
				p.Declaration()
			}

		case NinetailParserMETA:
			{
				p.SetState(7)
				p.MetaDeclaration()
			}

		default:
			p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			goto errorExit
		}

		p.SetState(12)
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

// IDeclarationContext is an interface to support dynamic dispatch.
type IDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	VAR() antlr.TerminalNode
	TYPE() antlr.TerminalNode
	INPUT() antlr.TerminalNode
	OUTPUT() antlr.TerminalNode

	// IsDeclarationContext differentiates from other interfaces.
	IsDeclarationContext()
}

type DeclarationContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDeclarationContext() *DeclarationContext {
	var p = new(DeclarationContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = NinetailParserRULE_declaration
	return p
}

func InitEmptyDeclarationContext(p *DeclarationContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = NinetailParserRULE_declaration
}

func (*DeclarationContext) IsDeclarationContext() {}

func NewDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DeclarationContext {
	var p = new(DeclarationContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = NinetailParserRULE_declaration

	return p
}

func (s *DeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *DeclarationContext) VAR() antlr.TerminalNode {
	return s.GetToken(NinetailParserVAR, 0)
}

func (s *DeclarationContext) TYPE() antlr.TerminalNode {
	return s.GetToken(NinetailParserTYPE, 0)
}

func (s *DeclarationContext) INPUT() antlr.TerminalNode {
	return s.GetToken(NinetailParserINPUT, 0)
}

func (s *DeclarationContext) OUTPUT() antlr.TerminalNode {
	return s.GetToken(NinetailParserOUTPUT, 0)
}

func (s *DeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NinetailParserListener); ok {
		listenerT.EnterDeclaration(s)
	}
}

func (s *DeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NinetailParserListener); ok {
		listenerT.ExitDeclaration(s)
	}
}

func (s *DeclarationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case NinetailParserVisitor:
		return t.VisitDeclaration(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *NinetailParser) Declaration() (localctx IDeclarationContext) {
	localctx = NewDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, NinetailParserRULE_declaration)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(13)
		_la = p.GetTokenStream().LA(1)

		if !(_la == NinetailParserINPUT || _la == NinetailParserOUTPUT) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	{
		p.SetState(14)
		p.Match(NinetailParserVAR)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(15)
		p.Match(NinetailParserTYPE)
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

// IMetaDeclarationContext is an interface to support dynamic dispatch.
type IMetaDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	META() antlr.TerminalNode
	VAR() antlr.TerminalNode
	ASSIGN() antlr.TerminalNode
	STRING() antlr.TerminalNode

	// IsMetaDeclarationContext differentiates from other interfaces.
	IsMetaDeclarationContext()
}

type MetaDeclarationContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMetaDeclarationContext() *MetaDeclarationContext {
	var p = new(MetaDeclarationContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = NinetailParserRULE_metaDeclaration
	return p
}

func InitEmptyMetaDeclarationContext(p *MetaDeclarationContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = NinetailParserRULE_metaDeclaration
}

func (*MetaDeclarationContext) IsMetaDeclarationContext() {}

func NewMetaDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MetaDeclarationContext {
	var p = new(MetaDeclarationContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = NinetailParserRULE_metaDeclaration

	return p
}

func (s *MetaDeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *MetaDeclarationContext) META() antlr.TerminalNode {
	return s.GetToken(NinetailParserMETA, 0)
}

func (s *MetaDeclarationContext) VAR() antlr.TerminalNode {
	return s.GetToken(NinetailParserVAR, 0)
}

func (s *MetaDeclarationContext) ASSIGN() antlr.TerminalNode {
	return s.GetToken(NinetailParserASSIGN, 0)
}

func (s *MetaDeclarationContext) STRING() antlr.TerminalNode {
	return s.GetToken(NinetailParserSTRING, 0)
}

func (s *MetaDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MetaDeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MetaDeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NinetailParserListener); ok {
		listenerT.EnterMetaDeclaration(s)
	}
}

func (s *MetaDeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NinetailParserListener); ok {
		listenerT.ExitMetaDeclaration(s)
	}
}

func (s *MetaDeclarationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case NinetailParserVisitor:
		return t.VisitMetaDeclaration(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *NinetailParser) MetaDeclaration() (localctx IMetaDeclarationContext) {
	localctx = NewMetaDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, NinetailParserRULE_metaDeclaration)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(17)
		p.Match(NinetailParserMETA)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(18)
		p.Match(NinetailParserVAR)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(19)
		p.Match(NinetailParserASSIGN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(20)
		p.Match(NinetailParserSTRING)
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
