package exchangerate

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sync"
	"time"

	"github.com/ervitis/freelancetools/config"
	"github.com/ervitis/freelancetools/exchangerate/client"
	"github.com/ervitis/freelancetools/exchangerate/client/rates"
	"github.com/go-openapi/strfmt"
)

type (
	ExchangeApi struct {
		URL        string
		APIKEY     string
		client     *client.FreecurrencyAPI
		currencies []string
		mtx        sync.Mutex
	}

	ConvertedCurrency struct {
		Currency string
		Value    float64
	}
)

var (
	currenciesFile = fmt.Sprintf(`env%scurrencies.json`, string(filepath.Separator))

	ErrNoApiKeyProvided  = fmt.Errorf("it is necessary an Api Key to do requests")
	ErrEmptyCurrencyFile = fmt.Errorf("currency file is empty")
)

func NewClient(config *config.AppConfigParameters) (*ExchangeApi, error) {
	if config.ExchangeRateApi == "" {
		return nil, ErrNoApiKeyProvided
	}

	c := &ExchangeApi{
		APIKEY: config.ExchangeRateApi,
		client: client.NewHTTPClient(strfmt.Default),
	}

	if err := c.loadCurrencies(); err != nil && !errors.Is(err, ErrEmptyCurrencyFile) {
		return nil, err
	}

	if err := c.getCurrencies(); err != nil {
		log.Println("There was an error trying exchange rate api", err)
	}

	if err := c.saveCurrencies(); err != nil {
		return nil, err
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

func (eapi *ExchangeApi) loadCurrencies() error {
	if _, err := os.Stat(currenciesFile); errors.Is(err, os.ErrNotExist) {
		return ErrEmptyCurrencyFile
	}

	eapi.mtx.Lock()
	defer eapi.mtx.Unlock()

	f, err := os.Open(currenciesFile)
	defer func() {
		_ = f.Close()
	}()
	if err != nil {
		return err
	}

	if st, err := f.Stat(); err != nil && st.Size() == 0 {
		return ErrEmptyCurrencyFile
	}

	if err := json.NewDecoder(f).Decode(&eapi.currencies); err != nil {
		return err
	}
	return nil
}

func (eapi *ExchangeApi) saveCurrencies() error {
	if len(eapi.currencies) == 0 {
		return nil
	}

	eapi.mtx.Lock()
	defer eapi.mtx.Unlock()

	f, err := os.OpenFile(currenciesFile, os.O_TRUNC|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		return err
	}

	defer func() {
		_ = f.Close()
	}()

	if err := json.NewEncoder(f).Encode(eapi.currencies); err != nil {
		return err
	}
	return nil
}

func (eapi *ExchangeApi) validateInput(input string) error {
	if len(eapi.currencies) == 0 {
		log.Println("currencies are empty due to api error, so let's skip")
		return nil
	}

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

func (eapi *ExchangeApi) ConvertCurrencyLatest(from string, to string) (*ConvertedCurrency, error) {
	if err := eapi.validateInput(from); err != nil {
		return nil, fmt.Errorf("convertCurrency: error validating from: %w", err)
	}

	if err := eapi.validateInput(to); err != nil {
		return nil, fmt.Errorf("convertCurrency: error validating to: %w", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	cc := &ConvertedCurrency{
		Currency: to,
		Value:    1.0,
	}

	resp, err := eapi.client.Rates.GetLatest(&rates.GetLatestParams{
		Apikey: eapi.APIKEY, BaseCurrency: &from, Context: ctx,
	})
	if err != nil {
		log.Println("convertCurrency: error doing request:", err)
	} else {
		q, ok := resp.GetPayload().Data[to]
		if !ok {
			log.Println(fmt.Sprintf("data for currency %s not exists, so you have to put the conversion rate manually", to))
			q = 1.0
		}
		cc.Value = q
	}

	return cc, nil
}
