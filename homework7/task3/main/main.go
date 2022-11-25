package main

import "fmt"

func DigitsCount(number int) int {
	result := 0
	for number > 0 {
		result++
		number /= 10
	}
	return result
}

func main() {
	fmt.Printf("number of digits of number 321: %d\n", DigitsCount(321))
	fmt.Printf("number of digits of number 1243: %d\n", DigitsCount(1243))
	fmt.Printf("number of digits of number 12543: %d", DigitsCount(12543))
}
