package lexer

var (
	String = str{}
)

type str struct{}

func (i str) Matches(content []rune) bool {
	return startWithFn(content, isRune('"'))
}

func (i str) Read(content []rune) (Token, []rune) {
	if len(content) == 0 {
		return NewTokenEOF(), content
	}

	if !startWithFn(content, isRune('"')) {
		return NewTokenIllegal("expecting string to start with double quote character"), content
	}

	str, remaining, found := readWhile(content[1:], func(c rune) bool { return c != '"' })

	if !found {
		return NewTokenIllegal("didn't detect any string"), content
	}

	if !startWithFn(remaining, func(c rune) bool { return c == '"' }) {
		return NewTokenIllegal("expecting string to end with a double quote character"), content
	}

	return NewTokenString(str), remaining[1:]
}
