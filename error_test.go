package vanerrors_test

import (
	"bytes"
	"testing"
	"time"

	"github.com/VandiKond/vanerrors"
)

func TestNew(t *testing.T) {
	testCases := []struct {
		Name    string
		ErrData vanerrors.ErrorData
		Options vanerrors.Options
		Expect  string
	}{
		{
			Name: "Basic values",
			ErrData: vanerrors.ErrorData{
				Name:     "error name",
				Message:  "unknown error in tests",
				Code:     500,
				Severity: 2,
			},
			Options: vanerrors.Options{
				ShowMessage:  true,
				ShowCode:     true,
				ShowSeverity: false,
			},
			Expect: "500 error name: unknown error in tests",
		},
		{
			Name: "Empty name",
			ErrData: vanerrors.ErrorData{
				Message:  "unknown error in tests",
				Code:     500,
				Severity: 2,
			},
			Options: vanerrors.Options{
				ShowMessage:  true,
				ShowCode:     true,
				ShowSeverity: false,
			},
			Expect: "500 unknown error: unknown error in tests",
		},
		{
			Name: "Invalid code",
			ErrData: vanerrors.ErrorData{
				Name:     "error name",
				Message:  "unknown error in tests",
				Code:     -1,
				Severity: 2,
			},
			Options: vanerrors.Options{
				ShowMessage:  true,
				ShowCode:     true,
				ShowSeverity: false,
			},
			Expect: "500 error name: unknown error in tests",
		},
		{
			Name: "Invalid severity",
			ErrData: vanerrors.ErrorData{
				Name:     "error name",
				Message:  "unknown error in tests",
				Code:     500,
				Severity: -1,
			},
			Options: vanerrors.Options{
				ShowMessage:  true,
				ShowCode:     true,
				ShowSeverity: true,
				IntSeverity:  false,
			},
			Expect: "level: 2, 500 error name: unknown error in tests",
		},
		{
			Name: "Cause error",
			ErrData: vanerrors.ErrorData{
				Name:     "error name",
				Message:  "unknown error in tests",
				Code:     500,
				Severity: 2,
				Cause:    vanerrors.New(vanerrors.ErrorData{Name: "cause error", Message: "cause error message"}, vanerrors.DefaultOptions, vanerrors.EmptyLoggerOptions),
			},
			Options: vanerrors.Options{
				ShowMessage:  true,
				ShowCode:     true,
				ShowSeverity: false,
				ShowCause:    true,
			},
			Expect: "500 error name: unknown error in tests, cause: 500 cause error: cause error message",
		},
		{
			Name: "Description error",
			ErrData: vanerrors.ErrorData{
				Name:        "error name",
				Message:     "unknown error in tests",
				Code:        500,
				Severity:    2,
				Description: bytes.NewReader([]byte("This is a description")),
			},
			Options: vanerrors.Options{
				ShowMessage: true, ShowCode: true,
				ShowSeverity:    false,
				ShowDescription: true,
			},
			Expect: "500 error name: unknown error in tests, description: This is a description",
		},
		{
			Name: "Log on creation",
			ErrData: vanerrors.ErrorData{
				Name:        "error name",
				Message:     "unknown error in tests",
				Code:        500,
				Severity:    2,
				Description: bytes.NewReader([]byte("This is a description")),
				Logger:      &bytes.Buffer{},
			},
			Options: vanerrors.Options{
				ShowMessage:     true,
				ShowCode:        true,
				ShowSeverity:    false,
				ShowDescription: true,
			},
			Expect: "500 error name: unknown error in tests, description: This is a description",
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.Name, func(t *testing.T) {
			err := vanerrors.New(testCase.ErrData, testCase.Options, vanerrors.EmptyLoggerOptions)
			result := err.Error()
			if result != testCase.Expect {
				t.Fatalf("expected err.Error() = '%s', but got '%s'", testCase.Expect, result)
			}
		})
	}
}

func TestNewWithLogger(t *testing.T) {
	testCases := []struct {
		Name          string
		ErrData       vanerrors.ErrorData
		Options       vanerrors.Options
		LoggerOptions vanerrors.LoggerOptions
		Expect        string
	}{
		{
			Name: "Log with logger",
			ErrData: vanerrors.ErrorData{
				Name:     "error name",
				Message:  "unknown error in tests",
				Code:     500,
				Severity: 2,
				Logger:   &bytes.Buffer{},
			},
			Options: vanerrors.Options{
				ShowMessage:  true,
				ShowCode:     true,
				ShowSeverity: false,
			},
			LoggerOptions: vanerrors.LoggerOptions{
				DoLog: true,
				LogBy: false,
			},
			Expect: "500 error name: unknown error in tests",
		},
		{
			Name: "Log with logger and custom log by",
			ErrData: vanerrors.ErrorData{
				Name:     "error name",
				Message:  "unknown error in tests",
				Code:     500,
				Severity: 2,
				Logger:   &bytes.Buffer{},
			},
			Options: vanerrors.Options{
				ShowMessage:  true,
				ShowCode:     true,
				ShowSeverity: false,
			},
			LoggerOptions: vanerrors.LoggerOptions{
				DoLog: true,
				LogBy: true,
			},
			Expect: "500 error name: unknown error in tests",
		},
		{
			Name: "Log with logger and custom log by",
			ErrData: vanerrors.ErrorData{
				Name:     "error name",
				Message:  "unknown error in tests",
				Code:     500,
				Severity: 2,
				Logger:   nil,
			},
			Options: vanerrors.Options{
				ShowMessage:  true,
				ShowCode:     true,
				ShowSeverity: false,
			},
			LoggerOptions: vanerrors.LoggerOptions{
				DoLog: true,
				LogBy: true,
			},
			Expect: "500 error name: unknown error in tests",
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.Name, func(t *testing.T) {
			err := vanerrors.New(testCase.ErrData, testCase.Options, testCase.LoggerOptions)
			result := err.Error()
			if result != testCase.Expect {
				t.Fatalf("expected err.Error() = '%s', but got '%s'", testCase.Expect, result)
			}
		})
	}
}

func TestVanError_Error(t *testing.T) {
	now := time.Now()
	testCases := []struct {
		Name     string
		VanError vanerrors.VanError
		Expect   string
	}{
		{
			Name: "Basic values",
			VanError: vanerrors.VanError{
				Name:          "error name",
				Message:       "unknown error in tests",
				Code:          500,
				Severity:      2,
				Date:          now,
				Options:       vanerrors.Options{ShowMessage: true, ShowCode: true, ShowSeverity: false},
				LoggerOptions: vanerrors.EmptyLoggerOptions,
			},
			Expect: "500 error name: unknown error in tests",
		},
		{
			Name: "Show severity as string",
			VanError: vanerrors.VanError{
				Name:          "error name",
				Message:       "unknown error in tests",
				Code:          500,
				Severity:      2,
				Date:          now,
				Options:       vanerrors.Options{ShowMessage: true, ShowCode: true, ShowSeverity: true, IntSeverity: false},
				LoggerOptions: vanerrors.EmptyLoggerOptions,
			}, Expect: "level: 2, 500 error name: unknown error in tests",
		},
		{
			Name: "Show severity as int",
			VanError: vanerrors.VanError{
				Name:          "error name",
				Message:       "unknown error in tests",
				Code:          500,
				Severity:      2,
				Date:          now,
				Options:       vanerrors.Options{ShowMessage: true, ShowCode: true, ShowSeverity: true, IntSeverity: true},
				LoggerOptions: vanerrors.EmptyLoggerOptions,
			},
			Expect: "level: error, 500 error name: unknown error in tests",
		},
		{
			Name: "Show date",
			VanError: vanerrors.VanError{
				Name:          "error name",
				Message:       "unknown error in tests",
				Code:          500,
				Severity:      2,
				Date:          now,
				Options:       vanerrors.Options{ShowMessage: true, ShowCode: true, ShowSeverity: false, ShowDate: true},
				LoggerOptions: vanerrors.EmptyLoggerOptions,
			},
			Expect: now.Format("2006-01-02 15:04:05 ") + "500 error name: unknown error in tests",
		},
		{
			Name: "Show description",
			VanError: vanerrors.VanError{
				Name:          "error name",
				Message:       "unknown error in tests",
				Code:          500,
				Severity:      2,
				Date:          now,
				Options:       vanerrors.Options{ShowMessage: true, ShowCode: true, ShowSeverity: false, ShowDescription: true},
				Description:   bytes.NewReader([]byte("This is a description")),
				LoggerOptions: vanerrors.EmptyLoggerOptions,
			},
			Expect: "500 error name: unknown error in tests, description: This is a description",
		},
		{
			Name: "Show cause",
			VanError: vanerrors.VanError{
				Name:          "error name",
				Message:       "unknown error in tests",
				Code:          500,
				Severity:      2,
				Date:          now,
				Options:       vanerrors.Options{ShowMessage: true, ShowCode: true, ShowSeverity: false, ShowCause: true},
				Cause:         vanerrors.New(vanerrors.ErrorData{Name: "cause error", Message: "cause error message"}, vanerrors.DefaultOptions, vanerrors.EmptyLoggerOptions),
				LoggerOptions: vanerrors.EmptyLoggerOptions,
			},
			Expect: "500 error name: unknown error in tests, cause: 500 cause error: cause error message",
		},
		{
			Name: "Log on Error",
			VanError: vanerrors.VanError{
				Name:     "error name",
				Message:  "unknown error in tests",
				Code:     500,
				Severity: 2,
				Date:     now,
				Options:  vanerrors.Options{ShowMessage: true, ShowCode: true, ShowSeverity: false},
				LoggerOptions: vanerrors.LoggerOptions{
					DoLog: true,
					LogBy: false,
				},
			},
			Expect: "500 error name: unknown error in tests",
		},
		{
			Name: "Log on Error with custom log by",
			VanError: vanerrors.VanError{
				Name:     "error name",
				Message:  "unknown error in tests",
				Code:     500,
				Severity: 2,
				Date:     now,
				Options:  vanerrors.Options{ShowMessage: true, ShowCode: true, ShowSeverity: false},
				LoggerOptions: vanerrors.LoggerOptions{
					DoLog: true,
					LogBy: true,
				},
			},
			Expect: "500 error name: unknown error in tests",
		},
		{
			Name: "Log on Error with custom log by and nil logger",
			VanError: vanerrors.VanError{
				Name:     "error name",
				Message:  "unknown error in tests",
				Code:     500,
				Severity: 2,
				Date:     now,
				Options:  vanerrors.Options{ShowMessage: true, ShowCode: true, ShowSeverity: false},
				LoggerOptions: vanerrors.LoggerOptions{
					DoLog: true,
					LogBy: true,
				},
			},
			Expect: "500 error name: unknown error in tests",
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.Name, func(t *testing.T) {
			result := testCase.VanError.Error()
			if result != testCase.Expect {
				t.Fatalf("expected err.Error() = '%s', but got '%s'", testCase.Expect, result)
			}
		})
	}
}

func TestNew_Panic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()
	var logWriter bytes.Buffer
	vanerrors.New(vanerrors.ErrorData{
		Name:     "error name",
		Message:  "unknown error in tests",
		Code:     500,
		Severity: 3,
		Logger:   &logWriter,
	}, vanerrors.DefaultOptions, vanerrors.LoggerOptions{
		DoLog: true,
	})
}

func TestNew_Panic_On_Error(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()
	var logWriter bytes.Buffer
	_ = vanerrors.New(vanerrors.ErrorData{
		Name:     "error name",
		Message:  "unknown error in tests",
		Code:     500,
		Severity: 3,
		Logger:   &logWriter,
	}, vanerrors.DefaultOptions, vanerrors.LoggerOptions{
		DoLog: true, LogBy: true,
	}).Error()
}

func TestNew_Panic_On_Log(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()
	var logWriter bytes.Buffer
	vanerrors.New(vanerrors.ErrorData{
		Name:     "error name",
		Message:  "unknown error in tests",
		Code:     500,
		Severity: 3,
		Logger:   &logWriter,
	}, vanerrors.DefaultOptions, vanerrors.EmptyLoggerOptions).Log()
}

func TestGetName(t *testing.T) {
	err := vanerrors.NewName("only name", nil)
	name := "only name"
	gotName := vanerrors.GetName(err)
	if name != gotName {
		t.Errorf("error names are different: expect: %s, got %s", name, gotName)
	}
}
