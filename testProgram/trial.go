package main

import (
"fmt"
)

func main() {
	size := 0
	fmt.Print("Number of elements n=")
	fmt.Scanln(&size)
	fmt.Println("Enter the numbers")
	elements := make([]int, size)
	for i := 0; i < size; i++ {
		fmt.Scanln(&elements[i])
	}
	fmt.Println("Entered Array of elements:", elements)
	result := 0

	for i := 0; i < size; i++ {
		result += elements[i]

	}
	fmt.Println("Sum of elements of array:", result)
}
