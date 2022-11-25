package main

import (
	"fmt"
)

func main() {
	count := 1
	seconds := 19

	for i := 1; i <= seconds; i++ {
		count *= 2
	}

	fmt.Printf("number of bacteries after 19 seconds: %d", count)
}
