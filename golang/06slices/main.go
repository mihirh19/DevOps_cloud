package main

import "fmt"

func main() {
	
	numbers := []int{1, 2, 3, 4, 5}
	fmt.Println(numbers)
	
	fmt.Printf("Length: %d, Capacity: %d\n", len(numbers), cap(numbers))
	
	numbers = append(numbers, 6) // not added
	fmt.Println(numbers)
	
	numbers = make([]int, 3, 5)
	fmt.Println(numbers)
	
	fmt.Println("Length: %d, Capacity: %d\n", len(numbers), cap(numbers))
	numbers = append(numbers, 4)
	numbers = append(numbers, 5)
	numbers = append(numbers, 6)
	
	fmt.Println("Length: %d, Capacity: %d\n", len(numbers), cap(numbers))
	
}