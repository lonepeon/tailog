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

type FieldString struct {
	name  string
	value string
}

type FieldNumber struct {
	name  string
	value float64
}

func ScanFieldString(name string, content []byte) (FieldString, error) {
	value, err := scanString(content)
	if err != nil {
		return FieldString{}, err
	}

	return FieldString{name: name, value: value}, nil
}

func ScanFieldNumber(name string, content []byte) (FieldNumber, error) {
	value, err := scanNumber(content)
	if err != nil {
		return FieldNumber{}, err
	}

	return FieldNumber{name: name, value: value}, nil
}

func (f FieldString) Equal(s interface{}) bool {
	value, ok := s.(string)
	if !ok {
		return false

	}

	return f.value == value
}

func (f FieldString) String() string {
	return f.value
}

func (f FieldNumber) Equal(s interface{}) bool {
	var value float64
	switch v := s.(type) {
	case float64:
		value = v
	case int:
		value = float64(v)
	default:
		return false
	}

	return f.value == value
}

func (f FieldNumber) Number() float64 {
	return f.value
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
