package entity

import "time"

type MovieTitle string
type MovieReleaseStatus string

type Movie struct {
	ID            UUID
	Title         MovieTitle
	ReleaseDate   time.Time
	ReleaseStatus MovieReleaseStatus
}

const (
	MovieReleaseStatusComingSoon MovieReleaseStatus = "COMING SOON"
	MovieReleaseStatusNowOpen    MovieReleaseStatus = "NOW OPEN"
	MovieReleaseStatusReleased   MovieReleaseStatus = "RELEASED"
)

var releaseStatuses = []MovieReleaseStatus{
	MovieReleaseStatusComingSoon,
	MovieReleaseStatusNowOpen,
	MovieReleaseStatusReleased,
}

func NewMovie(title MovieTitle, releaseDate time.Time, releaseStatus MovieReleaseStatus) *Movie {
	return &Movie{
		ID:            NewUUID(),
		Title:         title,
		ReleaseDate:   releaseDate,
		ReleaseStatus: releaseStatus,
	}
}

func (m *Movie) ToNowOpen() {
	m.ReleaseStatus = MovieReleaseStatusNowOpen
}

func (m *Movie) ToReleased() {
	m.ReleaseStatus = MovieReleaseStatusReleased
}

func NewMovieTitle(title string) (MovieTitle, error) {
	if title == "" || len([]rune(title)) > 255 {
		return "", ErrInvalidLengthMovieTitle
	}
	return MovieTitle(title), nil
}

func (mt MovieTitle) String() string {
	return string(mt)
}

func NewMovieReleaseStatus(releaseStatus string) (MovieReleaseStatus, error) {
	for _, rs := range releaseStatuses {
		if rs.String() == releaseStatus {
			return rs, nil
		}
	}
	return "", ErrInvalidMovieReleaseStatus
}

func (mrs MovieReleaseStatus) String() string {
	return string(mrs)
}
