// Generated from /home/ankan/Documents/git/me/9tail/client/antlr/NinetailParser.g4 by ANTLR 4.13.2
// noinspection ES6UnusedImports,JSUnusedGlobalSymbols,JSUnusedLocalSymbols

import {
	ATN,
	ATNDeserializer, DecisionState, DFA, FailedPredicateException,
	RecognitionException, NoViableAltException, BailErrorStrategy,
	Parser, ParserATNSimulator,
	RuleContext, ParserRuleContext, PredictionMode, PredictionContextCache,
	TerminalNode, RuleNode,
	Token, TokenStream,
	Interval, IntervalSet
} from 'antlr4';
import NinetailParserListener from "./NinetailParserListener.js";
import NinetailParserVisitor from "./NinetailParserVisitor.js";

// for running tests with parameters, TODO: discuss strategy for typed parameters in CI
// eslint-disable-next-line no-unused-vars
type int = number;

export default class NinetailParser extends Parser {
	public static readonly INPUT = 1;
	public static readonly OUTPUT = 2;
	public static readonly META = 3;
	public static readonly TYPE = 4;
	public static readonly VAR = 5;
	public static readonly ASSIGN = 6;
	public static readonly STRING = 7;
	public static readonly WS = 8;
	public static override readonly EOF = Token.EOF;
	public static readonly RULE_program = 0;
	public static readonly RULE_declaration = 1;
	public static readonly RULE_metaDeclaration = 2;
	public static readonly literalNames: (string | null)[] = [ null, "'@input'", 
                                                            "'@output'", 
                                                            "'@meta'", null, 
                                                            null, "'='" ];
	public static readonly symbolicNames: (string | null)[] = [ null, "INPUT", 
                                                             "OUTPUT", "META", 
                                                             "TYPE", "VAR", 
                                                             "ASSIGN", "STRING", 
                                                             "WS" ];
	// tslint:disable:no-trailing-whitespace
	public static readonly ruleNames: string[] = [
		"program", "declaration", "metaDeclaration",
	];
	public get grammarFileName(): string { return "NinetailParser.g4"; }
	public get literalNames(): (string | null)[] { return NinetailParser.literalNames; }
	public get symbolicNames(): (string | null)[] { return NinetailParser.symbolicNames; }
	public get ruleNames(): string[] { return NinetailParser.ruleNames; }
	public get serializedATN(): number[] { return NinetailParser._serializedATN; }

	protected createFailedPredicateException(predicate?: string, message?: string): FailedPredicateException {
		return new FailedPredicateException(this, predicate, message);
	}

	constructor(input: TokenStream) {
		super(input);
		this._interp = new ParserATNSimulator(this, NinetailParser._ATN, NinetailParser.DecisionsToDFA, new PredictionContextCache());
	}
	// @RuleVersion(0)
	public program(): ProgramContext {
		let localctx: ProgramContext = new ProgramContext(this, this._ctx, this.state);
		this.enterRule(localctx, 0, NinetailParser.RULE_program);
		let _la: number;
		try {
			this.enterOuterAlt(localctx, 1);
			{
			this.state = 10;
			this._errHandler.sync(this);
			_la = this._input.LA(1);
			while ((((_la) & ~0x1F) === 0 && ((1 << _la) & 14) !== 0)) {
				{
				this.state = 8;
				this._errHandler.sync(this);
				switch (this._input.LA(1)) {
				case 1:
				case 2:
					{
					this.state = 6;
					this.declaration();
					}
					break;
				case 3:
					{
					this.state = 7;
					this.metaDeclaration();
					}
					break;
				default:
					throw new NoViableAltException(this);
				}
				}
				this.state = 12;
				this._errHandler.sync(this);
				_la = this._input.LA(1);
			}
			}
		}
		catch (re) {
			if (re instanceof RecognitionException) {
				localctx.exception = re;
				this._errHandler.reportError(this, re);
				this._errHandler.recover(this, re);
			} else {
				throw re;
			}
		}
		finally {
			this.exitRule();
		}
		return localctx;
	}
	// @RuleVersion(0)
	public declaration(): DeclarationContext {
		let localctx: DeclarationContext = new DeclarationContext(this, this._ctx, this.state);
		this.enterRule(localctx, 2, NinetailParser.RULE_declaration);
		let _la: number;
		try {
			this.enterOuterAlt(localctx, 1);
			{
			this.state = 13;
			_la = this._input.LA(1);
			if(!(_la===1 || _la===2)) {
			this._errHandler.recoverInline(this);
			}
			else {
				this._errHandler.reportMatch(this);
			    this.consume();
			}
			this.state = 14;
			this.match(NinetailParser.VAR);
			this.state = 15;
			this.match(NinetailParser.TYPE);
			}
		}
		catch (re) {
			if (re instanceof RecognitionException) {
				localctx.exception = re;
				this._errHandler.reportError(this, re);
				this._errHandler.recover(this, re);
			} else {
				throw re;
			}
		}
		finally {
			this.exitRule();
		}
		return localctx;
	}
	// @RuleVersion(0)
	public metaDeclaration(): MetaDeclarationContext {
		let localctx: MetaDeclarationContext = new MetaDeclarationContext(this, this._ctx, this.state);
		this.enterRule(localctx, 4, NinetailParser.RULE_metaDeclaration);
		try {
			this.enterOuterAlt(localctx, 1);
			{
			this.state = 17;
			this.match(NinetailParser.META);
			this.state = 18;
			this.match(NinetailParser.VAR);
			this.state = 19;
			this.match(NinetailParser.ASSIGN);
			this.state = 20;
			this.match(NinetailParser.STRING);
			}
		}
		catch (re) {
			if (re instanceof RecognitionException) {
				localctx.exception = re;
				this._errHandler.reportError(this, re);
				this._errHandler.recover(this, re);
			} else {
				throw re;
			}
		}
		finally {
			this.exitRule();
		}
		return localctx;
	}

	public static readonly _serializedATN: number[] = [4,1,8,23,2,0,7,0,2,1,
	7,1,2,2,7,2,1,0,1,0,5,0,9,8,0,10,0,12,0,12,9,0,1,1,1,1,1,1,1,1,1,2,1,2,
	1,2,1,2,1,2,1,2,0,0,3,0,2,4,0,1,1,0,1,2,21,0,10,1,0,0,0,2,13,1,0,0,0,4,
	17,1,0,0,0,6,9,3,2,1,0,7,9,3,4,2,0,8,6,1,0,0,0,8,7,1,0,0,0,9,12,1,0,0,0,
	10,8,1,0,0,0,10,11,1,0,0,0,11,1,1,0,0,0,12,10,1,0,0,0,13,14,7,0,0,0,14,
	15,5,5,0,0,15,16,5,4,0,0,16,3,1,0,0,0,17,18,5,3,0,0,18,19,5,5,0,0,19,20,
	5,6,0,0,20,21,5,7,0,0,21,5,1,0,0,0,2,8,10];

	private static __ATN: ATN;
	public static get _ATN(): ATN {
		if (!NinetailParser.__ATN) {
			NinetailParser.__ATN = new ATNDeserializer().deserialize(NinetailParser._serializedATN);
		}

		return NinetailParser.__ATN;
	}


	static DecisionsToDFA = NinetailParser._ATN.decisionToState.map( (ds: DecisionState, index: number) => new DFA(ds, index) );

}

export class ProgramContext extends ParserRuleContext {
	constructor(parser?: NinetailParser, parent?: ParserRuleContext, invokingState?: number) {
		super(parent, invokingState);
    	this.parser = parser;
	}
	public declaration_list(): DeclarationContext[] {
		return this.getTypedRuleContexts(DeclarationContext) as DeclarationContext[];
	}
	public declaration(i: number): DeclarationContext {
		return this.getTypedRuleContext(DeclarationContext, i) as DeclarationContext;
	}
	public metaDeclaration_list(): MetaDeclarationContext[] {
		return this.getTypedRuleContexts(MetaDeclarationContext) as MetaDeclarationContext[];
	}
	public metaDeclaration(i: number): MetaDeclarationContext {
		return this.getTypedRuleContext(MetaDeclarationContext, i) as MetaDeclarationContext;
	}
    public get ruleIndex(): number {
    	return NinetailParser.RULE_program;
	}
	public enterRule(listener: NinetailParserListener): void {
	    if(listener.enterProgram) {
	 		listener.enterProgram(this);
		}
	}
	public exitRule(listener: NinetailParserListener): void {
	    if(listener.exitProgram) {
	 		listener.exitProgram(this);
		}
	}
	// @Override
	public accept<Result>(visitor: NinetailParserVisitor<Result>): Result {
		if (visitor.visitProgram) {
			return visitor.visitProgram(this);
		} else {
			return visitor.visitChildren(this);
		}
	}
}


export class DeclarationContext extends ParserRuleContext {
	constructor(parser?: NinetailParser, parent?: ParserRuleContext, invokingState?: number) {
		super(parent, invokingState);
    	this.parser = parser;
	}
	public VAR(): TerminalNode {
		return this.getToken(NinetailParser.VAR, 0);
	}
	public TYPE(): TerminalNode {
		return this.getToken(NinetailParser.TYPE, 0);
	}
	public INPUT(): TerminalNode {
		return this.getToken(NinetailParser.INPUT, 0);
	}
	public OUTPUT(): TerminalNode {
		return this.getToken(NinetailParser.OUTPUT, 0);
	}
    public get ruleIndex(): number {
    	return NinetailParser.RULE_declaration;
	}
	public enterRule(listener: NinetailParserListener): void {
	    if(listener.enterDeclaration) {
	 		listener.enterDeclaration(this);
		}
	}
	public exitRule(listener: NinetailParserListener): void {
	    if(listener.exitDeclaration) {
	 		listener.exitDeclaration(this);
		}
	}
	// @Override
	public accept<Result>(visitor: NinetailParserVisitor<Result>): Result {
		if (visitor.visitDeclaration) {
			return visitor.visitDeclaration(this);
		} else {
			return visitor.visitChildren(this);
		}
	}
}


export class MetaDeclarationContext extends ParserRuleContext {
	constructor(parser?: NinetailParser, parent?: ParserRuleContext, invokingState?: number) {
		super(parent, invokingState);
    	this.parser = parser;
	}
	public META(): TerminalNode {
		return this.getToken(NinetailParser.META, 0);
	}
	public VAR(): TerminalNode {
		return this.getToken(NinetailParser.VAR, 0);
	}
	public ASSIGN(): TerminalNode {
		return this.getToken(NinetailParser.ASSIGN, 0);
	}
	public STRING(): TerminalNode {
		return this.getToken(NinetailParser.STRING, 0);
	}
    public get ruleIndex(): number {
    	return NinetailParser.RULE_metaDeclaration;
	}
	public enterRule(listener: NinetailParserListener): void {
	    if(listener.enterMetaDeclaration) {
	 		listener.enterMetaDeclaration(this);
		}
	}
	public exitRule(listener: NinetailParserListener): void {
	    if(listener.exitMetaDeclaration) {
	 		listener.exitMetaDeclaration(this);
		}
	}
	// @Override
	public accept<Result>(visitor: NinetailParserVisitor<Result>): Result {
		if (visitor.visitMetaDeclaration) {
			return visitor.visitMetaDeclaration(this);
		} else {
			return visitor.visitChildren(this);
		}
	}
}
