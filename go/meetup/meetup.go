package meetup

import "time"

// Define the WeekSchedule type here.
type WeekSchedule int

const (
	First  WeekSchedule = 0
	Second WeekSchedule = 7
	Third  WeekSchedule = 14
	Fourth WeekSchedule = 21
	Fifth  WeekSchedule = 28
	Last   WeekSchedule = iota
	Teenth
)

var incDays = map[WeekSchedule]WeekSchedule{
	First:  First,
	Second: Second,
	Third:  Third,
	Fourth: Fourth,
	Fifth:  Fifth,
	Last:   Last,
}

var Weekdays = map[string]time.Weekday{
	"Monday":    time.Monday,
	"Tuesday":   time.Tuesday,
	"Wednesday": time.Wednesday,
	"Thursday":  time.Thursday,
	"Friday":    time.Friday,
	"Saturday":  time.Saturday,
	"Sunday":    time.Sunday,
}

func numDaysInMonth(m time.Month, y int) int {
	return time.Date(y, m+1, 0, 0, 0, 0, 0, time.UTC).Day()
}

func Day(week WeekSchedule, dayOfWeek time.Weekday, month time.Month, year int) int {
	if week == Teenth {
		for d := 13; d < 20; d++ {
			genDay := time.Date(year, month, d, 0, 0, 0, 0, time.UTC)
			if genDay.Weekday() == Weekdays[dayOfWeek.String()] {
				return genDay.Day()
			}
		}
	} else {
		if _, ok := incDays[week]; ok {
			for d := 1; d < 8; d++ {
				numDays := numDaysInMonth(month, year)
				genDay := time.Date(year, month, d, 0, 0, 0, 0, time.UTC)
				if genDay.Weekday() == Weekdays[dayOfWeek.String()] {
					if week == Last {
						is5thTheFinalWeek := d+int(Fifth) <= numDays
						if is5thTheFinalWeek {
							return time.Date(year, month, d+int(Fifth), 0, 0, 0, 0, time.UTC).Day()
						} else {
							return time.Date(year, month, d+int(Fourth), 0, 0, 0, 0, time.UTC).Day()
						}
					} else {
						addDaysRemainsTheSameMonth := d+int(incDays[week]) <= numDays
						if addDaysRemainsTheSameMonth {
							return time.Date(year, month, d+int(incDays[week]), 0, 0, 0, 0, time.UTC).Day()
						}
					}
				}
			}
		}
	}
	return 0
}
