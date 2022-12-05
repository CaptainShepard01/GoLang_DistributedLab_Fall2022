package main

import (
	"fmt"
	"math"
)

func reverseNumber(number int) int {
	tempInt := 0.
	for number > 0 {
		remainder := math.Mod(float64(number), 10.)
		tempInt *= 10
		tempInt += remainder
		number /= 10
	}

	return int(tempInt)
}

func main() {
	number := 12345

	fmt.Printf("reverse of %d: %d", number, reverseNumber(number))
}
