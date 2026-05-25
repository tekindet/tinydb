package main

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
	Statements []Statement
}

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}

type SelectStatement struct {
	Token Token
	Name  *Identifier
	Value Expression
}

func (s *SelectStatement) statementNode() {}

func (s *SelectStatement) TokenLiteral() string {
	return s.Token.Literal
}

type Identifier struct {
	Token Token
	Value string
}

func (i *Identifier) expressionNode() {}

func (i *Identifier) TokenLiteral() string {
	return i.Token.Literal
}
