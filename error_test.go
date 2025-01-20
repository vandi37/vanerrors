package vanerrors_test

import (
	"testing"
	"time"

	"github.com/vandi37/vanerrors"
)

func TestAdd(t *testing.T) {
	err := vanerrors.Add()
	date := time.Now()

	if err.Error() != "got an error" {
		t.Fatal("different errors")
	}

	if err.Date().Error() != "got an error: "+date.Format(vanerrors.DateFormat) {
		t.Fatal("different errors")
	}
}

func TestSimple(t *testing.T) {
	err := vanerrors.Simple("error")
	date := time.Now()

	if err.Error() != "error" {
		t.Fatal("different errors")
	}

	if err.Date().Error() != "error: "+date.Format(vanerrors.DateFormat) {
		t.Fatal("different errors")
	}
}

func TestNew(t *testing.T) {
	err := vanerrors.New("error", "message")
	date := time.Now()

	if err.Error() != "error - message" {
		t.Fatal("different errors")
	}

	if err.Date().Error() != "error - message: "+date.Format(vanerrors.DateFormat) {
		t.Fatal("different errors")
	}
}

func TestWrap(t *testing.T) {
	err := vanerrors.Wrap("error", vanerrors.Add())
	date := time.Now()

	if err.Error() != "error it is all because of got an error" {
		t.Fatal("different errors")
	}

	if err.Date().Error() != "error it is all because of got an error: "+date.Format(vanerrors.DateFormat) {
		t.Fatal("different errors")
	}
}
