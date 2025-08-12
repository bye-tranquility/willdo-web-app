package utils

import (
	"encoding/json"
	"io"
)

// FromJSON deserializes the object from JSON string
// in an io.Reader to the given interface
func FromJSON(i any, r io.Reader) error {
	dec := json.NewDecoder(r)
	return dec.Decode(i)
}

// ToJSON serializes the given interface into a string based JSON format
func ToJSON(i any, w io.Writer) error {
	enc := json.NewEncoder(w)
	return enc.Encode(i)
}
