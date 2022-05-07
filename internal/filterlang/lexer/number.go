package lexer

var (
	Number = number{}
)

type number struct{}

func (n number) Matches(content []rune) bool {
	return startWithFn(content, isNum)
}

func (n number) Read(content []rune) (Token, []rune) {
	if len(content) == 0 {
		return NewTokenEOF(), content
	}

	integer, remaining, found := readWhile(content, func(c rune) bool {
		return isNum(c)
	})

	if !found {
		return NewTokenIllegal("didn't detect any integer number"), content
	}

	if !startWith(remaining, []rune(".")) {
		return NewTokenNumber(integer), remaining
	}

	fraction, remaining, found := readWhile(remaining[1:], func(c rune) bool {
		return isNum(c)
	})

	if !found {
		return NewTokenIllegal("didn't detect any number after the decimal separator"), content
	}

	return NewTokenNumber(integer + "." + fraction), remaining
}
