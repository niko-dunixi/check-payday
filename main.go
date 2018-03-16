package main

import (
	"github.com/aodin/date"
	"os"
	"fmt"
	"github.com/paul-nelson-baker/check-payday/util"
)

func main() {
	firstPayDate := date.New(2018, 1, 5)
	// We don't include the end date, it incorrectly offsets the calculation
	// See this for example (be sure to update the second date):
	// https://www.timeanddate.com/date/durationresult.html?m1=1&d1=5&y1=2018&m2=3&d2=15&y2=2018
	daysUntilPayDay, isPayDay := util.CalculatePayDayInfo(firstPayDate, date.Today())
	if !isPayDay {
		var suffix string
		if daysUntilPayDay > 1 {
			suffix = "s"
		} else {
			suffix = ""
		}
		fmt.Printf("%d day%s until we get paid.\n", daysUntilPayDay, suffix)
	} else {
		fmt.Println("0 days until we get paid. IT'S PAYDAY!! :D")
	}
	// Only returns success (0) when it is payday
	// Useful for shell scripts and automation that
	// should be run only on this day.
	os.Exit(daysUntilPayDay)
}
