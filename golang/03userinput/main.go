package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Welcome to GoLang"
	fmt.Println(welcome)
	
	// var name string
	
	// fmt.Println("What is your name?")
	// fmt.Scan(&name)
	// fmt.Printf("Hello, %v\n", name)
	
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("What is your name? ")
	name, _ := reader.ReadString('\n')
	
	fmt.Printf("Hello, %v\n", name)
	
}