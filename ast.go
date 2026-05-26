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
	Statements []Statement `json:"statements"`
}

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}

type SelectStatement struct {
	Token   Token         `json:"token"`
	Columns []Expression  `json:"columnns"`
	From    *Identifier   `json:"from"`
	Tables  []*Identifier `json:"tables"`
	Where   Expression    `json:"where"`
	Limit   Expression    `json:"limit"`
}

func (s *SelectStatement) statementNode() {}

func (s *SelectStatement) TokenLiteral() string {
	return s.Token.Literal
}

type WildcardExpression struct {
	Token Token
}

func (w *WildcardExpression) expressionNode() {}
func (w *WildcardExpression) TokenLiteral() string {
	return w.Token.Literal
}

type Identifier struct {
	Token Token  `json:"token"`
	Value string `json:"value"`
}

func (i *Identifier) expressionNode() {}

func (i *Identifier) TokenLiteral() string {
	return i.Token.Literal
}
