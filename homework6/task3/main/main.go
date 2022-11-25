package main

import "fmt"

func main() {
	myArray := [8]int{48, 96, 86, 68, 57, 82, 63, 70}
	mySlice := myArray[len(myArray)-4:]
	length := len(mySlice)
	sum := mySlice[length-1]

	for i := 0; i < length-1; i++ {
		if mySlice[i] < mySlice[length-1] {
			mySlice[i] *= mySlice[i]
		}
		sum += mySlice[i]
	}

	fmt.Printf("sum of the result slice elements: %d", sum)
}
