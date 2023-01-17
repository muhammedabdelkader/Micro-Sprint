package jsonparser

import (
	"encoding/json"
	"io"
)

/*
The package defines a function ParseJSON that takes a reader r and a value v
* as a parameter and it uses the json.NewDecoder(r) function to create a new
* decoder that reads from the r reader.
It then uses the dec.Decode(v) function to decode json data from the reader and store the result in the value v.
It returns error if it fails otherwise it will return nil.

*/
// ParseJSON parses json data from a reader
func ParseJSON(r io.Reader, v interface{}) error {
	dec := json.NewDecoder(r)
	return dec.Decode(v)

}
