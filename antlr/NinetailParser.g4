parser grammar NinetailParser;

options { tokenVocab=NinetailLexer; }

program
    : (declaration | metaDeclaration)*
    ;

declaration
    : (INPUT | OUTPUT) VAR TYPE  // Matches @input or @output with variable name and type
    ;

metaDeclaration
    : META VAR ASSIGN STRING  // Matches meta annotations like @meta name = "fetch"
    ;
