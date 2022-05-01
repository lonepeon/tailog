package lexer

import "unicode"

func EatSpaces(content []rune) []rune {
	for i, c := range content {
		if unicode.IsSpace(c) {
			continue
		}

		return content[i:]
	}

	return nil
}
