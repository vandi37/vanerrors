package vanerrors

import (
	"strings"
	"time"
)

// Creates a new VanError
//
// sets default:
// 1. Name == "" -> "unknown error",
// 2. Code <= 0 -> 500,
// 3. Severity <= 0 or Severity > 3 -> 2,

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
		options.ShowCode = false
	}
	// Creating a vanError
	now := time.Now()
	vanError := VanError{
		Name:          errorData.Name,
		Message:       errorData.Message,
		Code:          errorData.Code,
		Cause:         errorData.Cause,
		Description:   errorData.Description,
		Logger:        errorData.Logger,
		LoggerOptions: loggerOptions,
		Options:       options,
		Date:          now,
	}
	// Logging if it needs to be when created
	if loggerOptions.DoLog && !loggerOptions.LogBy && errorData.Logger != nil {
		vanError.Log()
	}
	// Panic if needed
	if errorData.DoPanic {
		panic(vanError)
	}

	return vanError
}

// Returns default values for New
//
// You can use it
// err := New(DefaultValues(errorData))
//
// It does not provide any customization, it only sets default values
// If you want to have more settings use other methods
func DefaultValues(errorData ErrorData) (ErrorData, Options, LoggerOptions) {
	return errorData, DefaultOptions, DefaultLoggerOptions
}

// Creates a new default error
// With any error data in error data
// It does the same as
// err := New(DefaultValues(errorData))
func NewDefault(errorData ErrorData) VanError {
	return New(errorData, DefaultOptions, DefaultLoggerOptions)
}

// Creates a new error only with name and handler
// set handler as EmptyHandler if you don't need it
func NewName(Name string, Handler ErrorHandler) VanError {
	data, opt, logOpt := DefaultValues(ErrorData{Name: Name, ErrorHandler: Handler})
	opt.ShowCode = false
	opt.ShowMessage = false
	return New(data, opt, logOpt)
}

// Creates a new error with pair Name: Message
// set handler as EmptyHandler if you don't need it
func NewBasic(Name string, Message string, Handler ErrorHandler) VanError {
	data, opt, logOpt := DefaultValues(ErrorData{Name: Name, Message: Message, ErrorHandler: Handler})
	opt.ShowCode = false
	return New(data, opt, logOpt)
}

// Creates a new error with pair Code Name
// set handler as EmptyHandler if you don't need it
func NewHTTP(Name string, Code int, Handler ErrorHandler) VanError {
	data, opt, logOpt := DefaultValues(ErrorData{Name: Name, Code: Code, ErrorHandler: Handler})
	opt.ShowMessage = false
	return New(data, opt, logOpt)
}

// Creates a new error with Name, and error that is the error cause (it would not be shown it .Error())
// set handler as EmptyHandler if you don't need it
func NewWrap(Name string, Cause error, Handler ErrorHandler) VanError {
	data, opt, logOpt := DefaultValues(ErrorData{Name: Name, Cause: Cause, ErrorHandler: Handler})
	opt.ShowMessage = false
	opt.ShowCode = false
	opt.ShowCause = true
	return New(data, opt, logOpt)
}
