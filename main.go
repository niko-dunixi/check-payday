package main

import (
	"github.com/aodin/date"
	"os"
	"fmt"
	"github.com/paul-nelson-baker/check-payday/util"
)

func main() {
	firstPayDate := date.New(2018, 1, 5)
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
