// token/token.go

type TokenType string

type Token struct {
	Type TokenType
	Literal string
}

// tokentypes

const (
	ILLEGAL = "ILLEGAL"
	EOF = "EOF"
	
	// Identifiers + literals
	IDENT = "IDENT"
	INT = "INT"
	
	// Operators
	ASSIGN= "="
	PLUS= "+"
	
	// Delimiters
	COMMA = ","
	SEMICOLON = ";"
	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"
	
	// Keywords
	FUNCTION = "FUNCTION"
	LET = "LET"
)


