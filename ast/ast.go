package ast

import (
	"monkey/token"
)

type Node interface {
	TokenLiteral() string
}
type Statement interface {
	Node
	statementNode()
}
type Expression interface {
	Node
	expressionNode()
}

// Program node is going to be the root node of AST that holds all the statements in the program
type Program struct {
	Statements []Statement
}

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}

// LetStatement node is going to represent statements like let x = 5;
type LetStatement struct {
	Token token.Token // token.LET
	Name  *Identifier
	Value Expression
}

// LetStatement implements these 2 methods because Satement has them
func (ls *LetStatement) statementNode()       {}
func (ls *LetStatement) TokenLiteral() string { return ls.Token.Literal }

// Identifier node is going to represent identifiers like x, y, add, result, etc.
type Identifier struct {
	Token token.Token // the token.IDENT token
	Value string
}

func (i *Identifier) expressionNode()      {}
func (i *Identifier) TokenLiteral() string { return i.Token.Literal }
