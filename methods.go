package vanerrors

import "fmt"

// Formats the error to a string format
// Uses the options to customize the output
func (e VanError) Error() string {
	// Sets the return value
	var err string

	// Sets the options
	options := e.Options

	// Adding data
	if options.ShowDate {
		err += e.Date.Format("2006-01-02 15:04:05 ")
	}

	// Adding severity
	if options.ShowSeverity {
		if options.IntSeverity {
			err += "level: " + SeverityArray[e.Severity] + ","
		} else {
			err += fmt.Sprintf("level: %d,", e.Severity)
		}
	}

	// Adding code
	if options.ShowCode {
		err += fmt.Sprintf(" %d", e.Code)
	}

	// Adding name
	err += " " + e.Name

	// Adding message
	if options.ShowMessage {
		err += ": " + e.Message
	}
	// Adds ", "
	err += ", "

	// Adding description
	if options.ShowDescription && e.Description != nil {
		description := make([]byte, 4096)
		n, errRead := e.Description.Read(description)
		if errRead == nil {
			description = description[:n]
		}

		err += "description: " + string(description) + ", "
	}

	// Adding cause
	if options.ShowCause && e.Cause != nil {
		err += "cause: " + e.Cause.Error()
	}

	// Logging if it needs to be when created
	if e.LoggerOptions.DoLog && !e.LoggerOptions.LogBy && e.logger != nil {
		e.Log()
	}

	return err
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
func (e VanError) UnwrapOne() error {
	var err error
	err = e
	for {
		if getErr := Get(err); getErr != nil {
			err = getErr.Cause
		} else {
			return err
		}
	}
}

// Gets all errors in the VanError error stack
func (e VanError) Unwrap() []error {
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

// Gets the VanError if the error is it
//
// if error is a van error returns the van error
// if error is nol a van error returns nil
func Get(target error) *VanError {
	vanErr, ok := target.(*VanError)
	if !ok {
		return nil
	}
	return vanErr
}
