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
		{token.SCENARIO, "Scenario", 1},
		{token.STRING, "ecdh", 1},
		{token.COLON, ":", 1},
		{token.POUND, "#", 2},
		{token.POUND, "#", 3},
		{token.GIVEN, "Given", 4},
		{token.NOTHING, "nothing", 4},
		{token.POUND, "#", 6},
		{token.POUND, "#", 7},
		{token.POUND, "#", 8},
		{token.WHEN, "When", 9},
		{token.I, "I", 9},
		{token.CREATE, "create", 9},
		{token.THE, "the", 9},
		{token.IDENT, "ecdh", 9},
		{token.IDENT, "key", 9},
		{token.WHEN, "When", 10},
		{token.I, "I", 10},
		{token.CREATE, "create", 10},
		{token.THE, "the", 10},
		{token.IDENT, "ethereum", 10},
		{token.IDENT, "key", 10},
		{token.THEN, "Then", 12},
		{token.PRINT, "print", 12},
		{token.THE, "the", 12},
		{token.STRING, "keyring", 12},
		{token.POUND, "#", 14},
		{token.EOF, "", 14},
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
