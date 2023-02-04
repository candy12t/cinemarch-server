package entity

import "time"

type ScreenMovie struct {
	ID        UUID
	CinemaID  UUID
	MovieID   UUID
	StartTime time.Time
	EndTime   time.Time
}

func NewScreenMovie(cinemaID, movieID UUID, startTime, endTime time.Time) *ScreenMovie {
	return &ScreenMovie{
		ID:        NewUUID(),
		CinemaID:  cinemaID,
		MovieID:   movieID,
		StartTime: startTime,
		EndTime:   endTime,
	}
}
