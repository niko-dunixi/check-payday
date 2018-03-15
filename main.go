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
	currentDayInPayPeriod := howManyDays % 14
	fmt.Printf("%d days until we get paid.", 14-currentDayInPayPeriod)
	if currentDayInPayPeriod == 0 {
		fmt.Print(" PAYDAY!! :D")
	}
	fmt.Println()
	// Only returns success (0) when it is payday
	// Useful for shell scripts and automation that
	// should be run only on this day.
	os.Exit(currentDayInPayPeriod)
}
