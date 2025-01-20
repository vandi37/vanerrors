package vanerrors

import "time"

type basic struct {
	time      time.Time
	show_time bool
	show      bool
}

var DateFormat string = "06^01^02 15:04:05"

func (e *basic) Error() (s string) {
	if e.show {
		s = "got an error"
	}
	if e.show_time {
		s += ": " + e.time.Format(DateFormat)
	}
	return
}

func (e basic) Date() error {
	e.show_time = true
	return &e
}

func Add() interface {
	error
	Date() error
} {
	return &basic{time: time.Now(), show: true}
}
