package lexer_test

import (
	"testing"

	"github.com/lonepeon/golib/testutils"
	"github.com/lonepeon/tailog/internal/filterlang/lexer"
)

func TestEqualSpecialCharacterMatches(t *testing.T) {
	type TestCase struct {
		Input   []rune
		Matches bool
	}

	runner := func(name string, tc TestCase) {
		t.Run(name, func(t *testing.T) {
			match := lexer.EqualSpecialCharacter.Matches(tc.Input)
			testutils.AssertEqualBool(t, tc.Matches, match, "unexpected content match")
		})
	}

	runner("startWithAlphaCharacter", TestCase{
		Input:   []rune(`== 12`),
		Matches: true,
	})

	runner("onlyOneCharacter", TestCase{
		Input:   []rune(`= 12`),
		Matches: false,
	})

	runner("onlyOneMatchingCharacter", TestCase{
		Input:   []rune(`=~ 12`),
		Matches: false,
	})

	runner("startWithAnotherSign", TestCase{
		Input:   []rune(`!= 12`),
		Matches: false,
	})
}

func TestEqualSpecialCharacterRead(t *testing.T) {
	type TestCase struct {
		Input      []rune
		TokenType  lexer.TokenType
		TokenValue string
		Remaining  []rune
	}

	runner := func(name string, tc TestCase) {
		t.Run(name, func(t *testing.T) {
			token, remaining := lexer.EqualSpecialCharacter.Read(tc.Input)
			testutils.AssertEqualString(t, tc.TokenType.String(), token.Type.String(), "unexpected token type")
			testutils.AssertEqualString(t, tc.TokenValue, token.Value, "unexpected token value")
			testutils.AssertEqualString(t, string(tc.Remaining), string(remaining), "unexpected remaining input")
		})
	}

	runner("withDoubleEquals", TestCase{
		Input:      []rune("== 12"),
		TokenType:  lexer.TokenTypeEqual,
		TokenValue: "",
		Remaining:  []rune(" 12"),
	})

	runner("withDoubleEqualsNoSpace", TestCase{
		Input:      []rune("==12"),
		TokenType:  lexer.TokenTypeEqual,
		TokenValue: "",
		Remaining:  []rune("12"),
	})

	runner("withEmpty", TestCase{
		Input:      []rune(""),
		TokenType:  lexer.TokenTypeEOF,
		TokenValue: "",
		Remaining:  []rune(""),
	})

	runner("withOneEqual", TestCase{
		Input:      []rune("= 12"),
		TokenType:  lexer.TokenTypeIllegal,
		TokenValue: "expecting characters to be ==",
		Remaining:  []rune("= 12"),
	})

	runner("withOneMatchingEqual", TestCase{
		Input:      []rune("=~ 12"),
		TokenType:  lexer.TokenTypeIllegal,
		TokenValue: "expecting characters to be ==",
		Remaining:  []rune("=~ 12"),
	})
}

func TestNotEqualSpecialCharacterMatches(t *testing.T) {
	type TestCase struct {
		Input   []rune
		Matches bool
	}

	runner := func(name string, tc TestCase) {
		t.Run(name, func(t *testing.T) {
			match := lexer.NotEqualSpecialCharacter.Matches(tc.Input)
			testutils.AssertEqualBool(t, tc.Matches, match, "unexpected content match")
		})
	}

	runner("startWithProperSign", TestCase{
		Input:   []rune(`!= 12`),
		Matches: true,
	})

	runner("onlyOneCharacter", TestCase{
		Input:   []rune(`! 12`),
		Matches: false,
	})

	runner("onlyOneMatchingCharacter", TestCase{
		Input:   []rune(`!~ 12`),
		Matches: false,
	})

	runner("startWithAnotherSign", TestCase{
		Input:   []rune(`== 12`),
		Matches: false,
	})
}

func TestNotEqualSpecialCharacterRead(t *testing.T) {
	type TestCase struct {
		Input      []rune
		TokenType  lexer.TokenType
		TokenValue string
		Remaining  []rune
	}

	runner := func(name string, tc TestCase) {
		t.Run(name, func(t *testing.T) {
			token, remaining := lexer.NotEqualSpecialCharacter.Read(tc.Input)
			testutils.AssertEqualString(t, tc.TokenType.String(), token.Type.String(), "unexpected token type")
			testutils.AssertEqualString(t, tc.TokenValue, token.Value, "unexpected token value")
			testutils.AssertEqualString(t, string(tc.Remaining), string(remaining), "unexpected remaining input")
		})
	}

	runner("withNotEquals", TestCase{
		Input:      []rune("!= 12"),
		TokenType:  lexer.TokenTypeNotEqual,
		TokenValue: "",
		Remaining:  []rune(" 12"),
	})

	runner("withNotEqualsNoSpace", TestCase{
		Input:      []rune("!=12"),
		TokenType:  lexer.TokenTypeNotEqual,
		TokenValue: "",
		Remaining:  []rune("12"),
	})

	runner("withEmpty", TestCase{
		Input:      []rune(""),
		TokenType:  lexer.TokenTypeEOF,
		TokenValue: "",
		Remaining:  []rune(""),
	})

	runner("withOneEqual", TestCase{
		Input:      []rune("= 12"),
		TokenType:  lexer.TokenTypeIllegal,
		TokenValue: "expecting characters to be !=",
		Remaining:  []rune("= 12"),
	})

	runner("withOneBang", TestCase{
		Input:      []rune("! 12"),
		TokenType:  lexer.TokenTypeIllegal,
		TokenValue: "expecting characters to be !=",
		Remaining:  []rune("! 12"),
	})

	runner("withOneMatchingBang", TestCase{
		Input:      []rune("!~ 12"),
		TokenType:  lexer.TokenTypeIllegal,
		TokenValue: "expecting characters to be !=",
		Remaining:  []rune("!~ 12"),
	})
}

func TestAndSpecialCharacterMatches(t *testing.T) {
	type TestCase struct {
		Input   []rune
		Matches bool
	}

	runner := func(name string, tc TestCase) {
		t.Run(name, func(t *testing.T) {
			match := lexer.AndSpecialCharacter.Matches(tc.Input)
			testutils.AssertEqualBool(t, tc.Matches, match, "unexpected content match")
		})
	}

	runner("startWithDoubuleAmpersand", TestCase{
		Input:   []rune(`&& 12`),
		Matches: true,
	})

	runner("onlyOneCharacter", TestCase{
		Input:   []rune(`& 12`),
		Matches: false,
	})

	runner("onlyOneMatchingCharacter", TestCase{
		Input:   []rune(`&~ 12`),
		Matches: false,
	})

	runner("startWithAnotherSign", TestCase{
		Input:   []rune(`!& 12`),
		Matches: false,
	})
}

func TestAndSpecialCharacterRead(t *testing.T) {
	type TestCase struct {
		Input      []rune
		TokenType  lexer.TokenType
		TokenValue string
		Remaining  []rune
	}

	runner := func(name string, tc TestCase) {
		t.Run(name, func(t *testing.T) {
			token, remaining := lexer.AndSpecialCharacter.Read(tc.Input)
			testutils.AssertEqualString(t, tc.TokenType.String(), token.Type.String(), "unexpected token type")
			testutils.AssertEqualString(t, tc.TokenValue, token.Value, "unexpected token value")
			testutils.AssertEqualString(t, string(tc.Remaining), string(remaining), "unexpected remaining input")
		})
	}

	runner("withDoubleAmpersands", TestCase{
		Input:      []rune("&& 12"),
		TokenType:  lexer.TokenTypeAnd,
		TokenValue: "",
		Remaining:  []rune(" 12"),
	})

	runner("withDoubleAmpersandsNoSpace", TestCase{
		Input:      []rune("&&12"),
		TokenType:  lexer.TokenTypeAnd,
		TokenValue: "",
		Remaining:  []rune("12"),
	})

	runner("withEmpty", TestCase{
		Input:      []rune(""),
		TokenType:  lexer.TokenTypeEOF,
		TokenValue: "",
		Remaining:  []rune(""),
	})

	runner("withOneAmpersand", TestCase{
		Input:      []rune("& 12"),
		TokenType:  lexer.TokenTypeIllegal,
		TokenValue: "expecting characters to be &&",
		Remaining:  []rune("& 12"),
	})

	runner("withOneMatchingAmpersand", TestCase{
		Input:      []rune("&~ 12"),
		TokenType:  lexer.TokenTypeIllegal,
		TokenValue: "expecting characters to be &&",
		Remaining:  []rune("&~ 12"),
	})
}

func TestOrSpecialCharacterMatches(t *testing.T) {
	type TestCase struct {
		Input   []rune
		Matches bool
	}

	runner := func(name string, tc TestCase) {
		t.Run(name, func(t *testing.T) {
			match := lexer.OrSpecialCharacter.Matches(tc.Input)
			testutils.AssertEqualBool(t, tc.Matches, match, "unexpected content match")
		})
	}

	runner("startWithDoubulePipe", TestCase{
		Input:   []rune(`|| 12`),
		Matches: true,
	})

	runner("onlyOneCharacter", TestCase{
		Input:   []rune(`| 12`),
		Matches: false,
	})

	runner("onlyOneMatchingCharacter", TestCase{
		Input:   []rune(`|~ 12`),
		Matches: false,
	})

	runner("startWithAnotherSign", TestCase{
		Input:   []rune(`=| 12`),
		Matches: false,
	})
}

func TestOrSpecialCharacterRead(t *testing.T) {
	type TestCase struct {
		Input      []rune
		TokenType  lexer.TokenType
		TokenValue string
		Remaining  []rune
	}

	runner := func(name string, tc TestCase) {
		t.Run(name, func(t *testing.T) {
			token, remaining := lexer.OrSpecialCharacter.Read(tc.Input)
			testutils.AssertEqualString(t, tc.TokenType.String(), token.Type.String(), "unexpected token type")
			testutils.AssertEqualString(t, tc.TokenValue, token.Value, "unexpected token value")
			testutils.AssertEqualString(t, string(tc.Remaining), string(remaining), "unexpected remaining input")
		})
	}

	runner("withDoublePipe", TestCase{
		Input:      []rune("|| 12"),
		TokenType:  lexer.TokenTypeOr,
		TokenValue: "",
		Remaining:  []rune(" 12"),
	})

	runner("withDoublePipeNoSpace", TestCase{
		Input:      []rune("||12"),
		TokenType:  lexer.TokenTypeOr,
		TokenValue: "",
		Remaining:  []rune("12"),
	})

	runner("withEmpty", TestCase{
		Input:      []rune(""),
		TokenType:  lexer.TokenTypeEOF,
		TokenValue: "",
		Remaining:  []rune(""),
	})

	runner("withOnePipe", TestCase{
		Input:      []rune("| 12"),
		TokenType:  lexer.TokenTypeIllegal,
		TokenValue: "expecting characters to be ||",
		Remaining:  []rune("| 12"),
	})

	runner("withOneMatchingPipe", TestCase{
		Input:      []rune("|~ 12"),
		TokenType:  lexer.TokenTypeIllegal,
		TokenValue: "expecting characters to be ||",
		Remaining:  []rune("|~ 12"),
	})
}
