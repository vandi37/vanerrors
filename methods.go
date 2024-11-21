package vanerrors

import (
	"fmt"
	"io"
)

// Formats the error to a string format
// Uses the options to customize the output
func (e VanError) Error() string {
	// Sets the return value
	var err string

	// Sets the options
	options := e.Options

	// Adding data
	if options.ShowDate {
		err += e.Date.Format("2006/01/02 15:04:05 ")
	}

	// Adding severity
	if options.ShowSeverity {
		if !options.IntSeverity {
			err += SeverityArray[e.Severity] + ", "
		} else {
			err += fmt.Sprintf("level: %d, ", e.Severity)
		}
	}

	// Adding code
	if options.ShowCode {
		err += fmt.Sprintf("%d ", e.Code)
	}

	// Adding name
	err += e.Name

	// Adding message
	if options.ShowMessage {
		err += ": " + e.Message
	}

	if options.ShowDescription || options.ShowCause {
		// Adds ", "
		err += ", "

		// Adding description
		if options.ShowDescription && e.Description != nil {
			description, _ := io.ReadAll(e.Description)
			if len(description) > 0 {
				err += "description: " + string(description)
				if options.ShowCause {
					err += ", "
				}
			}
		}

		// Adding cause
		if options.ShowCause && e.Cause != nil {
			err += "cause: " + e.Cause.Error()
		}
	}
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
	return vanTarget.Code == e.Code && vanTarget.Name == e.Name && e.Options == vanTarget.Options && vanTarget.Severity == e.Severity

}

// Implements the touchable error interface
//
// Touches the error
func (e *VanError) Touch(name string) {
	bufErr := *e
	*e = NewWrap(name, bufErr, bufErr.Logger)
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
