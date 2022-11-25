package main

import "fmt"

func SumRange(a, b int) int {
	if a > b {
		return 0
	}

	sum := 0
	for i := a; i <= b; i++ {
		sum += i
	}

	return sum
}

func main() {
	fmt.Printf("sum from 10 to 20 inclusively: %d", SumRange(10, 20))
}
