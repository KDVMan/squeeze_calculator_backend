package services_quote_repository

import (
	"backend/internal/enums"
	models_quote "backend/internal/models/quote"
	services_helper "backend/pkg/services/helper"
	"github.com/adshao/go-binance/v2/futures"
)

func (object *quoteRepositoryServiceImplementation) Add(symbol string, quoteModel []*models_quote.QuoteModel) {
	object.mutex.Lock()
	defer object.mutex.Unlock()

	object.longData[symbol] = quoteModel
	object.shortData[symbol] = models_quote.InvertAll(quoteModel)
}

func (object *quoteRepositoryServiceImplementation) UpdateQuote(symbol string, interval enums.Interval, trade *futures.WsAggTradeEvent) {
	object.mutex.Lock()
	defer object.mutex.Unlock()

	longQuotesModels, exists := object.longData[symbol]
	if !exists || len(longQuotesModels) == 0 {
		return
	}

	shortQuotesModels, exists := object.shortData[symbol]
	if !exists || len(shortQuotesModels) == 0 {
		return
	}

	price := services_helper.MustConvertStringToFloat64(trade.Price)
	intervalMs := enums.IntervalMilliseconds(interval)
	tradeTime := trade.TradeTime
	candleOpenTime := (tradeTime / intervalMs) * intervalMs
	candleCloseTime := candleOpenTime + intervalMs - 1

	var currentQuote *models_quote.QuoteModel
	totalLongLen := len(longQuotesModels)

	if longQuotesModels[totalLongLen-1].TimeOpen == candleOpenTime {
		currentQuote = longQuotesModels[totalLongLen-1]
		currentQuote.IsClosed = false
	} else {
		longQuotesModels[totalLongLen-1].IsClosed = true
		shortQuotesModels[len(shortQuotesModels)-1].IsClosed = true

		if len(longQuotesModels) >= 1440 {
			longQuotesModels = longQuotesModels[1:]
			shortQuotesModels = shortQuotesModels[1:]
		}

		currentQuote = &models_quote.QuoteModel{
			Symbol:             symbol,
			Interval:           interval,
			TimeOpen:           candleOpenTime,
			TimeClose:          candleCloseTime,
			PriceOpen:          price,
			PriceHigh:          price,
			PriceLow:           price,
			PriceClose:         price,
			VolumeLeft:         0,
			Trades:             0,
			IsClosed:           false,
			TimeOpenFormatted:  services_helper.MustConvertUnixMillisecondsToString(candleOpenTime),
			TimeCloseFormatted: services_helper.MustConvertUnixMillisecondsToString(candleCloseTime),
		}

		longQuotesModels = append(longQuotesModels, currentQuote)
		shortQuotesModels = append(shortQuotesModels, models_quote.Invert(currentQuote))
	}

	if price > currentQuote.PriceHigh {
		currentQuote.PriceHigh = price
	}

	if price < currentQuote.PriceLow {
		currentQuote.PriceLow = price
	}

	currentQuote.PriceClose = price

	object.longData[symbol] = longQuotesModels

	currentShortQuote := *currentQuote
	currentShortQuote.PriceHigh = -currentQuote.PriceLow
	currentShortQuote.PriceLow = -currentQuote.PriceHigh
	currentShortQuote.PriceOpen = -currentQuote.PriceOpen
	currentShortQuote.PriceClose = -currentQuote.PriceClose

	object.shortData[symbol] = shortQuotesModels

	// log.Printf(
	// 	"current, symbol: %s | time: %s | openTime: %s, closeTime: %s, closed: %v\n",
	// 	symbol,
	// 	services_helper.MustConvertUnixMillisecondsToString(trade.TradeTime),
	// 	services_helper.MustConvertUnixMillisecondsToString(currentQuote.TimeOpen),
	// 	services_helper.MustConvertUnixMillisecondsToString(currentQuote.TimeClose),
	// 	currentQuote.IsClosed,
	// )
	//
	// log.Printf(
	// 	"prev long, symbol: %s | openTime: %s, closeTime: %s, closed: %v\n",
	// 	symbol,
	// 	services_helper.MustConvertUnixMillisecondsToString(longQuotesModels[len(longQuotesModels)-2].TimeOpen),
	// 	services_helper.MustConvertUnixMillisecondsToString(longQuotesModels[len(longQuotesModels)-2].TimeClose),
	// 	longQuotesModels[len(longQuotesModels)-2].IsClosed,
	// )
	//
	// log.Printf(
	// 	"prev short, symbol: %s | openTime: %s, closeTime: %s, closed: %v\n\n",
	// 	symbol,
	// 	services_helper.MustConvertUnixMillisecondsToString(shortQuotesModels[len(shortQuotesModels)-2].TimeOpen),
	// 	services_helper.MustConvertUnixMillisecondsToString(shortQuotesModels[len(shortQuotesModels)-2].TimeClose),
	// 	shortQuotesModels[len(shortQuotesModels)-2].IsClosed,
	// )
}

func (object *quoteRepositoryServiceImplementation) GetBySymbol(symbol string, tradeDirection enums.TradeDirection) []*models_quote.QuoteModel {
	object.mutex.Lock()
	defer object.mutex.Unlock()

	if tradeDirection == enums.TradeDirectionLong {
		return object.longData[symbol]
	} else if tradeDirection == enums.TradeDirectionShort {
		return object.shortData[symbol]
	}

	return nil
}

func (object *quoteRepositoryServiceImplementation) GetWindowBySymbol(symbol string, tradeDirection enums.TradeDirection, window int) []*models_quote.QuoteModel {
	object.mutex.Lock()
	defer object.mutex.Unlock()

	var data []*models_quote.QuoteModel
	if tradeDirection == enums.TradeDirectionLong {
		data = object.longData[symbol]
	} else if tradeDirection == enums.TradeDirectionShort {
		data = object.shortData[symbol]
	}

	if len(data) == 0 || window <= 0 {
		return nil
	}

	if len(data) < window {
		window = len(data)
	}

	result := make([]*models_quote.QuoteModel, window)
	copy(result, data[len(data)-window:])

	return result
}

func (object *quoteRepositoryServiceImplementation) Remove(symbol string) {
	object.mutex.Lock()
	defer object.mutex.Unlock()

	delete(object.longData, symbol)
	delete(object.shortData, symbol)
}
