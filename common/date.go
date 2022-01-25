package common

import (
	"log"
	"time"
)

type (
	date struct{}

	IDate interface {
		GetFirstDayAndLastDayCurrentMonth() (time.Time, time.Time)
	}
)

func NewDateTool() IDate {
	return &date{}
}

// GetFirstDayAndLastDayCurrentMonth returns first day of month and last day of current month
func (d *date) GetFirstDayAndLastDayCurrentMonth() (time.Time, time.Time) {
	loc, err := time.LoadLocation("Europe/Madrid")
	if err != nil {
		log.Fatalln(err)
	}

	now := time.Now().In(loc)
	firstDayMonth := time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, loc)
	return firstDayMonth, firstDayMonth.AddDate(0, 1, -1)
}
