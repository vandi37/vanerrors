package vanerrors

import (
	"encoding/json"
	"io"
	"strings"
)

const (
	DATE_FORMAT = "2006/01/02 15:04:05"
)

// The json construction of a vanerror
type JsonVanError struct {
	// The date
	Date string `json:"date,omitempty"`
	// The error code
	Code int `json:"code,omitempty"`
	// The main data
	Main string `json:"main"`
	// The description
	Description string `json:"description,omitempty"`
	// The cause
	Cause any `json:"cause,omitempty"`
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
