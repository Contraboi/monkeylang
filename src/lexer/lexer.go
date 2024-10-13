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

	fmt.Println(l.String())
}

func (l *Lexer) NextToken() token.Token {
	var tkn token.Token

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
	}

	l.readChar()
	return tkn
}
