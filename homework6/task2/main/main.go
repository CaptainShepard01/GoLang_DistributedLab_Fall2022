package main

import "fmt"

func main() {
	myArray := [8]int{48, 96, 86, 68, 57, 82, 63, 70}
	mySlice := myArray[:5]

	maxElement := mySlice[0]
	for _, element := range mySlice {
		if element > maxElement {
			maxElement = element
		}
	}

	for i := 0; i < len(mySlice); i++ {
		mySlice[i] += maxElement
	}

	fmt.Printf("result slice: %v", mySlice)
}
