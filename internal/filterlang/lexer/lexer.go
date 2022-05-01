package lexer

type Lexer struct {
	content []rune
}

func (l *Lexer) NextToken() Token {
	l.content = EatSpaces(l.content)

	switch l.content[0] {
	case '"':
		token, remaining := DoubleQuotesIdentifier(l.content)
		l.content = remaining
		return token
	}

	return Token{Type: TokenTypeIllegal, Value: "unparsable input"}
}
