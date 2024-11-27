package vanerrors

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
