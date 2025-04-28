package routes_calculator_formula_preset

import (
	models_calculator_formula_preset "backend/internal/models/calculator_formula_preset"
	"github.com/go-chi/render"
	"net/http"
)

func (object *calculatorFormulaPresetRouteImplementation) update() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var request models_calculator_formula_preset.UpdateRequestModel

		if err := object.requestService().Decode(w, r, &request); err != nil {
			return
		}

		if err := object.requestService().Validate(w, r, &request); err != nil {
			return
		}

		presetsModels, err := object.calculatorFormulaPresetService().Update(&request)
		if err != nil {
			message := "failed to update presetModel"
			object.loggerService().Error().Printf("%s: %v", message, err)

			render.Status(r, http.StatusInternalServerError)
			render.JSON(w, r, message)

			return
		}

		render.JSON(w, r, presetsModels)
	}
}
