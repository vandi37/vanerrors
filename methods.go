package vanerrors

// Formats the error to a string format
// Uses the options to customize the output
func (e VanError) Error() string {
	err := e.getView(false)
	// Logging if it needs to be when called
	if e.LoggerOptions.DoLog && e.LoggerOptions.LogBy && e.Logger != nil {
		e.Log()
	}

	return err
}

// Checks is the target a VanError
// If true replaces the VanError with the target
func (e VanError) As(target any) bool {
	p, ok := target.(*VanError)
	if !ok {
		return false
	}
	*p = e
	return true
}

// Unwraps the most deep error in the VanError error stack
func (e VanError) Unwrap() error {
	return e.Cause
}

// Gets all errors in the VanError error stack
func (e VanError) UnwrapAll() []error {
	var allErrors []error
	var err error
	err = e
	for {
		if getErr := Get(err); getErr != nil {
			err = getErr.Cause
			if err != nil {
				allErrors = append(allErrors, err)
			} else {
				break
			}
		} else {
			break
		}
	}

	return allErrors
}

// Checks is the target is the same as the vanError
// It does check date, description and log data (log options and logger)
func (e VanError) Is(target error) bool {
	vanTarget := Get(target)
	if vanTarget == nil {
		return false
	}
	return vanTarget.Code == e.Code && vanTarget.Name == e.Name

}

// Implements the touchable error interface
//
// Touches the error
func (e *VanError) Touch(name string) {
	bufErr := *e
	*e = NewWrap(name, bufErr, ErrorHandler{Logger: bufErr.Logger})
	e.Options.ShowCause = true
}

// Gets the VanError if the error is it
//
// if error is a van error returns the van error
// if error is nol a van error returns nil
func Get(target error) *VanError {
	vanErr, ok := target.(VanError)
	if !ok {
		return nil
	}
	return &vanErr
}
