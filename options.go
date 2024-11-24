package vanerrors

// The error options
//
// it describes the behavior of showing the error
type Options struct {
	// Do you need to show the error message
	//
	// recommended : true
	ShowMessage bool `json:"show_message"`
	// Do you need to show the error status code
	//
	// recommended : false
	ShowCode bool `json:"show_code"`
	// Do you want to show description
	//
	// recommended : false
	ShowDescription bool `json:"show_description"`
	// Do you want to show the cause error
	//
	// recommended : false
	ShowCause bool `json:"show_cause"`
	// Do you want to show the error date
	//
	// recommended : false
	ShowDate bool `json:"show_date"`
	// Do you need to show in json format
	//
	// recommended : false
	ShowAsJson bool `json:"show_as_json"`
}

// Default options
//
// Sets as showing:
// Message
var DefaultOptions Options = Options{
	ShowMessage: true,
}

// Sets your options as default for your program
func (o Options) SetAsDefault() {
	DefaultOptions = o
}
