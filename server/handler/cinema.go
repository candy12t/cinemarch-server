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
	cinema, err := h.cinemaUC.FindByID(ctx, cinemaID)
	if err != nil {
		if errors.Is(err, entity.ErrCinemaNotFound) {
			ResponseJSON(w, NewHTTPError(err.Error()), http.StatusNotFound)
			return
		}
		ResponseJSON(w, NewHTTPError(err.Error()), http.StatusInternalServerError)
		return
	}
	ResponseJSON(w, cinemaToResp(cinema), http.StatusOK)
}

func (h *CinemaHandler) Create(w http.ResponseWriter, r *http.Request) {
	type reqJSON struct {
		Name       string `json:"name"`
		Prefecture string `json:"prefecture"`
		Address    string `json:"address"`
		WebSite    string `json:"web_site"`
	}
	req := new(reqJSON)
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		ResponseJSON(w, NewHTTPError(err.Error()), http.StatusBadRequest)
		return
	}

	params := usecase.CreateCinemaParams{
		Name:       req.Name,
		Prefecture: req.Prefecture,
		Address:    req.Address,
		WebSite:    req.WebSite,
	}

	ctx := r.Context()
	cinema, err := h.cinemaUC.Create(ctx, params)
	if err != nil {
		ResponseJSON(w, NewHTTPError(err.Error()), http.StatusBadRequest)
		return
	}
	ResponseJSON(w, cinemaToResp(cinema), http.StatusOK)
}

type cinemaResp struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	Prefecture string `json:"prefecture"`
	Address    string `json:"address"`
	WebSite    string `json:"web_site"`
}

func cinemaToResp(cinema *usecase.CinemaDTO) *cinemaResp {
	return &cinemaResp{
		ID:         cinema.ID,
		Name:       cinema.Name,
		Prefecture: cinema.Prefecture,
		Address:    cinema.Address,
		WebSite:    cinema.WebSite,
	}
}
