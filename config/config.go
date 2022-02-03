package config

import (
	"github.com/sakirsensoy/genv"
	_ "github.com/sakirsensoy/genv/dotenv/autoload"
)

type (
	AppConfigParameters struct {
		ExchangeRateApi string
		DriveID         string
	}
)

var (
	AppConfig AppConfigParameters
)

func LoadConfigApp() {
	AppConfig.ExchangeRateApi = genv.Key("EXCHANGE_RATE_API").String()
	AppConfig.DriveID = genv.Key("DRIVE_ID").String()
}
