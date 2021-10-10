package models

import (
	"time"
)

var (
	ErrNoStartDate   = ValidationErr("Start date not found in query")
	ErrEndDateBefore = ValidationErr("End date before start date")
	ErrDateInFuture  = ValidationErr("Cannot search in the future")
)

type Date struct {
	From *time.Time `form:"from" time_format:"2006-01-02" time_utc:"1"`
	To   *time.Time `form:"to" time_format:"2006-01-02" time_utc:"1"`
}
type ValidationErr string

func (e ValidationErr) Error() string {
	return string(e)
}

type Url struct {
	Urls []string
}

// ValidateDate validates if dates are correct
// If date is empty
// If date is in the future
// If end date is before start date
func (d *Date) ValidateDate() error {
	if d.From == nil {
		return ErrNoStartDate
	}
	if d.From.After(time.Now()) {
		return ErrDateInFuture
	}
	if d.To != nil {
		if d.To.After(time.Now()) {
			return ErrDateInFuture
		}
		if d.From.After(*d.To) {
			return ErrEndDateBefore
		}
	} else {
		today := time.Now()
		d.To = &today
	}
	return nil
}
