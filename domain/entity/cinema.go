package entity

import "net/url"

type Cinema struct {
	ID      UUID
	Name    CinemaName
	Address CinemaAddress
	URL     CinemaURL
}

func NewCinema(name CinemaName, address CinemaAddress, _url CinemaURL) *Cinema {
	return &Cinema{
		ID:      NewUUID(),
		Name:    name,
		Address: address,
		URL:     _url,
	}
}

type CinemaName string

func NewCinemaName(name string) (CinemaName, error) {
	if name == "" || len([]rune(name)) > 255 {
		return "", ErrInvalidLengthCinemaName
	}
	return CinemaName(name), nil
}

func (cn CinemaName) String() string {
	return string(cn)
}

type CinemaAddress string

func NewCinemaAddress(address string) (CinemaAddress, error) {
	if address == "" || len([]rune(address)) > 255 {
		return "", ErrInvalidLengthCinemaAddress
	}
	return CinemaAddress(address), nil
}

func (ca CinemaAddress) String() string {
	return string(ca)
}

type CinemaURL string

func NewCinemaURL(urlStr string) (CinemaURL, error) {
	u, err := url.Parse(urlStr)
	if err != nil {
		return "", err
	}

	if u.String() == "" || len([]rune(u.String())) > 255 {
		return "", ErrInvalidLengthCinemaURL
	}
	return CinemaURL(u.String()), nil
}

func (c CinemaURL) String() string {
	return string(c)
}
