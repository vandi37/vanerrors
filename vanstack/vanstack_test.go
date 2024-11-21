package vanstack_test

import (
	"testing"
	"time"

	"github.com/VandiKond/vanerrors"
	"github.com/VandiKond/vanerrors/vanstack"
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
}

func TestVanStack(t *testing.T) {
	stack := vanstack.NewStack()
	call, _ := vanstack.NewCall("call")
	stack.Add(call)
	if len(stack) != 1 {
		t.Fatalf("expected stack length 1, but got %d", len(stack))
	}
	if stack[0].GetName() != "call" {
		t.Fatalf("expected call name 'call', but got '%s'", stack[0].GetName())
	}
	if stack.Period() != 0 {
		t.Fatalf("expected stack period 0, but got %d", stack.Period())
	}
	stack.Fill("call", 2)
	if len(stack) != 3 {
		t.Fatalf("expected stack length 3, but got %d", len(stack))
	}
}

func TestStackError(t *testing.T) {
	err := vanerrors.NewName("error", nil)
	stackError := vanstack.ToStackError(err)

	if stackError.Error() != err.Error() {
		t.Fatalf("expected error %s, but got %s", err.Error(), stackError.Error())
	}
}

func TestOutStack(t *testing.T) {
	err := vanerrors.NewName("error", nil)
	stackErr := vanstack.ToStackError(err)
	if !err.Is(vanstack.ErrorOutOfStack(stackErr)) {
		t.Fatalf("error out of the stack error is not the original error")
	}

	if !err.Is(vanstack.ErrorOutOfStack(err)) {
		t.Fatalf("error out of the stack error is not it self")
	}
}

func TestTouch(t *testing.T) {
	err := vanerrors.NewName("error", nil)
	stackErr := vanstack.ToStackError(err)

	stackErr.Touch("call")
	if len(stackErr.Stack) != 1 {
		t.Fatalf("expected stack length 1, but got %d", len(stackErr.Stack))
	}

	vanstack.Touch(&stackErr, "call")
	if len(stackErr.Stack) != 2 {
		t.Fatalf("expected stack length 2, but got %d", len(stackErr.Stack))
	}
}
