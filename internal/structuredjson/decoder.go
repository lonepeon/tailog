package structuredjson

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"

	"github.com/lonepeon/tailog/internal"
)

type Decoder struct {
	jsonDecoder *json.Decoder
}

func NewDecoder(input io.Reader) *Decoder {
	return &Decoder{jsonDecoder: json.NewDecoder(input)}
}

func (d *Decoder) Decode() (internal.Entry, error) {
	var entry Entry

	if err := d.jsonDecoder.Decode(&entry); err != nil {
		if errors.Is(err, io.EOF) {
			return nil, err
		}

		return nil, fmt.Errorf("can't decode structured JSON: %v", err)
	}

	return entry, nil
}
