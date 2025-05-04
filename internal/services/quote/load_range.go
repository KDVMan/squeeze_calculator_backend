package services_quote

import (
	"backend/internal/enums"
	enums_exchange_limit "backend/internal/enums/exchange_limit"
	enums_quote "backend/internal/enums/quote"
	models_quote "backend/internal/models/quote"
	models_websocket "backend/internal/models/websocket"
	"sort"
	"time"
)

func (object *quoteServiceImplementation) LoadRange(
	symbol string,
	quoteRange *models_quote.QuoteRangeModel,
	progressModel *models_websocket.ProgressChannelModel,
) ([]*models_quote.QuoteModel, error) {
	var quotes []*models_quote.QuoteModel
	timeSet := make(map[int64]bool)

	request := &models_quote.QuoteRequestModel{
		Symbol:      symbol,
		Interval:    enums.Interval1m,
		TimeEnd:     quoteRange.TimeTo,
		QuotesLimit: int(quoteRange.QuotesLimit),
		Type:        enums_quote.QuoteTypeRange,
	}

	for i := 0; i < quoteRange.Iterations; i++ {
		// if variables_calculator.Stop {
		// 	return nil, nil
		// }

		exchangeLimitModel, err := object.exchangeLimitService().Load()
		if err != nil {
			return nil, err
		}

		shouldWait := false

		for _, limit := range exchangeLimitModel {
			if limit.Type == enums_exchange_limit.RateTypeWeight && limit.TotalLeft <= 1000 {
				shouldWait = true
				break
			}
		}

		if shouldWait {
			now := time.Now()
			wait := 60 - now.Second()
			time.Sleep(time.Duration(wait+1) * time.Second)
		}

		result, err := object.Load(request)
		if err != nil {
			return nil, err
		}

		for _, quote := range result.Quotes {
			if !timeSet[quote.TimeOpen] && quote.TimeOpen >= quoteRange.TimeFrom && quote.TimeOpen <= quoteRange.TimeTo {
				timeSet[quote.TimeOpen] = true
				quotes = append(quotes, quote)
			}
		}

		request.TimeEnd -= quoteRange.TimeStep
		progressModel.Count++
		object.websocketService().GetProgressChannel() <- progressModel
	}

	sort.Slice(quotes, func(i, j int) bool {
		return quotes[i].TimeOpen < quotes[j].TimeOpen
	})

	return quotes, nil
}
