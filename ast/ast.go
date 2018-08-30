package ast

import "monkey_interpreter/token"

// Node interface to represent node
type Node interface {
	TokenLiteral() string
}

// Statement interface to represent statements
type Statement interface {
	TokenLiteral() string
	statementNode()
}

// Expression interface to represent expressions
type Expression interface {
	Node
	expressionNode()
}

// Program struct that represents the root node of every AST
type Program struct {
	Statements []Statement
}

// TokenLiteral implements the statement interface for Program
func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	}
	return ""
}

// LetStatement structure
type LetStatement struct {
	Token token.Token // the token.LET token
	Name  *Identifier
	Value Expression
}

func (ls *LetStatement) statementNode() {}

// TokenLiteral and statementNode implement Statement interface for LetStatement
func (ls *LetStatement) TokenLiteral() string {
	return ls.Token.Literal
}

// Identifier struct used to track identifiers
type Identifier struct {
	Token token.Token // the token.IDENT token
	Value string
}

func (i *Identifier) expressionNode() {}

// TokenLiteral and expressionNode implement Expression interface for Identifier
func (i *Identifier) TokenLiteral() string {
	return i.Token.Literal
}
