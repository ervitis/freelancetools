package common

import (
	"reflect"
	"testing"
	"time"
)

func TestGetFirstDayAndLastDayMonth(t *testing.T) {
	tests := []struct {
		name  string
		want  time.Time
		want1 time.Time
	}{
		{
			name:  "get dates success",
			want:  time.Time{},
			want1: time.Time{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := GetFirstDayAndLastDayCurrentMonth()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetFirstDayAndLastDayMonth() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("GetFirstDayAndLastDayMonth() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
