package decoding

type FieldComparison string

var (
	FieldComparisonGreaterThan FieldComparison = "gt"
	FieldComparisonLessThan    FieldComparison = "lt"
	FieldComparisonEqual       FieldComparison = "eq"
)

func (f FieldComparison) String() string {
	return string(f)
}

type Field interface {
	Name() string
	Compare(interface{}) FieldComparison
	Value() string
}

type Entry interface {
	Len() int
	Fields() []Field
	Field(string) (Field, bool)
}
