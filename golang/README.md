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

## 🌀 Arrays in GoLang
 - Arrays in Go provide a simple and efficient way to work with a fixed collection of elements. Unlike slices, arrays have a fixed size and cannot be resized.
 
	```go
	package main

	import "fmt"

	func main() {
		fmt.Println("🌟 Arrays in Golang")

		// Define an array of strings with a fixed size of 5
		var name [5]string

		// Assign values to the array
		name[0] = "John"
		name[1] = "Doe"

		// Print the array
		fmt.Println(name)

		// Define and initialize an array of integers
		var numbers = [5]int{1, 2, 3, 4, 5}

		// Print the array of numbers
		fmt.Println(numbers)
	}

	```
	
## 🚀 Slices in GoLang
 - In Go, a slice is a flexible and dynamic data structure that provides a more powerful alternative to arrays. Slices can grow and shrink in size, making them more versatile for handling collections of data.
 
```go
package main

import "fmt"

func main() {
    // Define and initialize a slice of integers
    numbers := []int{1, 2, 3, 4, 5}
    fmt.Println("🔢 Numbers:", numbers)

    // Print the length and capacity of the slice
    fmt.Printf("📏 Length: %d, 🚀 Capacity: %d\n", len(numbers), cap(numbers))

    // Append a new element to the slice
    numbers = append(numbers, 6)
    fmt.Println("➕ After Append:", numbers)

    // Create a new slice with make
    numbers = make([]int, 3, 5)
    fmt.Println("🛠️ New Slice:", numbers)

    fmt.Printf("📏 Length: %d, 🚀 Capacity: %d\n", len(numbers), cap(numbers))

    // Append more elements to the slice
    numbers = append(numbers, 4)
    numbers = append(numbers, 5)
    numbers = append(numbers, 6)

    // Print the updated length and capacity
    fmt.Printf("📏 Length: %d, 🚀 Capacity: %d\n", len(numbers), cap(numbers))
}

```
#### 🔍 Output:
```sh
	🔢 Numbers: [1 2 3 4 5]
	📏 Length: 5, 🚀 Capacity: 5
	➕ After Append: [1 2 3 4 5 6]	
	🛠️ New Slice: [0 0 0]
	📏 Length: 3, 🚀 Capacity: 5
	📏 Length: 6, 🚀 Capacity: 6
```