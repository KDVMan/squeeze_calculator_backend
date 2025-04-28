package services_interface_exchange_limit

import (
	models_exchange_limit "backend/internal/models/exchange_limit"
	"github.com/adshao/go-binance/v2/futures"
)

type ExchangeLimitService interface {
	Load() ([]*models_exchange_limit.ExchangeLimitModel, error)
	Create(limits []futures.RateLimit) error
	Update(limits map[string]int) error
}
