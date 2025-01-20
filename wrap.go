package vanerrors

import "time"

type wrap struct {
	wrap    error
	message string
	basic
}

func (e *wrap) Error() string {
	return e.message + " it is all because of " + e.wrap.Error() + e.basic.Error()
}

func (e wrap) Date() error {
	e.basic.show_time = true
	return &e
}

func Wrap(message string, err error) interface {
	error
	Date() error
} {
	return &wrap{
		basic:   basic{time: time.Now()},
		wrap:    err,
		message: message,
	}
}

func (e *wrap) Is(target error) bool {
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
