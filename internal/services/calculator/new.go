package services_calculator

import (
	models_calculate "backend/internal/models/calculate"
	services_interface_calculator "backend/internal/services/calculator/interface"
	services_interface_calculator_formula_preset "backend/internal/services/calculator_formula_preset/interface"
	services_interface_init "backend/internal/services/init/interface"
	services_interface_quote "backend/internal/services/quote/interface"
	services_interface_symbol "backend/internal/services/symbol/interface"
	services_interface_websocket "backend/internal/services/websocket/interface"
	services_interface_config "backend/pkg/services/config/interface"
	services_interface_dump "backend/pkg/services/dump/interface"
	services_interface_storage "backend/pkg/services/storage/interface"
)

type calculatorServiceImplementation struct {
	configService                  func() services_interface_config.ConfigService
	storageService                 func() services_interface_storage.StorageService
	websocketService               func() services_interface_websocket.WebsocketService
	dumpService                    func() services_interface_dump.DumpService
	initService                    func() services_interface_init.InitService
	symbolService                  func() services_interface_symbol.SymbolService
	quoteService                   func() services_interface_quote.QuoteService
	calculatorFormulaPresetService func() services_interface_calculator_formula_preset.CalculatorFormulaPresetService
	calculateSymbol                string
	calculateResult                []*models_calculate.CalculateResultModel
}

func NewCalculatorService(
	configService func() services_interface_config.ConfigService,
	storageService func() services_interface_storage.StorageService,
	websocketService func() services_interface_websocket.WebsocketService,
	dumpService func() services_interface_dump.DumpService,
	initService func() services_interface_init.InitService,
	symbolService func() services_interface_symbol.SymbolService,
	quoteService func() services_interface_quote.QuoteService,
	calculatorFormulaPresetService func() services_interface_calculator_formula_preset.CalculatorFormulaPresetService,
) services_interface_calculator.CalculatorService {
	return &calculatorServiceImplementation{
		configService:                  configService,
		storageService:                 storageService,
		websocketService:               websocketService,
		dumpService:                    dumpService,
		initService:                    initService,
		symbolService:                  symbolService,
		quoteService:                   quoteService,
		calculatorFormulaPresetService: calculatorFormulaPresetService,
	}
}
