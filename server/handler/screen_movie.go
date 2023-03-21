package handler

import (
	"encoding/json"
	"net/http"

	"github.com/candy12t/cinemarch-server/usecase"
)

type ScreenMovieHandler struct {
	screenMovieUC usecase.ScreenMovie
}

func NewScreenMovieHandler(screenMovieUC usecase.ScreenMovie) *ScreenMovieHandler {
	return &ScreenMovieHandler{
		screenMovieUC: screenMovieUC,
	}
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
