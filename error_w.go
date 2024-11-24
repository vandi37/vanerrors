package vanerrors

// Error writer
//
// has it own settings for not using default
type ErrorW struct {
	// OPtions
	Options Options `json:"options"`
	// LoggerOptions
	LoggerOptions LoggerOptions `json:"logger_options"`
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

// Creates a new error based on the error writer settings only with name and handler
// set handler as EmptyHandler if you don't need it
func (w ErrorW) NewName(Name string, Handler ErrorHandler) VanError {
	data := ErrorData{
		Name:         Name,
		ErrorHandler: Handler,
	}
	w.Options.ShowCode = false
	w.Options.ShowMessage = false
	return New(data, w.Options, w.LoggerOptions)
}

// Creates a new error based on the error writer settings with pair Name: Message
// set handler as EmptyHandler if you don't need it
func (w ErrorW) NewBasic(Name string, Message string, Handler ErrorHandler) VanError {
	data := ErrorData{Name: Name, Message: Message, ErrorHandler: Handler}
	w.Options.ShowCode = false
	return New(data, w.Options, w.LoggerOptions)
}

// Creates a new error based on the error writer settings with pair Code Name
// set handler as EmptyHandler if you don't need it
func (w ErrorW) NewHTTP(Name string, Code int, Handler ErrorHandler) VanError {
	data := ErrorData{Name: Name, Code: Code, ErrorHandler: Handler}
	w.Options.ShowMessage = false
	return New(data, w.Options, w.LoggerOptions)
}

// Creates a new error based on the error writer settings with Name, and error that is the error cause (it would not be shown it .Error())
// set handler as EmptyHandler if you don't need it
func (w ErrorW) NewWrap(Name string, Cause error, Handler ErrorHandler) VanError {
	data := ErrorData{Name: Name, Cause: Cause, ErrorHandler: Handler}
	w.Options.ShowMessage = false
	w.Options.ShowCode = false
	w.Options.ShowCause = true
	return New(data, w.Options, w.LoggerOptions)
}
