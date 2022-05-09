package lexer_test

import (
	"testing"

	"github.com/lonepeon/golib/testutils"
	"github.com/lonepeon/tailog/internal/filterlang/lexer"
)

// nolint:funlen
func TestLexStringMatches(t *testing.T) {
	type TestCase struct {
		Input   []rune
		Matches bool
	}

	runner := func(name string, tc TestCase) {
		t.Run(name, func(t *testing.T) {
			match := lexer.String.Matches(tc.Input)
			testutils.AssertEqualBool(t, tc.Matches, match, "unexpected content match")
		})
	}

	runner("startWithDoubleQuote", TestCase{
		Input:   []rune(`"name"`),
		Matches: true,
	})

	runner("startWithAlphaCharacter", TestCase{
		Input:   []rune(`name`),
		Matches: false,
	})

	runner("startWithUnderscore", TestCase{
		Input:   []rune(`_name`),
		Matches: false,
	})

	runner("startWithNumericCharacter", TestCase{
		Input:   []rune(`42name`),
		Matches: false,
	})

	runner("startWithDash", TestCase{
		Input:   []rune(`-name`),
		Matches: false,
	})

	runner("startWithSpace", TestCase{
		Input:   []rune(` name`),
		Matches: false,
	})
}

// nolint:funlen
func TestLexStringRead(t *testing.T) {
	type TestCase struct {
		Input      []rune
		TokenType  lexer.TokenType
		TokenValue string
		Remaining  []rune
	}

	runner := func(name string, tc TestCase) {
		t.Run(name, func(t *testing.T) {
			token, remaining := lexer.String.Read(tc.Input)
			testutils.AssertEqualString(t, tc.TokenType.String(), token.Type.String(), "unexpected token type")
			testutils.AssertEqualString(t, tc.TokenValue, token.Value, "unexpected token value")
			testutils.AssertEqualString(t, string(tc.Remaining), string(remaining), "unexpected remaining input")
		})
	}

	runner("onlyString", TestCase{
		Input:      []rune(`"name"`),
		TokenType:  lexer.TokenTypeString,
		TokenValue: "name",
		Remaining:  []rune(""),
	})

	runner("alphaLowercaseCharacters", TestCase{
		Input:      []rune(`"name" is great`),
		TokenType:  lexer.TokenTypeString,
		TokenValue: "name",
		Remaining:  []rune(" is great"),
	})

	runner("alphaUppercaseCharacters", TestCase{
		Input:      []rune(`"NAME" IS GREAT`),
		TokenType:  lexer.TokenTypeString,
		TokenValue: "NAME",
		Remaining:  []rune(" IS GREAT"),
	})

	runner("alphanumCharacters", TestCase{
		Input:      []rune(`"Name42" is great`),
		TokenType:  lexer.TokenTypeString,
		TokenValue: "Name42",
		Remaining:  []rune(" is great"),
	})

	runner("alphanumAndUnderscoreCharacters", TestCase{
		Input:      []rune(`"Name_42" is great`),
		TokenType:  lexer.TokenTypeString,
		TokenValue: "Name_42",
		Remaining:  []rune(" is great"),
	})

	runner("alphanumAndDashCharacters", TestCase{
		Input:      []rune(`"Name-42" is great`),
		TokenType:  lexer.TokenTypeString,
		TokenValue: "Name-42",
		Remaining:  []rune(" is great"),
	})

	runner("alphanumAndDotCharacters", TestCase{
		Input:      []rune(`"Name.42" is great`),
		TokenType:  lexer.TokenTypeString,
		TokenValue: "Name.42",
		Remaining:  []rune(" is great"),
	})

	runner("alphanumAndSpaceCharacters", TestCase{
		Input:      []rune(`"Name 42" is great`),
		TokenType:  lexer.TokenTypeString,
		TokenValue: "Name 42",
		Remaining:  []rune(" is great"),
	})

	runner("startsWithNumCharacter", TestCase{
		Input:      []rune(`"42Name" is great`),
		TokenType:  lexer.TokenTypeString,
		TokenValue: "42Name",
		Remaining:  []rune(" is great"),
	})

	runner("alphanumAndSpecialCharCharacters", TestCase{
		Input:      []rune(`"Name/42" is great`),
		TokenType:  lexer.TokenTypeString,
		TokenValue: "Name/42",
		Remaining:  []rune(" is great"),
	})

	runner("emptyContent", TestCase{
		Input:      []rune(""),
		TokenType:  lexer.TokenTypeEOF,
		TokenValue: "",
		Remaining:  []rune(""),
	})

	runner("startWithAlphaCharacter", TestCase{
		Input:      []rune(`n"ame" is not great`),
		TokenType:  lexer.TokenTypeIllegal,
		TokenValue: "expecting string to start with double quote character",
		Remaining:  []rune(`n"ame" is not great`),
	})

	runner("startWithNumericCharacter", TestCase{
		Input:      []rune(`42"Name" is not great`),
		TokenType:  lexer.TokenTypeIllegal,
		TokenValue: "expecting string to start with double quote character",
		Remaining:  []rune(`42"Name" is not great`),
	})

	runner("startWithSpecialCharacter", TestCase{
		Input:      []rune(`>"Name" is not great`),
		TokenType:  lexer.TokenTypeIllegal,
		TokenValue: "expecting string to start with double quote character",
		Remaining:  []rune(`>"Name" is not great`),
	})

	runner("startWithSpace", TestCase{
		Input:      []rune(` "Name" is not great`),
		TokenType:  lexer.TokenTypeIllegal,
		TokenValue: "expecting string to start with double quote character",
		Remaining:  []rune(` "Name" is not great`),
	})

	runner("neverEnds", TestCase{
		Input:      []rune(`"Name is not great`),
		TokenType:  lexer.TokenTypeIllegal,
		TokenValue: "expecting string to end with a double quote character",
		Remaining:  []rune(`"Name is not great`),
	})
}
