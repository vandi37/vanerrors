package vantouch

// This interface represents an error that could be touched
// a.k. add a new note about what is happening
//
// as for stack errors: adding a new call
//
// as for van errors: creating a new wrap
type TouchableError interface {
	// touch function
	Touch(name string)
	// Error interface
	error
}

// Function that can touch any error if possible
func Touch(err error, name string) {
	bufErr := err
	touchableError, ok := bufErr.(TouchableError)
	if !ok {
		return
	}
	touchableError.Touch(name)

	err = touchableError
}
