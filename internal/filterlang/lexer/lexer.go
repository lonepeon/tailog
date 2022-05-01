package lexer

type TokenReader interface {
	Matches([]rune) bool
	Read([]rune) (Token, []rune)
}

type Lexer struct {
	content  []rune
	registry []TokenReader
}

func NewLexer(content string) *Lexer {
	return &Lexer{
		content: []rune(content),
		registry: []TokenReader{
			EqualSpecialCharacter,
			DoubleQuotesIdentifier{},
			NoQuotesIdentifier{},
		},
	}
}

func (l *Lexer) NextToken() Token {
	l.content = EatSpaces(l.content)

	if len(l.content) == 0 {
		return newTokenEOF()
	}

	for i := range l.registry {
		if !l.registry[i].Matches(l.content) {
			continue
		}

		token, remaining := l.registry[i].Read(l.content)
		l.content = remaining
		return token
	}

	return newTokenIllegal("unparsable input")
}
