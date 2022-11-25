package main

import (
	"GoLang_DistributedLab_Fall2022/homework3/task1"
	"fmt"
)

func main() {
	var weight, height float64
	fmt.Scan(&weight, &height)

	result := task1.CheckWeight(weight, height)
	fmt.Printf("weight: %f, height: %f\nresult is %s", weight, height, result)
}
