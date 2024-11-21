# VanStack

## What is vanstack?

A simple package for working stack traces in errors
It has:

- Calls
- Stack
- Filling the stack
- Error with stack

## How to install the package

```bash
go get github.com/VandiKond/vanerrors/vanstack
```

## How to use?

Lets start with the call

- What is a call?

It is an interface for calls in the stack

You can create your own call type, but I will show on VanCall

### Call

```go
call, err := vanstack.NewCall("readme call" /*the call name*/, )
if err != nil {
    panic(err) // Some error handling 
}

fmt.Println(call.GetPath()) // README.md github.com/VandiKond/vanerrors/vanstack: 37 (line)
fmt.Println(call.GetName()) // readme call
fmt.Println(call.GetDate()) // time.Now()
```

You can create a stack of calls

### Stack

```go
stack := vanstack.NewStack()

call, err := vanstack.NewCall("readme call")
if err != nil {
    panic(err) // Some error handling 
}

stack.Add(call)

fmt.Println(len(stack)) // 1
```

What about filling the stack with the last function calls

### Fill Stack

```go
stack := vanstack.NewStack()

stack.Fill("readme call", 2 /* two last function calls*/)

fmt.Println(len(stack)) // 2
```

### More Stack methods

```go
stack := vanstack.NewStack()

// Gets the deference between the oldest and earliest calls
fmt.Println(stack.Period()) // 0 because the stack hasn't got any calls

// Making a string of the stack
fmt.Println(stack.ToString) // "" because the stack is empty
```

So what about errors with a stack?

### Stack Error

```go
err := vanerrors.NewName("readme error")

stackErr := vanstack.ToStackError(err)

fmt.Println(stackError.Stack.ToString()) // "" // because the stack is empty

// You can add a new call 
stackErr.Touch("readme call") 

// Or you can touch any error that is touchable
vanstack.Touch(&stackErr, "readme call") // Don't forget the pointer, the compiler won't say that it is wrong with out the pointer, but the error wouldn't change if you don't use the pointer

// Getting the error out of the stack
vanstack.ErrorOutOfStack(stackErr) // it will be the same as err
```

## Other information

### License

[MIT](../LICENSE)
