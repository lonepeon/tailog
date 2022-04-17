package structuredjson

import (
	"fmt"

	"github.com/lonepeon/tailog/internal"
)

type FieldType string

var (
	FieldTypeString FieldType = "string"
	FieldTypeNumber FieldType = "number"
)

func (f FieldType) String() string {
	return string(f)
}

func scanField(name string, rawValue []byte) (internal.Field, error) {
	commands := []func() (internal.Field, error){
		func() (internal.Field, error) { return ScanFieldString(name, rawValue) },
		func() (internal.Field, error) { return ScanFieldNumber(name, rawValue) },
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
