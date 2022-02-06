package common

import (
	"reflect"
	"testing"
	"time"

	"github.com/tenntenn/testtime"
)

func Test_date_GetNowSpainTime(t *testing.T) {
	loc, _ := time.LoadLocation("Europe/Madrid")
	stubTime := time.Date(2022, time.January, 10, 12, 23, 0, 0, loc)

	testtime.SetTime(t, stubTime)

	d := NewDateTool()
	got := d.GetNowSpainTime()
	if !stubTime.Equal(got) {
		t.Errorf("want %v, got %v", stubTime, got)
	}
}

func Test_date_GetFirstDayAndLastDayCurrentMonth(t *testing.T) {
	loc, _ := time.LoadLocation("Europe/Madrid")

	type fields struct {
		loc *time.Location
	}
	tests := []struct {
		name     string
		fields   fields
		stubTime time.Time
		firstDay time.Time
		lastDay  time.Time
	}{
		{
			name:     "get first and last days of February",
			fields:   fields{loc: loc},
			stubTime: time.Date(2022, time.February, 10, 12, 23, 0, 0, loc),
			firstDay: time.Date(2022, time.February, 1, 0, 0, 0, 0, loc),
			lastDay:  time.Date(2022, time.February, 28, 0, 0, 0, 0, loc),
		},
		{
			name:     "get first and last days of January",
			fields:   fields{loc: loc},
			stubTime: time.Date(2022, time.January, 10, 12, 23, 0, 0, loc),
			firstDay: time.Date(2022, time.January, 1, 0, 0, 0, 0, loc),
			lastDay:  time.Date(2022, time.January, 31, 0, 0, 0, 0, loc),
		},
		{
			name:     "get first and last days of June",
			fields:   fields{loc: loc},
			stubTime: time.Date(2022, time.June, 10, 12, 23, 0, 0, loc),
			firstDay: time.Date(2022, time.June, 1, 0, 0, 0, 0, loc),
			lastDay:  time.Date(2022, time.June, 30, 0, 0, 0, 0, loc),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			testtime.SetTime(t, tt.stubTime)

			d := &date{
				loc: tt.fields.loc,
			}
			got, got1 := d.GetFirstDayAndLastDayCurrentMonth()
			if !reflect.DeepEqual(got, tt.firstDay) {
				t.Errorf("GetFirstDayAndLastDayCurrentMonth() got = %v, want %v", got, tt.firstDay)
			}
			if !reflect.DeepEqual(got1, tt.lastDay) {
				t.Errorf("GetFirstDayAndLastDayCurrentMonth() got1 = %v, want %v", got1, tt.lastDay)
			}
		})
	}
}

func Test_date_GetNextLastDayOfMonth(t *testing.T) {
	loc, _ := time.LoadLocation("Europe/Madrid")

	type fields struct {
		loc *time.Location
	}
	tests := []struct {
		name     string
		fields   fields
		want     time.Time
		stubTime time.Time
	}{
		{
			name:     "get last day of February",
			fields:   fields{loc: loc},
			stubTime: time.Date(2022, time.January, 10, 12, 23, 0, 0, loc),
			want:     time.Date(2022, time.February, 28, 0, 0, 0, 0, loc),
		},
		{
			name:     "get last day of May",
			fields:   fields{loc: loc},
			stubTime: time.Date(2022, time.April, 10, 12, 23, 0, 0, loc),
			want:     time.Date(2022, time.May, 31, 0, 0, 0, 0, loc),
		},
		{
			name:     "get last day of June",
			fields:   fields{loc: loc},
			stubTime: time.Date(2022, time.May, 10, 12, 23, 0, 0, loc),
			want:     time.Date(2022, time.June, 30, 0, 0, 0, 0, loc),
		},
		{
			name:     "get last day of January 2021",
			fields:   fields{loc: loc},
			stubTime: time.Date(2021, time.December, 10, 12, 23, 0, 0, loc),
			want:     time.Date(2022, time.January, 31, 0, 0, 0, 0, loc),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			testtime.SetTime(t, tt.stubTime)

			d := &date{
				loc: tt.fields.loc,
			}
			if got := d.GetNextLastDayOfMonth(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetNextLastDayOfMonth() = %v, want %v", got, tt.want)
			}
		})
	}
}
