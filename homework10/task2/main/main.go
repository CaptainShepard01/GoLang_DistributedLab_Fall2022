package main

import (
	"fmt"
	"math"
)

func findSquaredTriangleSide(ac, bc, acb float64) float64 {
	return math.Pow(ac, 2.) + math.Pow(bc, 2.) - 2.*ac*bc*math.Cos(acb)
}

func degrees2radians(degrees float64) float64 {
	return degrees * (math.Pi / 180)
}

func main() {
	ac := 22.
	bc := 21.
	angle := degrees2radians(60)

	fmt.Printf("ab^2 = %f", findSquaredTriangleSide(ac, bc, angle))
}
