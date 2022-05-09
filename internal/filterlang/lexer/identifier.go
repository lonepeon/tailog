package lexer

var (
	DoubleQuotesIdentifier = doubleQuotesIdentifier{}
	NoQuotesIdentifier     = noQuotesIdentifier{}
)

type doubleQuotesIdentifier struct{}

func (i doubleQuotesIdentifier) Matches(content []rune) bool {
	return startWith(content, []rune(`lbl:"`))
}

func (i doubleQuotesIdentifier) Read(content []rune) (Token, []rune) {
	if len(content) == 0 {
		return NewTokenEOF(), content
	}

	if !startWith(content, []rune(`lbl:"`)) {
		return NewTokenIllegal("expecting identifier's opening double quote character"), content
	}

	identifier, remaining, found := readWhile(content[5:], func(c rune) bool { return c != '"' })

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
	return startWith(content, []rune(`lbl:`)) && startWithFn(content[4:], isAlpha)
}

func (i noQuotesIdentifier) Read(content []rune) (Token, []rune) {
	if len(content) == 0 {
		return NewTokenEOF(), content
	}

	if !startWith(content, []rune(`lbl:`)) {
		return NewTokenIllegal("expecting identifier to start with an alphabetic character"), content
	}

	if !startWithFn(content[4:], isAlpha) {
		return NewTokenIllegal("expecting identifier to start with an alphabetic character"), content
	}

	identifier, remaining, found := readWhile(content[4:], func(c rune) bool {
		return isAlphaNum(c) || c == '_'
	})

	if !found {
		return NewTokenIllegal("didn't detect any identifier"), content
	}

	return NewTokenIdentifier(identifier), remaining
}
