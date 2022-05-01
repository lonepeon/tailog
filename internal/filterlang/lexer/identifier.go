package lexer

type DoubleQuotesIdentifier struct{}

func (i DoubleQuotesIdentifier) Matches(content []rune) bool {
	return startWithFn(content, isRune('"'))
}

func (i DoubleQuotesIdentifier) Read(content []rune) (Token, []rune) {
	if len(content) == 0 {
		return newTokenEOF(), content
	}

	if !startWithFn(content, func(c rune) bool { return c == '"' }) {
		return newTokenIllegal("expecting identifier's opening double quote character"), content
	}

	identifier, remaining, found := readWhile(content[1:], func(c rune) bool { return c != '"' })

	if !found {
		return newTokenIllegal("didn't detect any identifier"), content
	}

	if !startWithFn(remaining, func(c rune) bool { return c == '"' }) {
		return newTokenIllegal("expecting identifier's closing double quote character"), content
	}

	return newTokenIdentifier(identifier), remaining[1:]
}

type NoQuotesIdentifier struct{}

func (i NoQuotesIdentifier) Matches(content []rune) bool {
	return startWithFn(content, isAlpha)
}

func (i NoQuotesIdentifier) Read(content []rune) (Token, []rune) {
	if len(content) == 0 {
		return newTokenEOF(), content
	}

	if !startWithFn(content, isAlpha) {
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
