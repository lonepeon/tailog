package lexer

import "fmt"

var (
	EqualSpecialCharacter    = SpecialCharacter{characters: []rune("=="), token: newTokenEqual()}
	NotEqualSpecialCharacter = SpecialCharacter{characters: []rune("!="), token: newTokenNotEqual()}
)

type SpecialCharacter struct {
	characters []rune
	token      Token
}

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

	return c.token, content[len(c.characters):]
}
