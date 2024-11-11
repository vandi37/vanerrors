package vanerrors

import "io"

// Default options for .Error()
//
// Sets as showing:
// Message
// Code
var DefaultOptions Options = Options{
	ShowMessage: true,
	ShowCode:    true,
}

// Default options for logger
//
// Sets as showing:
// DoLog
// Message
// Code
// Severity
// Description
// Cause
var DefaultLoggerOptions LoggerOptions = LoggerOptions{
	DoLog:           true,
	ShowMessage:     true,
	ShowCode:        true,
	ShowSeverity:    true,
	ShowDescription: true,
	ShowCause:       true,
}

// An empty logger settings
//
// It makes everything off
var EmptyLoggerOptions LoggerOptions = LoggerOptions{}

// 	Creates default ErrorData
//
// Name - error name
// Message - error message
// Code - status code
// Logger - writer for logs
//
// Returns ErrorData Based on the values
func DefaultData(Name string, Message string, Code int, Logger io.Writer) ErrorData {
	return ErrorData{
		Name:     Name,
		Message:  Message,
		Code:     Code,
		Logger:   Logger,
		Severity: 2,
	}
}

// Returns default values for New()
//
// Name - error name
// Message - error message
// Code - status code
// Logger - writer for logs
//
// Returns:
// ErrorData Based on the values
//Options DefaultOptions
// LoggerOptions DefaultLoggerOptions
//
// You can use it New(DefaultValues(Name, Message, Code, Logger))
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
// Name - error name
// Message - error message
// Code - status code
// Logger - writer for logs
//
// Returns:
// VanError based on the input data and base settings
//
// Use if you don't need to customize
func Default(Name string, Message string, Code int, Logger io.Writer) VanError {
	return New(DefaultValues(Name, Message, Code, Logger))
}
