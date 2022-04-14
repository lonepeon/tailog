package structuredjson_test

import (
	"testing"

	"github.com/lonepeon/golib/testutils"

	"github.com/lonepeon/tailog/internal/structuredjson"
)

func TestFieldNumberScanSuccessInt(t *testing.T) {
	field, err := structuredjson.ScanFieldNumber("msg", []byte(`12`))

	testutils.RequireNoError(t, err, "expecting to scan the log")
	testutils.AssertEqualString(t, "12", field.Value(), "invalid field value")
	testutils.AssertEqualString(t, "msg", field.Name(), "invalid field name")
}

func TestFieldNumberScanSuccessFloat(t *testing.T) {
	field, err := structuredjson.ScanFieldNumber("msg", []byte(`12.45`))

	testutils.RequireNoError(t, err, "expecting to scan the log")
	testutils.AssertEqualString(t, "12.45", field.Value(), "invalid field value")
	testutils.AssertEqualString(t, "msg", field.Name(), "invalid field name")
}

func TestFieldNumberScanError(t *testing.T) {
	_, err := structuredjson.ScanFieldNumber("msg", []byte(`"boom"`))

	testutils.RequireHasError(t, err, "expecting to not scan the log")
	testutils.AssertContainsString(t, `value="boom"`, err.Error(), "unexpected error message")
}

func TestFieldNumberEqualSuccess(t *testing.T) {
	field, err := structuredjson.ScanFieldNumber("msg", []byte(`12`))
	testutils.RequireNoError(t, err, "expecting to scan the log")

	type TestCase struct {
		OtherValue interface{}
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

	runner("equalInt", TestCase{
		OtherValue: 12,
		Expected:   structuredjson.FieldComparisonEqual,
	})

	runner("lessThanInt", TestCase{
		OtherValue: 14,
		Expected:   structuredjson.FieldComparisonLessThan,
	})

	runner("greaterThanInt", TestCase{
		OtherValue: 5,
		Expected:   structuredjson.FieldComparisonGreaterThan,
	})

	runner("equalFloat", TestCase{
		OtherValue: 12.0,
		Expected:   structuredjson.FieldComparisonEqual,
	})

	runner("lessThanFloat", TestCase{
		OtherValue: 14.5,
		Expected:   structuredjson.FieldComparisonLessThan,
	})

	runner("greaterThanFloat", TestCase{
		OtherValue: 5.8,
		Expected:   structuredjson.FieldComparisonGreaterThan,
	})
}
