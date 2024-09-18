// Code generated from /home/ankan/Documents/git/me/9tail/antlr/NinetailParser.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // NinetailParser

import "github.com/antlr4-go/antlr/v4"

// BaseNinetailParserListener is a complete listener for a parse tree produced by NinetailParser.
type BaseNinetailParserListener struct{}

var _ NinetailParserListener = &BaseNinetailParserListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseNinetailParserListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseNinetailParserListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseNinetailParserListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseNinetailParserListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterProgram is called when production program is entered.
func (s *BaseNinetailParserListener) EnterProgram(ctx *ProgramContext) {}

// ExitProgram is called when production program is exited.
func (s *BaseNinetailParserListener) ExitProgram(ctx *ProgramContext) {}

// EnterDeclaration is called when production declaration is entered.
func (s *BaseNinetailParserListener) EnterDeclaration(ctx *DeclarationContext) {}

// ExitDeclaration is called when production declaration is exited.
func (s *BaseNinetailParserListener) ExitDeclaration(ctx *DeclarationContext) {}

// EnterMetaDeclaration is called when production metaDeclaration is entered.
func (s *BaseNinetailParserListener) EnterMetaDeclaration(ctx *MetaDeclarationContext) {}

// ExitMetaDeclaration is called when production metaDeclaration is exited.
func (s *BaseNinetailParserListener) ExitMetaDeclaration(ctx *MetaDeclarationContext) {}
