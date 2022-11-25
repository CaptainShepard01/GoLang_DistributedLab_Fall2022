package main

import (
	"GoLang_DistributedLab_Fall2022/homework2/task3"
	"fmt"
)

func main() {
	var a int
	fmt.Scan(&a)

	reverseNumber := task3.ReverseNumber(a)
	fmt.Printf("reverse of the %d is %d", a, reverseNumber)
}
