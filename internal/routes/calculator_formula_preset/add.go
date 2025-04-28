package routes_calculator_formula_preset

import (
	models_calculator_formula_preset "backend/internal/models/calculator_formula_preset"
	"github.com/go-chi/render"
	"net/http"
)

func (object *calculatorFormulaPresetRouteImplementation) add() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var request models_calculator_formula_preset.AddRequestModel

		if err := object.requestService().Decode(w, r, &request); err != nil {
			return
		}

		if err := object.requestService().Validate(w, r, &request); err != nil {
			return
		}

		formulasPresetsModels, err := object.calculatorFormulaPresetService().Add(&request)

		if err != nil {
			message := "failed to add formulaPresetModel"
			object.loggerService().Error().Printf("%s: %v", message, err)

			render.Status(r, http.StatusInternalServerError)
			render.JSON(w, r, message)

			return
		}

		render.JSON(w, r, formulasPresetsModels)
	}
}
