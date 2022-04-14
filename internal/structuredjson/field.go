package structuredjson

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
