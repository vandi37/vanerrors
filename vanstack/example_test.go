package vanstack_test

import (
	"fmt"
	"time"

	"github.com/vandi37/vanerrors"
	"github.com/vandi37/vanerrors/vanstack"
)

func ExampleSettings_SetAsDefault() {
	mySettings := vanstack.Settings{10, true}

	fmt.Println(vanstack.DefaultSettings)

	mySettings.SetAsDefault()

	fmt.Println(vanstack.DefaultSettings)

	// Output:
	// {1 false}
	// {10 true}
}

func ExampleVanCall_GetPath() {
	call, err := vanstack.NewCall("call" /*the call name*/)
	if err != nil {
		panic(err) // Some error handling
	}

	fmt.Println(call.GetPath())

	// Output:
	// example_test.go:24
}

func ExampleVanCall_GetName() {
	call, err := vanstack.NewCall("call" /*the call name*/)
	if err != nil {
		panic(err) // Some error handling
	}

	fmt.Println(call.GetName())

	// Output:
	// call
}

func ExampleVanCall_GetDate() {
	call, err := vanstack.NewCall("call" /*the call name*/)
	if err != nil {
		panic(err) // Some error handling
	}

	fmt.Println(call.GetDate())
}

func ExampleVanCall_GetSettings() {
	call, err := vanstack.NewCall("call" /*the call name*/)
	if err != nil {
		panic(err) // Some error handling
	}

	fmt.Println(call.GetSettings())

	// Output:
	// {1 false}
}

func ExampleVanCall_SetSettings() {
	mySettings := vanstack.Settings{10, true}

	call, err := vanstack.NewCall("call" /*the call name*/)
	if err != nil {
		panic(err) // Some error handling
	}

	fmt.Println(call.GetSettings())

	call.SetSettings(mySettings)

	fmt.Println(call.GetSettings())

	// Output:
	// {1 false}
	// {10 true}

}

func ExampleVanCall_String() {
	call, err := vanstack.NewCall("call" /*the call name*/)
	if err != nil {
		panic(err) // Some error handling
	}

	fmt.Println(call.String())

	// Output:
	// example_test.go:89
}

func ExampleVanCall_SetShowName() {
	call, err := vanstack.NewCall("call" /*the call name*/)
	if err != nil {
		panic(err) // Some error handling
	}

	call.SetShowName(true)

	fmt.Println(call.String())

	// Output:
	// call example_test.go:101
}

func ExampleNewCall() {
	call, err := vanstack.NewCall("call" /*the call name*/)
	if err != nil {
		panic(err) // Some error handling
	}

	fmt.Println(call)

	// Output:
	// example_test.go:115
}

func ExampleNewStack() {
	stack := vanstack.NewStack()

	fmt.Println(stack)

	// Output:
	//
}

func ExampleVanStack_Add() {
	stack := vanstack.NewStack()

	call, err := vanstack.NewCall("readme call")
	if err != nil {
		panic(err) // Some error handling
	}

	stack.Add(call)

	fmt.Println(stack)

	// Output:
	// example_test.go:138
}

func ExampleVanStack_Fill() {
	stack := vanstack.NewStack()

	stack.Fill("call", 2 /* two last function calls*/)

	fmt.Println(stack)

	// Output:
	// run_example.go:63
	// example_test.go:154
}

func ExampleVanStack_Period() {
	stack := vanstack.NewStack()

	stack.Fill("call", 1)

	time.Sleep(time.Second)

	stack.Fill("call", 1)

	fmt.Println(stack.Period().Round(time.Millisecond))

	// Output:
	// 1s
}

func ExampleVanStack_String() {
	stack := vanstack.NewStack()

	stack.Fill("call", 1)

	fmt.Println(stack)

	// Output:
	// example_test.go:182
}

func ExampleVanStack_Len() {
	stack := vanstack.NewStack()

	stack.Fill("call", 5)

	fmt.Println(stack.Len())

	// Output:
	// 5
}

func ExampleVanStack_SetSeparator() {
	stack := vanstack.NewStack()

	stack.Fill("call", 2)

	stack.Fill("call", 2)

	stack.SetSeparator(", ")

	fmt.Println(stack)

	// Output:
	// run_example.go:63, example_test.go:206, run_example.go:63, example_test.go:204
}

func ExampleVanStack_GetCalls() {
	stack := vanstack.NewStack()

	stack.Fill("call", 2)

	stack.Fill("call", 2)

	fmt.Println(stack.GetCalls())

	// Output:
	// [run_example.go:63 example_test.go:222 run_example.go:63 example_test.go:220]
}

func ExampleToStackError() {
	err := vanstack.ToStackError(vanerrors.NewSimple("error"))

	err.Stack.Fill("call", 2)

	fmt.Println(err)

	// Output:
	// error, stack: run_example.go:63
	// example_test.go:233
}

func ExampleStackError_Touch() {
	err := vanstack.ToStackError(vanerrors.NewSimple("error"))

	err.Touch("call")

	fmt.Println(err)

	// Output:
	// error, stack: example_test.go:245
}

func ExampleOutOfStack() {
	err := vanstack.ToStackError(vanerrors.NewSimple("error"))

	fmt.Println(vanstack.OutOfStack(err))

	// Output:
	// error
}

func ExampleOutOfError() {
	err := vanstack.ToStackError(vanerrors.NewSimple("error"))

	fmt.Println(vanstack.OutOfError(err))

	// Output:
	//
}

func ExampleTouch() {
	err := vanstack.ToStackError(vanerrors.NewSimple("error"))

	vanstack.Touch(err, "call")

	fmt.Println(err)

	// Output:
	// error, stack: example_test.go:274
}
