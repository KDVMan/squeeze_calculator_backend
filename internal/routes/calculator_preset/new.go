package routes_calculator_preset

import (
	routes_interface_calculator_preset "backend/internal/routes/calculator_preset/interface"
	services_interface_calculator_preset "backend/internal/services/calculator_preset/interface"
	services_interface_logger "backend/pkg/services/logger/interface"
	services_interface_request "backend/pkg/services/request/interface"
	"github.com/go-chi/chi/v5"
)

type calculatorPresetRouteImplementation struct {
	router                  *chi.Mux
	loggerService           func() services_interface_logger.LoggerService
	requestService          func() services_interface_request.RequestService
	calculatorPresetService func() services_interface_calculator_preset.CalculatorPresetService
}

func NewCalculatorPresetRoute(
	loggerService func() services_interface_logger.LoggerService,
	requestService func() services_interface_request.RequestService,
	calculatorPresetService func() services_interface_calculator_preset.CalculatorPresetService,
) routes_interface_calculator_preset.CalculatorPresetRoute {
	route := &calculatorPresetRouteImplementation{
		router:                  chi.NewRouter(),
		loggerService:           loggerService,
		requestService:          requestService,
		calculatorPresetService: calculatorPresetService,
	}

	route.router.Get("/load", route.load())
	route.router.Post("/add", route.add())
	route.router.Post("/edit", route.edit())
	route.router.Post("/delete", route.delete())
	route.router.Post("/duplicate", route.duplicate())

	return route
}

func (object *calculatorPresetRouteImplementation) GetRouter() *chi.Mux {
	return object.router
}
