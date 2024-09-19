// Generated from /home/ankan/Documents/git/me/9tail/client/antlr/NinetailLexer.g4 by ANTLR 4.13.2
// noinspection ES6UnusedImports,JSUnusedGlobalSymbols,JSUnusedLocalSymbols
import {
	ATN,
	ATNDeserializer,
	CharStream,
	DecisionState, DFA,
	Lexer,
	LexerATNSimulator,
	RuleContext,
	PredictionContextCache,
	Token
} from "antlr4";
export default class NinetailLexer extends Lexer {
	public static readonly INPUT = 1;
	public static readonly OUTPUT = 2;
	public static readonly META = 3;
	public static readonly TYPE = 4;
	public static readonly VAR = 5;
	public static readonly ASSIGN = 6;
	public static readonly STRING = 7;
	public static readonly WS = 8;
	public static readonly EOF = Token.EOF;

	public static readonly channelNames: string[] = [ "DEFAULT_TOKEN_CHANNEL", "HIDDEN" ];
	public static readonly literalNames: (string | null)[] = [ null, "'@input'", 
                                                            "'@output'", 
                                                            "'@meta'", null, 
                                                            null, "'='" ];
	public static readonly symbolicNames: (string | null)[] = [ null, "INPUT", 
                                                             "OUTPUT", "META", 
                                                             "TYPE", "VAR", 
                                                             "ASSIGN", "STRING", 
                                                             "WS" ];
	public static readonly modeNames: string[] = [ "DEFAULT_MODE", ];

	public static readonly ruleNames: string[] = [
		"INPUT", "OUTPUT", "META", "TYPE", "VAR", "ASSIGN", "STRING", "WS",
	];


	constructor(input: CharStream) {
		super(input);
		this._interp = new LexerATNSimulator(this, NinetailLexer._ATN, NinetailLexer.DecisionsToDFA, new PredictionContextCache());
	}

	public get grammarFileName(): string { return "NinetailLexer.g4"; }

	public get literalNames(): (string | null)[] { return NinetailLexer.literalNames; }
	public get symbolicNames(): (string | null)[] { return NinetailLexer.symbolicNames; }
	public get ruleNames(): string[] { return NinetailLexer.ruleNames; }

	public get serializedATN(): number[] { return NinetailLexer._serializedATN; }

	public get channelNames(): string[] { return NinetailLexer.channelNames; }

	public get modeNames(): string[] { return NinetailLexer.modeNames; }

	public static readonly _serializedATN: number[] = [4,0,8,82,6,-1,2,0,7,
	0,2,1,7,1,2,2,7,2,2,3,7,3,2,4,7,4,2,5,7,5,2,6,7,6,2,7,7,7,1,0,1,0,1,0,1,
	0,1,0,1,0,1,0,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,2,1,2,1,2,1,2,1,2,1,2,1,
	3,1,3,1,3,1,3,1,3,1,3,1,3,1,3,1,3,1,3,1,3,1,3,1,3,1,3,1,3,3,3,54,8,3,1,
	4,1,4,5,4,58,8,4,10,4,12,4,61,9,4,1,5,1,5,1,6,1,6,1,6,1,6,5,6,69,8,6,10,
	6,12,6,72,9,6,1,6,1,6,1,7,4,7,77,8,7,11,7,12,7,78,1,7,1,7,0,0,8,1,1,3,2,
	5,3,7,4,9,5,11,6,13,7,15,8,1,0,5,3,0,65,90,95,95,97,122,4,0,48,57,65,90,
	95,95,97,122,2,0,34,34,92,92,7,0,34,34,92,92,98,98,102,102,110,110,114,
	114,116,116,3,0,9,10,13,13,32,32,87,0,1,1,0,0,0,0,3,1,0,0,0,0,5,1,0,0,0,
	0,7,1,0,0,0,0,9,1,0,0,0,0,11,1,0,0,0,0,13,1,0,0,0,0,15,1,0,0,0,1,17,1,0,
	0,0,3,24,1,0,0,0,5,32,1,0,0,0,7,53,1,0,0,0,9,55,1,0,0,0,11,62,1,0,0,0,13,
	64,1,0,0,0,15,76,1,0,0,0,17,18,5,64,0,0,18,19,5,105,0,0,19,20,5,110,0,0,
	20,21,5,112,0,0,21,22,5,117,0,0,22,23,5,116,0,0,23,2,1,0,0,0,24,25,5,64,
	0,0,25,26,5,111,0,0,26,27,5,117,0,0,27,28,5,116,0,0,28,29,5,112,0,0,29,
	30,5,117,0,0,30,31,5,116,0,0,31,4,1,0,0,0,32,33,5,64,0,0,33,34,5,109,0,
	0,34,35,5,101,0,0,35,36,5,116,0,0,36,37,5,97,0,0,37,6,1,0,0,0,38,39,5,115,
	0,0,39,40,5,116,0,0,40,41,5,114,0,0,41,42,5,105,0,0,42,43,5,110,0,0,43,
	54,5,103,0,0,44,45,5,110,0,0,45,46,5,117,0,0,46,47,5,109,0,0,47,48,5,98,
	0,0,48,49,5,101,0,0,49,54,5,114,0,0,50,51,5,109,0,0,51,52,5,97,0,0,52,54,
	5,112,0,0,53,38,1,0,0,0,53,44,1,0,0,0,53,50,1,0,0,0,54,8,1,0,0,0,55,59,
	7,0,0,0,56,58,7,1,0,0,57,56,1,0,0,0,58,61,1,0,0,0,59,57,1,0,0,0,59,60,1,
	0,0,0,60,10,1,0,0,0,61,59,1,0,0,0,62,63,5,61,0,0,63,12,1,0,0,0,64,70,5,
	34,0,0,65,69,8,2,0,0,66,67,5,92,0,0,67,69,7,3,0,0,68,65,1,0,0,0,68,66,1,
	0,0,0,69,72,1,0,0,0,70,68,1,0,0,0,70,71,1,0,0,0,71,73,1,0,0,0,72,70,1,0,
	0,0,73,74,5,34,0,0,74,14,1,0,0,0,75,77,7,4,0,0,76,75,1,0,0,0,77,78,1,0,
	0,0,78,76,1,0,0,0,78,79,1,0,0,0,79,80,1,0,0,0,80,81,6,7,0,0,81,16,1,0,0,
	0,6,0,53,59,68,70,78,1,6,0,0];

	private static __ATN: ATN;
	public static get _ATN(): ATN {
		if (!NinetailLexer.__ATN) {
			NinetailLexer.__ATN = new ATNDeserializer().deserialize(NinetailLexer._serializedATN);
		}

		return NinetailLexer.__ATN;
	}


	static DecisionsToDFA = NinetailLexer._ATN.decisionToState.map( (ds: DecisionState, index: number) => new DFA(ds, index) );
}