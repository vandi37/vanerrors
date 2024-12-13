package vanstack_test

import (
	"testing"
	"time"

	"github.com/vandi37/vanerrors"
	"github.com/vandi37/vanerrors/vanstack"
)

func TestVanCall(t *testing.T) {
	call, err := vanstack.NewCall("call")
	now := time.Now()
	if err != nil {
		t.Fatal(err)
	}
	if call.GetDate() != now {
		t.Fatalf("expected date %v but got %v", now, call.GetDate())
	}

	if call.GetName() != call.Name || call.GetName() != "call" {
		t.Fatalf("call name should be the same. expected 'call', but got '%s' and '%s'", call.Name, call.GetName())
	}
	path := "vanstack_test.go:24"
	if call.GetPath() != "vanstack_test.go:24" {
		t.Fatalf("call path should be the same. expected %s, but got '%s'", path, call.GetPath())
	}
}

func TestVanStack(t *testing.T) {
	stack := vanstack.NewStack()
	call, _ := vanstack.NewCall("call")
	stack.Add(call)
	if len(stack.GetCalls()) != 1 {
		t.Fatalf("expected stack length 1, but got %d", len(stack.GetCalls()))
	}
	if stack.GetCalls()[0].GetName() != "call" {
		t.Fatalf("expected call name 'call', but got '%s'", stack.GetCalls()[0].GetName())
	}
	if stack.Period() != 0 {
		t.Fatalf("expected stack period 0, but got %d", stack.Period())
	}
	stack.Fill("call", 2)
	if len(stack.GetCalls()) != 3 {
		t.Fatalf("expected stack length 3, but got %d", len(stack.GetCalls()))
	}
}

func TestStackError(t *testing.T) {
	err := vanerrors.NewName("error", vanerrors.EmptyHandler)
	stackError := vanstack.ToStackError(err)

	if stackError.Error() != err.Error() {
		t.Fatalf("expected error %s, but got %s", err.Error(), stackError.Error())
	}
}

func TestOutStack(t *testing.T) {
	err := vanerrors.NewName("error", vanerrors.EmptyHandler)
	stackErr := vanstack.ToStackError(err)
	if !err.Is(vanstack.OutOfStack(stackErr)) {
		t.Fatalf("error out of the stack error is not the original error")
	}

	if !err.Is(vanstack.OutOfStack(err)) {
		t.Fatalf("error out of the stack error is not it self")
	}
}

func TestTouch(t *testing.T) {
	err := vanerrors.NewName("error", vanerrors.EmptyHandler)
	stackErr := vanstack.ToStackError(err)

	stackErr.Touch("call")
	if len(stackErr.Stack.GetCalls()) != 1 {
		t.Fatalf("expected stack length 1, but got %d", len(stackErr.Stack.GetCalls()))
	}

	vanstack.Touch(&stackErr, "call")
	if len(stackErr.Stack.GetCalls()) != 2 {
		t.Fatalf("expected stack length 2, but got %d", len(stackErr.Stack.GetCalls()))
	}
}
