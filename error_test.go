package vanerrors_test

import (
	"bytes"
	"errors"
	"testing"
	"time"

	"github.com/VandiKond/vanerrors"
)

func TestNewName(t *testing.T) {
	var name = "name"
	var err = vanerrors.NewName(name, nil)
	if err.Name != name {
		t.Fatalf("name don't match, wait '%s', got '%s'", name, err.Name)
	}
}

func TestNewBasic(t *testing.T) {
	var name = "name"
	var message = "message"
	var err = vanerrors.NewBasic(name, message, nil)
	if err.Message != message {
		t.Fatalf("message don't match, wait '%s', got '%s'", message, err.Message)
	}
	if err.Name != name {
		t.Fatalf("name don't match, wait '%s', got '%s'", name, err.Name)
	}
}

func TestNewHTTP(t *testing.T) {
	var name = "name"
	var code = 400
	var err = vanerrors.NewHTTP(name, code, nil)
	if err.Name != name {
		t.Fatalf("name don't match, wait '%s', got '%s'", name, err.Name)
	}
	if err.Code != code {
		t.Fatalf("code don't match, wait '%d', got '%d'", code, err.Code)
	}
}

func TestNewChild(t *testing.T) {
	var name = "name"
	var cause = errors.New("cause")
	var err = vanerrors.NewChild(name, cause, nil)
	if err.Name != name {
		t.Fatalf("name don't match, wait '%s', got '%s'", name, err.Name)
	}
	if err.Cause != cause {
		t.Fatalf("code don't match, wait '%s', got '%s'", cause.Error(), err.Cause.Error())
	}
}

func TestNewDefault(t *testing.T) {
	var data = vanerrors.ErrorData{
		Name:        "name",
		Message:     "message",
		Code:        400,
		Cause:       errors.New("cause"),
		Description: bytes.NewReader([]byte("description")),
		Logger:      nil,
		Severity:    2,
	}
	var err = vanerrors.NewDefault(data)
	if err.Name != data.Name {
		t.Fatalf("name don't match, wait '%s', got '%s'", data.Name, err.Name)
	}
	if err.Message != data.Message {
		t.Fatalf("message don't match, wait '%s', got '%s'", data.Message, err.Message)
	}
	if err.Code != data.Code {
		t.Fatalf("code don't match, wait '%d', got '%d'", data.Code, err.Code)
	}
	if err.Cause != data.Cause {
		t.Fatalf("cause don't match, wait '%v', got '%v'", data.Cause, err.Cause)
	}
	if err.Description != data.Description {
		t.Fatalf("description don't match")
	}
	if err.Logger != data.Logger {
		t.Fatalf("logger don't match, wait '%v', got '%v'", data.Logger, err.Logger)
	}
	if err.Severity != data.Severity {
		t.Fatalf("severity don't match, wait '%d', got '%d'", data.Severity, err.Severity)
	}
}

func TestDefaultValues(t *testing.T) {
	var data = vanerrors.ErrorData{
		Name:        "name",
		Message:     "message",
		Code:        400,
		Cause:       errors.New("cause"),
		Description: bytes.NewReader([]byte("description")),
		Logger:      nil,
		Severity:    2,
	}
	var got, gotOptions, gotLoggerOptions = vanerrors.DefaultValues(data)
	if got != data {
		t.Fatalf("got data don't match, wait '%v', got '%v'", data, got)
	}
	if gotOptions != vanerrors.DefaultOptions {
		t.Fatalf("got options don't match, wait '%v', got '%v'", vanerrors.DefaultOptions, gotOptions)
	}
	if gotLoggerOptions != vanerrors.DefaultLoggerOptions {
		t.Fatalf("got options don't match, wait '%v', got '%v'", vanerrors.DefaultLoggerOptions, gotLoggerOptions)
	}
}

func TestNew(t *testing.T) {
	var data = vanerrors.ErrorData{
		Name:        "name",
		Message:     "message",
		Code:        400,
		Cause:       errors.New("cause"),
		Description: bytes.NewReader([]byte("description")),
		Logger:      nil,
		Severity:    2,
	}
	var err = vanerrors.New(data, vanerrors.DefaultOptions, vanerrors.DefaultLoggerOptions)
	if err.Name != data.Name {
		t.Fatalf("name don't match, wait '%s', got '%s'", data.Name, err.Name)
	}
	if err.Message != data.Message {
		t.Fatalf("message don't match, wait '%s', got '%s'", data.Message, err.Message)
	}
	if err.Code != data.Code {
		t.Fatalf("code don't match, wait '%d', got '%d'", data.Code, err.Code)
	}
	if err.Cause != data.Cause {
		t.Fatalf("cause don't match, wait '%v', got '%v'", data.Cause, err.Cause)
	}
	if err.Description != data.Description {
		t.Fatalf("description don't match")
	}
	if err.Logger != data.Logger {
		t.Fatalf("logger don't match, wait '%v', got '%v'", data.Logger, err.Logger)
	}
	if err.Severity != data.Severity {
		t.Fatalf("severity don't match, wait '%d', got '%d'", data.Severity, err.Severity)
	}
	if err.Options != vanerrors.DefaultOptions {
		t.Fatalf("got options don't match, wait '%v', got '%v'", vanerrors.DefaultOptions, err.Options)
	}
	if err.LoggerOptions != vanerrors.DefaultLoggerOptions {
		t.Fatalf("got options don't match, wait '%v', got '%v'", vanerrors.DefaultLoggerOptions, err.LoggerOptions)
	}
}

func TestGetters(t *testing.T) {
	var data = vanerrors.ErrorData{
		Name:        "name",
		Message:     "message",
		Code:        400,
		Cause:       errors.New("cause"),
		Description: bytes.NewReader([]byte("description")),
		Logger:      nil,
		Severity:    2,
	}
	var err = vanerrors.NewDefault(data)
	var date = err.Date
	if err.Name != vanerrors.GetName(err) {
		t.Fatalf("name don't match, wait '%s', got '%s'", err.Name, vanerrors.GetName(err))
	}
	if err.Message != vanerrors.GetMessage(err) {
		t.Fatalf("message don't match, wait '%s', got '%s'", err.Message, vanerrors.GetMessage(err))
	}
	if err.Code != vanerrors.GetCode(err) {
		t.Fatalf("code don't match, wait '%d', got '%d'", err.Code, vanerrors.GetCode(err))
	}
	if vanerrors.GetDescription(err) != err.Description {
		t.Fatalf("description don't match, wait '%s', got '%s'", err.Description, vanerrors.GetDescription(err))
	}
	if err.Severity != vanerrors.GetSeverityInt(err) {
		t.Fatalf("severity as int don't match, wait '%d', got '%d'", err.Severity, vanerrors.GetSeverityInt(err))
	}
	if vanerrors.SeverityArray[err.Severity] != vanerrors.GetSeverityStr(err) {
		t.Fatalf("severity as str don't match, wait '%s', got '%s'", vanerrors.SeverityArray[err.Severity], vanerrors.GetSeverityStr(err))
	}
	if date.Year() != vanerrors.GetDate(err).Year() ||
		date.Month() != vanerrors.GetDate(err).Month() ||
		date.Day() != vanerrors.GetDate(err).Day() ||
		date.Hour() != vanerrors.GetDate(err).Hour() ||
		date.Minute() != vanerrors.GetDate(err).Minute() ||
		date.Second() != vanerrors.GetDate(err).Second() {
		t.Fatalf("date don't match, wait '%v', got '%v'", date, vanerrors.GetDate(err))
	}

}
func TestLogger(t *testing.T) {
	var description = "description"
	var logger = bytes.Buffer{}
	var data = vanerrors.ErrorData{
		Name:        "name",
		Message:     "message",
		Code:        400,
		Severity:    2,
		Cause:       errors.New("cause"),
		Description: bytes.NewReader([]byte(description)),
		Logger:      &logger,
	}
	_ = vanerrors.NewDefault(data)

	expectedOutput := time.Now().Format("2006/01/02 15:04:05") + " level: error, 400 name: message, description: description, cause: cause\n"
	if logger.String() != expectedOutput {
		t.Fatalf("log output don't match, wait '%s', got '%s'", expectedOutput, logger.String())
	}
}

func TestLoggerOnError(t *testing.T) {
	var description = "description"
	var logger = bytes.Buffer{}
	var logger_options = vanerrors.DefaultLoggerOptions
	logger_options.LogBy = true
	var data = vanerrors.ErrorData{
		Name:        "name",
		Message:     "message",
		Code:        400,
		Severity:    2,
		Cause:       errors.New("cause"),
		Description: bytes.NewReader([]byte(description)),
		Logger:      &logger,
	}
	var err = vanerrors.New(data, vanerrors.DefaultOptions, logger_options)
	_ = err.Error()
	expectedOutput := time.Now().Format("2006/01/02 15:04:05") + " level: error, 400 name: message, description: description, cause: cause\n"
	if logger.String() != expectedOutput {
		t.Fatalf("log output don't match, wait '%s', got '%s'", expectedOutput, logger.String())
	}
}

func TestLoggerOnLog(t *testing.T) {
	var description = "description"
	var logger = bytes.Buffer{}
	var logger_options = vanerrors.DefaultLoggerOptions
	logger_options.DoLog = false
	var data = vanerrors.ErrorData{
		Name:        "name",
		Message:     "message",
		Code:        400,
		Severity:    2,
		Cause:       errors.New("cause"),
		Description: bytes.NewReader([]byte(description)),
		Logger:      &logger,
	}
	var err = vanerrors.New(data, vanerrors.DefaultOptions, logger_options)
	err.Log()
	expectedOutput := time.Now().Format("2006/01/02 15:04:05") + " level: error, 400 name: message, description: description, cause: cause\n"
	if logger.String() != expectedOutput {
		t.Fatalf("log output don't match, wait '%s', got '%s'", expectedOutput, logger.String())
	}
}

func TestNewPanic(t *testing.T) {
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
func TestNewPanicOnError(t *testing.T) {
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
func TestNewPanicOnLog(t *testing.T) {
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

func TestDefaultResults(t *testing.T) {
	var data = vanerrors.ErrorData{}
	var err = vanerrors.NewDefault(data)
	if err.Name != "unknown error" {
		t.Fatalf("name don't match, wait 'unknown error', got '%s'", err.Name)
	}
	if err.Code != 500 {
		t.Fatalf("code don't match, wait '500', got '%d'", err.Code)
	}
	if err.Severity != 2 {
		t.Fatalf("severity don't match, wait '2', got '%d'", err.Severity)
	}
}

func TestError(t *testing.T) {
	var description = "description"
	var options = vanerrors.Options{
		ShowMessage:     true,
		ShowCode:        true,
		ShowSeverity:    true,
		ShowDescription: true,
		ShowCause:       true,
		ShowDate:        true,
	}
	var data = vanerrors.ErrorData{
		Name:        "name",
		Message:     "message",
		Code:        400,
		Severity:    2,
		Cause:       errors.New("cause"),
		Description: bytes.NewReader([]byte(description)),
	}
	var err = vanerrors.New(data, options, vanerrors.DefaultLoggerOptions)
	var got = err.Error()
	expectedOutput := time.Now().Format("2006/01/02 15:04:05") + " level: error, 400 name: message, description: description, cause: cause"
	if got != expectedOutput {
		t.Fatalf("error output don't match, wait '%s', got '%s'", expectedOutput, got)
	}
}

func TestAsAndIs(t *testing.T) {
	var data = vanerrors.ErrorData{
		Name:        "name",
		Message:     "message",
		Code:        400,
		Severity:    2,
		Cause:       errors.New("cause"),
		Description: bytes.NewReader([]byte("description")),
		Logger:      nil,
	}
	var err = vanerrors.NewDefault(data)

	var errTrue vanerrors.VanError
	res := err.As(&errTrue)
	if !res {
		t.Fatalf("got result in as 'false', expected 'true'")
	}
	if !err.Is(errTrue) {
		t.Fatalf("got result in is 'false', expected 'true'")
	}

	var errFalse = errors.New("")
	res = err.As(&errFalse)
	if res {
		t.Fatalf("got result in as 'true', expected 'false'")
	}
	if err.Is(errFalse) {
		t.Fatalf("got result in is 'true', expected 'false'")
	}

	var errNotIs = vanerrors.NewName("not that data", nil)
	if err.Is(errNotIs) {
		t.Fatalf("got result in is 'true', expected 'false'")
	}
}

func TestGet(t *testing.T) {
	var errNil = errors.New("error")
	var VanErrNil = vanerrors.Get(errNil)
	if VanErrNil != nil {
		t.Fatalf("got result in get '%s', expected 'nil'", VanErrNil.Error())
	}

	var errNotNil = vanerrors.NewName("name", nil)
	var VanErrNotNil = vanerrors.Get(errNotNil)
	if VanErrNotNil == nil {
		t.Fatalf("got result in get '%s', expected 'name'", VanErrNotNil.Error())
	}
}
