package handler

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/candy12t/cinemarch-server/domain/entity"
	"github.com/candy12t/cinemarch-server/usecase"
	"github.com/go-chi/chi/v5"
)

type CinemaHandler struct {
	cinemaUC usecase.Cinema
}

func NewCinemaHandler(cinemaUC usecase.Cinema) *CinemaHandler {
	return &CinemaHandler{
		cinemaUC: cinemaUC,
	}
}

func (h *CinemaHandler) Show(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	cinemaID := chi.URLParam(r, "cinemaID")

	cinema, err := h.cinemaUC.Show(ctx, cinemaID)
	if err != nil {
		switch {
		case errors.Is(err, entity.ErrCinemaNotFound):
			ResponseJSON(w, NewHTTPError(err.Error()), http.StatusNotFound)
			return
		default:
			ResponseJSON(w, NewHTTPError(err.Error()), http.StatusInternalServerError)
			return
		}
	}
	ResponseJSON(w, cinemaToJSON(cinema), http.StatusOK)
}

func (h *CinemaHandler) Create(w http.ResponseWriter, r *http.Request) {
	type reqJSON struct {
		Name    string `json:"name"`
		Address string `json:"address"`
		URL     string `json:"url"`
	}
	req := new(reqJSON)
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		ResponseJSON(w, NewHTTPError(err.Error()), http.StatusBadRequest)
	}

	params := usecase.CreateCinemaParams{
		Name:    req.Name,
		Address: req.Address,
		URL:     req.URL,
	}

	ctx := r.Context()
	cinema, err := h.cinemaUC.Create(ctx, params)
	if err != nil {
		ResponseJSON(w, NewHTTPError(err.Error()), http.StatusInternalServerError)
		return
	}

	ResponseJSON(w, cinemaToJSON(cinema), http.StatusOK)
}

type CinemaJSON struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Address string `json:"address"`
	URL     string `json:"url"`
}

func cinemaToJSON(cinema *usecase.CinemaDTO) *CinemaJSON {
	return &CinemaJSON{
		ID:      cinema.ID,
		Name:    cinema.Name,
		Address: cinema.Address,
		URL:     cinema.URL,
	}
}
