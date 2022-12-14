package entity

import (
	"net/url"
	"time"
)

type CinemaName string

type Prefecture string

type CinemaAddress string

type Cinema struct {
	ID         UUID
	Name       CinemaName
	Prefecture Prefecture
	Address    CinemaAddress
	URL        *url.URL
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

func NewCinema(name CinemaName, prefecture Prefecture, address CinemaAddress, _url *url.URL) *Cinema {
	now := time.Now()
	return &Cinema{
		ID:         NewUUID(),
		Name:       name,
		Prefecture: prefecture,
		Address:    address,
		URL:        _url,
		CreatedAt:  now,
		UpdatedAt:  now,
	}
}

// TODO: validation
func NewCinemaName(name string) (CinemaName, error) {
	return CinemaName(name), nil
}

// TODO: validation
func NewCinemaPrefecture(prefecture string) (Prefecture, error) {
	return Prefecture(prefecture), nil
}

// TODO: validation
func NewCinemaAddress(address string) (CinemaAddress, error) {
	return CinemaAddress(address), nil
}

func NewCinemaURL(_url string) (*url.URL, error) {
	return url.Parse(_url)
}
