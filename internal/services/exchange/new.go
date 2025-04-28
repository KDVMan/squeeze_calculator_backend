package services_exchange

import (
	services_interface_exchange "backend/internal/services/exchange/interface"
	services_exchange_limit "backend/internal/services/exchange_limit"
	services_interface_exchange_limit "backend/internal/services/exchange_limit/interface"
	services_interface_storage "backend/pkg/services/storage/interface"
	"github.com/adshao/go-binance/v2"
	"github.com/adshao/go-binance/v2/futures"
	"net/http"
)

type exchangeServiceImplementation struct {
	storageService       func() services_interface_storage.StorageService
	exchangeLimitService func() services_interface_exchange_limit.ExchangeLimitService
	client               *futures.Client
	listenKey            string
	stopRenewListenKey   chan struct{}
}

func NewExchangeService(
	storageService func() services_interface_storage.StorageService,
	exchangeLimitService func() services_interface_exchange_limit.ExchangeLimitService,
) services_interface_exchange.ExchangeService {
	client := binance.NewFuturesClient("", "")

	client.HTTPClient = &http.Client{
		Transport: &services_exchange_limit.Transport{Value: http.DefaultTransport},
	}

	return &exchangeServiceImplementation{
		storageService:       storageService,
		exchangeLimitService: exchangeLimitService,
		client:               client,
		listenKey:            "",
		stopRenewListenKey:   nil,
	}
}
