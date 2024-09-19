// Generated from /home/ankan/Documents/git/me/9tail/client/antlr/NinetailParser.g4 by ANTLR 4.13.2

import {ParseTreeVisitor} from 'antlr4';


import { ProgramContext } from "./NinetailParser.js";
import { DeclarationContext } from "./NinetailParser.js";
import { MetaDeclarationContext } from "./NinetailParser.js";


/**
 * This interface defines a complete generic visitor for a parse tree produced
 * by `NinetailParser`.
 *
 * @param <Result> The return type of the visit operation. Use `void` for
 * operations with no return type.
 */
export default class inetailParserVisitor<Result> extends ParseTreeVisitor<Result> {
	/**
	 * Visit a parse tree produced by `NinetailParser.program`.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	visitProgram?: (ctx: ProgramContext) => Result;
	/**
	 * Visit a parse tree produced by `NinetailParser.declaration`.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	visitDeclaration?: (ctx: DeclarationContext) => Result;
	/**
	 * Visit a parse tree produced by `NinetailParser.metaDeclaration`.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	visitMetaDeclaration?: (ctx: MetaDeclarationContext) => Result;
}

