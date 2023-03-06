package entity

import "net/url"

type Cinema struct {
	ID         UUID
	Name       CinemaName
	Prefecture Prefecture
	Address    CinemaAddress
	WebSite    CinemaWebSite
}

type Cinemas []*Cinema

func NewCinema(name CinemaName, prefecture Prefecture, address CinemaAddress, webSite CinemaWebSite) *Cinema {
	return &Cinema{
		ID:         NewUUID(),
		Name:       name,
		Prefecture: prefecture,
		Address:    address,
		WebSite:    webSite,
	}
}

type CinemaName string

func NewCinemaName(name string) (CinemaName, error) {
	if name == "" || len([]rune(name)) > 128 {
		return "", ErrInvalidLengthCinemaName
	}
	return CinemaName(name), nil
}

func (c CinemaName) String() string {
	return string(c)
}

type CinemaAddress string

func NewCinemaAddress(address string) (CinemaAddress, error) {
	if address == "" || len([]rune(address)) > 128 {
		return "", ErrInvalidLengthCinemaAddress
	}
	return CinemaAddress(address), nil
}

func (c CinemaAddress) String() string {
	return string(c)
}

type CinemaWebSite string

func NewCinemaWebSite(webSite string) (CinemaWebSite, error) {
	if _, err := url.Parse(webSite); err != nil {
		return "", ErrInvalidFormatCinemaWebSite
	}
	if webSite == "" || len([]rune(webSite)) > 128 {
		return "", ErrInvalidLengthCinemaWebSite
	}
	return CinemaWebSite(webSite), nil
}

func (c CinemaWebSite) String() string {
	return string(c)
}
