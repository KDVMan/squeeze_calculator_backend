package services_exchange_limit

import (
	enums_exchange_limit "backend/internal/enums/exchange_limit"
	enums_websocket "backend/internal/enums/websocket"
	models_exchange_limit "backend/internal/models/exchange_limit"
	models_websocket "backend/internal/models/websocket"
	"gorm.io/gorm"
	"strings"
)

func (object *exchangeLimitServiceImplementation) Update(limits map[string]int) error {
	for key, used := range limits {
		data := map[string]interface{}{
			"total_left": gorm.Expr("total - ?", used),
		}

		err := object.storageService().DB().
			Model(&models_exchange_limit.ExchangeLimitModel{}).
			Where("type = ? AND interval = ?", getType(key), getInterval(key)).
			Updates(data).
			Error

		if err != nil {
			return err
		}
	}

	exchangeLimitModel, err := object.Load()

	if err != nil {
		return err
	}

	broadcastModel := models_websocket.BroadcastChannelModel{
		Event: enums_websocket.WebsocketEventExchangeLimits,
		Data:  exchangeLimitModel,
	}

	object.websocketService().GetBroadcastChannel() <- &broadcastModel

	return nil
}

func getType(key string) enums_exchange_limit.RateType {
	switch key {
	case "x-mbx-used-weight", "x-mbx-used-weight-1m":
		return enums_exchange_limit.RateTypeWeight
	case "x-mbx-order-count-1s", "x-mbx-order-count-1m", "x-mbx-order-count-1h", "x-mbx-order-count-1d":
		return enums_exchange_limit.RateTypeOrder
	default:
		return enums_exchange_limit.RateTypeUnknown
	}
}

func getInterval(key string) enums_exchange_limit.RateInterval {
	if strings.Contains(key, "1s") {
		return enums_exchange_limit.RateIntervalSecond
	} else if strings.Contains(key, "1m") {
		return enums_exchange_limit.RateIntervalMinute
	} else if strings.Contains(key, "1h") {
		return enums_exchange_limit.RateIntervalHour
	} else if strings.Contains(key, "1d") {
		return enums_exchange_limit.RateIntervalDay
	}

	return enums_exchange_limit.RateIntervalUnknown
}
