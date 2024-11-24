package vanerrors

import (
	"io"
	"time"
)

// It is an error with multiple data and view and log settings
type VanError struct {
	// The error name
	//
	// better to be short and easy to understand
	Name string `json:"name"`
	// The error message
	//
	// it could have more information about the error
	//
	// this information shouldn't be to long, however it should contain more information that the Name
	Message string `json:"message"`
	// The status code
	//
	// you can put http status codes
	Code int `json:"code"`
	// The error, because of what this error was created
	Cause error `json:"error"`
	// The error description
	//
	// it should contain the most information, that would be displaced in logs
	//
	// please add all error information here
	Description io.Reader `json:"description"`
	// The error logger
	//
	// it is a writer, where it will write the logs
	Logger io.Writer `json:"logger"`
	// The options for the logger behavior
	//
	// Use it to customize the program behavior
	LoggerOptions LoggerOptions `json:"logger_options"`
	// The options for the .Error()
	//
	// Use it to customize the program behavior
	Options Options `json:"options"`
	// The error date
	//
	// It controls the error date
	Date time.Time `json:"date"`
}
