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
	RULE     = "RULE"

	//Scenarios
	ECDH     = "ECDH"
	ETHEREUM = "ETHEREUM"

	//RULES
	CHECK   = "CHECK"
	VERSION = "VERSION"
	UNKNOWN = "UNKNOWN"
	IGNORE  = "IGNORE"
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
	"Rule":     RULE,
}

var Scenarios = map[string]TokenType{
	"ecdh":     ECDH,
	"ethereum": ETHEREUM,
}

var Rules = map[string]TokenType{
	"check":   CHECK,
	"version": VERSION,
	"unknown": UNKNOWN,
	"ignore":  IGNORE,
}

func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}

	if tok := LookupScenario(ident); tok != ILLEGAL {
		return tok
	}
	if tok := LookupRules(ident); tok != ILLEGAL {
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

func LookupRules(rule string) TokenType {
	if tok, ok := Rules[rule]; ok {
		return tok
	}
	return ILLEGAL
}
