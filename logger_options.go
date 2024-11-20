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
	// Do you need to show the error message
	//
	// recommended : true
	ShowMessage bool `json:"show_message"`
	// Do you need to show the error status code
	//
	// recommended : true
	ShowCode bool `json:"show_code"`
	// Do you need to show severity level
	//
	// recommended : true
	ShowSeverity bool `json:"show_severity"`
	// Do you need to show severity level as int or string
	// false : string
	// true : int
	//
	// recommended : false
	IntSeverity bool `json:"int_severity"`
	// Do you want to show description
	//
	// recommended : true
	ShowDescription bool `json:"show_description"`
	// Do you want to show the cause error
	//
	// recommended : true
	ShowCause bool `json:"show_cause"`
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

// Sets your logger options as default for your program
func (o LoggerOptions) SetAsDefault() {
	DefaultLoggerOptions = o
}
