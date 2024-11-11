package vanerrors

import (
	"strings"
	"time"
)

// Creates a new VanError
//
// errorData: the data to create the error
// options: options how to show the error
// loggerOptions: the logger options
//
// sets default:
// Name == "" -> "unknown error"
// Code <= 0 -> 500
// Severity < 0 | Severity > 3 -> 2
//
// # Does not check other values
//
// Returns: VanError
func New(errorData ErrorData, options Options, loggerOptions LoggerOptions) VanError {
	// Validating name
	if errorName := strings.TrimSpace(errorData.Name); errorName == "" {
		errorData.Name = "unknown error"
	}
	// Validating code
	if errorData.Code <= 0 {
		errorData.Code = 500
	}
	if errorData.Severity < 0 || errorData.Severity > 3 {
		errorData.Severity = 2
	}
	// Creating a vanError
	now := time.Now()
	vanError := VanError{
		Name:          errorData.Name,
		Message:       errorData.Message,
		Code:          errorData.Code,
		Cause:         errorData.Cause,
		Description:   errorData.Description,
		logger:        errorData.Logger,
		LoggerOptions: loggerOptions,
		Options:       options,
		Severity:      errorData.Severity,
		Date:          now,
	}
	// Logging if it needs to be when created
	if loggerOptions.DoLog && !loggerOptions.LogBy && errorData.Logger != nil {
		vanError.Log()
	}

	return vanError
}
