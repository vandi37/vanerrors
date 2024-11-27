package vanerrors

// Error writer
//
// has it own settings how to log and show the van error
type ErrorW struct {
	// Options
	Options Options `json:"options"`
	// LoggerOptions
	LoggerOptions LoggerOptions `json:"logger_options"`
	// Error Handler
	Handler ErrorHandler `json:"error_handler"`
}

// Creates a Error Writer
// It is based on the options and logger options
func NewW(options Options, loggerOptions LoggerOptions, Handler ErrorHandler) ErrorW {
	return ErrorW{
		Options:       options,
		LoggerOptions: loggerOptions,
		Handler:       Handler,
	}
}

// Creates a new error based on the error writer settings
func (w ErrorW) New(errorData ErrorData) VanError {
	errorData.ErrorHandler = w.Handler
	return New(errorData, w.Options, w.LoggerOptions)
}

// It creates a new van error only with a name
//
// You can set error handler data with auto panic and logger
func (w ErrorW) NewName(Name string) VanError {
	data := ErrorData{
		Name:         Name,
		ErrorHandler: w.Handler,
	}
	w.Options.ShowCode = false
	w.Options.ShowMessage = false
	return New(data, w.Options, w.LoggerOptions)
}

// It creates a new van error only with name and message
// The result would be shown as "Name: Message"
//
// You can set error handler data with auto panic and logger
func (w ErrorW) NewBasic(Name string, Message string) VanError {
	data := ErrorData{Name: Name, Message: Message, ErrorHandler: w.Handler}
	w.Options.ShowCode = false
	w.Options.ShowMessage = true
	return New(data, w.Options, w.LoggerOptions)
}

// It creates a new van error only with name and code
// The result would be shown as "Code Name"
//
// You can set error handler data with auto panic and logger
func (w ErrorW) NewHTTP(Name string, Code int) VanError {
	data := ErrorData{Name: Name, Code: Code, ErrorHandler: w.Handler}
	w.Options.ShowMessage = false
	w.Options.ShowCode = true
	return New(data, w.Options, w.LoggerOptions)
}

// It creates a new van error only with name and cause
// The result would be shown as "Name, cause: Cause"
//
// You can set error handler data with auto panic and logger
func (w ErrorW) NewWrap(Name string, Cause error) VanError {
	data := ErrorData{Name: Name, Cause: Cause, ErrorHandler: w.Handler}
	w.Options.ShowMessage = false
	w.Options.ShowCode = false
	w.Options.ShowCause = true
	return New(data, w.Options, w.LoggerOptions)
}

// Sets your error writer as default for your program
func (w ErrorW) SetAsDefault() {
	w.Options.SetAsDefault()
	w.LoggerOptions.SetAsDefault()
	w.Handler.SetAsDefault()
}
