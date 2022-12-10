package main

import "fmt"

type Shape interface {
	CalculateVolume(pi float64) float64
}

type Cylinder struct {
	r, h float64
}

type Cone struct {
	r, h float64
}

func (cylinder Cylinder) CalculateVolume(pi float64) float64 {
	return pi * cylinder.r * cylinder.r * cylinder.h
}

func (cone Cone) CalculateVolume(pi float64) float64 {
	return pi * cone.r * cone.r * cone.h / 3
}

func main() {
	myPi := 3.14
	myR := 10.
	myH := 6.
	var cylinder Shape = Cylinder{r: myR, h: myH}
	var cone Shape = Cone{r: myR, h: myH}

	fmt.Printf("volume of cylinder with R=%f and H=%f: %f\n", myR, myH, cylinder.CalculateVolume(myPi))
	fmt.Printf("volume of cone with R=%f and H=%f: %f", myR, myH, cone.CalculateVolume(myPi))
}
