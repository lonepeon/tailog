package decodingtest

import (
	"io"

	"github.com/lonepeon/tailog/internal/decoding"
)

type Decoder struct {
	entries []Entry
	index   int
}

func NewDecoder() *Decoder {
	return &Decoder{}
}

func (d *Decoder) AddEntry(e Entry) {
	d.entries = append(d.entries, e)
}

func (d *Decoder) More() bool {
	return d.index < len(d.entries)
}

func (d *Decoder) Decode() (decoding.Entry, error) {
	if !d.More() {
		return nil, io.EOF
	}

	entry := d.entries[d.index]
	d.index += 1

	return entry, nil
}
