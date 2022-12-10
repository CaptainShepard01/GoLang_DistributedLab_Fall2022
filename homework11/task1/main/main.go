package main

import "fmt"

type Shape interface {
	CalculateVolume() float64
}

type Prism struct {
	s, h float64
}

type Pyramid struct {
	s, h float64
}

func (prism Prism) CalculateVolume() float64 {
	return prism.s * prism.h
}

func (pyramid Pyramid) CalculateVolume() float64 {
	return pyramid.s * pyramid.h / 3.
}

func main() {
	myS := 12.
	myH := 6.
	var prism Shape = Prism{s: myS, h: myH}
	var pyramid Shape = Pyramid{s: myS, h: myH}

	fmt.Printf("volume of prism with S=%f and H=%f: %f\n", myS, myH, prism.CalculateVolume())
	fmt.Printf("volume of pyramid with S=%f and H=%f: %f", myS, myH, pyramid.CalculateVolume())
}
