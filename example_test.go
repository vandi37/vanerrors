package vanerrors_test

import (
	"bytes"
	"errors"
	"fmt"
	"os"

	"github.com/VandiKond/vanerrors"
)

func ExampleNewName() {
	err := vanerrors.NewName("readme error", vanerrors.EmptyHandler)

	fmt.Println(err.Error())

	// Output:
	// readme error
}

func ExampleNewBasic() {
	err := vanerrors.NewBasic("readme error", "here could be the error message", vanerrors.EmptyHandler)
	fmt.Println(err.Error())

	// Output:
	// readme error: here could be the error message
}

func ExampleNewHTTP() {
	err := vanerrors.NewHTTP("readme error", 500, vanerrors.EmptyHandler)
	fmt.Println(err.Error())

	// Output:
	// 500 readme error
}

func ExampleNewWrap() {
	err := vanerrors.NewWrap("readme error", errors.New("some cause"), vanerrors.EmptyHandler)
	fmt.Println(err.Error())

	// Output:
	// readme error, cause: some cause
}

func ExampleNewDefault() {
	err := vanerrors.NewDefault(vanerrors.ErrorData{
		Name:        "readme error",                                                                      // The error name
		Message:     "here could be the error message",                                                   // The error message
		Code:        500,                                                                                 // The error code
		Cause:       errors.New("some cause error"),                                                      // The cause error (it could be nil)
		Description: bytes.NewReader([]byte("you can add more information here. the more - the better")), // The error description (as io.Reader)
		ErrorHandler: vanerrors.ErrorHandler{
			Logger:  nil,   // No logger
			DoPanic: false, // Don't panic
		},
	})
	fmt.Println(err.Error()) // Would be shown only name and message because of the default settings

	// Output:
	// readme error: here could be the error message
}

func ExampleNew() {
	err := vanerrors.New(vanerrors.ErrorData{
		Name:        "readme error",
		Message:     "here could be the error message",
		Code:        500,
		Cause:       errors.New("some cause error"),
		Description: bytes.NewReader([]byte("you can add more information here. the more - the better")),
		ErrorHandler: vanerrors.ErrorHandler{
			Logger:  nil,
			DoPanic: false,
		},
	},
		vanerrors.Options{
			ShowMessage:     true,  // Do you need to show error message
			ShowCode:        true,  // Do you need to show error code
			ShowDescription: true,  // Do you need to show description
			ShowCause:       true,  // Do you need to show error case
			ShowDate:        false, // Do you need to show at which date the error was created
			ShowAsJson:      false, // Do you need to show in json format
		}, vanerrors.EmptyLoggerOptions, // Empty log settings
	)
	fmt.Println(err.Error())

	err = vanerrors.New(vanerrors.ErrorData{
		Name:        "readme error",
		Message:     "here could be the error message",
		Code:        500,
		Cause:       errors.New("some cause error"),
		Description: bytes.NewReader([]byte("you can add more information here. the more - the better")),
		ErrorHandler: vanerrors.ErrorHandler{
			Logger:  os.Stdout,
			DoPanic: false,
		},
	},
		vanerrors.Options{
			ShowMessage: true,
			ShowCode:    true,
		}, vanerrors.LoggerOptions{
			DoLog: true, // Do the program need to log data automatically
			Options: vanerrors.Options{ // The options
				ShowMessage:     true,
				ShowCode:        true,
				ShowDescription: true,
				ShowCause:       true,
				ShowDate:        false,
			},
			LogBy: true, // Do you need to log when the error is created or when it is called (false: created, true: called)
		},
	)
	err.Log()

	// Output:
	// 500 readme error: here could be the error message, description: you can add more information here. the more - the better, cause: some cause error
	// 500 readme error: here could be the error message, description: you can add more information here. the more - the better, cause: some cause error
}

func ExampleOptions_SetAsDefault() {
	options := vanerrors.Options{
		ShowCause: true,
	}

	// Setting options as default
	options.SetAsDefault()

	fmt.Println(vanerrors.DefaultOptions.ShowCause)
	fmt.Println(vanerrors.DefaultOptions.ShowMessage) // In default options it was enabled

	// Output:
	// true
	// false
}

func ExampleLoggerOptions_SetAsDefault() {
	loggerOptions := vanerrors.LoggerOptions{
		LogBy: true,
	}

	// Setting logger options as default
	loggerOptions.SetAsDefault()

	fmt.Println(vanerrors.DefaultLoggerOptions.LogBy)
	fmt.Println(vanerrors.DefaultLoggerOptions.ShowDate) // In default options it was enabled

	// Output:
	// true
	// false
}

func ExampleNewW() {
	errorW := vanerrors.NewW(
		vanerrors.Options{
			ShowMessage: true,
		},
		vanerrors.EmptyLoggerOptions,
	)
	err := errorW.NewBasic("readme error", "here could be the error message", vanerrors.EmptyHandler)
	fmt.Println(err.Error())

	// Output:
	// readme error: here could be the error message
}

func ExampleVanError_Error() {
	err := vanerrors.NewBasic("readme error", "here could be the error message", vanerrors.EmptyHandler)
	errorText := err.Error()
	fmt.Println(errorText)

	// Output:
	// readme error: here could be the error message
}

func ExampleVanError_Log() {
	vanerrors.LoggerOptions{
		Options: vanerrors.Options{
			ShowMessage: true,
		},
	}.SetAsDefault()
	err := vanerrors.NewBasic("readme error", "here could be the error message", vanerrors.ErrorHandler{
		Logger: os.Stdout,
	})
	err.Log()

	// Output:
	// readme error: here could be the error message
}

func ExampleVanError_As() {
	err := vanerrors.NewBasic("readme error", "here could be the error message", vanerrors.EmptyHandler)

	var targetErr error
	fmt.Println(err.As(&targetErr))

	var targetErr2 = vanerrors.NewName("target", vanerrors.EmptyHandler)
	fmt.Println(err.As(&targetErr2))
	fmt.Println(err == targetErr2)

	// Output:
	// false
	// true
	// true
}

func ExampleVanError_Unwrap() {
	err := vanerrors.NewBasic("readme error", "here could be the error message", vanerrors.EmptyHandler)

	err.Cause = vanerrors.NewName("some cause", vanerrors.EmptyHandler)
	fmt.Println(err.Unwrap())

	// Output:
	// some cause
}

func ExampleVanError_UnwrapAll() {
	// UnwrapAll
	err := vanerrors.NewBasic("readme error", "here could be the error message", vanerrors.EmptyHandler)
	err2 := vanerrors.NewWrap("some cause", vanerrors.NewName("some other cause", vanerrors.EmptyHandler), vanerrors.EmptyHandler)
	err.Cause = err2
	fmt.Println(err.UnwrapAll())

	// Output:
	// [some cause some other cause]
}

func ExampleVanError_Is() {
	err := vanerrors.NewBasic("readme error", "here could be the error message", vanerrors.EmptyHandler)
	err2 := vanerrors.NewBasic("readme error", "here could be the error message", vanerrors.EmptyHandler)
	fmt.Println(err.Is(err2))
	err3 := vanerrors.NewBasic("other readme error", "here could be the error message", vanerrors.EmptyHandler)
	fmt.Println(err3.Is(err))

	// Output:
	// true
	// false
}

func GetError() error {
	return vanerrors.NewName("readme error", vanerrors.EmptyHandler)
}

func GetOtherError() error {
	return errors.New("some error")
}

func ExampleGet() {
	fmt.Println(vanerrors.Get(GetError()), fmt.Sprintf("%T", vanerrors.Get(GetError())))
	fmt.Println(vanerrors.Get(GetOtherError()), fmt.Sprintf("%T", vanerrors.Get(GetOtherError())))

	// Output:
	// readme error *vanerrors.VanError
	// <nil> *vanerrors.VanError
}

func ExampleVanError_Touch() {
	orgErr := vanerrors.NewName("readme error", vanerrors.EmptyHandler)
	orgErr.Touch("readme touch")
	fmt.Println(orgErr)

	// Output:
	// readme touch, cause: readme error
}

func GetError2gGetters() error {
	return vanerrors.NewDefault(vanerrors.ErrorData{
		Name:        "readme error",
		Message:     "here could be the error message",
		Code:        500,
		Cause:       errors.New("some cause error"),
		Description: bytes.NewReader([]byte("you can add more information here. the more - the better")),
	})
}

func Example_getters() {
	err := GetError2gGetters()
	err2 := errors.New("not vandi error")

	fmt.Println(vanerrors.GetName(err))

	fmt.Println(vanerrors.GetMessage(err))
	fmt.Println(vanerrors.GetMessage(err2))

	fmt.Println(vanerrors.GetCode(err))
	fmt.Println(vanerrors.GetCode(err2))

	fmt.Printf("%T\n", vanerrors.GetDescription(err))
	fmt.Printf("%T\n", vanerrors.GetDescription(err2))

	// GetDate
	fmt.Println(vanerrors.GetDate(err2))

	// Output:
	// readme error
	// here could be the error message
	//
	// 500
	// 0
	// *bytes.Reader
	// <nil>
	// <nil>
}
