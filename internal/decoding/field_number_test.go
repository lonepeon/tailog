package decoding_test

import (
	"testing"

	"github.com/lonepeon/golib/testutils"

	"github.com/lonepeon/tailog/internal/decoding"
)

func TestFieldNumberEqualSuccess(t *testing.T) {
	type TestCase struct {
		Field      decoding.FieldNumber
		OtherValue interface{}
		Expected   decoding.FieldComparison
	}

	runner := func(name string, tc TestCase) {
		t.Run(name, func(t *testing.T) {
			testutils.AssertEqualString(
				t,
				tc.Expected.String(),
				tc.Field.Compare(tc.OtherValue).String(),
				"expecting successful comparison",
			)
		})
	}

	runner("equalInt", TestCase{
		Field:      decoding.NewFieldNumber("label", 12),
		OtherValue: 12,
		Expected:   decoding.FieldComparisonEqual,
	})

	runner("lessThanInt", TestCase{
		Field:      decoding.NewFieldNumber("label", 12),
		OtherValue: 14,
		Expected:   decoding.FieldComparisonLessThan,
	})

	runner("greaterThanInt", TestCase{
		Field:      decoding.NewFieldNumber("label", 12),
		OtherValue: 5,
		Expected:   decoding.FieldComparisonGreaterThan,
	})

	runner("equalFloat", TestCase{
		Field:      decoding.NewFieldNumber("label", 12),
		OtherValue: 12.0,
		Expected:   decoding.FieldComparisonEqual,
	})

	runner("lessThanFloat", TestCase{
		Field:      decoding.NewFieldNumber("label", 12),
		OtherValue: 14.5,
		Expected:   decoding.FieldComparisonLessThan,
	})

	runner("greaterThanFloat", TestCase{
		Field:      decoding.NewFieldNumber("label", 12),
		OtherValue: 5.8,
		Expected:   decoding.FieldComparisonGreaterThan,
	})
}
