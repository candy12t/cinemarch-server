package entity

import "time"

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

func (m *Movie) Release() {
	m.ReleaseStatus = ReleaseStatusNowOpen
}

func (m *Movie) Finish() {
	m.ReleaseStatus = ReleaseStatusReleased
}

type MovieTitle string

// TODO: validation
func NewMovieTitle(title string) (MovieTitle, error) {
	return MovieTitle(title), nil
}

type ReleaseStatus string

const (
	ReleaseStatusComingSoon ReleaseStatus = "COMING SOON"
	ReleaseStatusNowOpen    ReleaseStatus = "NOW OPEN"
	ReleaseStatusReleased   ReleaseStatus = "RELEASED"
)

var releaseStatuses = []ReleaseStatus{ReleaseStatusComingSoon, ReleaseStatusNowOpen, ReleaseStatusReleased}

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
