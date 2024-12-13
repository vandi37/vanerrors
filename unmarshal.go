package vanerrors

import (
	"encoding/json"
	"io"
	"strings"
)

// The json construction of a vanerror
type JsonVanError struct {
	// The date
	Date string `json:"date"`
	// The main data
	Main string `json:"main"`
	// The description
	Description string `json:"description"`
	// The cause
	Cause string `json:"cause"`
}

// it unmarshal with any reader
func Unmarshal(r io.Reader, t *JsonVanError) {
	json.NewDecoder(r).Decode(t)
}

// it unmarshal a string to a json van error
func UnmarshalString(s string, t *JsonVanError) {
	r := strings.NewReader(s)
	Unmarshal(r, t)
}
