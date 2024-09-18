package parser

import (
	"fmt"

	"github.com/antlr4-go/antlr/v4"
)

// Define the struct that will implement the NinetailParserVisitor interface
type MyNinetailVisitor struct {
	NinetailParserVisitor
}

// Implement the VisitProgram method
func (v *MyNinetailVisitor) VisitProgram(ctx *ProgramContext) interface{} {
	if ctx == nil {
		return nil
	}
	// Custom logic for visiting the ProgramContext here
	fmt.Println("Visiting Program")
	return v.VisitChildren(ctx) // Visits the children of the node
}

// Implement the VisitDeclaration method
func (v *MyNinetailVisitor) VisitDeclaration(ctx *DeclarationContext) interface{} {
	if ctx == nil {
		return nil
	}
	// Custom logic for visiting DeclarationContext
	fmt.Println("Visiting Declaration")
	return v.VisitChildren(ctx)
}

// Implement the VisitMetaDeclaration method
func (v *MyNinetailVisitor) VisitMetaDeclaration(ctx *MetaDeclarationContext) interface{} {
	if ctx == nil {
		return nil
	}
	// Custom logic for visiting MetaDeclarationContext
	fmt.Println("Visiting MetaDeclaration")
	return v.VisitChildren(ctx)
}

func Parse(input *antlr.InputStream) {
	// Create the lexer and token stream
	lexer := NewNinetailLexer(input)
	tokens := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	// Print tokens for debugging purposes
	tokens.Fill()
	fmt.Println("Total token", tokens.Size())
	for i := 0; i < tokens.Size(); i++ {
		fmt.Println("Tokens:", tokens.Get(i))
	}

	// Create the parser
	parser := NewNinetailParser(tokens)
	parser.RemoveErrorListeners()
	parser.AddErrorListener(antlr.NewDiagnosticErrorListener(true))

	// Parse the input starting from the 'program' rule
	programContext := parser.Program()

	// Perform a type assertion to convert IProgramContext to *ProgramContext
	concreteProgramCtx, ok := programContext.(*ProgramContext)
	if !ok || concreteProgramCtx == nil {
		fmt.Println("Error: Invalid program context")
		return
	}

	// Create the visitor and visit the parsed program
	visitor := &MyNinetailVisitor{}
	visitor.VisitProgram(concreteProgramCtx)
}
