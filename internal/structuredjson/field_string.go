package structuredjson

import (
	"encoding/json"
	"fmt"
	"strings"
)

type FieldString struct {
	name  string
	value string
}

func ScanFieldString(name string, content []byte) (FieldString, error) {
	value, err := scanString(content)
	if err != nil {
		return FieldString{}, err
	}

	return FieldString{name: name, value: value}, nil
}

func (f FieldString) Compare(other interface{}) FieldComparison {
	otherValue, ok := other.(string)
	if !ok {
		return FieldComparisonGreaterThan
	}

	rst := strings.Compare(f.value, otherValue)
	if rst > 0 {
		return FieldComparisonGreaterThan
	}

	if rst < 0 {
		return FieldComparisonLessThan
	}

	return FieldComparisonEqual
}

func (f FieldString) String() string {
	return f.value
}

func scanString(content []byte) (string, error) {
	var value string
	if err := json.Unmarshal(content, &value); err != nil {
		return "", fmt.Errorf("can't scan field (value=%s): %v", string(content), err)
	}

	return value, nil
}
