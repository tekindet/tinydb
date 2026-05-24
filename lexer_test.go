package main

import "testing"

func TestNextToken(t *testing.T) {
	input := `select*where=`

	tests := []struct {
		expectedType    TokenType
		expectedLiteral string
	}{
		{SELECT, "select"},
		{ASTERISK, "*"},
		{WHERE, "where"},
		{EQUALS, "="},
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
