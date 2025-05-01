package services_init

import (
	"backend/internal/enums"
	enums_symbol "backend/internal/enums/symbol"
	enums_websocket "backend/internal/enums/websocket"
	models_init "backend/internal/models/init"
	models_websocket "backend/internal/models/websocket"
)

func (object *initServiceImplementation) Update(request *models_init.UpdateRequestModel) (*models_init.InitModel, error) {
	initModel, err := object.Load()
	if err != nil {
		return nil, err
	}

	initModel.Symbol = request.Symbol
	initModel.Intervals = request.Intervals
	initModel.CalculateSortColumn = request.CalculateSortColumn
	initModel.CalculateSortDirection = request.CalculateSortDirection
	initModel.ExecActive = request.ExecActive

	symbolModel, err := object.symbolService().Load(request.Symbol, enums_symbol.SymbolStatusActive)
	if err != nil {
		return nil, err
	}

	initModel.Precision = symbolModel.Limit.Precision

	result := object.storageService().DB().Save(&initModel)
	if result.Error != nil {
		return nil, err
	}

	if initModel.ExecActive == enums.ExecActiveCalculate {
		object.websocketService().GetBroadcastChannel() <- &models_websocket.BroadcastChannelModel{
			Event: enums_websocket.WebsocketEventCalculateResult,
			Data:  object.calculatorService().LoadResult(request.Symbol),
		}
	} else if initModel.ExecActive == enums.ExecActiveBotList {
		object.websocketService().GetBroadcastChannel() <- &models_websocket.BroadcastChannelModel{
			Event: enums_websocket.WebsocketEventBotList,
			Data:  object.botService().LoadAll(),
		}
	}

	return initModel, nil
}
