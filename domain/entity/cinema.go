package entity

import "net/url"

type CinemaName string
type CinemaAddress string

type Cinema struct {
	ID      UUID
	Name    CinemaName
	Address CinemaAddress
	URL     *url.URL
}

func NewCinema(name CinemaName, address CinemaAddress, _url *url.URL) *Cinema {
	return &Cinema{
		ID:      NewUUID(),
		Name:    name,
		Address: address,
		URL:     _url,
	}
}

func NewCinemaName(name string) (CinemaName, error) {
	if name == "" || len([]rune(name)) > 255 {
		return "", ErrInvalidLengthCinemaName
	}
	return CinemaName(name), nil
}

func (cn CinemaName) String() string {
	return string(cn)
}

func NewCinemaAddress(address string) (CinemaAddress, error) {
	if address == "" || len([]rune(address)) > 255 {
		return "", ErrInvalidLengthCinemaAddress
	}
	return CinemaAddress(address), nil
}

func (ca CinemaAddress) String() string {
	return string(ca)
}

func NewCinemaURL(urlStr string) (*url.URL, error) {
	u, err := url.Parse(urlStr)
	if err != nil {
		return nil, err
	}

	if u.String() == "" || len([]rune(u.String())) > 255 {
		return nil, ErrInvalidLengthCinemaURL
	}
	return u, nil
}
