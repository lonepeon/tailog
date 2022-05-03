package lexer_test

import (
	"testing"

	"github.com/lonepeon/golib/testutils"
	"github.com/lonepeon/tailog/internal/filterlang/lexer"
)

// nolint:funlen
func TestLexNumberMatches(t *testing.T) {
	type TestCase struct {
		Input   []rune
		Matches bool
	}

	runner := func(name string, tc TestCase) {
		t.Run(name, func(t *testing.T) {
			match := lexer.Number.Matches(tc.Input)
			testutils.AssertEqualBool(t, tc.Matches, match, "unexpected content match")
		})
	}

	runner("startWithNumericCharacter", TestCase{
		Input:   []rune(`42name`),
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
func TestLexNumberRead(t *testing.T) {
	type TestCase struct {
		Input      []rune
		TokenType  lexer.TokenType
		TokenValue string
		Remaining  []rune
	}

	runner := func(name string, tc TestCase) {
		t.Run(name, func(t *testing.T) {
			token, remaining := lexer.Number.Read(tc.Input)
			testutils.AssertEqualString(t, tc.TokenType.String(), token.Type.String(), "unexpected token type")
			testutils.AssertEqualString(t, tc.TokenValue, token.Value, "unexpected token value")
			testutils.AssertEqualString(t, string(tc.Remaining), string(remaining), "unexpected remaining input")
		})
	}

	runner("integer", TestCase{
		Input:      []rune("42 is a great number"),
		TokenType:  lexer.TokenTypeNumber,
		TokenValue: "42",
		Remaining:  []rune(" is a great number"),
	})

	runner("float", TestCase{
		Input:      []rune("13.37 is a great number"),
		TokenType:  lexer.TokenTypeNumber,
		TokenValue: "13.37",
		Remaining:  []rune(" is a great number"),
	})

	runner("emptyContent", TestCase{
		Input:      []rune(""),
		TokenType:  lexer.TokenTypeEOF,
		TokenValue: "",
		Remaining:  []rune(""),
	})

	runner("missingFractionalPart", TestCase{
		Input:      []rune("13. is a great number"),
		TokenType:  lexer.TokenTypeIllegal,
		TokenValue: "didn't detect any number after the decimal separator",
		Remaining:  []rune("13. is a great number"),
	})

	runner("startWithAlphaCharacter", TestCase{
		Input:      []rune("Name is not great"),
		TokenType:  lexer.TokenTypeIllegal,
		TokenValue: "didn't detect any integer number",
		Remaining:  []rune("Name is not great"),
	})

	runner("startWithSpecialCharacter", TestCase{
		Input:      []rune(">Name is not great"),
		TokenType:  lexer.TokenTypeIllegal,
		TokenValue: "didn't detect any integer number",
		Remaining:  []rune(">Name is not great"),
	})

	runner("startWithSpace", TestCase{
		Input:      []rune(" Name is not great"),
		TokenType:  lexer.TokenTypeIllegal,
		TokenValue: "didn't detect any integer number",
		Remaining:  []rune(" Name is not great"),
	})
}
