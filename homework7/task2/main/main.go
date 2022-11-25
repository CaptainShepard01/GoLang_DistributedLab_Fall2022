package main

import "fmt"

func RingS(r1, r2 float64) float64 {
	if r1 > r2 {
		return 0
	}

	pi := 3.14
	return pi * (r2*r2 - r1*r1)
}

func main() {
	fmt.Printf("square of the ring between circles with r1=5, r2=10: %f", RingS(5, 10))
}
