package lexer_test

import (
	"testing"

	"github.com/lonepeon/golib/testutils"
	"github.com/lonepeon/tailog/internal/filterlang/lexer"
)

// nolint:funlen
func TestLexFieldNoQuotesMatches(t *testing.T) {
	type TestCase struct {
		Input   []rune
		Matches bool
	}

	runner := func(name string, tc TestCase) {
		t.Run(name, func(t *testing.T) {
			match := lexer.NoQuotesField.Matches(tc.Input)
			testutils.AssertEqualBool(t, tc.Matches, match, "unexpected content match")
		})
	}

	runner("startWithAlphaCharacter", TestCase{
		Input:   []rune(`field:name`),
		Matches: true,
	})

	runner("startWithUnderscore", TestCase{
		Input:   []rune(`field:_name`),
		Matches: false,
	})

	runner("startWithNumericCharacter", TestCase{
		Input:   []rune(`field:42name`),
		Matches: false,
	})

	runner("startWithDash", TestCase{
		Input:   []rune(`field:-name`),
		Matches: false,
	})

	runner("startWithSpace", TestCase{
		Input:   []rune(`field: name`),
		Matches: false,
	})

	runner("startWithEmpty", TestCase{
		Input:   []rune(``),
		Matches: false,
	})
}

// nolint:funlen
func TestLexFieldNoQuotesRead(t *testing.T) {
	type TestCase struct {
		Input      []rune
		TokenType  lexer.TokenType
		TokenValue string
		Remaining  []rune
	}

	runner := func(name string, tc TestCase) {
		t.Run(name, func(t *testing.T) {
			token, remaining := lexer.NoQuotesField.Read(tc.Input)
			testutils.AssertEqualString(t, tc.TokenType.String(), token.Type.String(), "unexpected token type")
			testutils.AssertEqualString(t, tc.TokenValue, token.Value, "unexpected token value")
			testutils.AssertEqualString(t, string(tc.Remaining), string(remaining), "unexpected remaining input")
		})
	}

	runner("onlyField", TestCase{
		Input:      []rune(`field:name`),
		TokenType:  lexer.TokenTypeField,
		TokenValue: "name",
		Remaining:  []rune(""),
	})

	runner("alphaLowercaseCharacters", TestCase{
		Input:      []rune(`field:name is great`),
		TokenType:  lexer.TokenTypeField,
		TokenValue: "name",
		Remaining:  []rune(" is great"),
	})

	runner("alphaUppercaseCharacters", TestCase{
		Input:      []rune(`field:NAME IS GREAT`),
		TokenType:  lexer.TokenTypeField,
		TokenValue: "NAME",
		Remaining:  []rune(" IS GREAT"),
	})

	runner("alphanumCharacters", TestCase{
		Input:      []rune(`field:Name42 is great`),
		TokenType:  lexer.TokenTypeField,
		TokenValue: "Name42",
		Remaining:  []rune(" is great"),
	})

	runner("alphanumAndUnderscoreCharacters", TestCase{
		Input:      []rune(`field:Name_42 is great`),
		TokenType:  lexer.TokenTypeField,
		TokenValue: "Name_42",
		Remaining:  []rune(" is great"),
	})

	runner("alphanumAndDashCharacters", TestCase{
		Input:      []rune(`field:Name-42 is great`),
		TokenType:  lexer.TokenTypeField,
		TokenValue: "Name",
		Remaining:  []rune("-42 is great"),
	})

	runner("alphanumAndDotCharacters", TestCase{
		Input:      []rune(`field:Name.42 is great`),
		TokenType:  lexer.TokenTypeField,
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
		Input:      []rune("field:42Name is not great"),
		TokenType:  lexer.TokenTypeIllegal,
		TokenValue: "expecting field to start with an alphabetic character",
		Remaining:  []rune("field:42Name is not great"),
	})

	runner("startWithSpecialCharacter", TestCase{
		Input:      []rune("field:>Name is not great"),
		TokenType:  lexer.TokenTypeIllegal,
		TokenValue: "expecting field to start with an alphabetic character",
		Remaining:  []rune("field:>Name is not great"),
	})

	runner("startWithSpace", TestCase{
		Input:      []rune("field: Name is not great"),
		TokenType:  lexer.TokenTypeIllegal,
		TokenValue: "expecting field to start with an alphabetic character",
		Remaining:  []rune("field: Name is not great"),
	})
}

// nolint:funlen
func TestLexFieldDoubleQuotesMatches(t *testing.T) {
	type TestCase struct {
		Input   []rune
		Matches bool
	}

	runner := func(name string, tc TestCase) {
		t.Run(name, func(t *testing.T) {
			match := lexer.DoubleQuotesField.Matches(tc.Input)
			testutils.AssertEqualBool(t, tc.Matches, match, "unexpected content match")
		})
	}

	runner("startWithDoubleQuote", TestCase{
		Input:   []rune(`field:"name"`),
		Matches: true,
	})

	runner("startWithAlphaCharacter", TestCase{
		Input:   []rune(`field:name`),
		Matches: false,
	})

	runner("startWithUnderscore", TestCase{
		Input:   []rune(`field:_name`),
		Matches: false,
	})

	runner("startWithNumericCharacter", TestCase{
		Input:   []rune(`field:42name`),
		Matches: false,
	})

	runner("startWithDash", TestCase{
		Input:   []rune(`field:-name`),
		Matches: false,
	})

	runner("startWithSpace", TestCase{
		Input:   []rune(`field: name`),
		Matches: false,
	})
}

// nolint:funlen
func TestLexFieldDoubleQuotesRead(t *testing.T) {
	type TestCase struct {
		Input      []rune
		TokenType  lexer.TokenType
		TokenValue string
		Remaining  []rune
	}

	runner := func(name string, tc TestCase) {
		t.Run(name, func(t *testing.T) {
			token, remaining := lexer.DoubleQuotesField.Read(tc.Input)
			testutils.AssertEqualString(t, tc.TokenType.String(), token.Type.String(), "unexpected token type")
			testutils.AssertEqualString(t, tc.TokenValue, token.Value, "unexpected token value")
			testutils.AssertEqualString(t, string(tc.Remaining), string(remaining), "unexpected remaining input")
		})
	}

	runner("onlyField", TestCase{
		Input:      []rune(`field:"name"`),
		TokenType:  lexer.TokenTypeField,
		TokenValue: "name",
		Remaining:  []rune(""),
	})

	runner("alphaLowercaseCharacters", TestCase{
		Input:      []rune(`field:"name" is great`),
		TokenType:  lexer.TokenTypeField,
		TokenValue: "name",
		Remaining:  []rune(" is great"),
	})

	runner("alphaUppercaseCharacters", TestCase{
		Input:      []rune(`field:"NAME" IS GREAT`),
		TokenType:  lexer.TokenTypeField,
		TokenValue: "NAME",
		Remaining:  []rune(" IS GREAT"),
	})

	runner("alphanumCharacters", TestCase{
		Input:      []rune(`field:"Name42" is great`),
		TokenType:  lexer.TokenTypeField,
		TokenValue: "Name42",
		Remaining:  []rune(" is great"),
	})

	runner("alphanumAndUnderscoreCharacters", TestCase{
		Input:      []rune(`field:"Name_42" is great`),
		TokenType:  lexer.TokenTypeField,
		TokenValue: "Name_42",
		Remaining:  []rune(" is great"),
	})

	runner("alphanumAndDashCharacters", TestCase{
		Input:      []rune(`field:"Name-42" is great`),
		TokenType:  lexer.TokenTypeField,
		TokenValue: "Name-42",
		Remaining:  []rune(" is great"),
	})

	runner("alphanumAndDotCharacters", TestCase{
		Input:      []rune(`field:"Name.42" is great`),
		TokenType:  lexer.TokenTypeField,
		TokenValue: "Name.42",
		Remaining:  []rune(" is great"),
	})

	runner("alphanumAndSpaceCharacters", TestCase{
		Input:      []rune(`field:"Name 42" is great`),
		TokenType:  lexer.TokenTypeField,
		TokenValue: "Name 42",
		Remaining:  []rune(" is great"),
	})

	runner("startsWithNumCharacter", TestCase{
		Input:      []rune(`field:"42Name" is great`),
		TokenType:  lexer.TokenTypeField,
		TokenValue: "42Name",
		Remaining:  []rune(" is great"),
	})

	runner("alphanumAndSpecialCharCharacters", TestCase{
		Input:      []rune(`field:"Name/42" is great`),
		TokenType:  lexer.TokenTypeField,
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
		Input:      []rune(`field:n"ame" is not great`),
		TokenType:  lexer.TokenTypeIllegal,
		TokenValue: "expecting field to start with double quote character",
		Remaining:  []rune(`field:n"ame" is not great`),
	})

	runner("startWithNumericCharacter", TestCase{
		Input:      []rune(`field:42"Name" is not great`),
		TokenType:  lexer.TokenTypeIllegal,
		TokenValue: "expecting field to start with double quote character",
		Remaining:  []rune(`field:42"Name" is not great`),
	})

	runner("startWithSpecialCharacter", TestCase{
		Input:      []rune(`field:>"Name" is not great`),
		TokenType:  lexer.TokenTypeIllegal,
		TokenValue: "expecting field to start with double quote character",
		Remaining:  []rune(`field:>"Name" is not great`),
	})

	runner("startWithSpace", TestCase{
		Input:      []rune(`field: "Name" is not great`),
		TokenType:  lexer.TokenTypeIllegal,
		TokenValue: "expecting field to start with double quote character",
		Remaining:  []rune(`field: "Name" is not great`),
	})

	runner("neverEnds", TestCase{
		Input:      []rune(`field:"Name is not great`),
		TokenType:  lexer.TokenTypeIllegal,
		TokenValue: "expecting field to end with a double quote character",
		Remaining:  []rune(`field:"Name is not great`),
	})
}
