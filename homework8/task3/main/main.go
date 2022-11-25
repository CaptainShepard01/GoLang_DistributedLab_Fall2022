package main

import "fmt"

func tripleValue(x *int) {
	*x *= 3
}

func main() {
	value := 2
	fmt.Println("value before: ", value)
	tripleValue(&value)
	fmt.Println("value after: ", value)
}
