package main

import "testing"

func TestSelectStatement(t *testing.T) {
	input := `
		select * from users
		select id,email,username from users
		select id,email,username from users;
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
		expectedIdentifier string
	}{
		{"*"}, {"id"}, {"email"}, {"username"}, {"users"},
	}

	for i, tt := range tests {
		stmt := program.Statements[i]
		if !testSelectStatement(t, stmt, tt.expectedIdentifier) {
			return
		}
	}
}

func testSelectStatement(t *testing.T, s Statement, name string) bool {
	if s.TokenLiteral() != "select" {
		t.Errorf("s.TokenLiteral not 'select'. got=%q",s.TokenLiteral())
		return false
	}

	selectStmt,ok := s.(*SelectStatement)
	if !ok {
		t.Errorf("s not *SelectStatement. got=%T",s)
		return false
	}

	if selectStmt.Name.Value != name {
		t.Errorf("selectStmt.Name.Value not '%s'. got=%s", name, selectStmt.Name.Value)
		return false
	}

	if selectStmt.Name.TokenLiteral() != name {
		t.Errorf("selectStmt.Name.TokenLiteral() not '%s'. got=%s", name, selectStmt.Name.TokenLiteral())
		return false
	}

	return true
}
