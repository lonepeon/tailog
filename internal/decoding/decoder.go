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

type Decoder interface {
	// Decode returns an entry if it can't parse one. It can return any errors,
	// they will considered as opaque errors and not handled except for io.EOF
	// to alert the stream is closed
	Decode() (Entry, error)
}
