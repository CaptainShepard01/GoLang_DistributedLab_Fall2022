package main

import "fmt"

type Shape interface {
	CalculateArea(pi float64) float64
}

type Circle struct {
	r float64
}

type Sphere struct {
	r float64
}

func (circle Circle) CalculateArea(pi float64) float64 {
	return pi * circle.r * circle.r
}

func (sphere Sphere) CalculateArea(pi float64) float64 {
	return 4 * pi * sphere.r * sphere.r
}

func main() {
	myPi := 3.14
	myR := 10.
	var circle Shape = Circle{r: myR}
	var sphere Shape = Sphere{r: myR}

	fmt.Printf("area of circle with R=%f: %f\n", myR, circle.CalculateArea(myPi))
	fmt.Printf("area of the surface of sphere with R=%f: %f", myR, sphere.CalculateArea(myPi))
}
