package main

import (
	"testing"
	"github.com/paul-nelson-baker/check-payday/util"
	"github.com/aodin/date"
)

func TestDaysCalculation(t *testing.T) {
	firstPayDate := date.New(2018, 1, 5)
	if days := util.GetDaysBetweenDates(firstPayDate, firstPayDate.AddDays(1)); days != 1 {
		t.Errorf("Incorrect number of days: %d", days)
	}
	if days := util.GetDaysBetweenDates(firstPayDate, firstPayDate.AddDays(7)); days != 7 {
		t.Errorf("Incorrect number of days: %d", days)
	}
	if days := util.GetDaysBetweenDates(firstPayDate, firstPayDate.AddDays(14)); days != 14 {
		t.Errorf("Incorrect number of days: %d", days)
	}
}

func TestCalculationOfRemainingDays(t *testing.T) {
	firstPayDate := date.New(2018, 1, 5)
	if days, _ := util.CalculatePayDayInfo(firstPayDate, firstPayDate.AddDays(1)); days != 13 {
		t.Errorf("Incorrect number of days: %d", days)
	}
	if days, _ := util.CalculatePayDayInfo(firstPayDate, firstPayDate.AddDays(7)); days != 7 {
		t.Errorf("Incorrect number of days: %d", days)
	}
	if days, _ := util.CalculatePayDayInfo(firstPayDate, firstPayDate.AddDays(14)); days != 0 {
		t.Errorf("Incorrect number of days: %d", days)
	}
}

func TestCalculationOfIsPayday(t *testing.T) {
	firstPayDate := date.New(2018, 1, 5)
	if _, isPayday := util.CalculatePayDayInfo(firstPayDate, firstPayDate.AddDays(1)); isPayday {
		t.Errorf("Should not have been payday, but was calculated as such")
	}
	if _, isPayday := util.CalculatePayDayInfo(firstPayDate, firstPayDate.AddDays(7)); isPayday {
		t.Errorf("Should not have been payday, but was calculated as such")
	}
	if _, isPayday := util.CalculatePayDayInfo(firstPayDate, firstPayDate.AddDays(14)); !isPayday {
		t.Errorf("Should not have been payday, but was calculated as such")
	}
}
