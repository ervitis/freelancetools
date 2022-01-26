package main

import (
	"github.com/ervitis/freelancetools/config"
	"github.com/ervitis/freelancetools/exchangerate"
	"log"
)

func init() {
	config.LoadConfigApp()
}

func main() {
	exchangeRateClient, err := exchangerate.NewClient(&config.AppConfig)
	if err != nil {
		log.Fatalln(err)
	}

	data, err := exchangeRateClient.ConvertCurrencyLatest("EUR", "JPY", 3987.0)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(data)
}
