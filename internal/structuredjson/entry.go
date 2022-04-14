package structuredjson

import (
	"encoding/json"
	"errors"
	"fmt"
)

var (
	ErrCannotParseEntry = errors.New("cannot parse entry")
)

type Entry struct {
	fields []Field
}

func ParseEntry(b []byte) (Entry, error) {
	jsonEntries := make(map[string]json.RawMessage)
	if err := json.Unmarshal(b, &jsonEntries); err != nil {
		return Entry{}, fmt.Errorf("can't unmarshal json: %w: %v", ErrCannotParseEntry, err)
	}

	fields := make([]Field, 0, len(jsonEntries))
	for name, rawValue := range jsonEntries {
		field, err := scanField(name, rawValue)
		if err != nil {
			return Entry{}, err
		}

		fields = append(fields, field)
	}

	return Entry{fields: fields}, nil
}

func (e Entry) Len() int {
	return len(e.fields)
}

func (e Entry) Fields() []Field {
	return e.fields
}

func (e Entry) Field(name string) (Field, bool) {
	for _, f := range e.fields {
		if f.Name() == name {
			return f, true
		}
	}

	return nil, false
}
