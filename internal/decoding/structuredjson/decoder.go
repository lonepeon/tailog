package structuredjson

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"

	"github.com/lonepeon/tailog/internal/decoding"
)

type Decoder struct {
	jsonDecoder *json.Decoder
}

func NewDecoder(input io.Reader) *Decoder {
	return &Decoder{jsonDecoder: json.NewDecoder(input)}
}

func (d *Decoder) More() bool {
	return d.jsonDecoder.More()
}

func (d *Decoder) Decode() (decoding.Entry, error) {
	var entry Entry

	if err := d.jsonDecoder.Decode(&entry); err != nil {
		if errors.Is(err, io.EOF) {
			return nil, err
		}

		return nil, fmt.Errorf("can't decode structured JSON: %v", err)
	}

	return entry, nil
}
