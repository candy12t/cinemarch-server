package entity_test

import (
	"errors"
	"testing"
	"time"

	"github.com/candy12t/cinema-search-server/domain/entity"
)

func TestUpdateReleaseStatusToNowOpen(t *testing.T) {
	movie := setupMovie()
	movie.Release()
	if movie.ReleaseStatus != entity.ReleaseStatusNowOpen {
		t.Errorf("Relase() got is %v, want is %v", movie.ReleaseStatus, entity.ReleaseStatusNowOpen)
	}
}

func TestUpdateReleaseStatusToReleased(t *testing.T) {
	movie := setupMovie()
	movie.Finish()
	if movie.ReleaseStatus != entity.ReleaseStatusReleased {
		t.Errorf("Finish() got is %v, want is %v", movie.ReleaseStatus, entity.ReleaseStatusReleased)
	}
}

func TestNewMovieTitle(t *testing.T) {
	tests := []struct {
		name       string
		movieTitle string
		wantErr    error
	}{
		{
			name:       "valid movie title",
			movieTitle: "valid movie title",
			wantErr:    nil,
		},
		{
			name:       "length is 0",
			movieTitle: "",
			wantErr:    entity.ErrInvalidLengthMovieTitle,
		},
		{
			name:       "length is 256",
			movieTitle: "I97IcsBUrALP1GiFLUVHXiUyvouJUgMQiJU0nK79lkRtudR7SK55pR0gx9SSwQyhb26BvYc7BV3MOI1nFaXdJFrO3vPvLSTEvXvwJgLHsbvjqsdNmtGSJcio8leAtFfQfOD6s7ZI8hjgoJed50TJCaqSd9I1YoD9SuRH9oeisOdPsc0aW4a5V5X3VfIRgtJLs01aEhByxLnAu0cgKUeD3rvIzQ9N3SY5wQeQLfWP4kqSR42e2rmxTlovGmyLAc5e",
			wantErr:    entity.ErrInvalidLengthMovieTitle,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := entity.NewMovieTitle(tt.movieTitle)
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("NewCinemaName() error is %v, wantErr is %v", err, tt.wantErr)
			}
		})
	}
}

func TestNewReleaseStatus(t *testing.T) {
	tests := []struct {
		name          string
		releaseStatus string
		want          entity.ReleaseStatus
	}{
		{
			name:          "coming soon",
			releaseStatus: "COMING SOON",
			want:          entity.ReleaseStatusComingSoon,
		},
		{
			name:          "now open",
			releaseStatus: "NOW OPEN",
			want:          entity.ReleaseStatusNowOpen,
		},
		{
			name:          "released",
			releaseStatus: "RELEASED",
			want:          entity.ReleaseStatusReleased,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := entity.NewReleaseStatus(tt.releaseStatus)
			if err != nil {
				t.Fatal(err)
			}
			if got != tt.want {
				t.Errorf("NewReleaseStatus() got is %v, want is %v", got, tt.want)
			}
		})
	}
}

func TestNewReleaseStatus_Invalid(t *testing.T) {
	tests := []struct {
		name          string
		releaseStatus string
		wantErr       error
	}{
		{
			name:          "invalid release status",
			releaseStatus: "INVALID",
			wantErr:       entity.ErrInvalidReleaseStatus,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := entity.NewReleaseStatus(tt.releaseStatus)
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("NewReleaseStatus() error is %v, wantErr is %v", err, tt.wantErr)
			}
		})
	}
}

func setupMovie() *entity.Movie {
	title, _ := entity.NewMovieTitle("hoge")
	now := time.Now()
	relaseStatus, _ := entity.NewReleaseStatus("COMING SOON")
	return entity.NewMovie(title, now, relaseStatus)
}
