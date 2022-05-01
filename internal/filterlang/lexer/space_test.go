package lexer_test

import (
	"testing"

	"github.com/lonepeon/golib/testutils"
	"github.com/lonepeon/tailog/internal/filterlang/lexer"
)

func TestEatSpaces(t *testing.T) {
	type TestCase struct {
		Input     []rune
		Remaining []rune
	}

	runner := func(name string, tc TestCase) {
		t.Run(name, func(t *testing.T) {
			remaining := lexer.EatSpaces(tc.Input)
			testutils.AssertEqualString(t, string(tc.Remaining), string(remaining), "unexpected remaining string")
		})
	}

	runner("withSomeSpaces", TestCase{
		Input:     []rune("  some input"),
		Remaining: []rune("some input"),
	})

	runner("withSomeTabs", TestCase{
		Input:     []rune("\t  some input"),
		Remaining: []rune("some input"),
	})

	runner("withSomeNewlines", TestCase{
		Input:     []rune("\n\t  some input"),
		Remaining: []rune("some input"),
	})

	runner("withNoSpace", TestCase{
		Input:     []rune("some input"),
		Remaining: []rune("some input"),
	})
}
