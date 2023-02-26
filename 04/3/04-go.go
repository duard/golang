package main

func main() {
	//  numbersToSum := []int{2,3}, []int{4,5}

	// fmt.Fprintf(SumAll(numbersToSum))
}

func SumAll(numbersToSum ...[]int) []int {
	lengthOfNumbers := len(numbersToSum)
	sums := make([]int, lengthOfNumbers)
	// fmt.Printf("numbersToSum %d , %d , %d \n", numbersToSum, lengthOfNumbers, sums)

	for i, numbers := range numbersToSum {
		sums[i] = Sum(numbers)
	}

	return sums
}

func Sum(numbers []int) int {
	sum := 0

	for _, number := range numbers {
		sum += number
	}

	return sum
}
