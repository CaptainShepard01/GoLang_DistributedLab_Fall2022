package task4

func PipesEfficiency(time1, time2 float64) float64 {
	return 1 / (1/time1 + 1/time2)
}
