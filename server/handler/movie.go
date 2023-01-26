package handler

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/candy12t/cinemarch-server/domain/entity"
	"github.com/candy12t/cinemarch-server/usecase"
	"github.com/go-chi/chi/v5"
)

type MovieHandler struct {
	movieUC usecase.Movie
}

func NewMovie(movieUC usecase.Movie) *MovieHandler {
	return &MovieHandler{
		movieUC: movieUC,
	}
}

func (h *MovieHandler) Show(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	movieID := chi.URLParam(r, "movieID")

	movie, err := h.movieUC.Show(ctx, movieID)
	if err != nil {
		switch {
		case errors.Is(err, entity.ErrMovieNotFound):
			ResponseJSON(w, NewHTTPError(err.Error()), http.StatusNotFound)
			return
		default:
			ResponseJSON(w, NewHTTPError(err.Error()), http.StatusInternalServerError)
			return
		}
	}

	ResponseJSON(w, movieToJSON(movie), http.StatusOK)
}

func (h *MovieHandler) Create(w http.ResponseWriter, r *http.Request) {
	type reqJSON struct {
		Title         string `json:"title"`
		ReleaseDate   string `json:"release_date"`
		ReleaseStatus string `json:"release_status"`
	}
	req := new(reqJSON)
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		ResponseJSON(w, NewHTTPError(err.Error()), http.StatusBadRequest)
		return
	}

	params := usecase.CreateMovieParams{
		Title:         req.Title,
		ReleaseDate:   req.ReleaseDate,
		ReleaseStatus: req.ReleaseStatus,
	}

	ctx := r.Context()
	movie, err := h.movieUC.Create(ctx, params)
	if err != nil {
		ResponseJSON(w, NewHTTPError(err.Error()), http.StatusInternalServerError)
		return
	}

	ResponseJSON(w, movieToJSON(movie), http.StatusOK)
}

func (h *MovieHandler) Update(w http.ResponseWriter, r *http.Request) {
	type reqJSON struct {
		ReleaseDate   string
		ReleaseStatus string
	}
	req := new(reqJSON)
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		ResponseJSON(w, NewHTTPError(err.Error()), http.StatusBadRequest)
		return
	}

	params := usecase.UpdateMovieParams{
		ReleaseDate:   req.ReleaseDate,
		ReleaseStatus: req.ReleaseStatus,
	}

	movieID := chi.URLParam(r, "movieID")
	ctx := r.Context()
	movie, err := h.movieUC.Update(ctx, movieID, params)
	if err != nil {
		ResponseJSON(w, NewHTTPError(err.Error()), http.StatusInternalServerError)
		return
	}

	ResponseJSON(w, movieToJSON(movie), http.StatusOK)
}

type movieJSON struct {
	ID            string `json:"id"`
	Title         string `json:"title"`
	ReleaseDate   string `json:"release_date"`
	ReleaseStatus string `json:"release_status"`
}

var format = "2006/01/02"

func movieToJSON(movie *usecase.MovieDTO) *movieJSON {
	releaseDate := movie.ReleaseDate.Format(format)
	return &movieJSON{
		ID:            movie.ID,
		Title:         movie.Title,
		ReleaseDate:   releaseDate,
		ReleaseStatus: movie.ReleaseStatus,
	}
}
