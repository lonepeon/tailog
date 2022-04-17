package structuredjson

import (
	"encoding/json"
	"fmt"
	"github.com/lonepeon/tailog/internal/decoding"
)

func ScanFieldNumber(name string, content []byte) (decoding.FieldNumber, error) {
	var value float64
	if err := json.Unmarshal(content, &value); err != nil {
		return decoding.FieldNumber{}, fmt.Errorf("can't scan field (value=%s): %v", string(content), err)
	}

	return decoding.NewFieldNumber(name, value), nil
}
