package vanerrors

import (
	"io"
	"time"
)

// gets the error name if it is a vanerror
func GetName(target error) string {
	vanError := Get(target)
	if vanError == nil {
		return ""
	}
	return vanError.Name
}

// gets the error message if it is a vanerror
func GetMessage(target error) string {
	vanError := Get(target)
	if vanError == nil {
		return ""
	}
	return vanError.Message
}

// gets the error code if it is a vanerror
func GetCode(target error) int {
	vanError := Get(target)
	if vanError == nil {
		return 0
	}
	return vanError.Code
}

// gets the error date if it is a vanerror
func GetDate(target error) *time.Time {
	vanError := Get(target)
	if vanError == nil {
		return nil
	}
	return &vanError.Date
}

// gets the description if it is a vanerror
func GetDescription(target error) io.Reader {
	vanError := Get(target)
	if vanError == nil {
		return nil
	}
	return vanError.Description
}
