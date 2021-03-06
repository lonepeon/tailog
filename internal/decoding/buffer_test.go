package decoding_test

import (
	"testing"

	"github.com/lonepeon/golib/testutils"
	"github.com/lonepeon/tailog/internal/decoding"
	"github.com/lonepeon/tailog/internal/decoding/decodingtest"
)

func TestCircularBufferLenNoData(t *testing.T) {
	buf := decoding.NewCircularBuffer(3)
	testutils.AssertEqualInt(t, 0, buf.Len(), "invalid buffer size")
}

func TestCircularBufferLenNotFilled(t *testing.T) {
	buf := decoding.NewCircularBuffer(3)

	buf.Push(decodingtest.NewEntry(t, map[string]interface{}{
		"label1": "value 1",
		"label2": "value 2",
	}))
	testutils.AssertEqualInt(t, 1, buf.Len(), "invalid buffer size")

	buf.Push(decodingtest.NewEntry(t, map[string]interface{}{
		"label1": "value 3",
		"label2": "value 4",
	}))
	testutils.AssertEqualInt(t, 2, buf.Len(), "invalid buffer size")
}

func TestCircularBufferLenFilled(t *testing.T) {
	buf := decoding.NewCircularBuffer(3)

	buf.Push(decodingtest.NewEntry(t, map[string]interface{}{
		"label1": "value 1",
		"label2": "value 2",
	}))
	testutils.AssertEqualInt(t, 1, buf.Len(), "invalid buffer size")

	buf.Push(decodingtest.NewEntry(t, map[string]interface{}{
		"label1": "value 3",
		"label2": "value 4",
	}))
	testutils.AssertEqualInt(t, 2, buf.Len(), "invalid buffer size")

	buf.Push(decodingtest.NewEntry(t, map[string]interface{}{
		"label1": "value 5",
		"label2": "value 6",
	}))
	testutils.AssertEqualInt(t, 3, buf.Len(), "invalid buffer size")

	buf.Push(decodingtest.NewEntry(t, map[string]interface{}{
		"label1": "value 7",
		"label2": "value 8",
	}))
	testutils.AssertEqualInt(t, 3, buf.Len(), "invalid buffer size")
}

func TestCircularBufferAtNotNegative(t *testing.T) {
	buf := decoding.NewCircularBuffer(3)
	buf.Push(decodingtest.NewEntry(t, map[string]interface{}{"label1": "value 1"}))
	buf.Push(decodingtest.NewEntry(t, map[string]interface{}{"label2": "value 2"}))

	entry, found := buf.At(-1)
	testutils.AssertEqualBool(t, false, found, "unexpected entry at index %s", entry)
}

func TestCircularBufferAtTooBig(t *testing.T) {
	buf := decoding.NewCircularBuffer(3)
	buf.Push(decodingtest.NewEntry(t, map[string]interface{}{"label1": "value 1"}))
	buf.Push(decodingtest.NewEntry(t, map[string]interface{}{"label2": "value 2"}))

	entry, found := buf.At(2)
	testutils.AssertEqualBool(t, false, found, "unexpected entry at index %s", entry)
}

func TestCircularBufferAtValidNotFilled(t *testing.T) {
	buf := decoding.NewCircularBuffer(3)
	buf.Push(decodingtest.NewEntry(t, map[string]interface{}{"label1": "value 1"}))
	buf.Push(decodingtest.NewEntry(t, map[string]interface{}{"label2": "value 2"}))

	entry, found := buf.At(1)
	testutils.RequireEqualBool(t, true, found, "expecting entry to be found")

	field, found := entry.Field("label2")
	testutils.RequireEqualBool(t, true, found, "expecting field to be found")
	testutils.AssertEqualString(t, "value 2", field.Value(), "unexpected field value")
}

func TestCircularBufferAtValidFilled(t *testing.T) {
	buf := decoding.NewCircularBuffer(3)
	buf.Push(decodingtest.NewEntry(t, map[string]interface{}{"label1": "value 1"}))
	buf.Push(decodingtest.NewEntry(t, map[string]interface{}{"label2": "value 2"}))
	buf.Push(decodingtest.NewEntry(t, map[string]interface{}{"label3": "value 3"}))
	buf.Push(decodingtest.NewEntry(t, map[string]interface{}{"label4": "value 4"}))

	entry, found := buf.At(2)
	testutils.RequireEqualBool(t, true, found, "expecting entry to be found")

	field, found := entry.Field("label4")
	testutils.RequireEqualBool(t, true, found, "expecting field to be found")
	testutils.AssertEqualString(t, "value 4", field.Value(), "unexpected field value")
}

func TestCircularBufferPush(t *testing.T) {
	buf := decoding.NewCircularBuffer(3)

	evicted := buf.Push(decodingtest.NewEntry(t, map[string]interface{}{"label1": "value 1"}))
	testutils.AssertEqualBool(t, false, evicted, "didn't expect value to be evicted")

	evicted = buf.Push(decodingtest.NewEntry(t, map[string]interface{}{"label2": "value 2"}))
	testutils.AssertEqualBool(t, false, evicted, "didn't expect value to be evicted")

	evicted = buf.Push(decodingtest.NewEntry(t, map[string]interface{}{"label3": "value 3"}))
	testutils.AssertEqualBool(t, false, evicted, "didn't expect value to be evicted")

	evicted = buf.Push(decodingtest.NewEntry(t, map[string]interface{}{"label4": "value 4"}))
	testutils.AssertEqualBool(t, true, evicted, "didn't expect value to be evicted")

	evicted = buf.Push(decodingtest.NewEntry(t, map[string]interface{}{"label4": "value 5"}))
	testutils.AssertEqualBool(t, true, evicted, "didn't expect value to be evicted")
}

func TestCircularBufferPushAddingFilter(t *testing.T) {
	buf := decoding.NewCircularBuffer(3)
	buf.Push(decodingtest.NewEntry(t, map[string]interface{}{"label": "keep-me"}))
	buf.Push(decodingtest.NewEntry(t, map[string]interface{}{"label": "hide-me"}))
	buf.Push(decodingtest.NewEntry(t, map[string]interface{}{"label": "keep-me"}))
	testutils.RequireEqualInt(t, 3, buf.Len(), "unexpected buffer size")

	buf.ReplaceFilter(func(e decoding.Entry) bool {
		field, ok := e.Field("label")
		testutils.RequireEqualBool(t, true, ok, "didn't expect an entry with no label field")
		return field.Compare("keep-me") == decoding.FieldComparisonEqual
	})

	testutils.RequireEqualInt(t, 2, buf.Len(), "unexpected buffer size after filter")

	entry, found := buf.At(0)
	testutils.RequireEqualBool(t, true, found, "expecting entry to be found")
	field, found := entry.Field("label")
	testutils.RequireEqualBool(t, true, found, "expecting field to be found")
	testutils.AssertEqualString(t, "keep-me", field.Value(), "unexpected field value")

	entry, found = buf.At(1)
	testutils.RequireEqualBool(t, true, found, "expecting entry to be found")
	field, found = entry.Field("label")
	testutils.RequireEqualBool(t, true, found, "expecting field to be found")
	testutils.AssertEqualString(t, "keep-me", field.Value(), "unexpected field value")
}
