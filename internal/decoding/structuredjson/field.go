package structuredjson

import (
	"fmt"

	"github.com/lonepeon/tailog/internal/decoding"
)

type FieldType string

var (
	FieldTypeString FieldType = "string"
	FieldTypeNumber FieldType = "number"
)

func (f FieldType) String() string {
	return string(f)
}

func scanField(name string, rawValue []byte) (decoding.Field, error) {
	commands := []func() (decoding.Field, error){
		func() (decoding.Field, error) { return ScanFieldString(name, rawValue) },
		func() (decoding.Field, error) { return ScanFieldNumber(name, rawValue) },
	}

	for i := range commands {
		field, err := commands[i]()
		if err != nil {
			continue
		}

		return field, nil
	}

	return nil, fmt.Errorf("can't parse field '%s': %w", name, ErrCannotParseEntry)
}
