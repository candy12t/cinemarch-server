package entity

type ScreenMovieScreenType struct {
	ID               UUID
	ScreenMovieID UUID
	ScreenTypeID  UUID
}

func NewScreenMovieScreenType(screenMovieID, screenTypeID UUID) *ScreenMovieScreenType {
	return &ScreenMovieScreenType{
		ID:               NewUUID(),
		ScreenMovieID: screenMovieID,
		ScreenTypeID:  screenTypeID,
	}
}
