package services_interface_exchange

import (
	"github.com/adshao/go-binance/v2/futures"
)

type ExchangeService interface {
	ExchangeInfo() ([]futures.Symbol, error)
	Kline(string, string, int64, int) ([]*futures.Kline, error)
}
