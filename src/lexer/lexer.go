package lexer

import (
	"fmt"
	"monkeylang/src/token"
)

// TODO: support full unicode and UTF-8 range

type Lexer struct {
	input        string
	position     int
	readPosition int
	ch           byte
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

func (l *Lexer) String() string {
	return fmt.Sprintf("input: %s | position: %d | readPosition: %d | ch: %s", l.input, l.position, l.readPosition, string(l.ch))
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0 // EOF
	} else {
		l.ch = l.input[l.readPosition]
	}

	l.position = l.readPosition
	l.readPosition++
}

func (l *Lexer) NextToken() token.Token {
	var tkn token.Token
	l.skipWhitespace()

	switch l.ch {
	case '=':
		tkn = token.NewToken(token.ASSIGN, l.ch)
	case '+':
		tkn = token.NewToken(token.PLUS, l.ch)
	case ';':
		tkn = token.NewToken(token.SEMICOLON, l.ch)
	case ',':
		tkn = token.NewToken(token.COMMA, l.ch)
	case '(':
		tkn = token.NewToken(token.LPAREN, l.ch)
	case ')':
		tkn = token.NewToken(token.RPAREN, l.ch)
	case '{':
		tkn = token.NewToken(token.LBRACE, l.ch)
	case '}':
		tkn = token.NewToken(token.RBRACE, l.ch)
	case 0:
		tkn.Literal = ""
		tkn.Type = token.EOF
	default:
		if isLetter(l.ch) {
			tkn.Literal = l.readIndetifier()
			tkn.Type = token.LookupIdent(tkn.Literal)
			return tkn
		} else if isDigit(l.ch) {
			tkn.Literal = l.readNumber()
			tkn.Type = token.INT
			return tkn
		} else {
			tkn = token.NewToken(token.ILLEGAL, l.ch)
		}
	}

	l.readChar()
	return tkn
}

func (l *Lexer) readIndetifier() string {
	start := l.position
	for isLetter(l.ch) {
		l.readChar()
	}
	end := l.position
	return l.input[start:end]

}

func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

func (l *Lexer) readNumber() string {
	start := l.position
	for isDigit(l.ch) {
		l.readChar()
	}
	end := l.position
	return l.input[start:end]
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}
func isLetter(ch byte) bool {
	isLowerCaseAtoZ := 'a' <= ch && ch <= 'z'
	isUpperCaseAtoZ := 'A' <= ch && ch <= 'Z'
	return isLowerCaseAtoZ || isUpperCaseAtoZ || ch == '_'
}
