package parser

import (
	"monkey_interpreter/ast"
	"monkey_interpreter/lexer"
	"monkey_interpreter/token"
)

// Parser struct
type Parser struct {
	l *lexer.Lexer

	curToken  token.Token
	peekToken token.Token
}

// New creates a new Parser object
func New(l *lexer.Lexer) *Parser {
	p := &Parser{l: l}

	// Read two tokens, so curToken and peekToken are both set
	p.nextToken()
	p.nextToken()

	return p
}

func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

// ParseProgram does exactly what it says on the label
func (p *Parser) ParseProgram() *ast.Program {
	return nil
}
