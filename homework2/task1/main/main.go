package main

import (
	"GoLang_DistributedLab_Fall2022/homework2/task1"
	"fmt"
)

func main() {
	var a, b, c, d int
	fmt.Scan(&a, &b, &c, &d)

	sum := task1.Sum(a, b, c, d)
	fmt.Printf("%d + %d + %d + %d = %d", a, b, c, d, sum)
}
