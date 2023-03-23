package dto

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
)

type ScreenMovie struct {
	ID            string          `db:"id"`
	CinemaName    string          `db:"cinema_name"`
	MovieTitle    string          `db:"movie_title"`
	ScreenType    string          `db:"screen_type"`
	TranslateType string          `db:"translate_type"`
	ThreeD        bool            `db:"three_d"`
	Schedules     ScreenSchedules `db:"schedules"`
}

type ScreenMovies []*ScreenMovie

type ScreenSchedule struct {
	ID            string `db:"id" json:"id"`
	ScreenMovieID string `db:"screen_movie_id" json:"screen_movie_id"`
	StartTime     string `db:"start_time" json:"start_time"`
	EndTime       string `db:"end_time" json:"end_time"`
}

type ScreenSchedules []*ScreenSchedule

func (s *ScreenSchedules) Scan(val any) error {
	switch v := val.(type) {
	case []byte:
		json.Unmarshal(v, &s)
		return nil
	default:
		return fmt.Errorf("Unsupported type: %T", v)
	}
}

func (s *ScreenSchedules) Value() (driver.Value, error) {
	return json.Marshal(s)
}
