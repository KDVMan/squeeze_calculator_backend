package services_quote

import (
	"backend/internal/enums"
	enums_websocket "backend/internal/enums/websocket"
	models_quote "backend/internal/models/quote"
	models_websocket "backend/internal/models/websocket"
	services_helper "backend/pkg/services/helper"
	"fmt"
	"github.com/adshao/go-binance/v2/futures"
	"gorm.io/gorm"
)

func (object *quoteServiceImplementation) CurrentPrice(symbol string, interval enums.Interval, kline futures.WsKline) error {
	hash := services_helper.MustConvertStringToMd5(fmt.Sprintf("hash | symbol:%s | interval:%s", symbol, interval.String()))
	quote := models_quote.WsKlineToQuote(hash, symbol, interval, kline)

	if quote.IsClosed {
		err := object.storageService().DB().Transaction(func(tx *gorm.DB) error {
			if err := tx.Where("hash = ? AND time_open = ?", quote.Hash, quote.TimeOpen).
				Assign(*quote).
				FirstOrCreate(&quote).Error; err != nil {
				return err
			}

			return nil
		})
		if err != nil {
			return err
		}
	}

	broadcastModel := models_websocket.BroadcastChannelModel{
		Event: enums_websocket.WebsocketEventCurrentPrice,
		Data:  quote,
	}

	object.websocketService().GetBroadcastChannel() <- &broadcastModel

	return nil
}
