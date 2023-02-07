package entity

type ScreenMovie struct {
	ID                UUID
	CinemaID          UUID
	MovieID           UUID
	ScreenType        ScreenType
	SubtitleOrDubbing SubtitleOrDubbing
	ThreeD            bool
	Schedules         []*ScreenSchedule
}

func NewScreenMovie(cinemaID, movieID UUID, screenType ScreenType, subOrDub SubtitleOrDubbing, threeD bool) *ScreenMovie {
	return &ScreenMovie{
		ID:                NewUUID(),
		CinemaID:          cinemaID,
		MovieID:           movieID,
		ScreenType:        screenType,
		SubtitleOrDubbing: subOrDub,
		ThreeD:            threeD,
	}
}

type ScreenType string

const (
	IMAX         ScreenType = "IMAX"
	IMAXLaser    ScreenType = "IMAX Laser"
	IMAXLaserGT  ScreenType = "IMAX Laser/GT Technology"
	DolbyAtmos   ScreenType = "DolbyAtmos"
	DolbyCinema  ScreenType = "DolbyCinema"
	FourDX       ScreenType = "4DX"
	MX4D         ScreenType = "MX4D"
	ScreenX      ScreenType = "ScreenX"
	FourDXScreen ScreenType = "4DXScreen"
	BESTIA       ScreenType = "BESTIA"
)

var screenTypes = []ScreenType{IMAX, IMAXLaser, IMAXLaserGT, DolbyAtmos, DolbyCinema, FourDX, MX4D, ScreenX, FourDXScreen, BESTIA}

func NewScreenType(screenType string) (ScreenType, error) {
	for _, st := range screenTypes {
		if st.String() == screenType {
			return st, nil
		}
	}
	return "", ErrInvalidScreenType
}

func (st ScreenType) String() string {
	return string(st)
}

type SubtitleOrDubbing string

const (
	Subtitle SubtitleOrDubbing = "Subtitle"
	Dubbing  SubtitleOrDubbing = "Dubbing"
	Neither  SubtitleOrDubbing = "Neither"
)

var subtitleOrDubbings = []SubtitleOrDubbing{Subtitle, Dubbing, Neither}

func NewSubtitleOrDubbing(subOrDub string) (SubtitleOrDubbing, error) {
	for _, sd := range subtitleOrDubbings {
		if sd.String() == subOrDub {
			return sd, nil
		}
	}
	return "", ErrInvalidSubtitleOrDubbing
}

func (sd SubtitleOrDubbing) String() string {
	return string(sd)
}
