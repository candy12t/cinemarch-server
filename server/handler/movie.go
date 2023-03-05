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

func NewMovieHandler(movieUC usecase.Movie) *MovieHandler {
	return &MovieHandler{
		movieUC: movieUC,
	}
}

func (h *MovieHandler) Show(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	movieID := chi.URLParam(r, "movieID")
	movie, err := h.movieUC.FindByID(ctx, movieID)
	if err != nil {
		if errors.Is(err, entity.ErrMovieNotFound) {
			ResponseJSON(w, NewHTTPError(err.Error()), http.StatusNotFound)
			return
		}
		ResponseJSON(w, NewHTTPError(err.Error()), http.StatusInternalServerError)
		return
	}
	ResponseJSON(w, movieToResp(movie), http.StatusOK)
}

func (h *MovieHandler) Search(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	title := r.URL.Query().Get("title")
	movies, err := h.movieUC.FindAllByTitle(ctx, title)
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

	movieJSONs := make([]*movieResp, 0, len(movies))
	for _, movie := range movies {
		movieJSONs = append(movieJSONs, movieToResp(movie))
	}
	ResponseJSON(w, movieJSONs, http.StatusOK)
}

func (h *MovieHandler) Upsert(w http.ResponseWriter, r *http.Request) {
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

	params := usecase.UpsertMovieParams{
		Title:         req.Title,
		ReleaseDate:   req.ReleaseDate,
		ReleaseStatus: req.ReleaseStatus,
	}

	ctx := r.Context()
	movie, err := h.movieUC.Upsert(ctx, params)
	if err != nil {
		ResponseJSON(w, NewHTTPError(err.Error()), http.StatusBadRequest)
		return
	}
	ResponseJSON(w, movieToResp(movie), http.StatusOK)
}

type movieResp struct {
	ID            string `json:"id"`
	Title         string `json:"title"`
	ReleaseDate   string `json:"release_date"`
	ReleaseStatus string `json:"release_status"`
}

func movieToResp(movie *usecase.MovieDTO) *movieResp {
	return &movieResp{
		ID:            movie.ID,
		Title:         movie.Title,
		ReleaseDate:   movie.ReleaseDate,
		ReleaseStatus: movie.ReleaseStatus,
	}
}
