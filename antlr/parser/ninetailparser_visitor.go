// Code generated from /home/ankan/Documents/git/me/9tail/antlr/NinetailParser.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // NinetailParser

import "github.com/antlr4-go/antlr/v4"

// A complete Visitor for a parse tree produced by NinetailParser.
type NinetailParserVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by NinetailParser#program.
	VisitProgram(ctx *ProgramContext) interface{}

	// Visit a parse tree produced by NinetailParser#declaration.
	VisitDeclaration(ctx *DeclarationContext) interface{}

	// Visit a parse tree produced by NinetailParser#metaDeclaration.
	VisitMetaDeclaration(ctx *MetaDeclarationContext) interface{}
}
