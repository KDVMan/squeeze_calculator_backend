package routes_calculator

import (
	models_calculator "backend/internal/models/calculator"
	"github.com/go-chi/render"
	"net/http"
)

func (object *calculatorRouteImplementation) calculate() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var request models_calculator.CalculatorRequestModel

		if err := object.requestService().Decode(w, r, &request); err != nil {
			return
		}

		if err := object.requestService().Validate(w, r, &request); err != nil {
			return
		}

		if err := object.calculatorService().Calculator(&request); err != nil {
			var message = "failed to calculate"
			object.loggerService().Error().Printf("%s: %v", message, err)

			render.Status(r, http.StatusInternalServerError)
			render.JSON(w, r, message)

			return
		}

		render.JSON(w, r, nil)
	}
}
