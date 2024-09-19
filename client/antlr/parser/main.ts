import { CharStreams, CommonTokenStream } from "antlr4";
import NinetailLexer from "./NinetailLexer";
import NinetailParser, { DeclarationContext, ProgramContext } from "./NinetailParser";
import NinetailParserVisitor from "./NinetailParserVisitor";

export function createHTMLfrom9Tail(input: string): string {
    const chars = CharStreams.fromString(input);
    const lexer = new NinetailLexer(chars);
    const tokens = new CommonTokenStream(lexer);
    const parser = new NinetailParser(tokens);
    const tree = parser.program();

    class Visitor extends NinetailParserVisitor<string> {
        visitChildren(node: ProgramContext): string {
            let fields: string[] = [];
            if (node.children)
                for (const i of node.children) {
                    if (i instanceof DeclarationContext) {
                        const name = i.VAR().getText();
                        const type = i.TYPE().getText();
                        fields.push(`<input name='${name}' type='${type}' />`)
                    }
                }
            return `<from>
            ${fields.join('\n')}
            </form>`
        }
    }

    const visitor = new Visitor();
    return tree.accept(visitor);
}

const input = "@input name string";
const htmlForm = createHTMLfrom9Tail(input);
console.log(htmlForm);
