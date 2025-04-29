package routes_calculator_formula_preset

import (
	models_calculator_formula_preset "backend/internal/models/calculator_formula_preset"
	"github.com/go-chi/render"
	"net/http"
)

func (object *calculatorFormulaPresetRouteImplementation) duplicate() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var request models_calculator_formula_preset.DuplicateRequestModel

		if err := object.requestService().Decode(w, r, &request); err != nil {
			return
		}

		if err := object.requestService().Validate(w, r, &request); err != nil {
			return
		}

		formulasPresetsModels, err := object.calculatorFormulaPresetService().Duplicate(&request)

		if err != nil {
			message := "failed to duplicate formulaPresetModel"
			object.loggerService().Error().Printf("%s: %v", message, err)

			render.Status(r, http.StatusInternalServerError)
			render.JSON(w, r, message)

			return
		}

		render.JSON(w, r, formulasPresetsModels)
	}
}
