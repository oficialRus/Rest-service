package remove

import (
	"errors"
	"log/slog"
	"net/http"
	"rest-service/internal/lib/api/response"
	"rest-service/internal/storage"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
)

type RemoveURL interface {
	DeleteURL(alias string) error
}

func New(log *slog.Logger, urlDelete RemoveURL) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "internal.http-server.handlers.url.delete"
		log = log.With(
			slog.String("op", op),
			slog.String("request_id", middleware.GetReqID(r.Context())),
		)
		alias := chi.URLParam(r, "alias")
		if alias == "" {
			log.Info("alias is empty")
			render.JSON(w, r, response.Error("internal error"))
			return
		}

		err := urlDelete.DeleteURL(alias)
		if errors.Is(err, storage.ErrURLNotFound) {
			log.Info("url not found")
			render.JSON(w, r, response.Error("url not found"))
			return
		}
		log.Info("Sucsesful delete url")
		render.JSON(w, r, response.OK())

	}

}
