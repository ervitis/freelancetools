package main

import (
	"context"
	"fmt"
	"github.com/ervitis/freelancetools/credentials"
	"google.golang.org/api/drive/v3"
	"google.golang.org/api/option"
	"google.golang.org/api/sheets/v4"
	"log"
	"time"
)

const (
	folderId = "" // TODO set folder ID

	invoiceDateLayout = "02-01-2006"

	templateNameInvoice = "template"
)

var (
	invoiceName = "" // TODO set invoice name
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	credentialManagerClient := credentials.New()
	if err := credentialManagerClient.SetConfigWithScopes(sheets.SpreadsheetsScope, drive.DriveScope, drive.DriveMetadataScope, drive.DriveFileScope, drive.DriveAppdataScope); err != nil {
		log.Fatalln(err)
	}

	sheetsService, err := sheets.NewService(ctx, option.WithHTTPClient(credentialManagerClient.GetClient()))
	if err != nil {
		log.Fatalln(err)
	}

	driveService, err := drive.NewService(ctx, option.WithHTTPClient(credentialManagerClient.GetClient()))
	if err != nil {
		log.Fatalln(err)
	}

	listFiles := driveService.Files.List().IncludeItemsFromAllDrives(true).SupportsAllDrives(true)

	listBillingModel, err := listFiles.Q(fmt.Sprintf(`"%s" in parents and name contains "MODELO FACTURA" and trashed = false`, folderId)).Do()
	if err != nil {
		log.Fatalln(err)
	}

	if len(listBillingModel.Files) == 0 {
		log.Println("files not found")
		return
	}

	listInvoices, err := listFiles.Q(fmt.Sprintf(`"%s" in parents and not name contains "MODELO FACTURA" and trashed = false`, folderId)).Do()
	if err != nil {
		log.Fatalln(err)
	}

	billingModel := listBillingModel.Files[0]
	invoiceName = fmt.Sprintf(invoiceName, len(listInvoices.Files)+1, time.Now().Format(invoiceDateLayout))

	copiedFile, err := driveService.Files.Copy(billingModel.Id, &drive.File{
		MimeType: "application/vnd.google-apps.spreadsheet",
		Name:     invoiceName,
	}).Do()
	if err != nil {
		log.Fatalln(err)
	}

	invoice, err := sheetsService.Spreadsheets.Get(copiedFile.Id).Do()
	if err != nil {
		log.Fatalln(err)
	}

	valueRange := make([]*sheets.ValueRange, 0)
	row := map[string]interface{}{
		"H3":  fmt.Sprintf("%d", len(listInvoices.Files)+1),
		"H4":  time.Now().Format(invoiceDateLayout),
		"H5":  time.Now().Format(invoiceDateLayout), // TODO set last day of month
		"H8":  "",                                   // TODO company name
		"H9":  "",                                   // TODO company address
		"B15": 15.0,                                 // TODO quantity of hours
		"C15": "",                                   // TODO Description
		"E15": 1.0,                                  // TODO set price in yen
	}

	for k, v := range row {
		data := make([][]interface{}, 0)
		data2 := make([]interface{}, 0)
		data2 = append(data2, v)
		data = append(data, data2)
		vr := &sheets.ValueRange{
			Range:  fmt.Sprintf("%s!%s", templateNameInvoice, k),
			Values: data,
		}
		valueRange = append(valueRange, vr)
	}

	resp, err := sheetsService.Spreadsheets.Values.BatchUpdate(invoice.SpreadsheetId, &sheets.BatchUpdateValuesRequest{
		Data:             valueRange,
		ValueInputOption: "RAW",
	}).Do()
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(resp)
}
