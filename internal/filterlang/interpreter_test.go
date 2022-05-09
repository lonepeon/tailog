package filterlang_test

import (
	"testing"

	"github.com/lonepeon/golib/testutils"
	"github.com/lonepeon/tailog/internal/decoding"
	"github.com/lonepeon/tailog/internal/decoding/decodingtest"
	"github.com/lonepeon/tailog/internal/filterlang"
)

// nolint:funlen
func TestExecute(t *testing.T) {
	type Testcase struct {
		Entry          decoding.Entry
		ExpectedResult bool
	}

	source := `field:"http.method" == "POST" && field:"http.path" == "/plop" || field:userid == 42`
	interpreter, err := filterlang.Parse(source)
	testutils.RequireNoError(t, err, "expecting to parse source code: %v", source)

	runner := func(name string, tc Testcase) {
		t.Run(name, func(t *testing.T) {
			match := interpreter.Execute(tc.Entry)
			testutils.AssertEqualBool(t, tc.ExpectedResult, match, "unexpected execution result")
		})
	}

	runner("allIsExact", Testcase{
		Entry: decodingtest.NewEntry(t, map[string]interface{}{
			"http.method": "POST",
			"http.path":   "/plop",
			"userid":      42,
		}),
		ExpectedResult: true,
	})

	runner("orHasGreaterPrecedence", Testcase{
		Entry: decodingtest.NewEntry(t, map[string]interface{}{
			"http.method": "GET",
			"http.path":   "/plop",
			"userid":      42,
		}),
		ExpectedResult: true,
	})

	runner("orHasGreaterPrecedence2", Testcase{
		Entry: decodingtest.NewEntry(t, map[string]interface{}{
			"http.method": "POST",
			"http.path":   "/something-else",
			"userid":      42,
		}),
		ExpectedResult: true,
	})

	runner("orHasGreaterPrecedence3", Testcase{
		Entry: decodingtest.NewEntry(t, map[string]interface{}{
			"http.method": "POST",
			"http.path":   "/plop",
			"userid":      1337,
		}),
		ExpectedResult: true,
	})

	runner("orHasGreaterPrecedence4", Testcase{
		Entry: decodingtest.NewEntry(t, map[string]interface{}{
			"http.method": "POST",
			"http.path":   "/plop",
		}),
		ExpectedResult: true,
	})

	runner("conditionDontMach", Testcase{
		Entry: decodingtest.NewEntry(t, map[string]interface{}{
			"http.method": "GET",
			"http.path":   "/something-else",
		}),
		ExpectedResult: false,
	})
}
