package entity

import "time"

type ScreeningMovieScreeningType struct {
	ID               UUID
	CinemaID         UUID
	ScreeningMovieID UUID
	ScreeningTypeID  UUID
	CreatedAt        time.Time
	UpdatedAt        time.Time
}

func NewScreeningMovieScreeningType(cinemaID, screeningMovieID, screeningTypeID UUID) *ScreeningMovieScreeningType {
	now := time.Now()
	return &ScreeningMovieScreeningType{
		ID:               NewUUID(),
		CinemaID:         cinemaID,
		ScreeningMovieID: screeningMovieID,
		ScreeningTypeID:  screeningTypeID,
		CreatedAt:        now,
		UpdatedAt:        now,
	}
}
