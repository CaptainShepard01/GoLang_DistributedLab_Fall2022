package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	err := os.Mkdir("homework12/Test2", 0o750)
	if err != nil {
		log.Fatal(err)
	}

	sourceFile, err := os.Open("homework12/Test/test.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer sourceFile.Close()

	newFile, err := os.Create("homework12/Test2/test.txt")
	if err != nil {
		log.Fatal(err)
	}

	_, err = io.Copy(newFile, sourceFile)
	if err != nil {
		log.Fatal(err)
	}
	newFile.Close()

	err = os.Rename("homework12/Test2/test.txt", "homework12/Test2/test2.txt")
	if err != nil {
		log.Fatal(err)
	}

	test3, err := os.Create("homework12/Test2/test3.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer test3.Close()
	test4, err := os.Create("homework12/Test2/test4.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer test4.Close()

	dir, err := os.Open("homework12/Test2")
	if err != nil {
		log.Fatal(err)
	}
	defer dir.Close()
	files, err := dir.Readdir(0)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		fmt.Println(file.Name())
	}
}
