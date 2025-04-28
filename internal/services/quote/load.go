package services_quote

import (
	enums_quote "backend/internal/enums/quote"
	models_quote "backend/internal/models/quote"
	services_helper "backend/pkg/services/helper"
	"fmt"
)

func (object *quoteServiceImplementation) Load(request *models_quote.QuoteRequestModel) (*models_quote.QuoteResponseModel, error) {
	var err error
	var quotes []*models_quote.QuoteModel
	hash := services_helper.MustConvertStringToMd5(fmt.Sprintf("hash | symbol:%s | interval:%s", request.Symbol, request.Interval.String()))
	response := &models_quote.QuoteResponseModel{}

	if request.Type == enums_quote.QuoteTypeInit {
		object.exchangeWebsocketService().SubscribeCurrentPrice(request.Symbol, request.Interval)
	} else if request.Type == enums_quote.QuoteTypeCalculate {
		calculateResultsModels := object.calculatorService().LoadResult(request.Symbol)

		calculatorPresetModel, err := object.calculatorPresetService().LoadSelected()
		if err != nil {
			return nil, err
		}

		request.TimeStart = calculatorPresetModel.TimeFrom
		request.TimeEnd = calculatorPresetModel.TimeTo

		response.Deals = calculateResultsModels[request.Index].Deals
	}

	if request.TimeEnd > 0 {
		quotes, err = object.loadLocal(hash, request)
		if err != nil {
			return nil, err
		}

		if request.Type != enums_quote.QuoteTypeCalculate && len(quotes) < request.QuotesLimit {
			quotes, err = object.loadRemote(hash, request)
			if err != nil {
				return nil, err
			}
		}
	} else {
		quotes, err = object.loadRemote(hash, request)
		if err != nil {
			return nil, err
		}
	}

	response.TimeFrom = request.TimeStart
	response.TimeTo = request.TimeEnd
	response.Quotes = quotes

	return response, nil
}
