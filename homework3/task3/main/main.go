package main

import (
	"GoLang_DistributedLab_Fall2022/homework3/task3"
	"fmt"
)

func main() {
	var day, month int
	fmt.Scan(&day, &month)

	zodiac := task3.DefineZodiac(day, month)
	fmt.Printf("day: %d, month: %d\nzodiac sign is %s", day, month, zodiac)
}
