package vanerrors

import (
	"fmt"
	"io"
	"log"
)

// Creates a log of VanError based on it settings
//
// The method could be used inside some methods (New(), VanError.Error()) and outside
// err := Default(Name, Message, Code, Logger)
// err.Log()
//
// It is a basic logger, using the standard log package
// If you want to have better log, set logs of.
func (e VanError) Log() {
	// Setting the logger output
	options := e.LoggerOptions
	var result string

	// Adding severity
	if options.ShowSeverity {
		if !options.IntSeverity {
			result += SeverityArray[e.Severity] + ", "
		} else {
			result += fmt.Sprintf("level %d, ", e.Severity)
		}
	}

	// Adding code
	if options.ShowCode {
		result += fmt.Sprintf("%d ", e.Code)
	}

	// Adding name
	result += e.Name

	// Adding message
	if options.ShowMessage {
		result += ": " + e.Message
	}

	if options.ShowDescription || options.ShowCause {
		// Adds ", "
		result += ", "

		// Adding description
		if options.ShowDescription && e.Description != nil {
			description, _ := io.ReadAll(e.Description)
			if len(description) > 0 {
				result += "description: " + string(description)
				if options.ShowCause {
					result += ", "
				}
			}
		}

		// Adding cause
		if options.ShowCause && e.Cause != nil {
			result += "cause: " + e.Cause.Error()
		}
	}

	// Getting the result string
	logger := log.New(e.Logger, "", log.LstdFlags)
	if e.Severity == 3 {
		logger.Panicln(result)
	}
	logger.Println(result)
}
