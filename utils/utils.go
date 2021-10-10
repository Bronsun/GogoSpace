package utils

import (
	"errors"
	"time"
)

var (
	ErrEndBeforeStart = errors.New("End date before start")
	ErrDateInFuture   = errors.New("Date is in the future")
)

// GetDatesFromQuery saves all dates between from and to an array
func GetDatesFromQuery(start, end time.Time) ([]string, error) {
	var dateStringList []string
	if end.Before(start) {
		return nil, ErrEndBeforeStart
	}
	if start.After(time.Now()) || end.After(time.Now()) {
		return nil, ErrDateInFuture
	}
	for days := DaysBetween(start, end); ; {
		date := days()
		if date.IsZero() {
			break
		}
		dateStringList = append(dateStringList, date.Format("2006-01-02"))
	}
	return dateStringList, nil
}

// DaysBetween finds all dates between From and To query
func DaysBetween(start, end time.Time) func() time.Time {
	y, m, d := start.Date()
	start = time.Date(y, m, d, 0, 0, 0, 0, time.UTC)
	y, m, d = end.Date()
	end = time.Date(y, m, d, 0, 0, 0, 0, time.UTC)

	return func() time.Time {
		if start.After(end) {
			return time.Time{}
		}
		date := start
		start = start.AddDate(0, 0, 1)
		return date
	}
}
