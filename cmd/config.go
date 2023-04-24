package cmd

func returnDaysInMont(month string) int {
	daysInMonth := map[string]int{
		"January":   31,
		"February":  28,
		"March":     31,
		"April":     30,
		"May":       31,
		"June":      30,
		"July":      31,
		"August":    31,
		"September": 30,
		"October":   31,
		"November":  30,
		"December":  31,
	}
	return daysInMonth[month]

}

func returnWeakDayNumber(weekDay string) int {
	weakDayNumber := map[string]int{
		"Monday":    1,
		"Thursday":  2,
		"Thearsday": 3,
		"Wendsdey":  4,
		"Friday":    5,
		"Suturday":  6,
		"Sunday":    7,
	}
	return weakDayNumber[weekDay]

}

func returnFirstWeakDayNumber(weekDay int, currentDay int) int {

	for currentDay <= 8 {
		currentDay -= 7
	}
	if currentDay <= 7 {
		return 1
	}
	weakDayNumb := 7 - weekDay - 1

	return weakDayNumb
}
