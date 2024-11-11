package vanerrors

import (
	"fmt"
	"log"
)

// Logs the error
//
// vanError - a VanError, which we are going to log
func Log(vanError VanError) error {
	// Setting the logger output
	options := vanError.LoggerOptions
	var result string

	// Adding severity
	if options.ShowSeverity {
		if options.IntSeverity {
			result += "level: " + SeverityArray[vanError.Severity] + ","
		} else {
			result += fmt.Sprintf("level: %d,", vanError.Severity)
		}
	}

	// Adding code
	if options.ShowCode {
		result += fmt.Sprintf(" %d", vanError.Code)
	}

	// Adding name
	result += " " + vanError.Name

	// Adding message
	if options.ShowMessage {
		result += ": " + vanError.Message
	}

	// Adding , to show the next data
	result += ", "

	// Adding description
	if options.ShowDescription && vanError.Description != nil {
		description := make([]byte, 4096)
		n, err := vanError.Description.Read(description)
		if err == nil {
			description = description[:n]
		}

		result += "description: " + string(description) + ", "
	}

	// Adding cause
	if options.ShowCause && vanError.Cause != nil {
		result += "cause: " + vanError.Cause.Error()
	}

	// Getting the result string
	logger := log.New(vanError.logger, "", log.LstdFlags|log.Llongfile|log.Lshortfile)
	if vanError.Severity == 3 {
		logger.Fatalln(result)
	}
	logger.Println(result)

	return nil
}
