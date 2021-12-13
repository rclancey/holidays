package holidays

import (
	"testing"
	"time"
)

func TestObserved(t *testing.T) {
	loc, _ := time.LoadLocation("America/Los_Angeles")
	when := time.Date(2021, time.December, 25, 12, 0, 0, 0, loc)
	y, m, d := Observed(when).Date()
	if y != 2021 || m != 12 || d != 24 {
		t.Errorf("expected (2021, 12, 24), got (%d, %d, %d)", y, m, d)
	}
}

func TestNthDayOfMonth(t *testing.T) {
	loc, _ := time.LoadLocation("America/Los_Angeles")
	when := NthDayOfMonth(2021, time.November, time.Thursday, 4, loc)
	exp := time.Date(2021, time.November, 25, 12, 0, 0, 0, loc)
	if !when.Equal(exp) {
		t.Errorf("expected %s, got %s", exp, when)
	}
}

func TestHolidays(t *testing.T) {
	loc, _ := time.LoadLocation("America/Los_Angeles")
	year := 2021
	exp := []time.Time{
		time.Date(year, time.January, 1, 12, 0, 0, 0, loc),
		time.Date(year, time.January, 18, 12, 0, 0, 0, loc),
		time.Date(year, time.February, 15, 12, 0, 0, 0, loc),
		time.Date(year, time.May, 31, 12, 0, 0, 0, loc),
		time.Date(year, time.June, 18, 12, 0, 0, 0, loc),
		time.Date(year, time.July, 5, 12, 0, 0, 0, loc),
		time.Date(year, time.September, 6, 12, 0, 0, 0, loc),
		time.Date(year, time.October, 11, 12, 0, 0, 0, loc),
		time.Date(year, time.November, 11, 12, 0, 0, 0, loc),
		time.Date(year, time.November, 25, 12, 0, 0, 0, loc),
		time.Date(year, time.November, 26, 12, 0, 0, 0, loc),
		time.Date(year, time.December, 24, 12, 0, 0, 0, loc),
		time.Date(year, time.December, 24, 12, 0, 0, 0, loc),
		time.Date(year, time.December, 31, 12, 0, 0, 0, loc),
	}
	act := Holidays(year, loc)
	if len(act) != len(exp) {
		t.Errorf("expected %d holidays, got %d", len(exp), len(act))
		return
	}
	for i, h := range act {
		if !h.Equal(exp[i]) {
			t.Errorf("holiday %d: expected %s, got %s", i, exp[i], h)
		}
	}
}
