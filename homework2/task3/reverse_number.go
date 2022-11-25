package task3

func ReverseNumber(number int) int {
	hundreds := number / 100
	tens := number % 100 / 10
	units := number % 10

	return units*100 + tens*10 + hundreds
}
