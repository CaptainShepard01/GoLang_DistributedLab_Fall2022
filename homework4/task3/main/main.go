package main

import "fmt"

func main() {
	count := 0
	maxNumber := 1200000

	for i := 1; i <= maxNumber; i++ {
		if i%2 == 0 || i%5 == 0 {
			count++
		}
	}

	fmt.Printf("percentage of numbers from 1 to 1 200 000, that are divided by 2 or 5: %f%%",
		float64(count)/float64(maxNumber)*100)
}
