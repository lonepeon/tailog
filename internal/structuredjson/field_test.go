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

func TestFieldStringEqualSuccess(t *testing.T) {
	field, err := structuredjson.ScanFieldString("msg", []byte(`"something"`))
	testutils.RequireNoError(t, err, "expecting to scan the log")

	testutils.AssertEqualBool(t, true, field.Equal("something"), "expecting equality")
}

func TestFieldStringEqualFailure(t *testing.T) {
	field, err := structuredjson.ScanFieldString("msg", []byte(`"something"`))
	testutils.RequireNoError(t, err, "expecting to scan the log")

	testutils.AssertEqualBool(t, false, field.Equal("nope"), "expecting no value equality")
}

func TestFieldNumberScanSuccessInt(t *testing.T) {
	field, err := structuredjson.ScanFieldNumber("msg", []byte(`12`))

	testutils.RequireNoError(t, err, "expecting to scan the log")
	testutils.AssertEqualFloat64(t, 12, field.Number(), "invalid field value")
}

func TestFieldNumberScanSuccessFloat(t *testing.T) {
	field, err := structuredjson.ScanFieldNumber("msg", []byte(`12.45`))

	testutils.RequireNoError(t, err, "expecting to scan the log")
	testutils.AssertEqualFloat64(t, 12.45, field.Number(), "invalid field value")
}

func TestFieldNumberScanError(t *testing.T) {
	_, err := structuredjson.ScanFieldNumber("msg", []byte(`"boom"`))

	testutils.RequireHasError(t, err, "expecting to not scan the log")
	testutils.AssertContainsString(t, `value="boom"`, err.Error(), "unexpected error message")
}

func TestFieldNumberEqualIntSuccess(t *testing.T) {
	field, err := structuredjson.ScanFieldNumber("msg", []byte(`12`))
	testutils.RequireNoError(t, err, "expecting to scan the log")

	testutils.AssertEqualBool(t, true, field.Equal(12), "expecting equality")
}

func TestFieldNumberEqualIntFailure(t *testing.T) {
	field, err := structuredjson.ScanFieldNumber("msg", []byte(`12`))
	testutils.RequireNoError(t, err, "expecting to scan the log")

	testutils.AssertEqualBool(t, false, field.Equal(150), "expecting no value equality")
}

func TestFieldNumberEqualFloatSuccess(t *testing.T) {
	field, err := structuredjson.ScanFieldNumber("msg", []byte(`12.42`))
	testutils.RequireNoError(t, err, "expecting to scan the log")

	testutils.AssertEqualBool(t, true, field.Equal(12.42), "expecting equality")
}

func TestFieldNumberEqualFloatFailure(t *testing.T) {
	field, err := structuredjson.ScanFieldNumber("msg", []byte(`12.42`))
	testutils.RequireNoError(t, err, "expecting to scan the log")

	testutils.AssertEqualBool(t, false, field.Equal(150.12), "expecting no value equality")
}
