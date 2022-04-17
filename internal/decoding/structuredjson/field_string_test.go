package structuredjson_test

import (
	"testing"

	"github.com/lonepeon/golib/testutils"

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
