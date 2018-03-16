package util

import (
	"github.com/aodin/date"
)

func GetDaysBetweenDates(start, end date.Date) int {
	// We don't include the end date, it incorrectly offsets the calculation
	// See this for example (be sure to update the second date):
	// https://www.timeanddate.com/date/durationresult.html?m1=1&d1=5&y1=2018&m2=3&d2=15&y2=2018
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
