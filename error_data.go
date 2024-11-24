package vanerrors

import "io"

// The data to create an error
//
// Everything is optional, if you don't want to create one of the data just slip it
type ErrorData struct {
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
	// it would not be used on .Error() func, however you will see in logs it
	Cause error `json:"error"`
	// The error description
	//
	// it should contain the most information, that would be displaced in logs
	//
	// please add all error information here
	Description io.Reader `json:"description"`
	// The handler of the error data
	ErrorHandler
}

// The handler for logger and does it need to panic
type ErrorHandler struct {
	// The error logger
	//
	// it is a writer, where it will write the logs
	Logger io.Writer `json:"logger"`
	// Does it need to panic
	DoPanic bool `json:"do_panic"`
}

// An empty handler
var EmptyHandler ErrorHandler = ErrorHandler{}
