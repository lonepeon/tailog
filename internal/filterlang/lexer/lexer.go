package lexer

import (
	"errors"
	"unicode"
)

var (
	EOF                     = errors.New("end of stream")
	ErrUnexpectedIdentifier = errors.New("expecting identifier [:alpha:][alphanum]*")
)

func SimpleIdentifier(content string) (string, string, error) {
	if content == "" {
		return "", "", EOF
	}

	if !isAlpha([]rune(content)[0]) {
		return "", "", ErrUnexpectedIdentifier
	}

	var position int
	for i, c := range content {
		if !isAlphaNum(c) && c != '_' {
			break
		}

		position = i + 1
	}

	if position == 0 {
		return "", "", ErrUnexpectedIdentifier
	}

	return content[:position], content[position:], nil
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
