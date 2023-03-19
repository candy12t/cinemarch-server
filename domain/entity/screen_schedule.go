package entity

import "time"

type ScreenSchedule struct {
	ID            UUID
	ScreenMovieID UUID
	StartTime     time.Time
	EndTime       time.Time
}

type ScreenSchedules []*ScreenSchedule

func NewScreenSchedule(screenMovieID UUID, startTime, endTime time.Time) *ScreenSchedule {
	return &ScreenSchedule{
		ID:            NewUUID(),
		ScreenMovieID: screenMovieID,
		StartTime:     startTime,
		EndTime:       endTime,
	}
}
