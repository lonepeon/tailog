package structuredjson

import "fmt"

type FieldComparison string

var (
	FieldComparisonGreaterThan FieldComparison = "gt"
	FieldComparisonLessThan    FieldComparison = "lt"
	FieldComparisonEqual       FieldComparison = "eq"
)

func (f FieldComparison) String() string {
	return string(f)
}

type FieldType string

var (
	FieldTypeString FieldType = "string"
	FieldTypeNumber FieldType = "number"
)

func (f FieldType) String() string {
	return string(f)
}

type Field interface {
	Name() string
	Compare(interface{}) FieldComparison
	Value() string
}

func scanField(name string, rawValue []byte) (Field, error) {
	commands := []func() (Field, error){
		func() (Field, error) { return ScanFieldString(name, rawValue) },
		func() (Field, error) { return ScanFieldNumber(name, rawValue) },
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
