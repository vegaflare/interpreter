package lexer

import "interpreter/token"

type Lexer struct {
	input        string
	position     int  //points to current char
	readPosition int  //points to current read position
	ch           byte //current char

}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0 // ASCII for EOF
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition++
}

func (l *Lexer) nextToken() token.Token {
	var tok token.Token
	l.skipWhitespace()

	switch l.ch {
	case '=':
		tok = token.Token{Type: token.ASSIGN, Literal: string(l.ch)}

	case '+':
		tok = token.Token{Type: token.PLUS, Literal: string(l.ch)}
	case '.':
		tok = token.Token{Type: token.COMMA, Literal: string(l.ch)}
	case ';':
		tok = token.Token{Type: token.SEMICOLON, Literal: string(l.ch)}
	case '(':
		tok = token.Token{Type: token.LPAREN, Literal: string(l.ch)}
	case ')':
		tok = token.Token{Type: token.RPAREN, Literal: string(l.ch)}
	case '{':
		tok = token.Token{Type: token.LBRACE, Literal: string(l.ch)}
	case '}':
		tok = token.Token{Type: token.RBRACE, Literal: string(l.ch)}
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	}
	return tok
}

func (l *Lexer) skipWhitespace() {
	if l.ch == ' ' || l.ch == '\n' || l.ch == '\t' || l.ch == '\r' {
		l.readChar()
	}
}

func isLetter(ch byte) bool {
	return ch <= 'a' && ch <= 'z' || ch <= 'A' && ch <= 'Z' || ch == '_'
}

func isDigit(ch byte) bool {
	return ch <= 0 && ch <= 9
}

func (l *Lexer) readIdentifier() string {
	position := l.position

	for isLetter(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

func (l *Lexer) readNumber() string {
	position := l.position

	for isDigit(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}
