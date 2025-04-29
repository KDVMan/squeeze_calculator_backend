package routes_calculator_formula_preset

import (
	routes_interface_calculator_formula_preset "backend/internal/routes/calculator_formula_preset/interface"
	services_interface_calculator_formula_preset "backend/internal/services/calculator_formula_preset/interface"
	services_interface_logger "backend/pkg/services/logger/interface"
	services_interface_request "backend/pkg/services/request/interface"
	"github.com/go-chi/chi/v5"
)

type calculatorFormulaPresetRouteImplementation struct {
	loggerService                  func() services_interface_logger.LoggerService
	requestService                 func() services_interface_request.RequestService
	calculatorFormulaPresetService func() services_interface_calculator_formula_preset.CalculatorFormulaPresetService
	router                         *chi.Mux
}

func NewCalculatorFormulaPresetRoute(
	loggerService func() services_interface_logger.LoggerService,
	requestService func() services_interface_request.RequestService,
	calculatorFormulaPresetService func() services_interface_calculator_formula_preset.CalculatorFormulaPresetService,
) routes_interface_calculator_formula_preset.CalculatorFormulaPresetRoute {
	route := &calculatorFormulaPresetRouteImplementation{
		router:                         chi.NewRouter(),
		loggerService:                  loggerService,
		requestService:                 requestService,
		calculatorFormulaPresetService: calculatorFormulaPresetService,
	}

	route.router.Get("/load", route.load())
	route.router.Post("/add", route.add())
	route.router.Post("/edit", route.edit())
	route.router.Post("/delete", route.delete())
	route.router.Post("/update", route.update())
	route.router.Post("/duplicate", route.duplicate())

	return route
}

func (object *calculatorFormulaPresetRouteImplementation) GetRouter() *chi.Mux {
	return object.router
}
