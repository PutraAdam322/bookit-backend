package helper

import "time"

func thisWeekendRange() (satStart time.Time, sunEnd time.Time) {
	loc, err := time.LoadLocation("Asia/Kuala_Lumpur")
	if err != nil {
		loc = time.Local
	}

	now := time.Now().In(loc)
	todayStart := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, loc)
	weekday := now.Weekday()
	daysToSat := (int(time.Saturday) - int(weekday) + 7) % 7
	satStart = todayStart.AddDate(0, 0, daysToSat) // Saturday 00:00:00
	sun := satStart.AddDate(0, 0, 1)               // Sunday 00:00:00
	sunEnd = time.Date(sun.Year(), sun.Month(), sun.Day(), 23, 59, 59, 999999999, loc)
	return satStart, sunEnd
}
