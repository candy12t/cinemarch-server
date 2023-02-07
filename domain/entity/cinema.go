package entity

import "net/url"

type Cinema struct {
	ID         UUID
	Name       CinemaName
	Prefecture Prefecture
	Address    Address
	WebSiteURL WebSiteURL
}

func NewCinema(name CinemaName, prefecure Prefecture, address Address, webSiteURL WebSiteURL) *Cinema {
	return &Cinema{
		ID:         NewUUID(),
		Name:       name,
		Prefecture: prefecure,
		Address:    address,
		WebSiteURL: webSiteURL,
	}
}

type CinemaName string

func NewCinemaName(name string) (CinemaName, error) {
	if name == "" || len([]rune(name)) > 64 {
		return "", ErrInvalidLengthCinemaName
	}
	return CinemaName(name), nil
}

func (c CinemaName) String() string {
	return string(c)
}

type Address string

func NewAddress(address string) (Address, error) {
	if address == "" || len([]rune(address)) > 255 {
		return "", ErrInvalidLengthAddress
	}
	return Address(address), nil
}

func (a Address) String() string {
	return string(a)
}

type WebSiteURL string

func NewWebSiteURL(urlStr string) (WebSiteURL, error) {
	u, err := url.Parse(urlStr)
	if err != nil {
		return "", err
	}

	if u.String() == "" || len([]rune(u.String())) > 255 {
		return "", ErrInvalidLengthWebSiteURL
	}
	return WebSiteURL(u.String()), nil
}

func (w WebSiteURL) String() string {
	return string(w)
}
