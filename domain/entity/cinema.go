package entity

import (
	"net/url"
)

type CinemaName string

type Prefecture string

type CinemaAddress string

type CinemaURL string

type Cinema struct {
	ID         UUID
	Name       CinemaName
	Prefecture Prefecture
	Address    CinemaAddress
	URL        CinemaURL
}

func NewCinema(name CinemaName, prefecture Prefecture, address CinemaAddress, _url CinemaURL) *Cinema {
	return &Cinema{
		ID:         NewUUID(),
		Name:       name,
		Prefecture: prefecture,
		Address:    address,
		URL:        _url,
	}
}

func NewCinemaName(name string) (CinemaName, error) {
	if name == "" || len([]rune(name)) > 255 {
		return "", ErrInvalidLengthMovieTitle
	}
	return CinemaName(name), nil
}

// TODO: define enum
func NewCinemaPrefecture(prefecture string) (Prefecture, error) {
	if prefecture == "" || len([]rune(prefecture)) > 255 {
		return "", ErrInvalidLengthPrefecture
	}
	return Prefecture(prefecture), nil
}

func NewCinemaAddress(address string) (CinemaAddress, error) {
	if address == "" || len([]rune(address)) > 255 {
		return "", ErrInvalidLengthAddress
	}
	return CinemaAddress(address), nil
}

func NewCinemaURL(_url string) (CinemaURL, error) {
	u, err := url.Parse(_url)
	if err != nil {
		return "", err
	}

	if u.String() == "" || len([]rune(u.String())) > 255 {
		return "", ErrInvalidLengthAddress
	}
	return CinemaURL(u.String()), nil
}
