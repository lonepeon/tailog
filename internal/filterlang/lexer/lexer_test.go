package lexer_test

import (
	"testing"

	"github.com/lonepeon/golib/testutils"
	"github.com/lonepeon/tailog/internal/filterlang/lexer"
)

func TestLexIdentifierNoQuotesSuccess(t *testing.T) {
	type TestCase struct {
		Input      string
		Identifier string
		Remaining  string
	}

	runner := func(name string, tc TestCase) {
		t.Run(name, func(t *testing.T) {
			identifier, remaining, err := lexer.SimpleIdentifier(tc.Input)
			testutils.RequireNoError(t, err, "can't parse identifier %s", tc.Input)
			testutils.AssertEqualString(t, tc.Identifier, identifier, "unexpected identifier")
			testutils.AssertEqualString(t, tc.Remaining, remaining, "unexpected remaining input")
		})
	}

	runner("alphaLowercaseCharacters", TestCase{
		Input:      "name is great",
		Identifier: "name",
		Remaining:  " is great",
	})

	runner("alphaUppercaseCharacters", TestCase{
		Input:      "NAME IS GREAT",
		Identifier: "NAME",
		Remaining:  " IS GREAT",
	})

	runner("alphanumCharacters", TestCase{
		Input:      "Name42 is great",
		Identifier: "Name42",
		Remaining:  " is great",
	})

	runner("alphanumAndUnderscoreCharacters", TestCase{
		Input:      "Name_42 is great",
		Identifier: "Name_42",
		Remaining:  " is great",
	})

	runner("alphanumAndDashCharacters", TestCase{
		Input:      "Name-42 is great",
		Identifier: "Name",
		Remaining:  "-42 is great",
	})

	runner("alphanumAndDotCharacters", TestCase{
		Input:      "Name.42 is great",
		Identifier: "Name",
		Remaining:  ".42 is great",
	})
}

func TestLexIdentifierNoQuotesErrors(t *testing.T) {
	type TestCase struct {
		Input string
		Error error
	}

	runner := func(name string, tc TestCase) {
		t.Run(name, func(t *testing.T) {
			identifier, remaining, err := lexer.SimpleIdentifier(tc.Input)
			testutils.RequireHasError(t, err, "shouldn't parse identifier %s", tc.Input)
			testutils.AssertErrorIs(t, tc.Error, err, "unexpected error")
			testutils.AssertEqualString(t, "", identifier, "unexpected identifier")
			testutils.AssertEqualString(t, "", remaining, "unexpected remaining input")
		})
	}

	runner("emptyContent", TestCase{
		Input: "",
		Error: lexer.EOF,
	})

	runner("startWithNumericCharacter", TestCase{
		Input: "42Name is not great",
		Error: lexer.ErrUnexpectedIdentifier,
	})

	runner("startWithSpecialCharacter", TestCase{
		Input: ">Name is not great",
		Error: lexer.ErrUnexpectedIdentifier,
	})

	runner("startWithSpace", TestCase{
		Input: " Name is not great",
		Error: lexer.ErrUnexpectedIdentifier,
	})
}
