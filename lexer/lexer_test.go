package lexer

import (
	"fmt"
	"os"
	"testing"

	"github.com/zen-lsp/token"
)

func getTestFromFile() string {
	file, _ := os.ReadFile("TestNewToken.zen")

	return string(file)
}

func TestNewToken(t *testing.T) {
	input := getTestFromFile()

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
		expectedLine    int
	}{
		{token.RULE, "Rule", 0},
		{token.CHECK, "check", 0},
		{token.STRING, "2.0.0", 0},
		{token.RULE, "Rule", 1},
		{token.UNKNOWN, "unknown", 1},
		{token.IGNORE, "ignore", 1},
		{token.SCENARIO, "Scenario", 3},
		{token.STRING, "ecdh", 3},
		{token.COLON, ":", 3},
		{token.POUND, "#", 4},
		{token.POUND, "#", 5},
		{token.GIVEN, "Given", 6},
		{token.NOTHING, "nothing", 6},
		{token.POUND, "#", 8},
		{token.POUND, "#", 9},
		{token.POUND, "#", 10},
		{token.WHEN, "When", 11},
		{token.I, "I", 11},
		{token.CREATE, "create", 11},
		{token.THE, "the", 11},
		{token.ECDH, "ecdh", 11},
		{token.IDENT, "key", 11},
		{token.WHEN, "When", 12},
		{token.I, "I", 12},
		{token.CREATE, "create", 12},
		{token.THE, "the", 12},
		{token.ETHEREUM, "ethereum", 12},
		{token.IDENT, "key", 12},
		{token.THEN, "Then", 14},
		{token.PRINT, "print", 14},
		{token.THE, "the", 14},
		{token.STRING, "keyring", 14},
		{token.POUND, "#", 16},
		{token.EOF, "", 16},
	}

	l := New(input)

	for i, tt := range tests {
		tok := l.NextToken()
		fmt.Printf("%+v\n", tok)

		if tok.Token.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q", i, tt.expectedType, tok.Token.Type)
		}

		if tok.Token.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q", i, tt.expectedLiteral, tok.Token.Literal)
		}

		if tok.Line != tt.expectedLine {
			t.Fatalf("tests[%d] - line wrong. expected=%d, got=%d", i, tt.expectedLine, tok.Line)
		}
	}
}
