package vanerrors

import (
	"time"
)

// It gets the error name
func GetName(target error) string {
	vanError := Get(target)
	if vanError == nil {
		return ""
	}
	return vanError.Name
}

// Gets the error message
func GetMessage(target error) string {
	vanError := Get(target)
	if vanError == nil {
		return ""
	}
	return vanError.Message
}

// Gets the error code
func GetCode(target error) int {
	vanError := Get(target)
	if vanError == nil {
		return 0
	}
	return vanError.Code
}

// Gets the date, when the error was created
func GetDate(target error) *time.Time {
	vanError := Get(target)
	if vanError == nil {
		return nil
	}
	return &vanError.Date
}

// Gets the error description
func GetDescription(target error) string {
	vanError := Get(target)
	if vanError == nil {
		return ""
	}
	return vanError.Description
}
