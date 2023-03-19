package entity

type ScreenMovie struct {
	ID              UUID
	CinemaID        UUID
	MovieID         UUID
	ScreenType      ScreenType
	TranslateType   TranslateType
	TreeD           bool
	ScreenSchedules ScreenSchedules
}

type ScreenMovies []*ScreenMovie

func NewScreenMovie(cinemaID, movieID UUID, screenType ScreenType, translateType TranslateType, treeD bool) *ScreenMovie {
	return &ScreenMovie{
		ID:            NewUUID(),
		CinemaID:      cinemaID,
		MovieID:       movieID,
		ScreenType:    screenType,
		TranslateType: translateType,
		TreeD:         treeD,
	}
}

type TranslateType string

const (
	Subtitle TranslateType = "Subtitle"
	Dubbing  TranslateType = "Dubbing"
	Original TranslateType = "Original"
)

var translateTypes = []TranslateType{Subtitle, Dubbing, Original}

func NewTranslateType(translateType string) (TranslateType, error) {
	for _, tt := range translateTypes {
		if tt.String() == translateType {
			return tt, nil
		}
	}
	return "", ErrInvalidTranslateType
}

func (tt TranslateType) String() string {
	return string(tt)
}
