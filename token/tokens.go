package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"
	// Identifiers
	IDENT = "IDENT"
	// literals
	STRING = "STRING"
	// Operators
	PLUS     = "+"
	MINUS    = "-"
	ASSIGN   = "="
	BANG     = "!"
	ASTERISK = "*"
	SLASH    = "/"
	LT       = "<"
	GT       = ">"

	EQ     = "=="
	NOT_EQ = "!="
	GT_EQ  = ">="
	LT_EQ  = "<="

	// Delimiters
	COLON = ":"
	POUND = "#"

	LPAREN   = "("
	RPAREN   = ")"
	LBRACE   = "{"
	RBRACE   = "}"
	LBRACKET = "["
	RBRACKET = "]"

	// Keywords
	SCENARIO = "SCENARIO"
	GIVEN    = "GIVEN"
	WHEN     = "WHEN"
	THEN     = "THEN"
	I        = "I"
	AM       = "AM"
	HAVE     = "HAVE"
	AND      = "AND"
	PRINT    = "PRINT"
	CREATE   = "CREATE"
	THE      = "THE"
	NOTHING  = "NOTHING"
	FROM     = "FROM"
	MY       = "MY"
	DATA     = "DATA"
	ALL      = "ALL"
	AS       = "AS"
	KNOWN    = "KNOWN"
	KEYRING  = "KEYRING"

	//Scenarios
	ECDH     = "ECDH"
	ETHEREUM = "ETHEREUM"
)

func NewToken(tokenType TokenType, ch byte) Token {
	return Token{Type: tokenType, Literal: string(ch)}
}

var keywords = map[string]TokenType{
	"Scenario": SCENARIO,
	"Given":    GIVEN,
	"When":     WHEN,
	"I":        I,
	"am":       AM,
	"have":     HAVE,
	"and":      AND,
	"Then":     THEN,
	"print":    PRINT,
	"create":   CREATE,
	"the":      THE,
	"nothing":  NOTHING,
	"from":     FROM,
	"my":       MY,
	"all":      ALL,
	"as":       AS,
	"known":    KNOWN,
	"data":     DATA,
	"keyring":  KEYRING,
}

var Scenarios = map[string]TokenType{
	"ecdh":     ECDH,
	"ethereum": ETHEREUM,
}

func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}

func LookupScenario(scenario string) TokenType {
	if tok, ok := Scenarios[scenario]; ok {
		return tok
	}
	return ILLEGAL
}
