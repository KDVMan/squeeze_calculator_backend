package routes_interface_calculator_preset

import "github.com/go-chi/chi/v5"

type CalculatorPresetRoute interface {
	GetRouter() *chi.Mux
}
