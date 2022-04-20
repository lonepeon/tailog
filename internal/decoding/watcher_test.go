package decoding_test

import (
	"testing"

	"github.com/lonepeon/golib/testutils"
	"github.com/lonepeon/tailog/internal/decoding"
	"github.com/lonepeon/tailog/internal/decoding/decodingtest"
)

func TestWatchMultipleLine(t *testing.T) {
	decoder := decodingtest.NewDecoder()
	decoder.AddEntry(decodingtest.NewEntry(map[string]string{"msg": "log 1"}))
	decoder.AddEntry(decodingtest.NewEntry(map[string]string{"msg": "log 2"}))
	watcher := decoding.NewWatcher(decoder)
	watcher.Start()
	defer watcher.Stop()

	entries := make(chan decoding.Entry, 2)
	watcher.Notify(func(e decoding.Entry, err error) {
		testutils.RequireNoError(t, err, "didn't expect decoding issue")
		entries <- e
	})

	entry := <-entries
	field, ok := entry.Field("msg")
	testutils.AssertEqualBool(t, true, ok, "can't find field msg on entry 0")
	testutils.AssertEqualString(t, "log 1", field.Value(), "unexpected field msg value on entry 0")

	entry = <-entries
	field, ok = entry.Field("msg")
	testutils.AssertEqualBool(t, true, ok, "can't find field msg on entry 1")
	testutils.AssertEqualString(t, "log 2", field.Value(), "unexpected field msg value on entry 1")

	testutils.RequireEqualInt(t, 0, len(entries), "unexpected number of remaining entries")
}
