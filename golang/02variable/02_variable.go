package main

import (
	"fmt"
)
const LoginToken string = "asjdhjashdjk"  //public variable
func main() {
	var username string = "John Doe"
	fmt.Println("Hello, ", username)
	
	var isverified bool = true
	fmt.Println("Is verified: ", isverified)
	
	var age int = 30
	fmt.Println("Age: ", age)
	
	var height float32 = 1.75
	fmt.Println("Height: ", height)
	
	
	number:= 100 // not allowd to use := outside of the function
	fmt.Println("Number: ", number)
	
	fmt.Println("Login Token: ", LoginToken)
}