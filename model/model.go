package model

type (
	companyInvoice struct {
		Name    cellPairData `json:"name"`
		Address cellPairData `json:"address"`

		Description cellPairData `json:"description"`
		UnitPrice   float64      `json:"unitPrice"`
		MoneySymbol string       `json:"moneySymbol"`
	}

	cellUnitData struct {
		NumberInvoice string `json:"numberInvoice"`
		DateInvoice   string `json:"dateInvoice"`
		DatePayment   string `json:"datePayment"`
		TotalHours    string `json:"totalHours"`
		Quantity      string `json:"quantity"`
	}

	cellPairData struct {
		Data string `json:"data"`
		Cell string `json:"cell"`
	}

	InvoicesData struct {
		Name                  string           `json:"name"`
		SpreadSheetIDFromCopy string           `json:"spreadSheetIdFromCopy"`
		CellData              cellUnitData     `json:"cellData"`
		Companies             []companyInvoice `json:"companies"`
	}
)
