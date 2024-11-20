package vanerrors

import "io"

// Error writer
//
// has it own settings for not using default
type ErrorW struct {
	// OPtions
	Options Options
	// LoggerOptions
	LoggerOptions LoggerOptions
}

// Creates a Error Writer
func NewW(options Options, loggerOptions LoggerOptions) ErrorW {
	return ErrorW{
		Options:       options,
		LoggerOptions: loggerOptions,
	}
}

// Creates a new error based on the error writer settings
func (w ErrorW) New(errorData ErrorData) VanError {
	return New(errorData, w.Options, w.LoggerOptions)
}

// Creates a new error based on the error writer settings only with name and logger
// If you don't need the logger set it to nil
func (w ErrorW) NewName(Name string, Logger io.Writer) VanError {
	data := ErrorData{
		Name:   Name,
		Logger: Logger,
	}
	w.Options.ShowCode = false
	w.Options.ShowMessage = false
	return New(data, w.Options, w.LoggerOptions)
}

// Creates a new error based on the error writer settings with pair Name: Message
// If you don't need the logger set it to nil
func (w ErrorW) NewBasic(Name string, Message string, Logger io.Writer) VanError {
	data := ErrorData{Name: Name, Message: Message, Logger: Logger}
	w.Options.ShowCode = false
	return New(data, w.Options, w.LoggerOptions)
}

// Creates a new error based on the error writer settings with pair Code Name
// If you don't need the logger set it to nil
func (w ErrorW) NewHTTP(Name string, Code int, Logger io.Writer) VanError {
	data := ErrorData{Name: Name, Code: Code, Logger: Logger}
	w.Options.ShowMessage = false
	return New(data, w.Options, w.LoggerOptions)
}

// Creates a new error based on the error writer settings with Name, and error that is the error cause (it would not be shown it .Error())
// If you don't need the logger set it to nil
func (w ErrorW) NewWrap(Name string, Cause error, Logger io.Writer) VanError {
	data := ErrorData{Name: Name, Cause: Cause, Logger: Logger}
	w.Options.ShowMessage = false
	w.Options.ShowCode = false
	w.Options.ShowCause = true
	return New(data, w.Options, w.LoggerOptions)
}
