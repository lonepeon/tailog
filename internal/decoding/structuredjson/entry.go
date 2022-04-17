package structuredjson

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/lonepeon/tailog/internal/decoding"
)

var (
	ErrCannotParseEntry = errors.New("cannot parse entry")
)

type Entry struct {
	fields []decoding.Field
}

func (e *Entry) UnmarshalJSON(b []byte) error {
	jsonEntries := make(map[string]json.RawMessage)
	if err := json.Unmarshal(b, &jsonEntries); err != nil {
		return fmt.Errorf("can't unmarshal json: %w: %v", ErrCannotParseEntry, err)
	}

	fields := make([]decoding.Field, 0, len(jsonEntries))
	for name, rawValue := range jsonEntries {
		field, err := scanField(name, rawValue)
		if err != nil {
			return err
		}

		fields = append(fields, field)
	}

	e.fields = fields

	return nil
}

func (e Entry) Len() int {
	return len(e.fields)
}

func (e Entry) Fields() []decoding.Field {
	return e.fields
}

func (e Entry) Field(name string) (decoding.Field, bool) {
	for _, f := range e.fields {
		if f.Name() == name {
			return f, true
		}
	}

	return nil, false
}
