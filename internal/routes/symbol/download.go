package routes_symbol

import (
	"github.com/go-chi/render"
	"net/http"
)

func (object *symbolRouteImplementation) download() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		symbols, err := object.exchangeService().ExchangeInfo()

		if err != nil {
			message := "failed to load exchangeInfo"
			object.loggerService().Error().Printf("%s: %v", message, err)

			render.Status(r, http.StatusInternalServerError)
			render.JSON(w, r, message)

			return
		}

		if err = object.symbolService().Download(symbols); err != nil {
			message := "failed to download symbols"
			object.loggerService().Error().Printf("%s: %v", message, err)

			render.Status(r, http.StatusInternalServerError)
			render.JSON(w, r, message)

			return
		}

		object.loggerService().Info().Printf("downloaded")
	}
}
