package vanerrors

// An array of error levels.
//
// The index of the value is the number of the level
// The value is it's text
var SeverityArray [4]string = [4]string{
	"info",
	"warn",
	"error",
	"fatal",
}

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
