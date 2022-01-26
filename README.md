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
  "companies": [
    {
      "name": "Name of the company",
      "address": "Address of the company",
      "description": "Description of the job",
      "unitPrice": "Number of the unit price of the job",
      "moneySymbol": "Currency ISO symbol"
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