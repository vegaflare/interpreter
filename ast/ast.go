//The Abstract Syntax Tree (AST)

package ast

import "interpreter/token"

type Node interface {
	TokenLiteral() string
	String() string
}

// Statement nodes do not produce values (e.g., let x = 5;)
type Statement interface {
	Node
	statementNode()
}

// Expression nodes produce values (e.g., 5 + 5, x)
type Expression interface {
	Node
	expressionNode()
}

// Program will be the root node of every AST parsed by our interpreter
type Program struct {
	Statements []Statement
}

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	}
	return ""
}

//to get the whole program
func (p *Program) String() string {
	var out string
	for _, s := range p.Statements {
		out += s.String()
	}
	return out
}

type LetStatement struct {
	Token token.Token //token.LET
	Name  *Identifier
	Value Expression
}

func (ls *LetStatement) statementNode()       {}
func (ls *LetStatement) TokenLiteral() string { return ls.Token.Literal }

func (ls *LetStatement) String() string {
	var out string
	out += ls.TokenLiteral() + " "
	if ls.Name != nil {
		out += ls.Name.String()
	}
	out += "="
	if ls.Value != nil {
		out += ls.Value.String()
	}
	out += ";"
	return out
}

type Identifier struct {
	Token token.Token //token.Ident
	Value string
}

func (i *Identifier) expressionNode() {}
func (i *Identifier) TokenLiteral() string {
	return i.Token.Literal
}

func (i *Identifier) String() string { return i.Value }
