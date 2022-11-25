package task1

func CheckWeight(weight, height float64) string {
	estimatedWeight := height - 100
	acceptableDifference := estimatedWeight * 0.1
	if weight < estimatedWeight-acceptableDifference {
		return "Треба поправитися"
	} else if weight > estimatedWeight+acceptableDifference {
		return "Треба схуднути"
	} else {
		return "Норма"
	}
}
