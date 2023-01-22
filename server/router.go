package server

import (
	"fmt"
	"net/http"

	"github.com/candy12t/cinema-search-server/server/handler"
	"github.com/candy12t/cinema-search-server/usecase"
	"github.com/go-chi/chi/v5"
)

var uuidRegexpPattern = `[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}`

func NewRouter(movieUC usecase.Movie) http.Handler {
	r := chi.NewRouter()

	r.Get("/healthcheck", handler.Healthcheck)

	r.Route("/movies", func(r chi.Router) {
		h := handler.NewMovie(movieUC)
		r.Post("/", h.Create)
		r.Route(fmt.Sprintf("/{movieID:%s}", uuidRegexpPattern), func(r chi.Router) {
			r.Get("/", h.Show)
		})
	})

	return r
}
