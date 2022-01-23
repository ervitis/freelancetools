package main

import (
	"context"
	"github.com/ervitis/freelancetools/credentials"
	"google.golang.org/api/calendar/v3"
	"google.golang.org/api/option"
	"log"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	credentialManagerClient := credentials.New()
	if err := credentialManagerClient.SetConfigWithScopes(calendar.CalendarReadonlyScope); err != nil {
		log.Fatalln(err)
	}

	srv, err := calendar.NewService(ctx, option.WithHTTPClient(credentialManagerClient.GetClient()))
	if err != nil {
		log.Fatalf("Unable to retrieve Calendar client: %v", err)
	}

	loc, err := time.LoadLocation("Europe/Madrid")
	if err != nil {
		log.Fatalln(err)
	}

	now := time.Now().In(loc)
	firstDayMonth := time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, loc)
	lastDayMonth := firstDayMonth.AddDate(0, 1, -1)

	events, err := srv.Events.List("primary").ShowDeleted(true).SingleEvents(true).Q("Work hours").ShowDeleted(false).TimeMin(firstDayMonth.Format(time.RFC3339)).TimeMax(lastDayMonth.Format(time.RFC3339)).Do()
	if err != nil {
		log.Fatalf("Unable to retrieve next ten of the user's events: %v", err)
	}
	log.Println("Upcoming events:")
	if len(events.Items) == 0 {
		log.Println("No upcoming events found.")
		return
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
	log.Printf("Total hours in %s month were %.2f\n", now.Month(), totalHours)
}
