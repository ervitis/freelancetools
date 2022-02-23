package workinghours

import (
	"context"
	"fmt"
	"github.com/ervitis/freelancetools/common"
	"github.com/ervitis/freelancetools/credentials"
	"google.golang.org/api/calendar/v3"
	"google.golang.org/api/option"
	"log"
	"time"
)

type (
	workingHours struct {
		calendar *calendar.Service
	}

	WorkingData struct {
		Month      string
		TotalHours float64
	}

	IWorkingHours interface {
		GetWorkingHoursActualMonth() (*WorkingData, error)
	}
)

func New(ctx context.Context, credManager *credentials.Manager) (IWorkingHours, error) {
	srv, err := calendar.NewService(ctx, option.WithHTTPClient(credManager.GetClient()))
	if err != nil {
		return nil, fmt.Errorf("error creating client: %w", err)
	}
	return &workingHours{
		calendar: srv,
	}, nil
}

func (w *workingHours) GetWorkingHoursActualMonth() (*WorkingData, error) {
	dateSrv := common.NewDateTool()

	firstDayMonth, lastDayMonth := dateSrv.GetFirstDayAndLastDayCurrentMonth()

	events, err := w.calendar.
		Events.
		List("primary").
		ShowDeleted(false).
		SingleEvents(true).
		Q("Work hours").
		TimeMin(firstDayMonth.Format(time.RFC3339)).
		TimeMax(lastDayMonth.AddDate(0, 0, 1).Format(time.RFC3339)).
		Do()
	if err != nil {
		log.Fatalf("Unable to retrieve next ten of the user's events: %v", err)
	}

	if len(events.Items) == 0 {
		log.Println("No upcoming events found.")
		return &WorkingData{}, nil
	}

	totalHours := 0.0
	for _, item := range events.Items {
		tStart, err := time.Parse(time.RFC3339, item.Start.DateTime)
		if err != nil {
			panic(err)
		}
		tEnd, err := time.Parse(time.RFC3339, item.End.DateTime)
		if err != nil {
			panic(err)
		}
		totalHours += tEnd.Sub(tStart).Hours()
	}
	return &WorkingData{
		Month:      lastDayMonth.Month().String(),
		TotalHours: totalHours,
	}, nil
}
