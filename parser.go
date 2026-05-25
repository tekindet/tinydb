package main

type Parser struct {
	l *Lexer

	curToken  Token
	peekToken Token
}

func NewParser(l *Lexer) *Parser {
	p := &Parser{l: l}
	p.nextToken()
	p.nextToken()

	return p
}

func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

func (p *Parser) ParseProgram() *Program {
	program := &Program{}
	program.Statements = []Statement{}

	for p.curToken.Type != EOF {
		stmt := p.parseStatement()
		if stmt != nil {
			program.Statements = append(program.Statements, stmt)
		}
		p.nextToken()
	}

	return program
}

func (p *Parser) parseStatement() Statement {
	switch p.curToken.Type {
	case SELECT:
		return p.parseSelectStatement()
	}
	return nil
}

func (p *Parser) parseSelectStatement() *SelectStatement {
	stmt := &SelectStatement{Token: p.curToken}

	if !p.expectPeek(ASTERISK) && !p.expectPeek(IDENT) {
		return nil
	}

	stmt.Name = &Identifier{Token: p.curToken, Value: p.curToken.Literal}

	for p.peekToken.Type == COMMA {
		p.nextToken()
		p.nextToken()

		if !p.expectPeek(IDENT) {
			return nil
		}

		stmt.Name = &Identifier{Token: p.curToken, Value: p.curToken.Literal}
	}

	if !p.expectPeek(FROM) {
		return nil
	}

	if !p.expectPeek(IDENT) {
		return nil
	}

	stmt.Value = &Identifier{Token: p.curToken, Value: p.curToken.Literal}

	if !p.expectPeek(SEMICOLON) {
		return nil
	}

	return stmt
}

func (p *Parser) expectPeek(t TokenType) bool {
	if p.peekToken.Type == t {
		p.nextToken()
		return true
	} else {
		return false
	}
}
