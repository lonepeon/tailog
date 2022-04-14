package structuredjson_test

import (
	"testing"

	"github.com/lonepeon/golib/testutils"

	"github.com/lonepeon/tailog/internal/structuredjson"
)

func TestFieldStringScanSuccess(t *testing.T) {
	field, err := structuredjson.ScanFieldString("msg", []byte(`"a log message"`))

	testutils.RequireNoError(t, err, "expecting to scan the log")
	testutils.AssertEqualString(t, "a log message", field.String(), "invalid field value")
}

func TestFieldStringScanError(t *testing.T) {
	_, err := structuredjson.ScanFieldString("msg", []byte(`12`))

	testutils.RequireHasError(t, err, "expecting to not scan the log")
	testutils.AssertContainsString(t, "value=12", err.Error(), "unexpected error message")
}

func TestFieldStringCompareString(t *testing.T) {
	field, err := structuredjson.ScanFieldString("msg", []byte(`"something"`))
	testutils.RequireNoError(t, err, "expecting to scan the log")

	type TestCase struct {
		OtherValue string
		Expected   structuredjson.FieldComparison
	}

	runner := func(name string, tc TestCase) {
		t.Run(name, func(t *testing.T) {
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
		Expected:   structuredjson.FieldComparisonEqual,
	})

	runner("lessThan", TestCase{
		OtherValue: "xxxxx",
		Expected:   structuredjson.FieldComparisonLessThan,
	})

	runner("greaterThan", TestCase{
		OtherValue: "aaaaa",
		Expected:   structuredjson.FieldComparisonGreaterThan,
	})
}
