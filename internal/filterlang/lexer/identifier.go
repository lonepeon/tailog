package lexer

func DoubleQuotesIdentifier(content []rune) (Token, []rune) {
	if len(content) == 0 {
		return newTokenEOF(), content
	}

	if !startWith(content, func(c rune) bool { return c == '"' }) {
		return newTokenIllegal("expecting identifier's opening double quote character"), content
	}

	identifier, remaining, found := readWhile(content[1:], func(c rune) bool { return c != '"' })

	if !found {
		return newTokenIllegal("didn't detect any identifier"), content
	}

	if !startWith(remaining, func(c rune) bool { return c == '"' }) {
		return newTokenIllegal("expecting identifier's closing double quote character"), content
	}

	return newTokenIdentifier(identifier), remaining[1:]
}

func SimpleIdentifier(content []rune) (Token, []rune) {
	if len(content) == 0 {
		return newTokenEOF(), content
	}

	if !startWith(content, isAlpha) {
		return newTokenIllegal("expecting identifier to start with an alphabetic character"), content
	}

	identifier, remaining, found := readWhile(content, func(c rune) bool {
		return isAlphaNum(c) || c == '_'
	})

	if !found {
		return newTokenIllegal("didn't detect any identifier"), content
	}

	return newTokenIdentifier(identifier), remaining
}
