package ctrls

import (
	"encoding/json"
	"io"
)

func LoadBody(r io.Reader, val interface{}) error {
	decode := json.NewDecoder(r)
	return decode.Decode(&val)
}
