package lexer

import (
	"errors"
	"unicode"
)

var (
	ErrEOF                              = errors.New("end of stream")
	ErrUnexpectedSimpleIdentifier       = errors.New("expecting identifier [:alpha:][alphanum]*")
	ErrUnexpectedDoubleQuotesIdentifier = errors.New(`expecting identifier "[^"]+"`)
)

func EatSpaces(content string) (string, error) {
	for i, c := range content {
		if unicode.IsSpace(c) {
			continue
		}

		return content[i:], nil
	}

	return "", ErrEOF
}

func DoubleQuotesIdentifier(content string) (string, string, error) {
	if content == "" {
		return "", "", ErrEOF
	}

	if !startWith(content, func(c rune) bool { return c == '"' }) {
		return "", "", ErrUnexpectedDoubleQuotesIdentifier
	}

	identifier, remaining, found := readWhile(content[1:], func(c rune) bool {
		return c != '"'
	})

	if !found {
		return "", "", ErrUnexpectedDoubleQuotesIdentifier
	}

	if !startWith(remaining, func(c rune) bool { return c == '"' }) {
		return "", "", ErrUnexpectedDoubleQuotesIdentifier
	}

	return identifier, remaining[1:], nil
}

func SimpleIdentifier(content string) (string, string, error) {
	if content == "" {
		return "", "", ErrEOF
	}

	if !startWith(content, isAlpha) {
		return "", "", ErrUnexpectedSimpleIdentifier
	}

	identifier, remaining, found := readWhile(content, func(c rune) bool {
		return isAlphaNum(c) || c == '_'
	})

	if !found {
		return "", "", ErrUnexpectedSimpleIdentifier
	}

	return identifier, remaining, nil
}

func isAlphaNum(c rune) bool {
	return isAlpha(c) || isNum(c)
}

func isAlpha(c rune) bool {
	return unicode.IsLetter(c)
}

func isNum(c rune) bool {
	return unicode.IsDigit(c)
}

func startWith(s string, fn func(c rune) bool) bool {
	runes := []rune(s)
	if len(runes) == 0 {
		return false
	}

	return fn(runes[0])
}

func readWhile(s string, fn func(c rune) bool) (string, string, bool) {
	var position int
	for i, c := range s {
		if !fn(c) {
			break
		}

		position = i + 1
	}

	if position == 0 {
		return "", "", false
	}

	return s[:position], s[position:], true
}
