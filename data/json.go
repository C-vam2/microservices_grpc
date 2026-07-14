package data

import (
	"encoding/json"
	"io"
)

// ToJSON
func ToJSON(p *Products, w io.Writer) error {
	enc := json.NewEncoder(w)
	return enc.Encode(p)
}

func FronJSON(p *Products, r io.Reader) error {
	dec := json.NewDecoder(r)
	return dec.Decode(p)
}
