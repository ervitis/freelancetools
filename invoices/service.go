package invoices

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/ervitis/freelancetools/common"
	"github.com/ervitis/freelancetools/config"
	"github.com/ervitis/freelancetools/credentials"
	"github.com/ervitis/freelancetools/exchangerate"
	"github.com/ervitis/freelancetools/workinghours"
	"github.com/ervitis/gotransactions"
	"google.golang.org/api/drive/v3"
	"google.golang.org/api/option"
	"google.golang.org/api/sheets/v4"
	"io"
	"log"
	"os"
	"path/filepath"
	"strconv"
)

const (
	invoicesDataFileName = "invoices.json"
	invoiceDateLayout    = "02-01-2006"
	templateNameInvoice  = "template"
)

type (
	companyInvoice struct {
		Name    string `json:"name"`
		Address string `json:"address"`

		Description string  `json:"description"`
		UnitPrice   float64 `json:"unitPrice"`
		MoneySymbol string  `json:"moneySymbol"`
	}

	invoicesData struct {
		Name                  string           `json:"name"`
		SpreadSheetIDFromCopy string           `json:"spreadSheetIdFromCopy"`
		Companies             []companyInvoice `json:"companies"`
	}

	invoices struct {
		sheetService        *sheets.Service
		driveService        *drive.Service
		invoicesData        invoicesData
		exchangeRateService *exchangerate.ExchangeApi
	}

	IInvoices interface {
		CreateNewInvoice(workinghours.WorkingData) error
	}
)

func New(ctx context.Context, credManager *credentials.Manager) (IInvoices, error) {
	sheetsService, err := sheets.NewService(ctx, option.WithHTTPClient(credManager.GetClient()))
	if err != nil {
		return nil, fmt.Errorf("creating sheetsService: %w", err)
	}

	driveService, err := drive.NewService(ctx, option.WithHTTPClient(credManager.GetClient()))
	if err != nil {
		return nil, fmt.Errorf("creating driveService: %w", err)
	}

	f, err := os.Open(fmt.Sprintf("env%s%s", string(filepath.Separator), invoicesDataFileName))
	if err != nil {
		return nil, fmt.Errorf("opening file of invoices data: %w", err)
	}

	defer func() {
		_ = f.Close()
	}()

	b, _ := io.ReadAll(f)
	var d invoicesData
	_ = json.Unmarshal(b, &d)

	ex, err := exchangerate.NewClient(&config.AppConfig)
	if err != nil {
		return nil, fmt.Errorf("creating exchange client: %w", err)
	}

	return &invoices{
		sheetService:        sheetsService,
		driveService:        driveService,
		invoicesData:        d,
		exchangeRateService: ex,
	}, nil
}

func (i *invoices) CreateNewInvoice(workHoursData workinghours.WorkingData) error {
	listFiles := i.driveService.Files.List().IncludeItemsFromAllDrives(true).SupportsAllDrives(true)

	listBillingModel, err := listFiles.
		Q(fmt.Sprintf(`"%s" in parents and name contains "MODELO FACTURA" and trashed = false`, i.invoicesData.SpreadSheetIDFromCopy)).
		Do()
	if err != nil {
		return fmt.Errorf("get list billing model from copy: %w", err)
	}

	if len(listBillingModel.Files) == 0 {
		log.Println("files not found")
		return nil
	}

	listInvoices, err := listFiles.
		Q(fmt.Sprintf(`"%s" in parents and not name contains "MODELO FACTURA" and trashed = false`, i.invoicesData.SpreadSheetIDFromCopy)).
		Do()
	if err != nil {
		return fmt.Errorf("get list invoices: %w", err)
	}

	billingModel := listBillingModel.Files[0]

	dateSrv := common.NewDateTool()
	dayPayment := dateSrv.GetNextLastDayOfMonth()
	firstDayCurrentMonth, lastDayCurrentMonth := dateSrv.GetFirstDayAndLastDayCurrentMonth()

	for _, company := range i.invoicesData.Companies {
		copiedFile, err := i.driveService.Files.Copy(billingModel.Id, &drive.File{
			MimeType: "application/vnd.google-apps.spreadsheet",
			Name:     fmt.Sprintf(i.invoicesData.Name, len(listInvoices.Files)+1, lastDayCurrentMonth.Format(invoiceDateLayout), company.Name),
		}).Do()
		if err != nil {
			return fmt.Errorf("copy file from model error: %w", err)
		}

		cc, err := i.exchangeRateService.
			ConvertCurrencyLatest(company.MoneySymbol, "EUR", company.UnitPrice)
		if err != nil {
			log.Println("conversion exchange rate api", err)
		}

		valueRange := make([]*sheets.ValueRange, 0)
		q, err := strconv.ParseFloat(fmt.Sprintf("%.2f", cc.Value*company.UnitPrice), 64)
		if err != nil {
			log.Println("error parsing float value of price, check it:", err)
			q = 1.0
		}
		row := map[string]interface{}{
			"H3":  fmt.Sprintf("%d", len(listInvoices.Files)+1),
			"H4":  dateSrv.GetNowSpainTime().Format(invoiceDateLayout),
			"H5":  dayPayment.Format(invoiceDateLayout),
			"H8":  company.Name,
			"H9":  company.Address,
			"B15": workHoursData.TotalHours,
			"C15": fmt.Sprintf(company.Description, firstDayCurrentMonth.Format(invoiceDateLayout), lastDayCurrentMonth.Format(invoiceDateLayout)),
			"E15": q,
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

		invoice, err := i.sheetService.Spreadsheets.Get(copiedFile.Id).Do()
		if err != nil {
			return fmt.Errorf("sheetService get spreadshet by id %s: %w", copiedFile.Id, err)
		}

		onTransactionCopy := gotransactions.OnTransaction(func() error {
			_, err = i.sheetService.Spreadsheets.Values.BatchUpdate(invoice.SpreadsheetId, &sheets.BatchUpdateValuesRequest{
				Data:             valueRange,
				ValueInputOption: "RAW",
			}).Do()
			if err != nil {
				return fmt.Errorf("batch update sheetService error: %w", err)
			}
			return nil
		})

		onRollback := gotransactions.OnRollback(func() error {
			_ = i.driveService.Files.Delete(copiedFile.Id).Do()
			return nil
		})

		if err := gotransactions.New(onTransactionCopy, onRollback).ExecuteTransaction(); err != nil {
			return err
		}
	}

	return nil
}
