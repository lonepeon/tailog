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

func ScanFieldString(name string, content []byte) (Field, error) {
	value, err := scanString(content)
	if err != nil {
		return Field{}, err
	}

	return Field{name: name, kind: FieldTypeString, str: value}, nil
}

func ScanFieldNumber(name string, content []byte) (Field, error) {
	value, err := scanNumber(content)
	if err != nil {
		return Field{}, err
	}

	return Field{name: name, kind: FieldTypeNumber, number: value}, nil
}

func (f Field) String() string {
	return f.str
}

func (f Field) Number() float64 {
	return f.number
}

func scanString(content []byte) (string, error) {
	var value string
	if err := json.Unmarshal(content, &value); err != nil {
		return "", fmt.Errorf("can't scan field (value=%s): %v", string(content), err)
	}

	return value, nil
}

func scanNumber(content []byte) (float64, error) {
	var value float64
	if err := json.Unmarshal(content, &value); err != nil {
		return 0, fmt.Errorf("can't scan field (value=%s): %v", string(content), err)
	}

	return value, nil
}
