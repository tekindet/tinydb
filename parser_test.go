package main

import (
	"log"
	"testing"
)

func TestSelectStatement(t *testing.T) {
	input := `
		select * from users;

		select id from users;

		select email,username from users;

	`

	l := NewLexer(input)
	p := NewParser(l)

	program := p.ParseProgram()

	if program == nil {
		t.Fatalf("ParseProgram() returned nil")
	}

	if len(program.Statements) != 3 {
		t.Fatalf("program.Statements does not conntain 3 statements. got=%d", len(program.Statements))
	}

	tests := []struct {
		numberOfColumns int
		numberOfTables  int
	}{
		{
			numberOfColumns: 1,
			numberOfTables:  1,
		},
		{
			numberOfColumns: 1,
			numberOfTables:  1,
		},
		{
			numberOfColumns: 2,
			numberOfTables:  1,
		},
	}

	for i, tt := range tests {
		stmt := program.Statements[i]
		if !testSelectStatement(t, stmt, tt.numberOfColumns, tt.numberOfTables) {
			return
		}
	}
}

func testSelectStatement(t *testing.T, s Statement, cols, tables int) bool {

	selectStmt, ok := s.(*SelectStatement)
	if !ok {
		log.Fatal("here")
		t.Errorf("s not *SelectStatement. got=%T", s)
		return false
	}

	if len(selectStmt.Columns) != cols {
		t.Errorf("s.Columns not '%d'. got=%d", cols, len(selectStmt.Columns))
		return false
	}

	if len(selectStmt.Tables) != tables {
		t.Errorf("s.Tables not '%d'. got=%d", cols, len(selectStmt.Tables))
		return false
	}

	return true
}
