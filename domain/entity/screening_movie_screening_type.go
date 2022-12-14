package entity

type ScreeningMovieScreeningType struct {
	ID               UUID
	CinemaID         UUID
	ScreeningMovieID UUID
	ScreeningTypeID  UUID
}

func NewScreeningMovieScreeningType(cinemaID, screeningMovieID, screeningTypeID UUID) *ScreeningMovieScreeningType {
	return &ScreeningMovieScreeningType{
		ID:               NewUUID(),
		CinemaID:         cinemaID,
		ScreeningMovieID: screeningMovieID,
		ScreeningTypeID:  screeningTypeID,
	}
}
