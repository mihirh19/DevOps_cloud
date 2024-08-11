package main

import "fmt"

func main() {
	fmt.Println("Arrays in Golang")
	var name [5]string
	
	name[0] = "John"
	name[1] = "Doe"
	
	fmt.Println(name)
	
	var numbers = [5]int{1, 2, 3, 4, 5}
	
	fmt.Println(numbers)

}
