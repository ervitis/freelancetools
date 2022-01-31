# Freelance tools

For automating processes using following products:

- Google Drive to create invoices
- Google Calendar to retrieve schedule work
- fixer.io exchange api for exchange rates

## Installing

Generate swagger client

```bash
 swagger generate client
```

Create a folder named `env` in the root of the project where you download the API credentials from your Google account console.

Create a file named `invoices.json` inside it that follows this format:

```json
{
  "name": "Name of the invoice file without extension",
  "spreadSheetIdFromCopy": "The spreadsheetId from it copies the model",
  "cellData": {
    "numberInvoice": "A3",
    "dateInvoice": "A4",
    "datePayment": "A5",
    "totalHours": "D15",
    "quantity": "E15"
  },
  "companies": [
    {
      "name": {
        "data": "name",
        "cell": "H2"
      },
      "address": {
        "data": "address",
        "cell": "H3"
      },
      "description": {
        "data": "description, also we can use the %s",
        "cell": "C5"
      },
      "unitPrice": 20000,
      "moneySymbol": "USD"
    }
  ]
}
```

## List of features

- Get working hours from Google calendar events
- Get the latest exchange rate of a currency
- Generate spreadsheet invoice

## List of TODO

- Better logging
- Personalize better the spreadsheet cells inside the json file