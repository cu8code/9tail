// Generated from /home/ankan/Documents/git/me/9tail/client/antlr/NinetailParser.g4 by ANTLR 4.13.2

import {ParseTreeListener} from "antlr4";


import { ProgramContext } from "./NinetailParser.js";
import { DeclarationContext } from "./NinetailParser.js";
import { MetaDeclarationContext } from "./NinetailParser.js";


/**
 * This interface defines a complete listener for a parse tree produced by
 * `NinetailParser`.
 */
export default class NinetailParserListener extends ParseTreeListener {
	/**
	 * Enter a parse tree produced by `NinetailParser.program`.
	 * @param ctx the parse tree
	 */
	enterProgram?: (ctx: ProgramContext) => void;
	/**
	 * Exit a parse tree produced by `NinetailParser.program`.
	 * @param ctx the parse tree
	 */
	exitProgram?: (ctx: ProgramContext) => void;
	/**
	 * Enter a parse tree produced by `NinetailParser.declaration`.
	 * @param ctx the parse tree
	 */
	enterDeclaration?: (ctx: DeclarationContext) => void;
	/**
	 * Exit a parse tree produced by `NinetailParser.declaration`.
	 * @param ctx the parse tree
	 */
	exitDeclaration?: (ctx: DeclarationContext) => void;
	/**
	 * Enter a parse tree produced by `NinetailParser.metaDeclaration`.
	 * @param ctx the parse tree
	 */
	enterMetaDeclaration?: (ctx: MetaDeclarationContext) => void;
	/**
	 * Exit a parse tree produced by `NinetailParser.metaDeclaration`.
	 * @param ctx the parse tree
	 */
	exitMetaDeclaration?: (ctx: MetaDeclarationContext) => void;
}

