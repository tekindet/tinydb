package main

import (
	"log"
)

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
	case INSERT:
		return p.parseInsertStatement()
	default:
		return nil
	}
	return nil
}

func (p *Parser) parseInsertStatement() *InsertStatement {
	stmt := &InsertStatement{Token: p.curToken}
	return stmt
}

func (p *Parser) parseColumns() []Expression {

	columns := []Expression{}

	if p.curToken.Type == ASTERISK {
		columns = append(columns, &WildcardExpression{
			Token: p.curToken,
		})

		return columns
	}

	for p.curToken.Type == IDENT {
		columns = append(columns, &Identifier{
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

	return columns
}

func (p *Parser) parseSelectStatement() *SelectStatement {
	stmt := &SelectStatement{Token: p.curToken}

	p.nextToken()

	columns := p.parseColumns()

	stmt.Columns = append(stmt.Columns, columns...)

	stmt.From = p.parseFromClause()
	tables := p.parseTables()

	stmt.Tables = append(stmt.Tables, tables...)

	if !p.expectPeek(SEMICOLON) {
		return nil
	}

	return stmt
}

func (p *Parser) parseFromClause() *Identifier {

	if !p.expectPeek(FROM) {
		log.Fatalf("expect from clause got %s", p.peekToken.Literal)
	}

	p.nextToken()

	return &Identifier{Token: p.curToken, Value: p.curToken.Literal}
}

func (p *Parser) parseTables() []*Identifier {
	if !p.expectPeek(IDENT) {
		return nil
	}
	p.nextToken()

	tables := []*Identifier{}

	for p.curToken.Type == IDENT {
		tables = append(tables, &Identifier{
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

	return tables

}

func (p *Parser) expectPeek(t TokenType) bool {
	if p.peekToken.Type == t {
		p.nextToken()
		return true
	} else {
		return false
	}
}
