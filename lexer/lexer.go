package lexer

import (
	"github.com/zen-lsp/token"
)

type Lexer struct {
	input        string
	position     int // current position in input (current char)
	readPosition int // current reading position in input (after current char)
	ch           byte
	line         int
}

type TokenWithPos struct {
	Token token.Token
	Line  int
	Start int
	End   int
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.ReadChar()
	return l
}

func (l *Lexer) ReadChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		if l.ch == '\n' {
			l.line++
		}
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition += 1
}

func (l *Lexer) NextToken() TokenWithPos {
	var tok token.Token

	l.skipWhitespace()

	start := l.position

	switch l.ch {
	case '=':
		if l.peekChar() == '=' {
			ch := l.ch
			l.ReadChar()
			tok = token.Token{Type: token.EQ, Literal: string(ch) + string(l.ch)}
		} else {
			tok = token.NewToken(token.ASSIGN, l.ch)
		}
	case '(':
		tok = token.NewToken(token.LPAREN, l.ch)
	case ')':
		tok = token.NewToken(token.RPAREN, l.ch)
	case '+':
		tok = token.NewToken(token.PLUS, l.ch)
	case '-':
		tok = token.NewToken(token.MINUS, l.ch)
	case '{':
		tok = token.NewToken(token.LBRACE, l.ch)
	case '}':
		tok = token.NewToken(token.RBRACE, l.ch)
	case '[':
		tok = token.NewToken(token.LBRACKET, l.ch)
	case ']':
		tok = token.NewToken(token.RBRACKET, l.ch)
	case '/':
		tok = token.NewToken(token.SLASH, l.ch)
	case '*':
		tok = token.NewToken(token.ASTERISK, l.ch)
	case '!':
		if l.peekChar() == '=' {
			ch := l.ch
			l.ReadChar()
			tok = token.Token{Type: token.NOT_EQ, Literal: string(ch) + string(l.ch)}
		} else {
			tok = token.NewToken(token.BANG, l.ch)
		}
	case '<':
		if l.peekChar() == '=' {
			ch := l.ch
			l.ReadChar()
			tok = token.Token{Type: token.LT_EQ, Literal: string(ch) + string(l.ch)}
		} else {
			tok = token.NewToken(token.LT, l.ch)
		}
	case '>':
		if l.peekChar() == '=' {
			ch := l.ch
			l.ReadChar()
			tok = token.Token{Type: token.GT_EQ, Literal: string(ch) + string(l.ch)}
		} else {
			tok = token.NewToken(token.GT, l.ch)
		}
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	case '\'':
		tok.Type = token.STRING
		tok.Literal = l.readString()
		//comment
	case ':':
		tok = token.NewToken(token.COLON, l.ch)
		l.ignoreLine()

	case '#':
		tok = token.NewToken(token.POUND, l.ch)
		l.ignoreLine()
	default:
		if isLetter(l.ch) {
			tok.Literal = l.readIdentifier()
			tok.Type = token.LookupIdent(tok.Literal)
		} else {
			tok = token.NewToken(token.ILLEGAL, l.ch)
		}
	}

	wrap := TokenWithPos{
		Token: tok,
		Start: start,
		End:   l.readPosition,
		Line:  l.line,
	}
	l.ReadChar()
	return wrap
}

func (l *Lexer) readIdentifier() string {
	position := l.position
	for isLetter(l.ch) {
		l.ReadChar()
	}
	return l.input[position:l.position]
}

func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.ReadChar()
	}
}

func (l *Lexer) peekChar() byte {
	if l.readPosition >= len(l.input) {
		return 0
	} else {
		return l.input[l.readPosition]
	}
}

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func (l *Lexer) readString() string {
	position := l.position + 1
	for {
		l.ReadChar()
		if l.ch == '\'' || l.ch == 0 {
			break
		}
	}
	return l.input[position:l.position]
}

func (l *Lexer) ignoreLine() {
	end := len(l.input)
	pos := l.position
	for pos < end && l.peekChar() != '\n' {
		l.ReadChar()
		pos = l.position
	}
}
