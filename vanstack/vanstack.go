package vanstack

import (
	"fmt"
	"runtime"
	"strings"
	"time"

	"github.com/VandiKond/vanerrors"
)

// The errors
const (
	CouldNotGetPath = "could not get path"
)

// This interface represents a call to a function
//
// It could be any call that could be add to VanStack
// It should have a path, data, and  date
type Call interface {
	// Getting the path of the called place
	GetPath() string
	// Getting the name of the call
	GetName() string
	// Getting the call date
	GetDate() time.Time
}

// VanCall is a structure for Call Interface
//
// It has basic realization of this interface
type VanCall struct {
	// The path where the call was
	path string
	// The call data
	Name string
	// The call date
	date time.Time
}

// Gets the path to the call
func (c VanCall) GetPath() string {
	return c.path
}

// Gets the name of the call
func (c VanCall) GetName() string {
	return c.Name
}

// Gets the call date
func (c VanCall) GetDate() time.Time {
	return c.date
}

// It is a slice of calls
// It can be used as a stack trace
type VanStack []Call

// Creates a new stack
func NewStack() VanStack {
	return VanStack{}
}

// Ads a call to the stack
func (s *VanStack) Add(call Call) {
	var bufStack = NewStack()
	bufStack = append(bufStack, call)
	bufStack = append(bufStack, *s...)
	*s = bufStack
}

// Gets the last added call
func (s VanStack) GetLastCall() Call {
	return s[0]
}

// Gets the first added call
func (s VanStack) GetFirstCall() Call {
	return s[len(s)-1]
}

// Gets the period when the were the last and the first call
func (s VanStack) GetStackPeriod() time.Duration {
	var LastCall time.Time = s.GetLastCall().GetDate()
	var FirstCall time.Time = s.GetFirstCall().GetDate()
	for _, c := range s {
		if c.GetDate().Before(LastCall) {
			LastCall = c.GetDate()
		}
		if c.GetDate().After(FirstCall) {
			FirstCall = c.GetDate()
		}
	}
	return FirstCall.Sub(LastCall)
}

func (s VanStack) ToString() string {
	var result string
	for i, c := range s {
		result += fmt.Sprintf("%s in %s", c.GetName(), c.GetPath())
		if i < len(s)-1 {
			result += ", "
		}
	}
	return result
}

// Creates a van call
func NewCall(name string) (*VanCall, error) {

	pc, file, line, ok := runtime.Caller(0)
	runtimeFunc := runtime.FuncForPC(pc)
	for i := 1; ok; i++ {
		pc, file, line, ok = runtime.Caller(i)
		runtimeFunc := runtime.FuncForPC(pc)
		pathSlice := strings.Split(runtimeFunc.Name(), ".")
		if pathSlice[len(pathSlice)-1] != "Touch" {
			break
		}
	}
	if !ok {
		return nil, vanerrors.NewName(CouldNotGetPath, nil)
	}
	var pathSlice = strings.Split(file, "/")
	var shortPath = pathSlice[len(pathSlice)-1]
	return &VanCall{
		path: fmt.Sprintf("%s %v: %d", shortPath, runtimeFunc.Name(), line),
		date: time.Now(),
		Name: name,
	}, nil
}

// a error with stack
type StackError struct {
	// any error
	Err error
	// a stack
	Stack VanStack
}

func (e StackError) Error() string {
	if len(e.Stack) > 0 {
		return e.Err.Error() + ", stack: " + e.Stack.ToString()
	}
	return e.Err.Error()
}

// Makes from the error a stack error
func ToStackError(err error) StackError {
	result := StackError{
		Err:   err,
		Stack: NewStack(),
	}
	return result
}

// Touching function for stack error
func (e *StackError) Touch(name string) {
	call, err := NewCall(name)
	if err != nil {
		return
	}
	e.Stack.Add(call)
}
