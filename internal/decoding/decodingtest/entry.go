package decodingtest

import (
	"testing"

	"github.com/lonepeon/tailog/internal/decoding"
)

type Entry []decoding.Field

func NewEntry(t *testing.T, entry map[string]interface{}) Entry {
	fields := make([]decoding.Field, 0, len(entry))
	for label, value := range entry {
		switch v := value.(type) {
		case string:
			fields = append(fields, decoding.NewFieldString(label, v))
		case int:
			fields = append(fields, decoding.NewFieldNumber(label, float64(v)))
		case float64:
			fields = append(fields, decoding.NewFieldNumber(label, v))
		default:
			t.Fatalf("can't build an entry with a field of type '%T'", value)
		}
	}

	return Entry(fields)
}

func (e Entry) Len() int {
	return len(e)
}

func (e Entry) Fields() []decoding.Field {
	return e
}

func (e Entry) Field(name string) (decoding.Field, bool) {
	for _, field := range e {
		if field.Name() == name {
			return field, true
		}
	}

	return decoding.FieldString{}, false
}
