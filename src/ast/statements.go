package ast

import (
	"monkeylang/src/token"
)

type LetStatement struct {
	Token token.Token
	Name  *Identifier
	Value Expression
}

func (ls *LetStatement) statementNode()
func (ls *LetStatement) TokenLiteral() string {
	return ls.Token.Literal
}

type Identifier struct {
	Token token.Token
	Value string
}

func (ls *Identifier) statementNode()
func (ls *Identifier) TokenLiteral() string {
	return ls.Token.Literal
}
