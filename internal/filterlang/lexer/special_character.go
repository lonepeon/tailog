package lexer

import "fmt"

var (
	EqualSpecialCharacter    = SpecialCharacter{characters: []rune("=="), token: NewTokenEqual()}
	NotEqualSpecialCharacter = SpecialCharacter{characters: []rune("!="), token: NewTokenNotEqual()}
	AndSpecialCharacter      = SpecialCharacter{characters: []rune("&&"), token: NewTokenAnd()}
	OrSpecialCharacter       = SpecialCharacter{characters: []rune("||"), token: NewTokenOr()}
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
		return NewTokenEOF(), content
	}

	if !startWith(content, c.characters) {
		return NewTokenIllegal(fmt.Sprintf("expecting characters to be %s", string(c.characters))), content
	}

	return c.token, content[len(c.characters):]
}
