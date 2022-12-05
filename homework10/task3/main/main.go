package main

import (
	"fmt"
	"strings"
)

func isCharInArray(char rune, charArray []rune) bool {
	for _, element := range charArray {
		if char == element {
			return true
		}
	}
	return false
}

func getDistinctLetters(someString string) []rune {
	var resultArray []rune
	for _, char := range someString {
		if char == ' ' {
			continue
		}
		if isCharInArray(char, resultArray) {
			continue
		}
		resultArray = append(resultArray, char)
	}
	return resultArray
}

func findIndex(char rune, charArray []rune) int {
	for i, element := range charArray {
		if char == element {
			return i
		}
	}
	return -1
}

func countLetters(someString string) ([]rune, []int) {
	charsArray := getDistinctLetters(someString)
	occurrencesArray := make([]int, len(charsArray))

	for _, char := range someString {
		index := findIndex(char, charsArray)
		if index == -1 {
			continue
		}
		occurrencesArray[index]++
	}

	return charsArray, occurrencesArray
}

func main() {
	myString := "Я люблю програмування на мові C++"
	myString = strings.Replace(myString, "C++", "Go", 1)
	myString = strings.ToUpper(myString)
	myString = strings.Repeat(myString, 10)
	charsArray, occurrencesArray := countLetters(myString)

	for i, char := range charsArray {
		fmt.Printf("letter %c, number of occurrences: %d\n", char, occurrencesArray[i])
	}
}
