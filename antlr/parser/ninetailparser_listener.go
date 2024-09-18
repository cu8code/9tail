// Code generated from /home/ankan/Documents/git/me/9tail/antlr/NinetailParser.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // NinetailParser

import "github.com/antlr4-go/antlr/v4"

// NinetailParserListener is a complete listener for a parse tree produced by NinetailParser.
type NinetailParserListener interface {
	antlr.ParseTreeListener

	// EnterProgram is called when entering the program production.
	EnterProgram(c *ProgramContext)

	// EnterDeclaration is called when entering the declaration production.
	EnterDeclaration(c *DeclarationContext)

	// EnterMetaDeclaration is called when entering the metaDeclaration production.
	EnterMetaDeclaration(c *MetaDeclarationContext)

	// ExitProgram is called when exiting the program production.
	ExitProgram(c *ProgramContext)

	// ExitDeclaration is called when exiting the declaration production.
	ExitDeclaration(c *DeclarationContext)

	// ExitMetaDeclaration is called when exiting the metaDeclaration production.
	ExitMetaDeclaration(c *MetaDeclarationContext)
}
