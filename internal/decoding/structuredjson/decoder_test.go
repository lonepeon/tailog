package structuredjson_test

import (
	"bytes"
	"fmt"
	"io"
	"strings"
	"testing"

	"github.com/lonepeon/golib/testutils"
	"github.com/lonepeon/tailog/internal/decoding/structuredjson"
)

func TestDecoderWithSuccessfulEntry(t *testing.T) {
	decoder := structuredjson.NewDecoder(strings.NewReader(`
		{ "id": 1, "message": "message 1" }
		{ "id": 2, "message": "message 2" }
	`))

	entry, err := decoder.Decode()
	testutils.RequireNoError(t, err, "didn't expect decoding error")
	testutils.RequireEqualInt(t, 2, entry.Len(), "unexpected number of fields: %v", entry.Fields())

	id, ok := entry.Field("id")
	testutils.RequireEqualBool(t, true, ok, "cannot find field id")
	testutils.AssertEqualString(t, "1", id.Value(), "invalid value for field id")

	message, ok := entry.Field("message")
	testutils.RequireEqualBool(t, true, ok, "cannot find field message")
	testutils.AssertEqualString(t, "message 1", message.Value(), "invalid value for field message")

	entry, err = decoder.Decode()
	testutils.RequireNoError(t, err, "didn't expect decoding error")
	testutils.RequireEqualInt(t, 2, entry.Len(), "unexpected number of fields: %v", entry.Fields())

	id, ok = entry.Field("id")
	testutils.RequireEqualBool(t, true, ok, "cannot find field id")
	testutils.AssertEqualString(t, "2", id.Value(), "invalid value for field id")

	message, ok = entry.Field("message")
	testutils.RequireEqualBool(t, true, ok, "cannot find field message")
	testutils.AssertEqualString(t, "message 2", message.Value(), "invalid value for field message")
}

func TestDecoderWithEOF(t *testing.T) {
	decoder := structuredjson.NewDecoder(strings.NewReader(`
		{ "id": 1, "message": "message 1" }
		{ "id": 2, "message": "message 2" }
	`))

	_, err := decoder.Decode()
	testutils.RequireNoError(t, err, "didn't expect decoding error")

	_, err = decoder.Decode()
	testutils.RequireNoError(t, err, "didn't expect decoding error")

	_, err = decoder.Decode()
	testutils.RequireHasError(t, err, "expect decoding error")
	testutils.AssertErrorIs(t, io.EOF, err, "unexpected error")
}

func TestDecoderWithInvalidContent(t *testing.T) {
	decoder := structuredjson.NewDecoder(strings.NewReader(`
		{ "id": 1, "message": "message 1" }
		ERROR: this is not a valid JSON message
		{ "id": 2, "message": "message 2" }
	`))

	_, err := decoder.Decode()
	testutils.RequireNoError(t, err, "didn't expect decoding error")

	_, err = decoder.Decode()
	testutils.RequireHasError(t, err, "expect decoding error")
	testutils.AssertErrorContains(t, "decode structured", err, "unexpected error")
	testutils.AssertErrorContains(t, "invalid character", err, "unexpected error")
}

func TestDecoderMore(t *testing.T) {
	decoder := structuredjson.NewDecoder(strings.NewReader(`
		{ "id": 1, "message": "message 1" }
		{ "id": 2, "message": "message 2" }
	`))

	testutils.AssertEqualBool(t, true, decoder.More(), "expecting data to come")

	_, err := decoder.Decode()
	testutils.RequireNoError(t, err, "didn't expect decoding error")

	testutils.AssertEqualBool(t, true, decoder.More(), "expecting more data to come")

	_, err = decoder.Decode()
	testutils.RequireNoError(t, err, "didn't expect decoding error")

	testutils.AssertEqualBool(t, false, decoder.More(), "didn't expect more data to come")
}

func TestDecoderMoreWithPause(t *testing.T) {
	builder := bytes.NewBufferString("")
	decoder := structuredjson.NewDecoder(builder)

	testutils.AssertEqualBool(t, false, decoder.More(), "didn't expect data to come")

	fmt.Fprintln(builder, `{ "id": 1, "message": "message 1" }`)

	testutils.AssertEqualBool(t, true, decoder.More(), "expecting data to come")
}
