package common

import (
	"time"
)

type (
	date struct {
		loc *time.Location
	}

	IDate interface {
		GetFirstDayAndLastDayCurrentMonth() (time.Time, time.Time)
		GetNextLastDayOfMonth() time.Time
		GetNowSpainTime() time.Time
	}
)

func NewDateTool() IDate {
	l, _ := time.LoadLocation("Europe/Madrid")
	return &date{l}
}

// GetFirstDayAndLastDayCurrentMonth returns first day of month and last day of current month
func (d *date) GetFirstDayAndLastDayCurrentMonth() (time.Time, time.Time) {
	now := d.GetNowSpainTime()
	firstDayMonth := time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, d.loc)
	return firstDayMonth, firstDayMonth.AddDate(0, 1, -1)
}

func (d *date) GetNextLastDayOfMonth() time.Time {
	t := d.GetNowSpainTime()
	t = time.Date(t.Year(), t.Month(), 1, 0, 0, 0, 0, d.loc)
	return t.AddDate(0, 2, -1)
}

func (d *date) GetNowSpainTime() time.Time {
	return time.Now().In(d.loc)
}
