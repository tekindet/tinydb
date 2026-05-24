package main

import "testing"

func TestNextToken(t *testing.T) {
	input := `
		select employee.name,
		department.name 
		join employee on employee.id = department.employee_id 
		where employee.age < 30 `

	tests := []struct {
		expectedType    TokenType
		expectedLiteral string
	}{
		{SELECT, "select"},
		{IDENT, "employee"},
		{DOT, "."},
		{IDENT, "name"},
		{COMMA, ","},
		{IDENT, "department"},
		{DOT, "."},
		{IDENT, "name"},
		{JOIN, "join"},
		{IDENT, "employee"},
		{ON, "on"},
		{IDENT, "employee"},
		{DOT, "."},
		{IDENT, "id"},
		{EQUALS, "="},
		{IDENT, "department"},
		{DOT, "."},
		{IDENT, "employee_id"},
		{WHERE, "where"},
		{IDENT, "employee"},
		{DOT, "."},
		{IDENT, "age"},
		{LESS_THAN, "<"},
		{INT, "30"},
	}

	l := NewLexer(input)

	for i, tt := range tests {
		tok := l.NextToken()

		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got%q", i, tt.expectedType, tok.Type)
		}

		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got%q", i, tt.expectedLiteral, tok.Literal)
		}
	}
}
