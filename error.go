package vanerrors

import (
	"time"
)

type name struct {
	message string
	basic
}

func (e name) Date() error {
	e.basic.show_time = true
	return &e
}

func (e *name) Error() string {
	return e.message + e.basic.Error()
}

func Simple(message string) interface {
	error
	Date() error
} {
	return &name{message: message, basic: basic{time: time.Now()}}
}

type desc struct {
	message     string
	description string
	basic
}

func (e desc) Date() error {
	e.basic.show_time = true
	return &e
}

func (e *desc) Error() string {
	return e.message + " - " + e.description + e.basic.Error()
}
func New(message string, description string) interface {
	error
	Date() error
} {
	return &desc{message: message, description: description, basic: basic{time: time.Now()}}
}

func (e *desc) Is(target error) bool {
	d, ok := target.(*desc)
	if ok {

		return d.message == e.message
	}

	w, ok := target.(*wrap)
	if ok {
		return w.message == e.message
	}

	n, ok := target.(*name)
	if ok {
		return n.message == e.message
	}
	return false
}

func (e *name) Is(target error) bool {
	d, ok := target.(*desc)
	if ok {

		return d.message == e.message
	}

	w, ok := target.(*wrap)
	if ok {
		return w.message == e.message
	}

	n, ok := target.(*name)
	if ok {
		return n.message == e.message
	}
	return false
}
