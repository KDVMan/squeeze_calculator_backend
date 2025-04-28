package routes_calculator_preset

import (
	"github.com/go-chi/render"
	"net/http"
)

func (object *calculatorPresetRouteImplementation) load() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		presetsModels, err := object.calculatorPresetService().Load()

		if err != nil {
			var message = "failed to load presetsModels"
			object.loggerService().Error().Printf("%s: %v", message, err)

			render.Status(r, http.StatusInternalServerError)
			render.JSON(w, r, message)

			return
		}

		render.JSON(w, r, presetsModels)
	}
}
