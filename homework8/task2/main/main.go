package main

import "fmt"

func changeValue(x *int) {
	*x = (*x) * (*x) * (*x)
}

func main() {
	d := 2
	fmt.Println("d before:", d)
	changeValue(&d)
	fmt.Println("d after:", d)
}
