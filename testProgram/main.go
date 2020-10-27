package main

import "fmt"

func addNumbers(numbers []int) int {
	total := 0
	for _, value := range numbers {
		total = total + value
	}
	return total
}

func runCalc() {
	var numbers []int

	for {
		var input int
		_, err := fmt.Scan(&input)

		if err != nil{
			numberSum := addNumbers(numbers)
			fmt.Println("Sum of ", numbers, numberSum)
			break
		}

		numbers = append(numbers, input)
	}

}

func main() {
	runCalc()
}
