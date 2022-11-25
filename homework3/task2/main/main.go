package main

import (
	"GoLang_DistributedLab_Fall2022/homework3/task2"
	"fmt"
)

func main() {
	var number int
	fmt.Scan(&number)

	result := task2.GetNumberDescription(number)
	fmt.Printf("number: %d\nnumber is %s", number, result)
}
