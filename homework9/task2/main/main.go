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

func calculateRectangleSquare(points []Point) float64 {
	return calculateDistance(points[0], points[1]) * calculateDistance(points[1], points[2])
}

func main() {
	var points []Point

	points = append(points, Point{1, 2})
	points = append(points, Point{1, 4})
	points = append(points, Point{7, 4})

	fmt.Printf("square of the rectangle with 3 points %v: %f", points, calculateRectangleSquare(points))
}
