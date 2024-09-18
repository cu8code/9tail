// Code generated from /home/ankan/Documents/git/me/9tail/antlr/NinetailLexer.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser

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

type NinetailLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var NinetailLexerLexerStaticData struct {
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

func ninetaillexerLexerInit() {
	staticData := &NinetailLexerLexerStaticData
	staticData.ChannelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.ModeNames = []string{
		"DEFAULT_MODE",
	}
	staticData.LiteralNames = []string{
		"", "'@input'", "'@output'", "'@meta'", "", "", "'='",
	}
	staticData.SymbolicNames = []string{
		"", "INPUT", "OUTPUT", "META", "TYPE", "VAR", "ASSIGN", "STRING", "WS",
	}
	staticData.RuleNames = []string{
		"INPUT", "OUTPUT", "META", "TYPE", "VAR", "ASSIGN", "STRING", "WS",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 8, 82, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 1, 0, 1, 0, 1, 0, 1, 0, 1,
		0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 1,
		2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1,
		3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 3, 3, 54, 8, 3, 1, 4, 1, 4,
		5, 4, 58, 8, 4, 10, 4, 12, 4, 61, 9, 4, 1, 5, 1, 5, 1, 6, 1, 6, 1, 6, 1,
		6, 5, 6, 69, 8, 6, 10, 6, 12, 6, 72, 9, 6, 1, 6, 1, 6, 1, 7, 4, 7, 77,
		8, 7, 11, 7, 12, 7, 78, 1, 7, 1, 7, 0, 0, 8, 1, 1, 3, 2, 5, 3, 7, 4, 9,
		5, 11, 6, 13, 7, 15, 8, 1, 0, 5, 3, 0, 65, 90, 95, 95, 97, 122, 4, 0, 48,
		57, 65, 90, 95, 95, 97, 122, 2, 0, 34, 34, 92, 92, 7, 0, 34, 34, 92, 92,
		98, 98, 102, 102, 110, 110, 114, 114, 116, 116, 3, 0, 9, 10, 13, 13, 32,
		32, 87, 0, 1, 1, 0, 0, 0, 0, 3, 1, 0, 0, 0, 0, 5, 1, 0, 0, 0, 0, 7, 1,
		0, 0, 0, 0, 9, 1, 0, 0, 0, 0, 11, 1, 0, 0, 0, 0, 13, 1, 0, 0, 0, 0, 15,
		1, 0, 0, 0, 1, 17, 1, 0, 0, 0, 3, 24, 1, 0, 0, 0, 5, 32, 1, 0, 0, 0, 7,
		53, 1, 0, 0, 0, 9, 55, 1, 0, 0, 0, 11, 62, 1, 0, 0, 0, 13, 64, 1, 0, 0,
		0, 15, 76, 1, 0, 0, 0, 17, 18, 5, 64, 0, 0, 18, 19, 5, 105, 0, 0, 19, 20,
		5, 110, 0, 0, 20, 21, 5, 112, 0, 0, 21, 22, 5, 117, 0, 0, 22, 23, 5, 116,
		0, 0, 23, 2, 1, 0, 0, 0, 24, 25, 5, 64, 0, 0, 25, 26, 5, 111, 0, 0, 26,
		27, 5, 117, 0, 0, 27, 28, 5, 116, 0, 0, 28, 29, 5, 112, 0, 0, 29, 30, 5,
		117, 0, 0, 30, 31, 5, 116, 0, 0, 31, 4, 1, 0, 0, 0, 32, 33, 5, 64, 0, 0,
		33, 34, 5, 109, 0, 0, 34, 35, 5, 101, 0, 0, 35, 36, 5, 116, 0, 0, 36, 37,
		5, 97, 0, 0, 37, 6, 1, 0, 0, 0, 38, 39, 5, 115, 0, 0, 39, 40, 5, 116, 0,
		0, 40, 41, 5, 114, 0, 0, 41, 42, 5, 105, 0, 0, 42, 43, 5, 110, 0, 0, 43,
		54, 5, 103, 0, 0, 44, 45, 5, 110, 0, 0, 45, 46, 5, 117, 0, 0, 46, 47, 5,
		109, 0, 0, 47, 48, 5, 98, 0, 0, 48, 49, 5, 101, 0, 0, 49, 54, 5, 114, 0,
		0, 50, 51, 5, 109, 0, 0, 51, 52, 5, 97, 0, 0, 52, 54, 5, 112, 0, 0, 53,
		38, 1, 0, 0, 0, 53, 44, 1, 0, 0, 0, 53, 50, 1, 0, 0, 0, 54, 8, 1, 0, 0,
		0, 55, 59, 7, 0, 0, 0, 56, 58, 7, 1, 0, 0, 57, 56, 1, 0, 0, 0, 58, 61,
		1, 0, 0, 0, 59, 57, 1, 0, 0, 0, 59, 60, 1, 0, 0, 0, 60, 10, 1, 0, 0, 0,
		61, 59, 1, 0, 0, 0, 62, 63, 5, 61, 0, 0, 63, 12, 1, 0, 0, 0, 64, 70, 5,
		34, 0, 0, 65, 69, 8, 2, 0, 0, 66, 67, 5, 92, 0, 0, 67, 69, 7, 3, 0, 0,
		68, 65, 1, 0, 0, 0, 68, 66, 1, 0, 0, 0, 69, 72, 1, 0, 0, 0, 70, 68, 1,
		0, 0, 0, 70, 71, 1, 0, 0, 0, 71, 73, 1, 0, 0, 0, 72, 70, 1, 0, 0, 0, 73,
		74, 5, 34, 0, 0, 74, 14, 1, 0, 0, 0, 75, 77, 7, 4, 0, 0, 76, 75, 1, 0,
		0, 0, 77, 78, 1, 0, 0, 0, 78, 76, 1, 0, 0, 0, 78, 79, 1, 0, 0, 0, 79, 80,
		1, 0, 0, 0, 80, 81, 6, 7, 0, 0, 81, 16, 1, 0, 0, 0, 6, 0, 53, 59, 68, 70,
		78, 1, 6, 0, 0,
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

// NinetailLexerInit initializes any static state used to implement NinetailLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewNinetailLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func NinetailLexerInit() {
	staticData := &NinetailLexerLexerStaticData
	staticData.once.Do(ninetaillexerLexerInit)
}

// NewNinetailLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewNinetailLexer(input antlr.CharStream) *NinetailLexer {
	NinetailLexerInit()
	l := new(NinetailLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &NinetailLexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	l.channelNames = staticData.ChannelNames
	l.modeNames = staticData.ModeNames
	l.RuleNames = staticData.RuleNames
	l.LiteralNames = staticData.LiteralNames
	l.SymbolicNames = staticData.SymbolicNames
	l.GrammarFileName = "NinetailLexer.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// NinetailLexer tokens.
const (
	NinetailLexerINPUT  = 1
	NinetailLexerOUTPUT = 2
	NinetailLexerMETA   = 3
	NinetailLexerTYPE   = 4
	NinetailLexerVAR    = 5
	NinetailLexerASSIGN = 6
	NinetailLexerSTRING = 7
	NinetailLexerWS     = 8
)
