// Code generated from /home/ankan/Documents/git/me/9tail/antlr/NinetailParser.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // NinetailParser

import "github.com/antlr4-go/antlr/v4"

type BaseNinetailParserVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseNinetailParserVisitor) VisitProgram(ctx *ProgramContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseNinetailParserVisitor) VisitDeclaration(ctx *DeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseNinetailParserVisitor) VisitMetaDeclaration(ctx *MetaDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}
