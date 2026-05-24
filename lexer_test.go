package main

import "testing"

func TestNextToken(t *testing.T) {
	input := `
		select student.name,
		course.name 
		join student on student.id = course.student_id 
		where student.age > 20 `

	tests := []struct {
		expectedType    TokenType
		expectedLiteral string
	}{
		{SELECT, "select"},
		{IDENT, "student"},
		{DOT, "."},
		{IDENT, "name"},
		{COMMA, ","},
		{IDENT, "course"},
		{DOT, "."},
		{IDENT, "name"},
		{JOIN, "join"},
		{IDENT, "student"},
		{ON, "on"},
		{IDENT, "student"},
		{DOT, "."},
		{IDENT, "id"},
		{EQUALS, "="},
		{IDENT, "course"},
		{DOT, "."},
		{IDENT, "student_id"},
		{WHERE, "where"},
		{IDENT, "student"},
		{DOT, "."},
		{IDENT, "age"},
		{GREATER_THAN, ">"},
		{INT, "20"},
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
