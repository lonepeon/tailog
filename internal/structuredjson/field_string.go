package structuredjson

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/lonepeon/tailog/internal"
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

func (f FieldString) Name() string {
	return f.name
}

func (f FieldString) Compare(other interface{}) internal.FieldComparison {
	otherValue, ok := other.(string)
	if !ok {
		return internal.FieldComparisonGreaterThan
	}

	rst := strings.Compare(f.value, otherValue)
	if rst > 0 {
		return internal.FieldComparisonGreaterThan
	}

	if rst < 0 {
		return internal.FieldComparisonLessThan
	}

	return internal.FieldComparisonEqual
}

func (f FieldString) Value() string {
	return f.value
}

func scanString(content []byte) (string, error) {
	var value string
	if err := json.Unmarshal(content, &value); err != nil {
		return "", fmt.Errorf("can't scan field (value=%s): %v", string(content), err)
	}

	return value, nil
}
