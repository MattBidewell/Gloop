package ast

import "github.com/mattbidewell/gloop-terpreter/token"

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

type Program struct {
	Statments []Statement
}

type LetStatement struct {
	Token token.Token // token.LET token
	Name *Identifier
	Value Expression
}

type Identifier struct {
	Token token.Token // token.IDENT token
	Value string
}

func (p *Program) TokenLiteral() string {
	if len(p.Statments) > 0 {
		return p.Statments[0].TokenLiteral()
	} else {
		return ""
	}
}

func (ls *LetStatement) statementNode() {}
func (ls *LetStatement) TokenLiteral() string { return ls.Token.Literal }

func (i *Identifier) expressionNode() {}
func (i *Identifier) TokenLiteral() string { return i.Token.Literal }
