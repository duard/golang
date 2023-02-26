package main

import "fmt"

func main() {
	numbers := []int{1, 2}
	fmt.Println(Add(numbers))
}

func Add(numbers []int) int {
	sum := 0

	for _, number := range numbers {
		sum += number
	}
	return sum
}
