package vanstack

import (
	"fmt"
	"runtime"
	"strings"
	"time"

	"github.com/vandi37/vanerrors"
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
	// Sets the setting of showing the path
	SetSettings(Settings)
	// Sets showing name
	DoShowName(bool)
	// Call to string
	fmt.Stringer
}

// The path settings
type Settings struct {
	FileLen int
	ShowFn  bool
}

// Standard settings
var StdSettings = Settings{
	FileLen: 1,
	ShowFn:  false,
}

type path struct {
	line     int
	file     []string
	fn       string
	settings Settings
}

func newPath(line int, file string, fn string, settings Settings) path {
	return path{line: line, file: strings.Split(file, "/"), fn: fn, settings: settings}
}

// Convert the path to string
func (p path) String() string {
	set := p.settings
	if set.FileLen <= 0 || set.FileLen > len(p.file) {
		set.FileLen = len(p.file) - 1
	}
	var res string
	res += strings.Join(p.file[set.FileLen:], "/")
	if set.ShowFn {
		res += " " + p.fn
	}
	res += fmt.Sprint(":", p.line)
	return res
}

// VanCall is a structure for Call Interface
//
// It has basic realization of this interface
type VanCall struct {
	// The path where the call was
	path path
	// The call data
	Name string
	// The call date
	date time.Time
	// Show name
	ShowName bool
}

// Gets the path to the call
func (c *VanCall) GetPath() string {
	return c.path.file[len(c.path.file)-1]
}

// Gets the name of the call
func (c *VanCall) GetName() string {
	return c.Name
}

// Gets the call date
func (c *VanCall) GetDate() time.Time {
	return c.date
}

// Sets settings to the call
func (c *VanCall) SetSettings(s Settings) {
	c.path.settings = s
}

// Sets showing name of call
func (c *VanCall) DoShowName(b bool) {
	c.ShowName = b
}

// Converts van call to string
func (c *VanCall) String() string {
	var res string
	if c.ShowName {
		res += c.Name + " "
	}
	res += c.path.String()
	return res
}

// Creates a van call
func NewCall(name string) (*VanCall, error) {

	pc, file, line, ok := runtime.Caller(0)
	fn := runtime.FuncForPC(pc)
	for i := 1; ok; i++ {
		pc, file, line, ok = runtime.Caller(i)
		runtimeFunc := runtime.FuncForPC(pc)
		pathSlice := strings.Split(runtimeFunc.Name(), "/")
		if pathSlice[len(pathSlice)-1] != "Touch" {
			break
		}
	}
	if !ok {
		return nil, vanerrors.NewSimple(CouldNotGetPath)
	}
	return &VanCall{
		path: newPath(line, file, fn.Name(), StdSettings),
		date: time.Now(),
		Name: name,
	}, nil
}

// It is a slice of calls
// It can be used as a stack trace
type VanStack struct {
	calls     []Call
	Separator string
}

// Creates a new stack
func NewStack() *VanStack {
	return &VanStack{
		calls:     []Call{},
		Separator: "\n",
	}
}

// Ads a call to the stack
func (s *VanStack) Add(call Call) {
	s.calls = append([]Call{call}, s.calls...)
}

// Fils the stack
func (s *VanStack) Fill(name string, n int) {
	for i := 1; i <= n; i++ {
		pc, file, line, ok := runtime.Caller(i)
		if !ok {
			break
		}
		fn := runtime.FuncForPC(pc)
		s.Add(&VanCall{
			path: newPath(line, file, fn.Name(), StdSettings),
			date: time.Now(),
			Name: fmt.Sprintf("%s %d", name, len(s.calls)),
		})
	}
}

// Gets the period when the were the last and the first call
func (s *VanStack) Period() time.Duration {
	if len(s.calls) <= 1 {
		return 0
	}
	var LastCall time.Time = s.calls[0].GetDate()
	var FirstCall time.Time = s.calls[len(s.calls)-1].GetDate()
	for _, c := range s.calls {
		if c.GetDate().Before(LastCall) {
			LastCall = c.GetDate()
		}
		if c.GetDate().After(FirstCall) {
			FirstCall = c.GetDate()
		}
	}
	return FirstCall.Sub(LastCall)
}

// Prints the short version of the stack
func (s *VanStack) String() string {
	var result string
	for _, c := range s.calls {
		result += fmt.Sprintf("%v%s", c, s.Separator)
	}
	return result
}

// Gets the length of the stack
func (s *VanStack) Len() int {
	return len(s.calls)
}

// Sets the separator of the stack
func (s *VanStack) SetSeparator(str string) {
	s.Separator = str
}

// Gets all calls of the stack
func (s *VanStack) GetCalls() []Call {
	return s.calls
}

// a error with stack
type StackError struct {
	// any error
	error
	// a stack
	Stack *VanStack
	// Do you need to show the stack
	ShowStack bool
}

// Converts the stack error to string
func (e StackError) Error() string {
	if len(e.Stack.calls) > 0 && e.ShowStack {
		return e.error.Error() + ", stack: " + e.Stack.String()
	}
	return e.error.Error()
}

// Makes from the error a stack error
func ToStackError(err error) StackError {
	result := StackError{
		error:     err,
		Stack:     NewStack(),
		ShowStack: true,
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

// Gets the error out of the stack
func OutOfStack(err error) error {
	stackError, ok := err.(StackError)
	if !ok {
		return err
	}
	return stackError.error
}

// Gets stack out of error
func OutOfError(err error) *VanStack {
	stackError, ok := err.(StackError)
	if !ok {
		return nil
	}
	return stackError.Stack
}

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
//
// Use a pointer, when using touch, or the error wouldn't change
//
//	func main() {
//	    vanstack.Touch(err) // wrong
//	    vanstack.Touch(&err) // right
//	}
func Touch(err error, name string) {
	bufErr := err
	touchableError, ok := bufErr.(TouchableError)
	if !ok {
		return
	}
	touchableError.Touch(name)

	err = touchableError
}
