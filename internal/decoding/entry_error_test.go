package decoding_test

import (
	"errors"
	"testing"

	"github.com/lonepeon/golib/testutils"
	"github.com/lonepeon/tailog/internal/decoding"
)

func TestEntryLen(t *testing.T) {
	entry := decoding.NewEntryError("msg", errors.New("boom"))
	testutils.AssertEqualInt(t, 1, entry.Len(), "invalid entry length")
}

func TestEntryFieldExist(t *testing.T) {
	entry := decoding.NewEntryError("msg", errors.New("boom"))
	field, ok := entry.Field("msg")
	testutils.RequireEqualBool(t, true, ok, "expecting to find the field")
	testutils.AssertEqualString(t, "boom", field.Value(), "invalid entry value")
}

func TestEntryFieldNotExist(t *testing.T) {
	entry := decoding.NewEntryError("msg", errors.New("boom"))
	_, ok := entry.Field("nope")
	testutils.RequireEqualBool(t, false, ok, "didn't expect to find the field")
}

func TestFields(t *testing.T) {
	entry := decoding.NewEntryError("msg", errors.New("boom"))
	fields := entry.Fields()
	testutils.RequireEqualInt(t, 1, len(fields), "unexpected number of fields")
	field, ok := entry.Field("msg")
	testutils.RequireEqualBool(t, true, ok, "expecting to find the field")
	testutils.AssertEqualString(t, "boom", field.Value(), "invalid entry value")
}
