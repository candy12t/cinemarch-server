package entity

import "time"

type Movie struct {
	ID            UUID
	Title         MovieTitle
	ReleaseDate   time.Time
	ReleaseStatus ReleaseStatus
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

func NewMovie(title MovieTitle, releaseDate time.Time, releaseStatus ReleaseStatus) *Movie {
	now := time.Now()
	return &Movie{
		ID:            NewUUID(),
		Title:         title,
		ReleaseDate:   releaseDate,
		ReleaseStatus: releaseStatus,
		CreatedAt:     now,
		UpdatedAt:     now,
	}
}

func (m *Movie) Release() {
	m.ReleaseStatus = ReleaseStatusNowOpen
	m.UpdatedAt = time.Now()
}

func (m *Movie) Finish() {
	m.ReleaseStatus = ReleaseStatusReleased
	m.UpdatedAt = time.Now()
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
