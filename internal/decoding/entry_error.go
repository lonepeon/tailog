package decoding

type entryError struct {
	field Field
}

func NewEntryError(fieldName string, err error) Entry {
	return entryError{NewFieldString(fieldName, err.Error())}
}

func (e entryError) Len() int {
	return 1
}

func (e entryError) Field(name string) (Field, bool) {
	if name != e.field.Name() {
		return nil, false
	}

	return e.field, true
}

func (e entryError) Fields() []Field {
	return []Field{e.field}
}
