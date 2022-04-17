package decoding_test

import (
	"testing"

	"github.com/lonepeon/golib/testutils"

	"github.com/lonepeon/tailog/internal/decoding"
)

func TestFieldStringCompareString(t *testing.T) {

	type TestCase struct {
		OtherValue string
		Expected   decoding.FieldComparison
	}

	runner := func(name string, tc TestCase) {
		t.Run(name, func(t *testing.T) {
			field := decoding.NewFieldString("msg", "something")

			testutils.AssertEqualString(
				t,
				tc.Expected.String(),
				field.Compare(tc.OtherValue).String(),
				"expecting successful comparison",
			)
		})
	}

	runner("equal", TestCase{
		OtherValue: "something",
		Expected:   decoding.FieldComparisonEqual,
	})

	runner("lessThan", TestCase{
		OtherValue: "xxxxx",
		Expected:   decoding.FieldComparisonLessThan,
	})

	runner("greaterThan", TestCase{
		OtherValue: "aaaaa",
		Expected:   decoding.FieldComparisonGreaterThan,
	})
}
