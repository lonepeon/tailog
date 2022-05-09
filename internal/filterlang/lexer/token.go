package lexer

import "fmt"

type TokenType struct {
	name string
}

func (t TokenType) String() string {
	return t.name
}

var (
	TokenTypeIllegal  = TokenType{name: "Illegal"}
	TokenTypeEOF      = TokenType{name: "EOF"}
	TokenTypeField    = TokenType{name: "Field"}
	TokenTypeEqual    = TokenType{name: "Equal"}
	TokenTypeNotEqual = TokenType{name: "NotEqual"}
	TokenTypeNumber   = TokenType{name: "Number"}
	TokenTypeAnd      = TokenType{name: "And"}
	TokenTypeOr       = TokenType{name: "Or"}
)

type Token struct {
	Type  TokenType
	Value string
}

func (t Token) String() string {
	if t.Value == "" {
		return t.Type.String()
	}

	return fmt.Sprintf("%s(%q)", t.Type, t.Value)
}

func NewTokenEOF() Token {
	return Token{Type: TokenTypeEOF, Value: ""}
}

func NewTokenIllegal(reason string) Token {
	return Token{Type: TokenTypeIllegal, Value: reason}
}

func NewTokenField(name string) Token {
	return Token{Type: TokenTypeField, Value: name}
}

func NewTokenEqual() Token {
	return Token{Type: TokenTypeEqual, Value: ""}
}

func NewTokenNotEqual() Token {
	return Token{Type: TokenTypeNotEqual, Value: ""}
}

func NewTokenNumber(value string) Token {
	return Token{Type: TokenTypeNumber, Value: value}
}

func NewTokenAnd() Token {
	return Token{Type: TokenTypeAnd, Value: ""}
}

func NewTokenOr() Token {
	return Token{Type: TokenTypeOr, Value: ""}
}
