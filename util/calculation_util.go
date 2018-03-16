package util

import (
	"github.com/aodin/date"
)

func GetDaysBetweenDates(start, end date.Date) int {
	return date.NewRange(start, end).Days() - 1
}

func CalculatePayDayInfo(firstPayDay, checkedDate date.Date) (int, bool) {
	days := GetDaysBetweenDates(firstPayDay, checkedDate)
	var booleanMultiplier int
	if days%14 == 0 {
		booleanMultiplier = 0
	} else {
		booleanMultiplier = 1
	}
	daysUntilPayday := (14 - days%14) * booleanMultiplier
	return daysUntilPayday, daysUntilPayday == 0
}
