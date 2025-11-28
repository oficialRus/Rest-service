package remove

import (
	"log/slog"
	"net/http"

	"github.com/go-chi/chi/v5/middleware"
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
		resURL, err := urlDelete.DeleteURL(alias)
	}

}
