package main

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	IDENT = "IDENT"
	INT   = "INT"

	SELECT = "SELECT"
	INSERT = "INSERT"
	WHERE  = "WHERE"
	FROM   = "FROM"

	ASTERISK  = "*"
	COMMA     = ","
	SEMICOLON = ";"
	EQUALS    = "="

	AND = "AND"
	OR  = "OR"
	NOT = "NOT"
)

var keywords = map[string]TokenType{
	"select": SELECT,
	"insert": INSERT,
	"where":  WHERE,
	"from":   FROM,
	"and":    AND,
	"or":     OR,
	"not":    NOT,
}

func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}
