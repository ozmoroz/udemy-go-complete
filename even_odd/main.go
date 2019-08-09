package main

import "fmt"

var numbers []int

func main() {
	// Create a slice of ints from 0 through 10
	for i := 0; i <= 10; i++ {
		numbers = append(numbers, i)
	}

	for _, number := range numbers {
		if number%2 == 0 {
			fmt.Printf("%d is an even number\n", number)
		} else {
			fmt.Printf("%d is an odd number\n", number)
		}
	}
}
