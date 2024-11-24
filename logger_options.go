package vanerrors

// Logger error options
//
// it describes the behavior of logging the error
//
// You can save base options into a .json file for same settings
type LoggerOptions struct {
	// Do you want to log the error
	//
	// recommended : true
	DoLog bool `json:"do_log"`
	// Options of showing logs
	Options
	// Do you need to log at the calling .Error() or by creating the error
	//
	// false : New() function
	// true : .Error() function
	//
	// recommended : false
	LogBy bool `json:"log_by"`
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
	DoLog: true,
	Options: Options{
		ShowMessage:     true,
		ShowCode:        true,
		ShowDescription: true,
		ShowCause:       true,
		ShowDate:        true,
	},
}

// An empty logger settings
//
// It makes everything off
var EmptyLoggerOptions LoggerOptions = LoggerOptions{
	Options: Options{ShowDate: true},
}

// Sets your logger options as default for your program
func (o LoggerOptions) SetAsDefault() {
	DefaultLoggerOptions = o
}
