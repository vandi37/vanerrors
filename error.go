package vanerrors

import (
	"io"
	"time"
)

// An error structure
// it has different information about the error
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
	//
	// example:
	// if err != nil {
	//		panic(VanError{Cause: err})
	// }
	//
	// it would not be used on .Error() func, however you will see in logs it
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

// An array of error levels.
//
// The index of the value is the number of the level
// The value is it's text
var SeverityArray [4]string = [4]string{
	"unknown",
	"warn",
	"error",
	"fatal",
}

// The severity levels
const (
	WarnLevel  = 1
	ErrorLevel = 2
	FatalLevel = 3
)
