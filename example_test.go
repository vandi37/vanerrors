package vanerrors_test

import (
	"fmt"

	"github.com/vandi37/vanerrors"
)

func ExampleAdd() {
	err := vanerrors.Add()
	fmt.Println(err)

	// Output:
	// got an error
}

func ExampleSimple() {
	err := vanerrors.Simple("error")
	fmt.Println(err)

	// Output:
	// error
}

func ExampleNew() {
	err := vanerrors.New("error", "message")
	fmt.Println(err)

	// Output:
	// error - message
}

func ExampleWrap() {
	err := vanerrors.Wrap("error", vanerrors.Add())
	fmt.Println(err)

	// Output:
	// error it is all because of got an error
}