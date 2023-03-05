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

type Movies []*Movie

func NewMovie(title MovieTitle, releaseDate time.Time, releaseStatus ReleaseStatus) *Movie {
	return &Movie{
		ID:            NewUUID(),
		Title:         title,
		ReleaseDate:   releaseDate,
		ReleaseStatus: releaseStatus,
	}
}

// ComingSoon -> NowOpen
// NowOpen    -> Released
// Released   -> NowOpen

func (m *Movie) ToNowOpen() error {
	if !m.isToNowOpen() || m.isBeforeReleaseDate() {
		return ErrNotChangeReleaseStatus
	}
	m.ReleaseStatus = NowOpen
	return nil
}

func (m *Movie) ToReleased() error {
	if !m.isToReleased() {
		return ErrNotChangeReleaseStatus
	}
	m.ReleaseStatus = Released
	return nil
}

func (m *Movie) UpdateReleaseDate(releaseDate time.Time) {
	m.ReleaseDate = releaseDate
}

func (m *Movie) isToNowOpen() bool {
	return m.ReleaseStatus == ComingSoon || m.ReleaseStatus == Released
}

func (m *Movie) isToReleased() bool {
	return m.ReleaseStatus == NowOpen
}

func (m *Movie) isBeforeReleaseDate() bool {
	return lib.TimeNow().Before(m.ReleaseDate)
}

type MovieTitle string

func NewMovieTitle(title string) (MovieTitle, error) {
	if title == "" || len([]rune(title)) > 128 {
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
