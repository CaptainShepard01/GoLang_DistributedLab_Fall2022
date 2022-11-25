package main

import (
	"fmt"
	"math"
)

type Point struct {
	X int
	Y int
}

func calculateDistance(firstPoint, secondPoint Point) float64 {
	return math.Sqrt(math.Pow(float64(firstPoint.X-secondPoint.X), 2) + math.Pow(float64(firstPoint.Y-secondPoint.Y), 2))
}

func calculateTriangleSquare(points []Point) float64 {
	a := calculateDistance(points[0], points[1])
	b := calculateDistance(points[1], points[2])
	c := calculateDistance(points[2], points[0])
	p := (a + b + c) / 2

	return math.Sqrt(p * (p - a) * (p - b) * (p - c))
}

func main() {
	var points []Point

	points = append(points, Point{2, -4})
	points = append(points, Point{-5, -6})
	points = append(points, Point{1, 3})

	fmt.Printf("square of the triangle %v: %f", points, calculateTriangleSquare(points))
}
