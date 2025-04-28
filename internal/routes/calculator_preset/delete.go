package routes_calculator_preset

import (
	models_calculator_preset "backend/internal/models/calculator_preset"
	"github.com/go-chi/render"
	"net/http"
)

func (object *calculatorPresetRouteImplementation) delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var request models_calculator_preset.DeleteRequestModel

		if err := object.requestService().Decode(w, r, &request); err != nil {
			return
		}

		if err := object.requestService().Validate(w, r, &request); err != nil {
			return
		}

		presetsModels, err := object.calculatorPresetService().Delete(&request)

		if err != nil {
			message := "failed to delete presetsModels"
			object.loggerService().Error().Printf("%s: %v", message, err)

			render.Status(r, http.StatusInternalServerError)
			render.JSON(w, r, message)

			return
		}

		render.JSON(w, r, presetsModels)
	}
}
