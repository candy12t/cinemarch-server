package entity

// TODO: generate from config file
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
