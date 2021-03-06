package decoding

import (
	"strings"
)

type FieldString struct {
	name  string
	value string
}

func NewFieldString(name string, value string) FieldString {
	return FieldString{name: name, value: value}
}

func (f FieldString) Name() string {
	return f.name
}

func (f FieldString) Compare(other interface{}) FieldComparison {
	var otherValue string

	otherFieldString, ok := other.(FieldString)
	if ok {
		otherValue = otherFieldString.value
	} else {
		stringValue, ok := other.(string)
		if !ok {
			return FieldComparisonGreaterThan
		}

		otherValue = stringValue
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

func (f FieldString) Value() string {
	return f.value
}
