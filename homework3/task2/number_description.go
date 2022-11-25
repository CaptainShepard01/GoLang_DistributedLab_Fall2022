package task2

func countDigits(number int) int {
	count := 0
	for number != 0 {
		number /= 10
		count++
	}
	return count
}

func getParity(number int) string {
	resultString := "парне"
	if number%2 != 0 {
		return "не" + resultString
	}
	return resultString
}

func GetNumberDescription(number int) string {
	digitsCount := countDigits(number)
	parity := getParity(number)

	if digitsCount == 1 {
		return parity + " однозначне число"
	} else if digitsCount == 2 {
		return parity + " двозначне число"
	} else if digitsCount == 3 {
		return parity + " тризначне число"
	} else {
		return "Unknown number"
	}
}
