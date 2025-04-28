package routes_calculator

import (
	routes_interface_calculator "backend/internal/routes/calculator/interface"
	services_interface_calculator "backend/internal/services/calculator/interface"
	services_interface_calculator_preset "backend/internal/services/calculator_preset/interface"
	services_interface_logger "backend/pkg/services/logger/interface"
	services_interface_request "backend/pkg/services/request/interface"
	"github.com/go-chi/chi/v5"
)

type calculatorRouteImplementation struct {
	router                  *chi.Mux
	loggerService           func() services_interface_logger.LoggerService
	requestService          func() services_interface_request.RequestService
	calculatorService       func() services_interface_calculator.CalculatorService
	calculatorPresetService func() services_interface_calculator_preset.CalculatorPresetService
}

func NewCalculatorRoute(
	loggerService func() services_interface_logger.LoggerService,
	requestService func() services_interface_request.RequestService,
	calculatorService func() services_interface_calculator.CalculatorService,
	calculatorPresetService func() services_interface_calculator_preset.CalculatorPresetService,
) routes_interface_calculator.CalculatorRoute {
	route := &calculatorRouteImplementation{
		router:                  chi.NewRouter(),
		loggerService:           loggerService,
		requestService:          requestService,
		calculatorService:       calculatorService,
		calculatorPresetService: calculatorPresetService,
	}

	route.router.Post("/calculate", route.calculate())

	return route
}

func (object *calculatorRouteImplementation) GetRouter() *chi.Mux {
	return object.router
}
