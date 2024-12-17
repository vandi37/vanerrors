package vanerrors_test

import (
	"bytes"
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/vandi37/vanerrors"
)

func ExampleErrorHandler_SetAsDefault() {
	vanerrors.EmptyHandler.SetAsDefault()
	fmt.Println(vanerrors.DefaultHandler)
	// Output: {<nil> false}

}

func ExampleUpdateDefaultLogger() {
	log.SetOutput(nil)
	vanerrors.UpdateDefaultLogger()
	fmt.Println(vanerrors.DefaultHandler)
	//Output: {<nil> false}
}

func ExampleFileHandler() {
	tempFile, err := os.CreateTemp("", "example-log")
	if err != nil {
		panic(err)
	}
	defer os.Remove(tempFile.Name()) //Clean up
	handler := vanerrors.FileHandler(tempFile.Name())
	fmt.Println(handler.DoPanic)
	// Output: false
}

func ExampleFileHandlerPanic() {
	tempFile, err := os.CreateTemp("", "example-log")
	if err != nil {
		panic(err)
	}
	defer os.Remove(tempFile.Name()) //Clean up
	handler := vanerrors.FileHandlerPanic(tempFile.Name())
	fmt.Println(handler.DoPanic)
	// Output: true
}

func ExampleNewDefault() {
	err := vanerrors.NewDefault(vanerrors.ErrorData{
		Name:        "default error",
		Message:     "This is a default error message.",
		Description: "Some extra description.",
	})
	fmt.Println(err.Error())
	// Output:
	// default error: This is a default error message.
}

func ExampleNewName() {
	err := vanerrors.NewName("NamedError", vanerrors.ErrorHandler{DoPanic: false})
	fmt.Println(err.Error())
	// Output:
	// NamedError
}

func ExampleNewBasic() {
	err := vanerrors.NewBasic("BasicError", "This is a basic error.", vanerrors.ErrorHandler{})
	fmt.Println(err.Error())
	// Output:
	// BasicError: This is a basic error.
}

func ExampleNewHTTP() {
	err := vanerrors.NewHTTP("HTTPError", 404, vanerrors.ErrorHandler{})
	fmt.Println(err.Error())
	// Output:
	// 404 HTTPError
}

func ExampleNewWrap() {
	cause := fmt.Errorf("underlying cause")
	err := vanerrors.NewWrap("WrappedError", cause, vanerrors.ErrorHandler{})
	fmt.Println(err.Error())
	// Output:
	// WrappedError, cause: underlying cause
}

func ExampleNewSimple() {
	err := vanerrors.NewSimple("SimpleError", "Message 1", "Message 2", "Message 3")
	fmt.Println(err.Error())
	// Output:
	// SimpleError: Message 1, description: Message 2, Message 3
}

func ExampleNew() {
	err := vanerrors.New(vanerrors.ErrorData{
		Name:        "example error",
		Message:     "here could be the error message",
		Code:        500,
		Cause:       errors.New("some cause error"),
		Description: "you can add more information here. the more - the better",
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
		Description: "you can add more information here. the more - the better",
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
	// 500 example error: here could be the error message, description: you can add more information here. the more - the better, cause: some cause error
	// 500 readme error: here could be the error message, description: you can add more information here. the more - the better, cause: some cause error
}

func ExampleOptions_SetAsDefault() {
	newDefaults := vanerrors.Options{ShowMessage: true, ShowCode: true}
	newDefaults.SetAsDefault()
	fmt.Println(vanerrors.DefaultOptions)
	// Output: {true true false false false false}
}

func ExampleLoggerOptions_SetAsDefault() {
	newDefaults := vanerrors.LoggerOptions{DoLog: true, Options: vanerrors.Options{ShowMessage: true}}
	newDefaults.SetAsDefault()
	fmt.Println(vanerrors.DefaultLoggerOptions)
	// Output: {true {true false false false false false} false}
}

func ExampleNewW() {
	w := vanerrors.NewW(vanerrors.Options{ShowMessage: true}, vanerrors.EmptyLoggerOptions, vanerrors.EmptyHandler)
	fmt.Println(w)
	// Output: &{{true false false false false false} {false {false false false false true false} false} {<nil> false}}
}

func ExampleErrorW_New() {
	w := vanerrors.NewW(vanerrors.Options{ShowMessage: true}, vanerrors.EmptyLoggerOptions, vanerrors.EmptyHandler)
	err := w.New(vanerrors.ErrorData{Name: "TestError", Message: "This is a test."})
	fmt.Println(err.Error())
	// Output: TestError: This is a test.
}

func ExampleErrorW_NewName() {
	w := vanerrors.NewW(vanerrors.Options{}, vanerrors.EmptyLoggerOptions, vanerrors.EmptyHandler)
	err := w.NewName("NameOnlyError")
	fmt.Println(err.Error())
	// Output: NameOnlyError
}

func ExampleErrorW_NewBasic() {
	w := vanerrors.NewW(vanerrors.Options{}, vanerrors.EmptyLoggerOptions, vanerrors.EmptyHandler)
	err := w.NewBasic("BasicError", "This is basic.")
	fmt.Println(err.Error())
	// Output: BasicError: This is basic.
}

func ExampleErrorW_NewHTTP() {
	w := vanerrors.NewW(vanerrors.Options{}, vanerrors.EmptyLoggerOptions, vanerrors.EmptyHandler)
	err := w.NewHTTP("HTTPError", 404)
	fmt.Println(err.Error())
	// Output: 404 HTTPError
}

func ExampleErrorW_NewWrap() {
	w := vanerrors.NewW(vanerrors.Options{ShowCause: true}, vanerrors.EmptyLoggerOptions, vanerrors.EmptyHandler)
	cause := fmt.Errorf("cause error")
	err := w.NewWrap("WrappedError", cause)
	fmt.Println(err.Error())
	// Output: WrappedError, cause: cause error
}

func ExampleErrorW_SetAsDefault() {
	w := vanerrors.NewW(vanerrors.Options{ShowMessage: true}, vanerrors.LoggerOptions{DoLog: true}, vanerrors.StdoutHandler)
	w.SetAsDefault()
	err := vanerrors.NewSimple("test", "test message")
	fmt.Println(err.Error())
	// Output: test: test message

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

func GetError() error {
	return vanerrors.NewName("readme error", vanerrors.EmptyHandler)
}

func GetOtherError() error {
	return errors.New("some error")
}

func ExampleVanError_As() {
	err := vanerrors.NewSimple("TestError", "Original Message")
	var target = vanerrors.NewSimple("")
	if err.As(target) {
		fmt.Println(target.Message)
	}
	// Output: Original Message
}

func ExampleVanError_Unwrap() {
	cause := errors.New("Cause Error")
	err := vanerrors.NewWrap("WrappedError", cause, vanerrors.EmptyHandler)
	unwrapped := err.Unwrap()
	fmt.Println(unwrapped.Error())
	// Output: Cause Error
}

func ExampleVanError_UnwrapAll() {
	cause := errors.New("Cause 2 Error")
	vanCause := vanerrors.NewWrap("Cause 1 Error", cause, vanerrors.EmptyHandler)
	err := vanerrors.NewWrap("WrappedError", vanCause, vanerrors.EmptyHandler)
	allErrors := err.UnwrapAll()
	for i, e := range allErrors {
		fmt.Printf("%d: %s\n", i+1, e.Error())
	}
	// Output:
	// 1: Cause 1 Error
	// 2: Cause 2 Error
}

func ExampleVanError_Is() {
	err1 := vanerrors.NewSimple("TestError", "Message 1", "Description 1")
	err2 := vanerrors.NewSimple("TestError", "Message 2", "Description 2")
	fmt.Println(err1.Is(err2))
	// Output: true
}

func ExampleVanError_Touch() {
	err := vanerrors.NewSimple("OriginalError", "Original Message")
	err.Touch("TouchedError")
	fmt.Println(err.Error())
	// Output: TouchedError, cause: OriginalError: Original Message

}

func ExampleGet() {
	err := vanerrors.NewSimple("TestError", "Original Message")
	vanErr := vanerrors.Get(err)
	fmt.Println(vanErr.Message)
	// Output: Original Message
}

func GetError2gGetters() error {
	return vanerrors.NewDefault(vanerrors.ErrorData{
		Name:        "readme error",
		Message:     "here could be the error message",
		Code:        500,
		Cause:       errors.New("some cause error"),
		Description: "you can add more information here. the more - the better",
	})
}

func ExampleGetName() {
	err := vanerrors.NewSimple("MyError", "Some message")
	name := vanerrors.GetName(err)
	fmt.Println(name)
	// Output: MyError
}

func ExampleGetMessage() {
	err := vanerrors.NewSimple("MyError", "Some message")
	message := vanerrors.GetMessage(err)
	fmt.Println(message)
	// Output: Some message
}

func ExampleGetCode() {
	err := vanerrors.NewSimple("MyError", "Some message")
	code := vanerrors.GetCode(err)
	fmt.Println(code)
	// Output: 500
}

func ExampleGetDate() {
	err := vanerrors.NewSimple("MyError", "Some message")
	date := vanerrors.GetDate(err)
	fmt.Println(*date == err.Date)
	// Output: true

}

func ExampleGetDescription() {
	err := vanerrors.NewSimple("MyError", "Some message", "A longer description")
	description := vanerrors.GetDescription(err)
	fmt.Println(description)
	// Output: A longer description
}

func ExampleDecode() {
	jsonData := []byte(`{"date": "2024-07-26T10:30:00Z", "main": "MyError: Some message", "description": "More details here", "cause": "Underlying cause"}`)
	var vanError vanerrors.JsonVanError
	vanerrors.Decode(bytes.NewReader(jsonData), &vanError)
	fmt.Println(vanError.Main)
	// Output: MyError: Some message
}

func ExampleDecodeString() {
	jsonString := `{"date": "2024-07-26T10:30:00Z", "main": "MyError: Some message", "description": "More details here", "cause": "Underlying cause"}`
	var vanError vanerrors.JsonVanError
	vanerrors.DecodeString(jsonString, &vanError)
	fmt.Println(vanError.Description)
	// Output: More details here
}

func ExampleVanError_Error_json() {
	err := vanerrors.NewWrap("error", vanerrors.NewSimple("cause"), vanerrors.EmptyHandler)
	err.Options.ShowAsJson = true

	fmt.Println(err.Error())

	// Output:
	// {"main":"error"}
}
