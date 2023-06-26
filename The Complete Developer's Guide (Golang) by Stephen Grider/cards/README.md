# Go: The Complete Developer's Guide (Golang)

[Go: The Complete Developer's Guide (Golang)](https://www.udemy.com/course/go-the-complete-developers-guide/)

## CLI

| Command      | Description                                                |
| ------------ | ---------------------------------------------------------- |
| `go build`   | Compile a bunch of go source code file                     |
| `go run`     | Compile and executes one or tow files                      |
| `go fmt`     | Formats all the code in each file in the current directory |
| `go install` | Compiles and "installs" a package                          |
| `go get`     | Downloads the raw source code of someone else's package    |
| `go test`    | Runs any tests associated with the current project         |

## Package

What is **package**? Go programs are organized into packages. A package is a collection of source files in the same directory that are compiled together. Functions, types, variables, and constants defined in one source file are visible to all other source files within the same package. [link](https://github.com/golang/tools/blob/master/gopls/doc/workspace.md)

**Type of packages**

- **Executable** - Executable commands must always use `package main`. Defines a page that can be compiled and then **executed**. **Must have a func called `main`**
- **Reusable** - Code used as helpers good place to put reusable logic. Defines a package that can be used as a dependency [link](https://www.codementor.io/@inanc/understanding-go-packages-cgsx1hzrx)

## Standard Library

<https://pkg.go.dev/std>

## Function

Go's function uses `func` keyword to declare functions.

```go
func main() {
  ...
}

// return function
func foo() string {
  return "hello"
}
```

## Basic Types

| Type      | Description                      |
| --------- | -------------------------------- |
| `bool`    | Boolean can be `true` or `false` |
| `string`  | String e.g. `"hello"`            |
| `int`     | Integer                          |
| `float64` | Decimal number                   |

## OO Approach

Go is not OOP language. It can define custom type to define properties to an object.

In Go you can define alias type by using `type` keyboard.

```go
// Create a type of `deck` which is a slice of strings
// equal to 
// cards := []string{"one", "two"}
type deck []string 
cards := deck{"one", "two"}
```
