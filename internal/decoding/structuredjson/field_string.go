package structuredjson

import (
	"encoding/json"
	"fmt"

	"github.com/lonepeon/tailog/internal/decoding"
)

func ScanFieldString(name string, content []byte) (decoding.FieldString, error) {
	var value string
	if err := json.Unmarshal(content, &value); err != nil {
		return decoding.FieldString{}, fmt.Errorf("can't scan field (value=%s): %v", string(content), err)
	}

	return decoding.NewFieldString(name, value), nil
}
