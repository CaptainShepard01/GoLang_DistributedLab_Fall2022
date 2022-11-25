package main

import "fmt"

func main() {
	myArray := [8]int{48, 96, 86, 68, 57, 82, 63, 70}
	mySlice := []int{}

	for _, element := range myArray {
		if element%2 == 0 {
			mySlice = append(mySlice, element)
		}
	}

	for i := 0; i < len(mySlice); i++ {
		for j := i; j < len(mySlice); j++ {
			if mySlice[i] > mySlice[j] {
				tmp := mySlice[i]
				mySlice[i] = mySlice[j]
				mySlice[j] = tmp
			}
		}
	}

	fmt.Printf("sorted slice of even elements: %v", mySlice)
}
