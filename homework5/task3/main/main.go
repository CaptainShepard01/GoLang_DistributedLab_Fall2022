package main

import "fmt"

func main() {
	myArray := [22]int{4, 3, 1, 5, 3, 14, 3, 4, 7, 10, 11, 12, 4, 3, 3, 56, 3, 12, 122, 3, 2, 3}

	// Bubble sort
	for i := 0; i < len(myArray); i++ {
		for j := i; j < len(myArray); j++ {
			if myArray[i] > myArray[j] {
				tmp := myArray[i]
				myArray[i] = myArray[j]
				myArray[j] = tmp
			}
		}
	}

	numDiff := 1
	lastElement := myArray[0]
	for i := 1; i < len(myArray); i++ {
		if myArray[i] != lastElement {
			numDiff++
			lastElement = myArray[i]
		}
	}

	fmt.Printf("number of different elements: %d", numDiff)
}
