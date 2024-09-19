import { CharStreams, CommonTokenStream } from "antlr4";
import NinetailLexer from "./NinetailLexer";
import NinetailParser, { DeclarationContext, ProgramContext } from "./NinetailParser";
import NinetailParserVisitor from "./NinetailParserVisitor";

export function createHTMLfrom9Tail(input: string): string {
    try {
        const chars = CharStreams.fromString(input);
        const lexer = new NinetailLexer(chars);
        const tokens = new CommonTokenStream(lexer);
        const parser = new NinetailParser(tokens);
        parser.removeErrorListeners();
        const tree = parser.program();

        class Visitor extends NinetailParserVisitor<string> {
            visitChildren(node: ProgramContext): string {
                let fields: string[] = [];
                if (node.children)
                    for (const i of node.children) {
                        if (i instanceof DeclarationContext) {
                            try {
                                const name = i.VAR().getText();
                                const type = i.TYPE().getText();
                                if (i.INPUT().getText()) {
                                    fields.push(`<div><label>${name}</label><input name='${name}' type='${type}' /></div>`)
                                }
                            } catch (e) {}
                        }
                    }
                if (fields.length !== 0) {
                    return (`<form>${fields.join('\n')}</form>`)
                }
                return '';
            }
        }

        const visitor = new Visitor();
        return tree.accept(visitor);
    } catch (error) {
        console.log("Error parsing input:", error);
        return "";
    }
}