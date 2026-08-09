package parser

import (
	"interpreter/ast"
	"interpreter/lexer"
	"interpreter/token"
)

type Parser struct {
	l *lexer.Lexer

	curToken  token.Token
	peekToken token.Token
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{l: l}

	// Read two tokens, so that curtoken and peekToken are set
	p.nextToken()
	p.nextToken()

	return p
}

func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

// Parse the whole program in a loop
func (p *Parser) ParserProgram() *ast.Program {
	program := &ast.Program{}
	program.Statements = []ast.Statement{}

	for p.curToken.Type != token.EOF {
		stmt := p.parseStatement()
		if stmt != nil {
			program.Statements = append(program.Statements, stmt)
		}
		p.nextToken()
	}
	return program
}

// Parse a statement (LET)
func (p *Parser) parseStatement() ast.Statement {
	switch p.curToken.Type {
	case token.LET:
		return p.parseLetStatement()
	default:
		return nil
	}
}

func (p *Parser) parseLetStatement() *ast.LetStatement {
	// Expected Let statement token structure LET IDNET ASSIGN VALUE SEMICOLON
	stmt := &ast.LetStatement{Token: p.curToken}

	// Expect an IDENT next
	if !p.expectPeek(token.IDENT) {
		return nil
	}

	stmt.Name = &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}

	// Expect an assingment ASSIGN operator next ("=")
	if !p.expectPeek(token.ASSIGN) {
		return nil
	}

	// TODO: need to parse VALUE here later
	// For now skipping to SEMICOLON
	if !p.expectPeek(token.SEMICOLON) {
		return nil
	}

	return stmt
}

func (p *Parser) expectPeek(token token.TokenType) bool {
	if p.peekTokenIs(token) {
		p.nextToken()
		return true
	} else {
		return false
	}
}

func (p *Parser) curTokenIs(token token.TokenType) bool {
	return p.curToken.Type == token
}
func (p *Parser) peekTokenIs(token token.TokenType) bool {
	return p.peekToken.Type == token
}
