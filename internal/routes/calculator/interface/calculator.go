package routes_interface_calculator

import "github.com/go-chi/chi/v5"

type CalculatorRoute interface {
	GetRouter() *chi.Mux
}
