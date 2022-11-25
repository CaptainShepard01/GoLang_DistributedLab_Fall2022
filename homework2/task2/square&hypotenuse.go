package task2

import "math"

func Square(cathetus1, cathetus2 float64) float64 {
	return (cathetus1 * cathetus2) / 2
}

func Hypotenuse(cathetus1, cathetus2 float64) float64 {
	return math.Sqrt(math.Pow(cathetus1, 2) + math.Pow(cathetus2, 2))
}
