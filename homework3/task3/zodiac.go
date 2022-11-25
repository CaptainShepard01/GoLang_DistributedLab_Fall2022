package task3

func DefineZodiac(day, month int) string {
	switch month {
	case 1:
		if day < 20 {
			return "Козеріг"
		} else {
			return "Водолій"
		}
	case 2:
		if day < 19 {
			return "Водолій"
		} else {
			return "Риби"
		}
	case 3:
		if day < 21 {
			return "Риби"
		} else {
			return "Овен"
		}
	case 4:
		if day < 20 {
			return "Овен"
		} else {
			return "Телець"
		}
	case 5:
		if day < 21 {
			return "Телець"
		} else {
			return "Близнюки"
		}
	case 6:
		if day < 22 {
			return "Близнюки"
		} else {
			return "Рак"
		}
	case 7:
		if day < 23 {
			return "Рак"
		} else {
			return "Лев"
		}
	case 8:
		if day < 23 {
			return "Лев"
		} else {
			return "Діва"
		}
	case 9:
		if day < 23 {
			return "Діва"
		} else {
			return "Терези"
		}
	case 10:
		if day < 23 {
			return "Терези"
		} else {
			return "Скорпіон"
		}
	case 11:
		if day < 23 {
			return "Скорпіон"
		} else {
			return "Стрілець"
		}
	case 12:
		if day < 22 {
			return "Стрілець"
		} else {
			return "Козеріг"
		}
	default:
		return "Unknown zodiac"
	}
}
