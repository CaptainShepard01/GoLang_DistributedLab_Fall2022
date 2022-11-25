package main

import (
	"GoLang_DistributedLab_Fall2022/homework2/task2"
	"fmt"
)

func main() {
	var a, b float64
	fmt.Scan(&a, &b)

	square := task2.Square(a, b)
	fmt.Printf("square of a triangle with cathetuses %f and %f is %f\n", a, b, square)

	hypotenuse := task2.Hypotenuse(a, b)
	fmt.Printf("hypotenuse of a triangle with cathetuses %f and %f is %f", a, b, hypotenuse)
}
