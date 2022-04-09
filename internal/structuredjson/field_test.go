package structuredjson_test

import (
	"testing"

	"github.com/lonepeon/golib/testutils"

	"github.com/lonepeon/tailog/internal/structuredjson"
)

func TestFieldStringScanSuccess(t *testing.T) {
	field := structuredjson.NewFieldString("msg")

	err := field.Scan([]byte(`"a log message"`))

	testutils.RequireNoError(t, err, "expecting to scan the log")
	testutils.AssertEqualString(t, "a log message", field.String(), "invalid field value")
}

func TestFieldStringScanError(t *testing.T) {
	field := structuredjson.NewFieldString("msg")

	err := field.Scan([]byte(`12`))

	testutils.RequireHasError(t, err, "expecting to not scan the log")
	testutils.AssertContainsString(t, "value=12", err.Error(), "unexpected error message")
}

func TestFieldNumberScanSuccessInt(t *testing.T) {
	field := structuredjson.NewFieldNumber("msg")

	err := field.Scan([]byte(`12`))

	testutils.RequireNoError(t, err, "expecting to scan the log")
	testutils.AssertEqualFloat64(t, 12, field.Number(), "invalid field value")
}

func TestFieldNumberScanSuccessFloat(t *testing.T) {
	field := structuredjson.NewFieldNumber("msg")

	err := field.Scan([]byte(`12.45`))

	testutils.RequireNoError(t, err, "expecting to scan the log")
	testutils.AssertEqualFloat64(t, 12.45, field.Number(), "invalid field value")
}

func TestFieldNumberScanError(t *testing.T) {
	field := structuredjson.NewFieldNumber("msg")

	err := field.Scan([]byte(`"boom"`))

	testutils.RequireHasError(t, err, "expecting to not scan the log")
	testutils.AssertContainsString(t, `value="boom"`, err.Error(), "unexpected error message")
}