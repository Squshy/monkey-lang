package parser

import (
	"fmt"
	"monkey/ast"
	"monkey/lexer"
	"monkey/token"
)

type Parser struct {
	l *lexer.Lexer

	errors    []string
	curToken  token.Token
	peekToken token.Token
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{l: l,
		errors: []string{}}

	// Read two tokens, so `curToken` and `peekToken` are both set
	p.nextToken()
	p.nextToken()

	return p
}

func (p *Parser) Errors() []string { return p.errors }
func (p *Parser) peekError(t token.TokenType) {
	msg := fmt.Sprintf("expected next token to be %s, got %s instead", t, p.peekToken.Type)
	p.errors = append(p.errors, msg)
}

func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

func (p *Parser) parseStatement() ast.Statement {
	switch p.curToken.Type {
	case token.LET:
		return p.parseLetStatement()
	default:
		return nil
	}
}

func (p *Parser) parseLetStatement() *ast.LetStatement {
	// Currently a `token.LET` token
	stmt := &ast.LetStatement{Token: p.curToken}

	// We must first see an identifier
	// `let x = 5;`
	if !p.expectPeek(token.IDENT) {
		return nil
	}

	// Construct an identifier node as the current token is now `token.IDENT`
	stmt.Name = &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}

	if !p.expectPeek(token.ASSIGN) {
		return nil
	}

	// TODO: We're skipping the expressions until we encounter a semicolon
	for !p.curTokenIs(token.SEMICOLON) {
		p.nextToken()
	}

	return stmt
}

func (p *Parser) curTokenIs(t token.TokenType) bool {
	return p.curToken.Type == t
}

func (p *Parser) peekTokenIs(t token.TokenType) bool {
	return p.peekToken.Type == t
}

// Asserts the next token is as expected and when true, advances both `curToken`
// and `peekToken`.
func (p *Parser) expectPeek(t token.TokenType) bool {
	// This is an assertion function nearly all parsers share. Its purpose is to
	// enforce the correctness of the order of tokens by checking the type of
	// the next token.
	// We only advance the tokens if the next available token is as expected.

	if p.peekTokenIs(t) {
		p.nextToken()
		return true
	} else {
		p.peekError(t)
		return false
	}
}

func (p *Parser) ParseProgram() *ast.Program {
	// Construct the root node
	program := &ast.Program{}
	program.Statements = []ast.Statement{}

	for !p.curTokenIs(token.EOF) {
		stmt := p.parseStatement()

		if stmt != nil {
			program.Statements = append(program.Statements, stmt)
		}

		p.nextToken()
	}

	return program
}
