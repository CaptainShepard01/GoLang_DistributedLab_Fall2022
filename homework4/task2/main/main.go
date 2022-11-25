package main

import (
	"fmt"
)

func main() {
	sum := 0
	maxNumber := 1200000

	for i := 1; i <= maxNumber; i++ {
		if i%2 == 0 || i%3 == 0 {
			sum += i
		}
	}

	fmt.Printf("sum of numbers from 1 to 1 200 000, that are divided by 2 or 3: %d", sum)
}
