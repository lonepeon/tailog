package structuredjson

import (
	"encoding/json"
	"fmt"
)

type FieldType string

var (
	FieldTypeString FieldType = "string"
	FieldTypeNumber FieldType = "number"
)

type Field struct {
	name string

	kind   FieldType
	str    string
	number float64
}

func NewFieldString(name string) *Field {
	return &Field{name: name, kind: FieldTypeString}
}

func NewFieldNumber(name string) *Field {
	return &Field{name: name, kind: FieldTypeNumber}
}

func (f *Field) Scan(content []byte) error {
	switch f.kind {
	case FieldTypeNumber:
		return f.scanNumber(content)
	case FieldTypeString:
		return f.scanString(content)
	}

	return nil
}

func (f *Field) String() string {
	return f.str
}

func (f *Field) Number() float64 {
	return f.number
}

func (f *Field) scanString(content []byte) error {
	var value string
	if err := json.Unmarshal(content, &value); err != nil {
		return fmt.Errorf("can't scan field (value=%s): %v", string(content), err)
	}

	f.str = value

	return nil
}

func (f *Field) scanNumber(content []byte) error {
	var value float64
	if err := json.Unmarshal(content, &value); err != nil {
		return fmt.Errorf("can't scan field (value=%s): %v", string(content), err)
	}

	f.number = value

	return nil
}
