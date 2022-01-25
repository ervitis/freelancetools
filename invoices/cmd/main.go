package main

import (
	"context"
	"github.com/ervitis/freelancetools/credentials"
	"github.com/ervitis/freelancetools/invoices"
	"github.com/ervitis/freelancetools/workinghours"
	"google.golang.org/api/calendar/v3"
	"google.golang.org/api/drive/v3"
	"google.golang.org/api/sheets/v4"
	"log"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	credentialManagerClient := credentials.New()
	if err := credentialManagerClient.
		SetConfigWithScopes(
			sheets.SpreadsheetsScope,
			drive.DriveScope,
			calendar.CalendarEventsReadonlyScope,
		); err != nil {
		log.Fatalln(err)
	}

	invoiceService, err := invoices.New(ctx, credentialManagerClient)
	if err != nil {
		log.Fatalln(err)
	}

	workingHoursService, err := workinghours.New(ctx, credentialManagerClient)
	if err != nil {
		log.Fatalln(err)
	}

	data, err := workingHoursService.GetWorkingHoursActualMonth()
	if err != nil {
		log.Fatalln(err)
	}

	if err := invoiceService.CreateNewInvoice(*data); err != nil {
		log.Fatalln(err)
	}
}
