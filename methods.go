package vanerrors

import (
	"errors"
	"fmt"
)

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

// Creates a log of VanError based on it settings
// Uses the logger options to customize the output
func (e VanError) Log() {
	if e.Logger == nil {
		return
	}
	result := e.getView(true)

	// Getting the result string
	fmt.Fprintln(e.Logger, result)
}

// Checks is the target a VanError
// If true sets the target as the VanError
func (e VanError) As(target any) bool {
	p, ok := target.(*VanError)
	if !ok {
		return false
	}
	*p = e
	return true
}

// Gets the error cause
func (e VanError) Unwrap() error {
	return e.Cause
}

// Gets all errors in the VanError error stack
func (e VanError) UnwrapAll() []error {
	var allErrors []error
	var err error
	err = e
	for {
		err = errors.Unwrap(err)
		if err != nil {
			vanErr := Get(err)
			if vanErr != nil {
				vanErr.Options.ShowCause = false
				err = *vanErr
			}
			allErrors = append(allErrors, err)
		} else {
			break
		}

	}
	return allErrors
}

// Checks is the target is the same as the vanError
//
// Check parameters:
// - Code
// - Name
func (e VanError) Is(target error) bool {
	vanTarget := Get(target)
	if vanTarget == nil {
		return false
	}
	return vanTarget.Code == e.Code && vanTarget.Name == e.Name

}

// It send the original error deeper in the error stack
func (e *VanError) Touch(name string) {
	bufErr := *e
	*e = NewWrap(name, bufErr, ErrorHandler{Logger: bufErr.Logger})
	e.Options.ShowCause = true
}

// it gets the VanError out of the target
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
