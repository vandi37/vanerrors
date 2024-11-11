package main

import (
	"fmt"
	"strings"
	"time"
)

func New(errorData ErrorData, options Options, loggerOptions LoggerOptions) (*VanError, error) {
	// Validating name
	if errorName := strings.TrimSpace(errorData.Name); errorName == "" {
		errorData.Name = "unknown error"
	}
	// Validating code
	if errorData.Code <= 0 {
		errorData.Code = 500
	}
	if errorData.Severity < 0 || errorData.Severity > 3 {
		// Sending error
		err, _ := New(
			ErrorData{
				Name:     "severity is not valid",
				Message:  fmt.Sprintf("$d is les than 0 or bigger than 3", errorData.Severity),
				Code:     400,
				Severity: 2,
			},
			DefaultOptions,
			EmptyLoggerOptions,
		)
		return nil, err
	}
	// Creating a vanError
	now := time.Now()
	vanError := VanError{
		Name:          errorData.Name,
		Message:       errorData.Message,
		Code:          errorData.Code,
		Cause:         errorData.Cause,
		Description:   errorData.Description,
		logger:        errorData.Logger,
		LoggerOptions: loggerOptions,
		Options:       options,
		Severity:      errorData.Severity,
		Date:          now,
	}
	// Logging if it needs to be when created
	if loggerOptions.DoLog && !loggerOptions.LogBy && errorData.Logger != nil {
		err := Log(errorData, loggerOptions)
		if err != nil {
			return nil, err
		}
	}

	return &vanError, nil
}

func (e VanError) Error() string
