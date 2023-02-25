package main

import "fmt"

func main() {
	sum := Add(1, 5)
	fmt.Printf("The result is '%d'", sum)
}

func Add(x, y int) int {
	return x + y
}
