package entity

import "time"

type ScreeningMovie struct {
	ID        UUID
	CinemaID  UUID
	MovieID   UUID
	StartTime time.Time
	EndTime   time.Time
}

func NewScreeningMovie(cinemaID, movieID UUID, startTime, endTime time.Time) *ScreeningMovie {
	return &ScreeningMovie{
		ID:        NewUUID(),
		CinemaID:  cinemaID,
		MovieID:   movieID,
		StartTime: startTime,
		EndTime:   endTime,
	}
}
