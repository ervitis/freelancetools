package exchangerate

import (
	"context"
	"fmt"
	"github.com/ervitis/freelancetools/config"
	"github.com/ervitis/freelancetools/exchangerate/client"
	"github.com/ervitis/freelancetools/exchangerate/client/rates"
	"github.com/go-openapi/strfmt"
	"time"
)

type (
	ExchangeApi struct {
		URL        string
		APIKEY     string
		client     *client.FreecurrencyAPI
		currencies []string
	}

	ConvertedCurrency struct {
		Currency string
		Value    float64
	}
)

var (
	ErrorNoApiKeyProvided = fmt.Errorf("it is necessary an Api Key to do requests")
)

func NewClient(config *config.AppConfigParameters) (*ExchangeApi, error) {
	if config.ExchangeRateApi == "" {
		return nil, ErrorNoApiKeyProvided
	}

	c := &ExchangeApi{
		APIKEY: config.ExchangeRateApi,
		client: client.NewHTTPClient(strfmt.Default),
	}

	if err := c.getCurrencies(); err != nil {
		return nil, fmt.Errorf("error creating client: %w", err)
	}
	return c, nil
}

func (eapi *ExchangeApi) getCurrencies() error {
	if eapi.currencies != nil || len(eapi.currencies) != 0 {
		return nil
	}

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	resp, err := eapi.client.Rates.GetLatest(&rates.GetLatestParams{Apikey: eapi.APIKEY, Context: ctx})
	if err != nil {
		return fmt.Errorf("error when trying to retrieve symbols: %w", err)
	}

	eapi.currencies = make([]string, 0)
	for v := range resp.GetPayload().Data {
		eapi.currencies = append(eapi.currencies, fmt.Sprintf("%v", v))
	}
	return nil
}

func (eapi *ExchangeApi) validateInput(input string) error {
	if !contains(eapi.currencies, input) {
		return fmt.Errorf("input data is not valid for currency symbol %s", input)
	}
	return nil
}

func contains(l []string, e string) bool {
	for _, v := range l {
		if v == e {
			return true
		}
	}
	return false
}

func (eapi *ExchangeApi) ConvertCurrencyLatest(from string, to string, quantity float64) (*ConvertedCurrency, error) {
	if err := eapi.validateInput(from); err != nil {
		return nil, fmt.Errorf("convertCurrency: error validating from: %w", err)
	}

	if err := eapi.validateInput(to); err != nil {
		return nil, fmt.Errorf("convertCurrency: error validating to: %w", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	resp, err := eapi.client.Rates.GetLatest(&rates.GetLatestParams{
		Apikey: eapi.APIKEY, BaseCurrency: &from, Context: ctx,
	})
	if err != nil {
		return nil, fmt.Errorf("convertCurrency: error doing request: %w", err)
	}

	q, ok := resp.GetPayload().Data[to]
	if !ok {
		q = 0.0
	}

	return &ConvertedCurrency{
		Currency: to,
		Value:    q,
	}, nil
}
