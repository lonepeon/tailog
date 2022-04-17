package structuredjson_test

import (
	"testing"

	"github.com/lonepeon/golib/testutils"

	"github.com/lonepeon/tailog/internal/decoding"
	"github.com/lonepeon/tailog/internal/decoding/structuredjson"
)

func TestFieldStringScanSuccess(t *testing.T) {
	field, err := structuredjson.ScanFieldString("msg", []byte(`"a log message"`))

	testutils.RequireNoError(t, err, "expecting to scan the log")
	testutils.AssertEqualString(t, "a log message", field.Value(), "invalid field value")
	testutils.AssertEqualString(t, "msg", field.Name(), "invalid field name")
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
		Expected   decoding.FieldComparison
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
