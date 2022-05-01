package lexer

type TokenType struct {
	name string
}

func (t TokenType) String() string {
	return t.name
}

var (
	TokenTypeIllegal    = TokenType{name: "Illegal"}
	TokenTypeEOF        = TokenType{name: "EOF"}
	TokenTypeIdentifier = TokenType{name: "Identifier"}
)

type Token struct {
	Type  TokenType
	Value string
}

func newTokenEOF() Token {
	return Token{Type: TokenTypeEOF, Value: ""}
}

func newTokenIllegal(reason string) Token {
	return Token{Type: TokenTypeIllegal, Value: reason}
}

func newTokenIdentifier(name string) Token {
	return Token{Type: TokenTypeIdentifier, Value: name}
}
