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

	buf.Push(decodingtest.NewEntry(map[string]string{
		"label1": "value 1",
		"label2": "value 2",
	}))
	testutils.AssertEqualInt(t, 1, buf.Len(), "invalid buffer size")

	buf.Push(decodingtest.NewEntry(map[string]string{
		"label1": "value 3",
		"label2": "value 4",
	}))
	testutils.AssertEqualInt(t, 2, buf.Len(), "invalid buffer size")
}

func TestCircularBufferLenFilled(t *testing.T) {
	buf := decoding.NewCircularBuffer(3)

	buf.Push(decodingtest.NewEntry(map[string]string{
		"label1": "value 1",
		"label2": "value 2",
	}))
	testutils.AssertEqualInt(t, 1, buf.Len(), "invalid buffer size")

	buf.Push(decodingtest.NewEntry(map[string]string{
		"label1": "value 3",
		"label2": "value 4",
	}))
	testutils.AssertEqualInt(t, 2, buf.Len(), "invalid buffer size")

	buf.Push(decodingtest.NewEntry(map[string]string{
		"label1": "value 5",
		"label2": "value 6",
	}))
	testutils.AssertEqualInt(t, 3, buf.Len(), "invalid buffer size")

	buf.Push(decodingtest.NewEntry(map[string]string{
		"label1": "value 7",
		"label2": "value 8",
	}))
	testutils.AssertEqualInt(t, 3, buf.Len(), "invalid buffer size")
}

func TestCircularBufferAtNotNegative(t *testing.T) {
	buf := decoding.NewCircularBuffer(3)
	buf.Push(decodingtest.NewEntry(map[string]string{"label1": "value 1"}))
	buf.Push(decodingtest.NewEntry(map[string]string{"label2": "value 2"}))

	entry, found := buf.At(-1)
	testutils.AssertEqualBool(t, false, found, "unexpected entry at index %s", entry)
}

func TestCircularBufferAtTooBig(t *testing.T) {
	buf := decoding.NewCircularBuffer(3)
	buf.Push(decodingtest.NewEntry(map[string]string{"label1": "value 1"}))
	buf.Push(decodingtest.NewEntry(map[string]string{"label2": "value 2"}))

	entry, found := buf.At(2)
	testutils.AssertEqualBool(t, false, found, "unexpected entry at index %s", entry)
}

func TestCircularBufferAtValidNotFilled(t *testing.T) {
	buf := decoding.NewCircularBuffer(3)
	buf.Push(decodingtest.NewEntry(map[string]string{"label1": "value 1"}))
	buf.Push(decodingtest.NewEntry(map[string]string{"label2": "value 2"}))

	entry, found := buf.At(1)
	testutils.RequireEqualBool(t, true, found, "expecting entry to be found")

	field, found := entry.Field("label2")
	testutils.RequireEqualBool(t, true, found, "expecting field to be found")
	testutils.AssertEqualString(t, "value 2", field.Value(), "unexpected field value")
}

func TestCircularBufferAtValidFilled(t *testing.T) {
	buf := decoding.NewCircularBuffer(3)
	buf.Push(decodingtest.NewEntry(map[string]string{"label1": "value 1"}))
	buf.Push(decodingtest.NewEntry(map[string]string{"label2": "value 2"}))
	buf.Push(decodingtest.NewEntry(map[string]string{"label3": "value 3"}))
	buf.Push(decodingtest.NewEntry(map[string]string{"label4": "value 4"}))

	entry, found := buf.At(2)
	testutils.RequireEqualBool(t, true, found, "expecting entry to be found")

	field, found := entry.Field("label4")
	testutils.RequireEqualBool(t, true, found, "expecting field to be found")
	testutils.AssertEqualString(t, "value 4", field.Value(), "unexpected field value")
}

func TestCircularBufferEntriesNoData(t *testing.T) {
	buf := decoding.NewCircularBuffer(3)
	entries := buf.Entries()
	testutils.AssertEqualInt(t, 0, len(entries), "invalid entries size")
}

func TestCircularBufferEntriesNotFilled(t *testing.T) {
	buf := decoding.NewCircularBuffer(3)
	buf.Push(decodingtest.NewEntry(map[string]string{"label1": "value 1"}))
	buf.Push(decodingtest.NewEntry(map[string]string{"label2": "value 2"}))

	entries := buf.Entries()
	testutils.RequireEqualInt(t, 2, len(entries), "invalid entries size")

	_, ok := entries[0].Field("label1")
	testutils.AssertEqualBool(t, true, ok, "expecting to find label1 field")

	_, ok = entries[1].Field("label2")
	testutils.AssertEqualBool(t, true, ok, "expecting to find label2 field")
}

func TestCircularBufferEntriesFilled(t *testing.T) {
	buf := decoding.NewCircularBuffer(3)
	buf.Push(decodingtest.NewEntry(map[string]string{"label1": "value 1"}))
	buf.Push(decodingtest.NewEntry(map[string]string{"label2": "value 2"}))
	buf.Push(decodingtest.NewEntry(map[string]string{"label3": "value 3"}))
	buf.Push(decodingtest.NewEntry(map[string]string{"label4": "value 4"}))

	entries := buf.Entries()
	testutils.RequireEqualInt(t, 3, len(entries), "invalid entries size")

	_, ok := entries[0].Field("label2")
	testutils.AssertEqualBool(t, true, ok, "expecting to find label2 field")

	_, ok = entries[1].Field("label3")
	testutils.AssertEqualBool(t, true, ok, "expecting to find label3 field")

	_, ok = entries[2].Field("label4")
	testutils.AssertEqualBool(t, true, ok, "expecting to find label4 field")
}

func TestCircularBufferPush(t *testing.T) {
	buf := decoding.NewCircularBuffer(3)

	evicted := buf.Push(decodingtest.NewEntry(map[string]string{"label1": "value 1"}))
	testutils.AssertEqualBool(t, false, evicted, "didn't expect value to be evicted")

	evicted = buf.Push(decodingtest.NewEntry(map[string]string{"label2": "value 2"}))
	testutils.AssertEqualBool(t, false, evicted, "didn't expect value to be evicted")

	evicted = buf.Push(decodingtest.NewEntry(map[string]string{"label3": "value 3"}))
	testutils.AssertEqualBool(t, false, evicted, "didn't expect value to be evicted")

	evicted = buf.Push(decodingtest.NewEntry(map[string]string{"label4": "value 4"}))
	testutils.AssertEqualBool(t, true, evicted, "didn't expect value to be evicted")

	evicted = buf.Push(decodingtest.NewEntry(map[string]string{"label4": "value 5"}))
	testutils.AssertEqualBool(t, true, evicted, "didn't expect value to be evicted")
}
