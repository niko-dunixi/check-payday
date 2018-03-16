package main

import (
	"github.com/aodin/date"
	"os"
	"fmt"
	"github.com/paul-nelson-baker/check-payday/util"
	"github.com/jessevdk/go-flags"
)

type options struct {
	DateString string `short:"d" long:"date" description:"A given date that is a payday, all calculations will be based on this date. Must be in format YYYY-MM-DD with leading 0's. (First payday of the year is advisable, but the most recent will suffice.)"`
}

func main() {
	firstPayDate := getDate()
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

func getDate() date.Date {
	options := options{}
	_, err := flags.Parse(&options)
	if err != nil {
		os.Exit(-1)
	}
	// If no one has specified a pay-day, it's probably me using the tool.
	// Set the default to my first payday of 2018. Deal wit' it, yo!
	if options.DateString == "" {
		options.DateString = "2018-01-05"
	}
	parsedDate, err := date.Parse(options.DateString)
	if err != nil {
		panic(err)
	}
	return parsedDate
}
