# 🌟 Go: Next Gen Language 🚀

## Why Golang? 🤔

- **Compiled**: Go tool can run files directly, without needing a virtual machine (VM).
- **Cross-Platform**: Go executables are different for each OS, making it versatile.

![Golang Logo](https://golang.org/lib/godoc/images/go-logo-blue.svg)

## What is Go? 🌐

- Go is used for everything from **system apps** to **web apps** and **cloud services**.
- **Widely Used**: It's mostly used in production environments for its efficiency and performance.

## OOP in Go 🧩

- **Yes and No**: Go doesn't have classes, but it has structures, making it a unique approach to Object-Oriented Programming.

## Missing Features 🚫

- No **try-catch** mechanism.
- The **lexer** does a lot of the heavy lifting behind the scenes.

## Types in Go 📚

- **Case Insensitive**: Almost entirely.
- **Static Typing**: Variable types should be known in advance.
- **Everything is a Type**: Go has a rich set of basic types.

  - `String`
  - `bool`
  - `integer`
  - `floating`
  - `complex`

```go
package main

import (
	"fmt"
)

const LoginToken string = "asjdhjashdjk"  // Public variable

func main() {
	var username string = "John Doe"
	fmt.Println("Hello, ", username)

	var isverified bool = true
	fmt.Println("Is verified: ", isverified)

	var age int = 30
	fmt.Println("Age: ", age)

	var height float32 = 1.75
	fmt.Println("Height: ", height)

	number := 100 // := is not allowed outside of a function
	fmt.Println("Number: ", number)

	fmt.Println("Login Token: ", LoginToken)
}
```

## User Input in Go 🎤

```go
welcome := "Welcome to GoLang"
fmt.Println(welcome)

var name string

fmt.Println("What is your name?")
fmt.Scan(&name)
fmt.Printf("Hello, %v\n", name)

```

Note: This is not the correct way to handle user input. Go has the bufio package for this purpose.

```go
reader := bufio.NewReader(os.Stdin)
fmt.Print("What is your name? ")
name, _ := reader.ReadString('\n')

fmt.Printf("Hello, %v\n", name)

```

A buffered reader is a type of reader that reads data from an underlying source, such as a file or standard input, and stores that data in a buffer. The buffer is a temporary storage area in memory.

## Functions in Go 🔧

```go
func simpleFunction() {
	fmt.Println("This is a simple function")
}
```

```go
func add(a int, b int) int {
	return a + b
}

```

Function with Named Return Variable 🎯
go

```go
func add(a int, b int) (result int) {
	result = a + b
	return
}
```
