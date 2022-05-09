package lexer

import "fmt"

var (
	DoubleQuotesField = doubleQuotesField{}
	NoQuotesField     = noQuotesField{}

	fieldPrefix = []rune("field:")
)

type doubleQuotesField struct{}

func (i doubleQuotesField) Matches(content []rune) bool {
	return startWith(content, fieldPrefix) && startWithFn(content[len(fieldPrefix):], isRune('"'))
}

func (i doubleQuotesField) Read(content []rune) (Token, []rune) {
	if len(content) == 0 {
		return NewTokenEOF(), content
	}

	if !startWith(content, fieldPrefix) {
		return NewTokenIllegal(fmt.Sprintf("expecting field to start with '%s' prefix", string(fieldPrefix))), content
	}

	if !startWithFn(content[len(fieldPrefix):], isRune('"')) {
		return NewTokenIllegal("expecting field to start with double quote character"), content
	}

	field, remaining, found := readWhile(content[len(fieldPrefix)+1:], func(c rune) bool { return c != '"' })

	if !found {
		return NewTokenIllegal("didn't detect any field"), content
	}

	if !startWithFn(remaining, func(c rune) bool { return c == '"' }) {
		return NewTokenIllegal("expecting field to end with a double quote character"), content
	}

	return NewTokenField(field), remaining[1:]
}

type noQuotesField struct{}

func (i noQuotesField) Matches(content []rune) bool {
	return startWith(content, fieldPrefix) && startWithFn(content[len(fieldPrefix):], isAlpha)
}

func (i noQuotesField) Read(content []rune) (Token, []rune) {
	if len(content) == 0 {
		return NewTokenEOF(), content
	}

	if !startWith(content, fieldPrefix) {
		return NewTokenIllegal(fmt.Sprintf("expecting field's to start with %s prefix", string(fieldPrefix))), content
	}

	if !startWithFn(content[len(fieldPrefix):], isAlpha) {
		return NewTokenIllegal("expecting field to start with an alphabetic character"), content
	}

	field, remaining, found := readWhile(content[len(fieldPrefix):], func(c rune) bool {
		return isAlphaNum(c) || c == '_'
	})

	if !found {
		return NewTokenIllegal("didn't detect any field"), content
	}

	return NewTokenField(field), remaining
}
