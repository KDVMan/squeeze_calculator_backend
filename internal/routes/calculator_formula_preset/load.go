package routes_calculator_formula_preset

import (
	"github.com/go-chi/render"
	"net/http"
)

func (object *calculatorFormulaPresetRouteImplementation) load() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		formulasPresetsModels, err := object.calculatorFormulaPresetService().Load()

		if err != nil {
			var message = "failed to load formulaPresetModel"
			object.loggerService().Error().Printf("%s: %v", message, err)

			render.Status(r, http.StatusInternalServerError)
			render.JSON(w, r, message)

			return
		}

		render.JSON(w, r, formulasPresetsModels)
	}
}
