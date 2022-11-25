package main

import "fmt"

func main() {
	myArray := [22]int{4, 3, 1, 5, 3, 14, 3, 4, 7, 10, 11, 12, 4, 3, 3, 56, 3, 12, 122, 3, 2, 3}

	minIndex, maxIndex := 0, 0
	minElement, maxElement := myArray[0], myArray[0]

	for i, element := range myArray {
		if element < minElement {
			minIndex = i
			minElement = element
		}
		if element > maxElement {
			maxIndex = i
			maxElement = element
		}
	}

	myArray[minIndex] = maxElement
	myArray[maxIndex] = minElement

	fmt.Printf("updated array: %v", myArray)
}
