package vanerrors

import (
	"io"
	"strings"
	"time"
)

// Creates a new VanError
//
// sets default:
// 1. Name == "" -> "unknown error"
// 2. Code <= 0 -> 500
// 3. Severity < 0 | Severity > 3 -> 2

// It is the most customizable function
// You can edit the error data, options and log options
// Be careful with severity. if it is equal to 3 it will create panic
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

// Returns default values for New
//
// You can use it
// err := New(DefaultValues(Name, Message, Code, Logger))
//
// It does not provide any customization, it only sets default values
// If you want to have more settings use other methods
func DefaultValues(Name string, Message string, Code int, Logger io.Writer) (ErrorData, Options, LoggerOptions) {
	return DefaultData(
		Name,
		Message,
		Code,
		Logger,
	), DefaultOptions, DefaultLoggerOptions
}

// Returns default value of VanError
//
// You can use it instead of
// err := New(DefaultValues(Name, Message, Code, Logger))
// // only
// err = Default(Name, Message, Code, Logger)
func Default(Name string, Message string, Code int, Logger io.Writer) VanError {
	return New(DefaultValues(Name, Message, Code, Logger))
}
