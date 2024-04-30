package handlers_symbol_calculator

import (
	"backend/core/services/app"
	"backend/core/services/logger"
	"backend/core/services/response"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	"log/slog"
	"net/http"
)

func LoadHandler(appService *core_services_app.AppService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		logger := appService.LoggerService.With(
			slog.String("label", "handlers.symbol_calculator.LoadHandler"),
			slog.String("request_id", middleware.GetReqID(r.Context())),
		)

		symbolCalculatorModel, err := appService.SymbolCalculatorService.Load()

		if err != nil {
			message := "failed to load symbolCalculatorModel"
			logger.Error(message, core_services_logger.Err(err))

			render.Status(r, http.StatusInternalServerError)
			render.JSON(w, r, core_services_response.Error(message))

			return
		}

		// logger.Info("symbolCalculatorModel loaded", slog.Any("symbolCalculatorModel", symbolCalculatorModel))
		render.JSON(w, r, symbolCalculatorModel)
	}
}
