package decoding

import (
	"math"
	"strconv"
)

type FieldNumber struct {
	name  string
	value float64
}

func NewFieldNumber(name string, value float64) FieldNumber {
	return FieldNumber{name: name, value: value}
}

func (f FieldNumber) Name() string {
	return f.name
}

func (f FieldNumber) Compare(other interface{}) FieldComparison {
	var otherValue float64

	otherFieldNumber, ok := other.(FieldNumber)
	if ok {
		otherValue = otherFieldNumber.value
	} else {
		numberValue, ok := toFloat64(other)
		if !ok {
			return FieldComparisonGreaterThan
		}

		otherValue = numberValue
	}

	if f.value > otherValue {
		return FieldComparisonGreaterThan
	}

	if f.value < otherValue {
		return FieldComparisonLessThan
	}

	return FieldComparisonEqual
}

func (f FieldNumber) Value() string {
	return strconv.FormatFloat(f.value, 'f', -1, 64)
}

func toFloat64(v interface{}) (float64, bool) {
	switch value := v.(type) {
	case float64:
		return value, true
	case int:
		return float64(value), true
	default:
		return math.NaN(), false
	}
}
