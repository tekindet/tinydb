package main

import (
	"testing"
)

func TestSelectStatement(t *testing.T) {
	input := `
		select * from users;
	`

	l := NewLexer(input)
	p := NewParser(l)

	program := p.ParseProgram()

	if program == nil {
		t.Fatalf("ParseProgram() returned nil")
	}

	if len(program.Statements) != 1 {
		t.Fatalf("program.Statements does not conntain 3 statements. got=%d", len(program.Statements))
	}

	tests := []struct {
		Columnns []string
		Tables   []string
	}{
		{
			Columnns: []string{"*"},
			Tables:   []string{"users"},
		},
	}

	for i, tt := range tests {
		stmt, ok := program.Statements[i].(*SelectStatement)
		if !ok {
			t.Fatalf("not a select statement,should be")
		}

		if len(tt.Columnns) != len(stmt.Columns) {
			t.Fatalf("want %d columns got %d columns", len(tt.Columnns), len(stmt.Columns))
		}
	}
}
