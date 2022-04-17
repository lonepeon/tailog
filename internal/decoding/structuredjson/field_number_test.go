package structuredjson_test

import (
	"testing"

	"github.com/lonepeon/golib/testutils"

	"github.com/lonepeon/tailog/internal/decoding/structuredjson"
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
