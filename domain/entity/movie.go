package entity

import (
	"time"

	"github.com/candy12t/cinemarch-server/lib"
)

type Movie struct {
	ID            UUID
	Title         MovieTitle
	ReleaseDate   time.Time
	ReleaseStatus ReleaseStatus
}

func NewMovie(title MovieTitle, releaseDate time.Time, releaseStatus ReleaseStatus) *Movie {
	return &Movie{
		ID:            NewUUID(),
		Title:         title,
		ReleaseDate:   releaseDate,
		ReleaseStatus: releaseStatus,
	}
}

func (m *Movie) ToNowOpen() error {
	if m.isBeforeReleaseDate() {
		return ErrNotAllowChangeMovieReleaseStatus
	}
	m.ReleaseStatus = NowOpen
	return nil
}

func (m *Movie) ToReleased() error {
	if m.ReleaseStatus != NowOpen {
		return ErrNotAllowChangeMovieReleaseStatus
	}
	m.ReleaseStatus = Released
	return nil
}

func (m *Movie) UpdateReleaseDate(releaseDate time.Time) {
	m.ReleaseDate = releaseDate
}

func (m *Movie) isBeforeReleaseDate() bool {
	return lib.TimeNow().Before(m.ReleaseDate)
}

type MovieTitle string

func NewMovieTitle(title string) (MovieTitle, error) {
	if title == "" || len([]rune(title)) > 255 {
		return "", ErrInvalidLengthMovieTitle
	}
	return MovieTitle(title), nil
}

func (m MovieTitle) String() string {
	return string(m)
}

type ReleaseStatus string

const (
	ComingSoon ReleaseStatus = "COMING SOON"
	NowOpen    ReleaseStatus = "NOW OPEN"
	Released   ReleaseStatus = "RELEASED"
)

var releaseStatuses = []ReleaseStatus{ComingSoon, NowOpen, Released}

func NewReleaseStatus(releaseStatus string) (ReleaseStatus, error) {
	for _, rs := range releaseStatuses {
		if rs.String() == releaseStatus {
			return rs, nil
		}
	}
	return "", ErrInvalidReleaseStatus
}

func (rs ReleaseStatus) String() string {
	return string(rs)
}

var dateFormat = "2006-01-02"

func NewMovieReleaseDate(releaseDate string) (time.Time, error) {
	date, err := parseDate(releaseDate)
	if err != nil {
		return time.Time{}, err
	}
	return date, nil
}

func parseDate(date string) (time.Time, error) {
	return time.Parse(dateFormat, date)
}
