package structuredjson

import (
	"encoding/json"
	"fmt"
	"math"
	"strconv"

	"github.com/lonepeon/tailog/internal/decoding"
)

type FieldNumber struct {
	name  string
	value float64
}

func ScanFieldNumber(name string, content []byte) (FieldNumber, error) {
	value, err := scanNumber(content)
	if err != nil {
		return FieldNumber{}, err
	}

	return FieldNumber{name: name, value: value}, nil
}

func (f FieldNumber) Name() string {
	return f.name
}

func (f FieldNumber) Compare(other interface{}) decoding.FieldComparison {
	otherValue, ok := toFloat64(other)
	if !ok {
		return decoding.FieldComparisonGreaterThan
	}

	if f.value > otherValue {
		return decoding.FieldComparisonGreaterThan
	}

	if f.value < otherValue {
		return decoding.FieldComparisonLessThan
	}

	return decoding.FieldComparisonEqual
}

func (f FieldNumber) Value() string {
	return strconv.FormatFloat(f.value, 'f', -1, 64)
}

func scanNumber(content []byte) (float64, error) {
	var value float64
	if err := json.Unmarshal(content, &value); err != nil {
		return 0, fmt.Errorf("can't scan field (value=%s): %v", string(content), err)
	}

	return value, nil
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
