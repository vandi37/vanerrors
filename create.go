package vanerrors

import (
	"strings"
	"time"
)

// Creates a new VanError based on the settings
//
// sets default:
// 1. errorData.Name == "" -> "unknown error",
// 2. errorData.Code <= 0 -> 500,
//
// It is the most customizable function
// You can edit the error data, options and log options
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

// IT returns default values for function vanerrors.New()
//
// It sets the default options and default logger options. You can customize only the error data
func DefaultValues(errorData ErrorData) (ErrorData, Options, LoggerOptions) {
	return errorData, DefaultOptions, DefaultLoggerOptions
}

// It creates a new default van error
//
// It has default options and logger options, however you can change the error data like you want
func NewDefault(errorData ErrorData) VanError {
	return New(errorData, DefaultOptions, DefaultLoggerOptions)
}

// It creates a new van error only with a name
//
// You can set error handler data with auto panic and logger
func NewName(Name string, Handler ErrorHandler) VanError {
	data, opt, logOpt := DefaultValues(ErrorData{Name: Name, ErrorHandler: Handler})
	opt.ShowCode = false
	opt.ShowMessage = false
	return New(data, opt, logOpt)
}

// It creates a new van error only with name and message
// The result would be shown as "Name: Message"
//
// You can set error handler data with auto panic and logger
func NewBasic(Name string, Message string, Handler ErrorHandler) VanError {
	data, opt, logOpt := DefaultValues(ErrorData{Name: Name, Message: Message, ErrorHandler: Handler})
	opt.ShowCode = false
	opt.ShowMessage = true
	return New(data, opt, logOpt)
}

// It creates a new van error only with name and code
// The result would be shown as "Code Name"
//
// You can set error handler data with auto panic and logger
func NewHTTP(Name string, Code int, Handler ErrorHandler) VanError {
	data, opt, logOpt := DefaultValues(ErrorData{Name: Name, Code: Code, ErrorHandler: Handler})
	opt.ShowMessage = false
	opt.ShowCode = true
	return New(data, opt, logOpt)
}

// It creates a new van error only with name and cause
// The result would be shown as "Name, cause: Cause"
//
// You can set error handler data with auto panic and logger
func NewWrap(Name string, Cause error, Handler ErrorHandler) VanError {
	data, opt, logOpt := DefaultValues(ErrorData{Name: Name, Cause: Cause, ErrorHandler: Handler})
	opt.ShowMessage = false
	opt.ShowCode = false
	opt.ShowCause = true
	return New(data, opt, logOpt)
}

// Creates a new error with no additional settings
func NewSimple(Name string, Message ...string) VanError {
	var opt = Options{
		ShowAsJson: DefaultOptions.ShowAsJson,
	}
	var data = ErrorData{
		Name: Name,
	}
	if len(Message) >= 1 {
		opt.ShowMessage = true
		data.Message = Message[0]
	}
	if len(Message) > 1 {
		opt.ShowDescription = true
		data.Description = strings.Join(Message[:1], ", ")
	}
	return New(data, opt, EmptyLoggerOptions)
}
