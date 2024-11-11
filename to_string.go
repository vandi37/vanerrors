package vanerrors

import "fmt"

// Formats the error to a string format
//
// Is a method of VanError
//
// returns a string
func (e VanError) Error() string {
	// Sets the return value
	var err string

	// Sets the options
	options := e.Options

	// Adding data
	if options.ShowDate {
		err += e.Date.Format("2006-01-02 15:04:05 ")
	}

	// Adding severity
	if options.ShowSeverity {
		if options.IntSeverity {
			err += "level: " + SeverityArray[e.Severity] + ","
		} else {
			err += fmt.Sprintf("level: %d,", e.Severity)
		}
	}

	// Adding code
	if options.ShowCode {
		err += fmt.Sprintf(" %d", e.Code)
	}

	// Adding name
	err += " " + e.Name

	// Adding message
	if options.ShowMessage {
		err += ": " + e.Message
	}
	// Adds ", "
	err += ", "

	// Adding description
	if options.ShowDescription && e.Description != nil {
		description := make([]byte, 4096)
		n, errRead := e.Description.Read(description)
		if errRead == nil {
			description = description[:n]
		}

		err += "description: " + string(description) + ", "
	}

	// Adding cause
	if options.ShowCause && e.Cause != nil {
		err += "cause: " + e.Cause.Error()
	}

	return err
}
