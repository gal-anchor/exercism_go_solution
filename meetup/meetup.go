package meetup

import "time"

// WeekSchedule represents which occurrence of a weekday in a month.
type WeekSchedule int

const (
	First WeekSchedule = iota
	Second
	Third
	Fourth
	Last
	Teenth
)

// Day returns the day of the month for a given weekday and schedule.
func Day(wSched WeekSchedule, wDay time.Weekday, month time.Month, year int) int {
	// Start from the first day of the month
	t := time.Date(year, month, 1, 0, 0, 0, 0, time.UTC)

	// Collect all days in the month that match the given weekday
	var days []int
	for t.Month() == month {
		if t.Weekday() == wDay {
			days = append(days, t.Day())
		}
		t = t.AddDate(0, 0, 1) // move to next day
	}

	switch wSched {
	case First:
		return days[0]
	case Second:
		return days[1]
	case Third:
		return days[2]
	case Fourth:
		return days[3]
	case Last:
		return days[len(days)-1]
	case Teenth:
		// find the first day in the 13–19 range
		for _, d := range days {
			if d >= 13 && d <= 19 {
				return d
			}
		}
	}

	panic("Invalid schedule")
}
