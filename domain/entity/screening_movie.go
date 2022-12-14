package entity

import "time"

type ScreeningMovie struct {
	ID        UUID
	CinemaID  UUID
	MovieID   UUID
	StartTime time.Time
	EndTime   time.Time
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewScreeningMovie(cinemaID, movieID UUID, startTime, endTime time.Time) *ScreeningMovie {
	now := time.Now()
	return &ScreeningMovie{
		ID:        NewUUID(),
		CinemaID:  cinemaID,
		MovieID:   movieID,
		StartTime: startTime,
		EndTime:   endTime,
		CreatedAt: now,
		UpdatedAt: now,
	}
}
