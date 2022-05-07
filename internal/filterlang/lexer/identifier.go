package lexer

var (
	DoubleQuotesIdentifier = doubleQuotesIdentifier{}
	NoQuotesIdentifier     = noQuotesIdentifier{}
)

type doubleQuotesIdentifier struct{}

func (i doubleQuotesIdentifier) Matches(content []rune) bool {
	return startWithFn(content, isRune('"'))
}

func (i doubleQuotesIdentifier) Read(content []rune) (Token, []rune) {
	if len(content) == 0 {
		return NewTokenEOF(), content
	}

	if !startWithFn(content, func(c rune) bool { return c == '"' }) {
		return NewTokenIllegal("expecting identifier's opening double quote character"), content
	}

	identifier, remaining, found := readWhile(content[1:], func(c rune) bool { return c != '"' })

	if !found {
		return NewTokenIllegal("didn't detect any identifier"), content
	}

	if !startWithFn(remaining, func(c rune) bool { return c == '"' }) {
		return NewTokenIllegal("expecting identifier's closing double quote character"), content
	}

	return NewTokenIdentifier(identifier), remaining[1:]
}

type noQuotesIdentifier struct{}

func (i noQuotesIdentifier) Matches(content []rune) bool {
	return startWithFn(content, isAlpha)
}

func (i noQuotesIdentifier) Read(content []rune) (Token, []rune) {
	if len(content) == 0 {
		return NewTokenEOF(), content
	}

	if !startWithFn(content, isAlpha) {
		return NewTokenIllegal("expecting identifier to start with an alphabetic character"), content
	}

	identifier, remaining, found := readWhile(content, func(c rune) bool {
		return isAlphaNum(c) || c == '_'
	})

	if !found {
		return NewTokenIllegal("didn't detect any identifier"), content
	}

	return NewTokenIdentifier(identifier), remaining
}
