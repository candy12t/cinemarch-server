package handler

import (
	"encoding/json"
	"net/http"
	"net/url"

	"github.com/candy12t/cinemarch-server/query/dto"
	"github.com/candy12t/cinemarch-server/usecase"
	"github.com/go-chi/chi/v5"
)

type ScreenMovieHandler struct {
	screenMovieUC usecase.ScreenMovie
}

func NewScreenMovieHandler(screenMovieUC usecase.ScreenMovie) *ScreenMovieHandler {
	return &ScreenMovieHandler{
		screenMovieUC: screenMovieUC,
	}
}

func (h *ScreenMovieHandler) List(w http.ResponseWriter, r *http.Request) {
	searchCondition := h.parseQuery(r.URL.Query())
	searchCondition[dto.MovieID] = chi.URLParam(r, "movieID")

	c := r.Context()
	screenMovies, err := h.screenMovieUC.Search(c, searchCondition)
	if err != nil {
		ResponseJSON(w, NewHTTPError(err.Error()), http.StatusInternalServerError)
		return
	}

	screenMovieJSONs := make([]*screenMovieResp, 0, len(screenMovies))
	for _, screenMovie := range screenMovies {
		screenMovieJSONs = append(screenMovieJSONs, screenMovieToResp(screenMovie))
	}
	ResponseJSON(w, screenMovieJSONs, http.StatusOK)
}

func (h *ScreenMovieHandler) parseQuery(params url.Values) dto.ScreenMovieSearchCondition {
	searchCondition := make(dto.ScreenMovieSearchCondition, len(dto.ScreenMovieSearchKeys))
	for k, v := range dto.ScreenMovieSearchKeys {
		if params.Has(k.String()) {
			searchCondition[k] = params.Get(k.String())
		} else if !params.Has(k.String()) && len(v) != 0 {
			searchCondition[k] = v
		}
	}
	return searchCondition
}

func (h *ScreenMovieHandler) Create(w http.ResponseWriter, r *http.Request) {
	type screenSchedule struct {
		StartTime string `json:"start_time"`
		EndTime   string `json:"end_time"`
	}
	type reqJSON struct {
		CinemaName      string            `json:"cinema_name"`
		MovieTitle      string            `json:"movie_title"`
		ScreenType      string            `json:"screen_type"`
		TranslateType   string            `json:"translate_type"`
		ThreeD          bool              `json:"three_d"`
		ScreenSchedules []*screenSchedule `json:"schedules"`
	}
	req := new(reqJSON)
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		ResponseJSON(w, NewHTTPError(err.Error()), http.StatusBadRequest)
		return
	}
	schedules := make([]*usecase.CreateScreenScheduleParams, 0, len(req.ScreenSchedules))
	for _, schedule := range req.ScreenSchedules {
		schedules = append(schedules, &usecase.CreateScreenScheduleParams{
			StartTime: schedule.StartTime,
			EndTime:   schedule.EndTime,
		})
	}
	ctx := r.Context()
	params := usecase.CreateScreenMovieParams{
		CinemaName:      req.CinemaName,
		MovieTitle:      req.MovieTitle,
		ScreenType:      req.ScreenType,
		TranslateType:   req.TranslateType,
		ThreeD:          req.ThreeD,
		ScreenSchedules: schedules,
	}
	screenMovie, err := h.screenMovieUC.Create(ctx, params)
	if err != nil {
		ResponseJSON(w, NewHTTPError(err.Error()), http.StatusInternalServerError)
		return
	}
	ResponseJSON(w, screenMovieToResp(screenMovie), http.StatusOK)
}

type screenMovieResp struct {
	CinemaName    string                `json:"cinema_name"`
	MovieTitle    string                `json:"movie_title"`
	ScreenType    string                `json:"screen_type"`
	TranslateType string                `json:"translate_type"`
	TreeD         bool                  `json:"three_d"`
	Schedules     []*screenScheduleJSON `jon:"schedules"`
}

type screenScheduleJSON struct {
	StartTime string `json:"start_time"`
	EndTime   string `json:"end_time"`
}

func screenMovieToResp(screenMovie *usecase.ScreenMovieDTO) *screenMovieResp {
	screenSchedules := make([]*screenScheduleJSON, 0, len(screenMovie.ScreenSchedules))
	for _, ss := range screenMovie.ScreenSchedules {
		screenSchedules = append(screenSchedules, &screenScheduleJSON{
			StartTime: ss.StartTime,
			EndTime:   ss.EndTime,
		})
	}
	return &screenMovieResp{
		CinemaName:    screenMovie.CinemaName,
		MovieTitle:    screenMovie.MovieTitle,
		ScreenType:    screenMovie.ScreenType,
		TranslateType: screenMovie.TranslateType,
		TreeD:         screenMovie.ThreeD,
		Schedules:     screenSchedules,
	}
}
