package entity

type ScreeningMovieScreeningType struct {
	ID               UUID
	ScreeningMovieID UUID
	ScreeningTypeID  UUID
}

func NewScreeningMovieScreeningType(screeningMovieID, screeningTypeID UUID) *ScreeningMovieScreeningType {
	return &ScreeningMovieScreeningType{
		ID:               NewUUID(),
		ScreeningMovieID: screeningMovieID,
		ScreeningTypeID:  screeningTypeID,
	}
}
