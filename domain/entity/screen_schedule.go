package entity

import "time"

type ScreenSchedule struct {
	ID            UUID
	ScreenMovieID UUID
	StartTime     time.Time
	EndTime       time.Time
}

func NewScreenSchedule(screen_movie_id UUID, start_time, end_time time.Time) *ScreenSchedule {
	return &ScreenSchedule{
		ID:            NewUUID(),
		ScreenMovieID: screen_movie_id,
		StartTime:     start_time,
		EndTime:       end_time,
	}
}
