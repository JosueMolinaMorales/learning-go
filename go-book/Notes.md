# Notes From The Book

## Chapter 1
* Go code is organized into packages
* `Package main` is special. It defines a standalone executable program, not a library.
* Within `Package main` is the function main. The entrypoint of the program
* The symbol `:=` is a short variable declaration. A statement that declares one or more variabels and gives them a type based on the value

### For Loop
A for loop has the following syntax

```
for initalization; condition; post {
    // Zero or more statements
}
```

### Named Types
A type declaration makes it possible to give a name to an existing type.

```go
type Point struct {
    X, Y int
}
var p Point
```

### Pointers
Go provides pointers, values that contain the address of a variable

Pointers are explicitly visible. The `&` operator yields  the address of a variable

The `*` operator retrieves the variable that the pointer refers to, but there is no pointer arithmetic

### Method and Interfaces
A method is a function associated with a named type

Interfaces are abstract types that let us treat different concrete types in the same way based on what methods they have, not how they are represented or implemented

## Chapter 2

### Go Types
Constants: true, false, iota, nil

Types: int int8 int16 int32 int64, uint..., uintptr, float32, float64, complex128, complex64, bool, byte, rune, string, error

Functions: make len cap new append copy close delete complex real imag panic recover

### The New Function
One way to create a variable is to use the built-in function `new`. The expression `new(T)` creates an *unnamed variable* of type T.

It initalizes it to the zero valyue of T, and returns its address, which is a value of type `*T`

```go
p := new(int)   // p, of type *int, points to an unamed int variable
fmt.Println(p)  // "0"
*p = 2          // sets the unnamed int to 2
fmt.Println(*p) // "2"
```

### Introduction to Goroutines
A *goroutine* is a concurrent function execution.

A *channel* is a communication mechanism that allows one goroutine to pass values of a specified type to another goroutine.

The function `main` runs in a goroutine and the `go` statement creates additional goroutines.

```go
func main() {
    ch := make(chan string)
}
```

The main function creates a channel of string using `make` 
