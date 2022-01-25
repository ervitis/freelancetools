package main

import (
	"context"
	"fmt"
	"github.com/ervitis/freelancetools/credentials"
	"github.com/ervitis/freelancetools/workinghours"
	"google.golang.org/api/calendar/v3"
	"log"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	credentialManagerClient := credentials.New()
	if err := credentialManagerClient.SetConfigWithScopes(calendar.CalendarReadonlyScope); err != nil {
		log.Fatalln(err)
	}

	whours, err := workinghours.New(ctx, credentialManagerClient)
	if err != nil {
		panic(err)
	}
	fmt.Println(whours.GetWorkingHoursActualMonth())
}
