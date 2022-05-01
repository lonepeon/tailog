package lexer

import "unicode"

func isAlphaNum(c rune) bool {
	return isAlpha(c) || isNum(c)
}

func isRune(expected rune) func(rune) bool {
	return func(c rune) bool { return c == expected }
}

func isAlpha(c rune) bool {
	return unicode.IsLetter(c)
}

func isNum(c rune) bool {
	return unicode.IsDigit(c)
}

func startWith(runes []rune, fn func(c rune) bool) bool {
	if len(runes) == 0 {
		return false
	}

	return fn(runes[0])
}

func readWhile(runes []rune, fn func(c rune) bool) (string, []rune, bool) {
	var position int
	for i, c := range runes {
		if !fn(c) {
			break
		}

		position = i + 1
	}

	if position == 0 {
		return "", runes, false
	}

	return string(runes[:position]), runes[position:], true
}
