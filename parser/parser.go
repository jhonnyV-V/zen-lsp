package parser

import (
	"github.com/zen-lsp/ast"
	"github.com/zen-lsp/lexer"
	"github.com/zen-lsp/token"
)

type Parser struct {
	l            *lexer.Lexer
	currentToken lexer.TokenWithPos
	peekToken    lexer.TokenWithPos
	errors       []string
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{
		l:              l,
		errors:         []string{},
		prefixParseFns: make(map[token.TokenType]prefixParseFn),
	}
	//prefix
	p.registerPrefix(token.IDENT, p.parseIdentifier)
	p.registerPrefix(token.INT, p.parseIntegerLiteral)
	p.registerPrefix(token.FALSE, p.parseBoolean)
	p.registerPrefix(token.TRUE, p.parseBoolean)
	p.registerPrefix(token.BANG, p.parsePrefixExpression)
	p.registerPrefix(token.MINUS, p.parsePrefixExpression)
	p.registerPrefix(token.DECREMENT, p.parsePrefixExpression)
	p.registerPrefix(token.INCREMENT, p.parsePrefixExpression)
	p.registerPrefix(token.LPAREN, p.parseGroupedExpression)
	p.registerPrefix(token.IF, p.parseIfExpression)
	p.registerPrefix(token.FOR, p.parseForLoop)
	p.registerPrefix(token.FUNCTION, p.parseFunctionLiteral)
	p.registerPrefix(token.STRING, p.parseStringLiteral)
	p.registerPrefix(token.LBRACKET, p.parseArrayLiteral)
	p.registerPrefix(token.LBRACE, p.parseHashLiteral)

	p.NextToken()
	return p
}

func (p *Parser) NextToken() {
	if p.peekToken.Token.Type == token.EOF {
		p.currentToken = p.peekToken
		p.peekToken.Token.Type = token.EOF
	} else {
		p.currentToken = p.peekToken
		p.peekToken = p.l.NextToken()
	}
}

// TODO: maybe tokenize the end of lines and use it as a delimiter

func (p *Parser) parseExpression(precedence int) ast.Node {
	if p.currentToken.Token.Type == "" {
		return nil
	}
	prefix := p.prefixParseFns[p.currentToken.Token.Type]
	if prefix == nil {
		p.noPrefixParseFnError(p.currentToken)
		return nil
	}
	leftExp := prefix()

	for !p.peekTokenIs(token.SEMICOLON) && precedence < p.peekPrecedence() {
		infix := p.infixParseFns[p.peekToken.Type]
		if infix == nil {
			return leftExp
		}
		p.NextToken()
		leftExp = infix(leftExp)
	}
	return leftExp
}

// func (p *Parser) ParseStatement() ast.Node {
func (p *Parser) ParseStatement() {
	switch p.currentToken.Token.Type {
	// case token.LET:
	// 	return p.ParseLetStatement()
	// case token.RETURN:
	// 	return p.ParseReturnStatement()
	default:
		return p.parseStatement()
		// return p.parseExpressionStatement()
	}
}

func (p Parser) ParseProgram() *ast.Program {
	program := &ast.Program{}
	program.Statements = []ast.Node{}
	for p.currentToken.Token.Type != token.EOF {
		var statement ast.Node
		if p.currentToken.Token.Type != "" {
			statement = p.ParseStatement()
		}
		if statement != nil && statement.TokenLiteral() != "" {
			program.Statements = append(program.Statements, statement)
		}
		p.NextToken()
	}
	return program
}
