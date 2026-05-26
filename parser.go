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
	case FROM:
		// todo : implement this to see if the tests
		// pass
		return nil
	case INSERT:
		return nil
	default:
		return nil
	}
	return nil
}

func (p *Parser) parseSelectStatement() *SelectStatement {
	// select * from users or select id,name .... from users;
	stmt := &SelectStatement{Token: p.curToken}

	if !p.expectPeek(ASTERISK) && !p.expectPeek(IDENT) {
		return nil
	}

	p.nextToken()

	if p.curToken.Type == ASTERISK {
		stmt.Columns = append(stmt.Columns, &WildcardExpression{
			Token: p.curToken,
		})
	}

	for p.curToken.Type == IDENT {
		stmt.Columns = append(stmt.Columns, &Identifier{
			Token: p.curToken,
			Value: p.curToken.Literal,
		})
		if p.peekToken.Type == COMMA {
			p.nextToken()
			p.nextToken()
		} else {
			p.nextToken()
		}
	}

	if !p.expectPeek(FROM) {
		return nil
	}

	p.nextToken()

	stmt.From = &Identifier{Token: p.curToken, Value: p.curToken.Literal}

	if !p.expectPeek(IDENT) {
		return nil
	}
	p.nextToken()

	for p.curToken.Type == IDENT {
		stmt.Tables = append(stmt.Tables, &Identifier{
			Token: p.curToken,
			Value: p.curToken.Literal,
		})

		if p.peekToken.Type == COMMA {
			p.nextToken()
			p.nextToken()
		} else {
			p.nextToken()
		}

	}

	if !p.expectPeek(SEMICOLON) {
		return nil
	}

	return stmt
}

func (p *Parser) parseFromClause() Expression {
	if !p.expectPeek(IDENT) {
		return nil
	}

	return &Identifier{Token: p.curToken, Value: p.curToken.Literal}
}

func (p *Parser) expectPeek(t TokenType) bool {
	if p.peekToken.Type == t {
		p.nextToken()
		return true
	} else {
		return false
	}
}
