package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	err := os.Mkdir("homework12/Test", 0o750)
	if err != nil {
		log.Fatal(err)
	}

	file, err := os.Create("homework12/Test/test.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	_, err = file.WriteString("Я люблю програмування\n")
	if err != nil {
		log.Fatal(err)
	}
	_, err = file.WriteString("Мова програмування Golang\n")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Операція виконана успішно")
}
