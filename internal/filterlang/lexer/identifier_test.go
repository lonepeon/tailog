package lexer_test

import (
	"testing"

	"github.com/lonepeon/golib/testutils"
	"github.com/lonepeon/tailog/internal/filterlang/lexer"
)

func TestLexIdentifierEatSpaces(t *testing.T) {
	type TestCase struct {
		Input     []rune
		Remaining []rune
	}

	runner := func(name string, tc TestCase) {
		t.Run(name, func(t *testing.T) {
			remaining := lexer.EatSpaces(tc.Input)
			testutils.AssertEqualString(t, string(tc.Remaining), string(remaining), "unexpected remaining input")
		})
	}

	runner("oneSpace", TestCase{
		Input:     []rune(" a text  "),
		Remaining: []rune("a text  "),
	})

	runner("severalSpaces", TestCase{
		Input:     []rune("      \t \n a text  "),
		Remaining: []rune("a text  "),
	})

	runner("noSpaces", TestCase{
		Input:     []rune("a text  "),
		Remaining: []rune("a text  "),
	})
}

// nolint:funlen
func TestLexIdentifierNoQuotes(t *testing.T) {
	type TestCase struct {
		Input      []rune
		TokenType  lexer.TokenType
		TokenValue string
		Remaining  []rune
	}

	runner := func(name string, tc TestCase) {
		t.Run(name, func(t *testing.T) {
			token, remaining := lexer.SimpleIdentifier(tc.Input)
			testutils.AssertEqualString(t, tc.TokenType.String(), token.Type.String(), "unexpected token type")
			testutils.AssertEqualString(t, tc.TokenValue, token.Value, "unexpected token value")
			testutils.AssertEqualString(t, string(tc.Remaining), string(remaining), "unexpected remaining input")
		})
	}

	runner("onlyIdentifier", TestCase{
		Input:      []rune(`name`),
		TokenType:  lexer.TokenTypeIdentifier,
		TokenValue: "name",
		Remaining:  []rune(""),
	})

	runner("alphaLowercaseCharacters", TestCase{
		Input:      []rune(`name is great`),
		TokenType:  lexer.TokenTypeIdentifier,
		TokenValue: "name",
		Remaining:  []rune(" is great"),
	})

	runner("alphaUppercaseCharacters", TestCase{
		Input:      []rune(`NAME IS GREAT`),
		TokenType:  lexer.TokenTypeIdentifier,
		TokenValue: "NAME",
		Remaining:  []rune(" IS GREAT"),
	})

	runner("alphanumCharacters", TestCase{
		Input:      []rune(`Name42 is great`),
		TokenType:  lexer.TokenTypeIdentifier,
		TokenValue: "Name42",
		Remaining:  []rune(" is great"),
	})

	runner("alphanumAndUnderscoreCharacters", TestCase{
		Input:      []rune(`Name_42 is great`),
		TokenType:  lexer.TokenTypeIdentifier,
		TokenValue: "Name_42",
		Remaining:  []rune(" is great"),
	})

	runner("alphanumAndDashCharacters", TestCase{
		Input:      []rune(`Name-42 is great`),
		TokenType:  lexer.TokenTypeIdentifier,
		TokenValue: "Name",
		Remaining:  []rune("-42 is great"),
	})

	runner("alphanumAndDotCharacters", TestCase{
		Input:      []rune(`Name.42 is great`),
		TokenType:  lexer.TokenTypeIdentifier,
		TokenValue: "Name",
		Remaining:  []rune(".42 is great"),
	})

	runner("emptyContent", TestCase{
		Input:      []rune(""),
		TokenType:  lexer.TokenTypeEOF,
		TokenValue: "",
		Remaining:  []rune(""),
	})

	runner("startWithNumericCharacter", TestCase{
		Input:      []rune("42Name is not great"),
		TokenType:  lexer.TokenTypeIllegal,
		TokenValue: "expecting identifier to start with an alphabetic character",
		Remaining:  []rune("42Name is not great"),
	})

	runner("startWithSpecialCharacter", TestCase{
		Input:      []rune(">Name is not great"),
		TokenType:  lexer.TokenTypeIllegal,
		TokenValue: "expecting identifier to start with an alphabetic character",
		Remaining:  []rune(">Name is not great"),
	})

	runner("startWithSpace", TestCase{
		Input:      []rune(" Name is not great"),
		TokenType:  lexer.TokenTypeIllegal,
		TokenValue: "expecting identifier to start with an alphabetic character",
		Remaining:  []rune(" Name is not great"),
	})
}

// nolint:funlen
func TestLexIdentifierDoubleQuotes(t *testing.T) {
	type TestCase struct {
		Input      []rune
		TokenType  lexer.TokenType
		TokenValue string
		Remaining  []rune
	}

	runner := func(name string, tc TestCase) {
		t.Run(name, func(t *testing.T) {
			token, remaining := lexer.DoubleQuotesIdentifier(tc.Input)
			testutils.AssertEqualString(t, tc.TokenType.String(), token.Type.String(), "unexpected token type")
			testutils.AssertEqualString(t, tc.TokenValue, token.Value, "unexpected token value")
			testutils.AssertEqualString(t, string(tc.Remaining), string(remaining), "unexpected remaining input")
		})
	}

	runner("onlyIdentifier", TestCase{
		Input:      []rune(`"name"`),
		TokenType:  lexer.TokenTypeIdentifier,
		TokenValue: "name",
		Remaining:  []rune(""),
	})

	runner("alphaLowercaseCharacters", TestCase{
		Input:      []rune(`"name" is great`),
		TokenType:  lexer.TokenTypeIdentifier,
		TokenValue: "name",
		Remaining:  []rune(" is great"),
	})

	runner("alphaUppercaseCharacters", TestCase{
		Input:      []rune(`"NAME" IS GREAT`),
		TokenType:  lexer.TokenTypeIdentifier,
		TokenValue: "NAME",
		Remaining:  []rune(" IS GREAT"),
	})

	runner("alphanumCharacters", TestCase{
		Input:      []rune(`"Name42" is great`),
		TokenType:  lexer.TokenTypeIdentifier,
		TokenValue: "Name42",
		Remaining:  []rune(" is great"),
	})

	runner("alphanumAndUnderscoreCharacters", TestCase{
		Input:      []rune(`"Name_42" is great`),
		TokenType:  lexer.TokenTypeIdentifier,
		TokenValue: "Name_42",
		Remaining:  []rune(" is great"),
	})

	runner("alphanumAndDashCharacters", TestCase{
		Input:      []rune(`"Name-42" is great`),
		TokenType:  lexer.TokenTypeIdentifier,
		TokenValue: "Name-42",
		Remaining:  []rune(" is great"),
	})

	runner("alphanumAndDotCharacters", TestCase{
		Input:      []rune(`"Name.42" is great`),
		TokenType:  lexer.TokenTypeIdentifier,
		TokenValue: "Name.42",
		Remaining:  []rune(" is great"),
	})

	runner("alphanumAndSpaceCharacters", TestCase{
		Input:      []rune(`"Name 42" is great`),
		TokenType:  lexer.TokenTypeIdentifier,
		TokenValue: "Name 42",
		Remaining:  []rune(" is great"),
	})

	runner("startsWithNumCharacter", TestCase{
		Input:      []rune(`"42Name" is great`),
		TokenType:  lexer.TokenTypeIdentifier,
		TokenValue: "42Name",
		Remaining:  []rune(" is great"),
	})

	runner("alphanumAndSpecialCharCharacters", TestCase{
		Input:      []rune(`"Name/42" is great`),
		TokenType:  lexer.TokenTypeIdentifier,
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
		TokenValue: "expecting identifier's opening double quote character",
		Remaining:  []rune(`n"ame" is not great`),
	})

	runner("startWithNumericCharacter", TestCase{
		Input:      []rune(`42"Name" is not great`),
		TokenType:  lexer.TokenTypeIllegal,
		TokenValue: "expecting identifier's opening double quote character",
		Remaining:  []rune(`42"Name" is not great`),
	})

	runner("startWithSpecialCharacter", TestCase{
		Input:      []rune(`>"Name" is not great`),
		TokenType:  lexer.TokenTypeIllegal,
		TokenValue: "expecting identifier's opening double quote character",
		Remaining:  []rune(`>"Name" is not great`),
	})

	runner("startWithSpace", TestCase{
		Input:      []rune(` "Name" is not great`),
		TokenType:  lexer.TokenTypeIllegal,
		TokenValue: "expecting identifier's opening double quote character",
		Remaining:  []rune(` "Name" is not great`),
	})

	runner("neverEnds", TestCase{
		Input:      []rune(`"Name is not great`),
		TokenType:  lexer.TokenTypeIllegal,
		TokenValue: "expecting identifier's closing double quote character",
		Remaining:  []rune(`"Name is not great`),
	})
}
