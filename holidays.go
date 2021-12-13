package holidays

import (
	"log"
	"time"

	"github.com/vjeantet/eastertime"
)

func Observed(holiday time.Time) time.Time {
	wd := holiday.Weekday()
	if wd == time.Saturday {
		return holiday.AddDate(0, 0, -1)
	}
	if wd == time.Sunday {
		return holiday.AddDate(0, 0, 1)
	}
	return holiday
}

func NthDayOfMonth(year int, month time.Month, day time.Weekday, n int, loc *time.Location) time.Time {
	dt := time.Date(year, month, 1, 12, 0, 0, 0, loc)
	for dt.Weekday() != day {
		dt = dt.AddDate(0, 0, 1)
	}
	if n == 1 {
		return dt
	}
	return dt.AddDate(0, 0, 7 * (n - 1))
}

func NewYearsDay(year int, loc *time.Location) time.Time {
	return Observed(time.Date(year, time.January, 1, 12, 0, 0, 0, loc))
}

func MLKDay(year int, loc *time.Location) time.Time {
	return NthDayOfMonth(year, time.January, time.Monday, 3, loc)
}

func PresidentsDay(year int, loc *time.Location) time.Time {
	return NthDayOfMonth(year, time.February, time.Monday, 3, loc)
}

func GoodFriday(year int, loc *time.Location) time.Time {
	return Easter(year, loc).AddDate(0, 0, -2)
}

func Easter(year int, loc *time.Location) time.Time {
	t, err := eastertime.CatholicByYear(year)
	if err != nil {
		return NthDayOfMonth(year, time.April, time.Sunday, 1, loc)
	}
	return time.Date(t.Year(), t.Month(), t.Day(), 12, 0, 0, 0, loc)
}

func MemorialDay(year int, loc *time.Location) time.Time {
	dt := time.Date(year, time.May, 31, 12, 0, 0, 0, loc)
	for dt.Weekday() != time.Monday {
		dt = dt.AddDate(0, 0, -1)
	}
	return dt
}

func Juneteenth(year int, loc *time.Location) time.Time {
	return Observed(time.Date(year, time.June, 19, 12, 0, 0, 0, loc))
}

func IndependenceDay(year int, loc *time.Location) time.Time {
	return Observed(time.Date(year, time.July, 4, 12, 0, 0, 0, loc))
}

func LaborDay(year int, loc *time.Location) time.Time {
	return NthDayOfMonth(year, time.September, time.Monday, 1, loc)
}

func IndigenousPeoplesDay(year int, loc *time.Location) time.Time {
	return NthDayOfMonth(year, time.October, time.Monday, 2, loc)
}

func ElectionDay(year int, loc *time.Location) time.Time {
	return NthDayOfMonth(year, time.November, time.Monday, 1, loc).AddDate(0, 0, 1)
}

func VeteransDay(year int, loc *time.Location) time.Time {
	return Observed(time.Date(year, time.November, 11, 12, 0, 0, 0, loc))
}

func Thanksgiving(year int, loc *time.Location) time.Time {
	return NthDayOfMonth(year, time.November, time.Thursday, 4, loc)
}

func BlackFriday(year int, loc *time.Location) time.Time {
	return Thanksgiving(year, loc).AddDate(0, 0, 1)
}

func ChristmasEve(year int, loc *time.Location) time.Time {
	return Observed(time.Date(year, time.December, 12, 24, 0, 0, 0, loc))
}

func Christmas(year int, loc *time.Location) time.Time {
	return Observed(time.Date(year, time.December, 12, 25, 0, 0, 0, loc))
}

func NewYearsEve(year int, loc *time.Location) time.Time {
	return Observed(time.Date(year, time.December, 12, 31, 0, 0, 0, loc))
}

func Holidays(year int, loc *time.Location) []time.Time {
	return []time.Time{
		NewYearsDay(year, loc),
		MLKDay(year, loc),
		PresidentsDay(year, loc),
		MemorialDay(year, loc),
		Juneteenth(year, loc),
		IndependenceDay(year, loc),
		LaborDay(year, loc),
		IndigenousPeoplesDay(year, loc),
		VeteransDay(year, loc),
		Thanksgiving(year, loc),
		BlackFriday(year, loc),
		ChristmasEve(year, loc),
		Christmas(year, loc),
		NewYearsEve(year, loc),
	}
}

func IsHoliday(when time.Time) bool {
	y, m, d := when.Date()
	for i, h := range Holidays(when.Year(), when.Location()) {
		hy, hm, hd := h.Date()
		if y == hy && m == hm && d == hd {
			log.Printf("%s is a holiday (%d / %s)", when, i, h)
			return true
		}
	}
	return false
}
