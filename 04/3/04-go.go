package main

import (
	"fmt"
	"os"
)

func main() {
	SumAll([]int{1, 2}, []int{0, 9})
	os.Exit(0)
}

func SumAll(numbersToSum ...[]int) []int {
	lengthOfNumbers := len(numbersToSum)
	sums := make([]int, lengthOfNumbers)

	for i, numbers := range numbersToSum {
		sums[i] = Sum(numbers)
	}

	fmt.Printf("numbersToSum %d , %d , %d \n", numbersToSum, lengthOfNumbers, sums)

	return sums
}

func Sum(numbers []int) int {
	sum := 0

	for _, number := range numbers {
		sum += number
	}

	return sum
}
