package routes_interface_calculator_formula_preset

import "github.com/go-chi/chi/v5"

type CalculatorFormulaPresetRoute interface {
	GetRouter() *chi.Mux
}
