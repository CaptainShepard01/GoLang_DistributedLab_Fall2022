package main

import "fmt"

func swap(x, y *int) {
	*x, *y = *y, *x
}

func main() {
	x := 1
	y := 2
	fmt.Printf("before: x = %d; y = %d\n", x, y)
	swap(&x, &y)
	fmt.Printf("after: x = %d; y = %d", x, y)
}
