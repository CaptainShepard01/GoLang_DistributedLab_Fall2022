package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	fileStat, err := os.Stat("homework12/Test/test.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("File Name:", fileStat.Name())
	fmt.Println("Size:", fileStat.Size())
	fmt.Println("Permissions:", fileStat.Mode())
	fmt.Println("Last Modified:", fileStat.ModTime())
	fmt.Println("Is Directory: ", fileStat.IsDir())
}
