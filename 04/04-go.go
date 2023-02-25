package main

import (
	"fmt"
)

func main() {
	numbers := [5]int{1, 2, 3, 4, 5}
	fmt.Printf("Total : %d", Sum(numbers))
}

func Sum(numbers [5]int) int {
	sum := 0

	for i, num := range numbers {
		fmt.Println("index:", i, numbers[i], num)
		sum += num
	}

	return sum
}
