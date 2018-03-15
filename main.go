package main

import (
	"github.com/aodin/date"
	"os"
	"fmt"
)

func main() {
	firstPayDate := date.New(2018, 1, 5)
	currentDateRange := date.NewRange(firstPayDate, date.Today())
	// We don't include the end date, it incorrectly offsets the calculation
	// See this for example (be sure to update the second date):
	// https://www.timeanddate.com/date/durationresult.html?m1=1&d1=5&y1=2018&m2=3&d2=15&y2=2018
	howManyDays := currentDateRange.Days() - 1
	daysUntilPayDay := 14 - howManyDays%14
	fmt.Printf("%d day(s) until we get paid.", daysUntilPayDay)
	if daysUntilPayDay == 0 {
		fmt.Print(" IT'S PAYDAY!! :D")
	}
	fmt.Println()
	// Only returns success (0) when it is payday
	// Useful for shell scripts and automation that
	// should be run only on this day.
	os.Exit(daysUntilPayDay)
}
