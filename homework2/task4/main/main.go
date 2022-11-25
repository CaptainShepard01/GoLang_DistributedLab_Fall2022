package main

import (
	"GoLang_DistributedLab_Fall2022/homework2/task4"
	"fmt"
)

func main() {
	var a, b float64
	fmt.Scan(&a, &b)

	cooperativeTime := task4.PipesEfficiency(a, b)
	fmt.Printf("time of filling a pool when working together for pipes with times %f and %f is\n%f", a, b, cooperativeTime)
}
