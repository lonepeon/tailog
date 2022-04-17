package structuredjson_test

import (
	"testing"

	"github.com/lonepeon/golib/testutils"
	"github.com/lonepeon/tailog/internal/decoding/structuredjson"
)

func TestParseEntrySuccess(t *testing.T) {
	var entry structuredjson.Entry
	err := entry.UnmarshalJSON([]byte(`{
		"ts": 1649965739,
		"level": "INFO",
		"httpStatus": 201,
		"message": "the foo resource has been created",
		"durationMs": 1.532
	}`))

	testutils.RequireNoError(t, err, "didn't expect any parsing error")

	testutils.RequireEqualInt(t, 5, entry.Len(), "unexpected number of fields: %v", entry.Fields())

	field, ok := entry.Field("ts")
	testutils.RequireEqualBool(t, true, ok, "cannot find field ts")
	testutils.AssertEqualString(t, "eq", field.Compare(1649965739).String(), "wrong ts field value")

	field, ok = entry.Field("level")
	testutils.RequireEqualBool(t, true, ok, "cannot find field level")
	testutils.AssertEqualString(t, "eq", field.Compare("INFO").String(), "wrong level field value")

	field, ok = entry.Field("httpStatus")
	testutils.RequireEqualBool(t, true, ok, "cannot find field httpStatus")
	testutils.AssertEqualString(t, "eq", field.Compare(201).String(), "wrong httpStatus field value")

	field, ok = entry.Field("message")
	testutils.RequireEqualBool(t, true, ok, "cannot find field message")
	testutils.AssertEqualString(t, "eq", field.Compare("the foo resource has been created").String(), "wrong message field value")

	field, ok = entry.Field("durationMs")
	testutils.RequireEqualBool(t, true, ok, "cannot find field durationMs")
	testutils.AssertEqualString(t, "eq", field.Compare(1.532).String(), "wrong durationMs field value")
}

func TestParseEntryMalformedJSON(t *testing.T) {
	var entry structuredjson.Entry
	err := entry.UnmarshalJSON([]byte(`{ "this is not a JSON" }`))
	testutils.RequireHasError(t, err, "expecting a parsing error")
	testutils.AssertContainsString(t, "unmarshal json", err.Error(), "expecting a parsing error")
}

func TestParseEntryUnsupportedValue(t *testing.T) {
	var entry structuredjson.Entry
	err := entry.UnmarshalJSON([]byte(`{ "array": [1,2,3] }`))
	testutils.RequireHasError(t, err, "expecting a parsing error")
	testutils.AssertContainsString(t, "parse field", err.Error(), "expecting a parsing error")
	testutils.AssertContainsString(t, "array", err.Error(), "expecting a parsing error")
}
