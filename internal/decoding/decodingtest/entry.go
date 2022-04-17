package decodingtest

import "github.com/lonepeon/tailog/internal/decoding"

type Entry []decoding.Field

func NewEntry(entry map[string]string) Entry {
	fields := make([]decoding.Field, 0, len(entry))
	for label, value := range entry {
		fields = append(fields, decoding.NewFieldString(label, value))
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
