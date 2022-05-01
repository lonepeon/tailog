package lexer

import "fmt"

type SpecialCharacter struct {
	characters []rune
	tokenType  TokenType
}

var (
	EqualSpecialCharacter = SpecialCharacter{characters: []rune("=="), tokenType: TokenTypeEqual}
)

func (c SpecialCharacter) Matches(content []rune) bool {
	return startWith(content, c.characters)
}

func (c SpecialCharacter) Read(content []rune) (Token, []rune) {
	if len(content) == 0 {
		return newTokenEOF(), content
	}

	if !startWith(content, c.characters) {
		return newTokenIllegal(fmt.Sprintf("expecting characters to be %s", string(c.characters))), content
	}

	return newTokenEqual(), content[len(c.characters):]
}
