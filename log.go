package vanerrors

import (
	"fmt"
	"log"
)

// Logs the error
//
// a method for the vanError
func (e VanError) Log() error {
	// Setting the logger output
	options := e.LoggerOptions
	var result string

	// Adding severity
	if options.ShowSeverity {
		if options.IntSeverity {
			result += "level: " + SeverityArray[e.Severity] + ","
		} else {
			result += fmt.Sprintf("level: %d,", e.Severity)
		}
	}

	// Adding code
	if options.ShowCode {
		result += fmt.Sprintf(" %d", e.Code)
	}

	// Adding name
	result += " " + e.Name

	// Adding message
	if options.ShowMessage {
		result += ": " + e.Message
	}

	// Adding , to show the next data
	result += ", "

	// Adding description
	if options.ShowDescription && e.Description != nil {
		description := make([]byte, 4096)
		n, err := e.Description.Read(description)
		if err == nil {
			description = description[:n]
		}

		result += "description: " + string(description) + ", "
	}

	// Adding cause
	if options.ShowCause && e.Cause != nil {
		result += "cause: " + e.Cause.Error()
	}

	// Getting the result string
	logger := log.New(e.logger, "", log.LstdFlags|log.Llongfile|log.Lshortfile)
	if e.Severity == 3 {
		logger.Fatalln(result)
	}
	logger.Println(result)

	return nil
}
