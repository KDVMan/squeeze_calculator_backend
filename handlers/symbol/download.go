package handlers_symbol

import (
	"backend/core/services/app"
	"backend/core/services/logger"
	"backend/core/services/response"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	"log/slog"
	"net/http"
)

func DownloadHandler(appService *core_services_app.AppService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		logger := appService.LoggerService.With(
			slog.String("label", "handlers.symbol.DownloadHandler"),
			slog.String("request_id", middleware.GetReqID(r.Context())),
		)

		symbols, err := appService.ExchangeService.ExchangeInfo()

		if err != nil {
			message := "failed to download symbols [ExchangeInfo]"
			logger.Error(message, core_services_logger.Err(err))

			render.Status(r, http.StatusInternalServerError)
			render.JSON(w, r, core_services_response.Error(message))

			return
		}

		if err = appService.SymbolService.Download(symbols); err != nil {
			message := "failed to download symbols [Download]"
			logger.Error(message, core_services_logger.Err(err))

			render.Status(r, http.StatusInternalServerError)
			render.JSON(w, r, core_services_response.Error(message))

			return
		}

		logger.Info("downloaded")
	}
}
