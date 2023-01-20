package entity_test

import (
	"errors"
	"reflect"
	"testing"
	"time"

	"github.com/candy12t/cinema-search-server/domain/entity"
)

func TestNewMovieTitle(t *testing.T) {
	tests := []struct {
		name       string
		movieTitle string
		wantErr    error
	}{
		{
			name:       "valid movie title",
			movieTitle: "RRR",
			wantErr:    nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := entity.NewMovieTitle(tt.movieTitle)
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("NewMovieTitle(): error is %v, want error is %v", err, tt.wantErr)
			}
		})
	}
}

func TestNewReleaseStatus(t *testing.T) {
	tests := []struct {
		name          string
		releaseStatus string
		want          entity.MovieReleaseStatus
	}{
		{
			name:          "coming soon",
			releaseStatus: "COMING SOON",
			want:          entity.MovieReleaseStatusComingSoon,
		},
		{
			name:          "now open",
			releaseStatus: "NOW OPEN",
			want:          entity.MovieReleaseStatusNowOpen,
		},
		{
			name:          "released",
			releaseStatus: "RELEASED",
			want:          entity.MovieReleaseStatusReleased,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := entity.NewMovieReleaseStatus(tt.releaseStatus)
			if err != nil {
				t.Fatal(err)
			}
			if got != tt.want {
				t.Errorf("NewReleaseStatus(): got is %v, want is %v", got, tt.want)
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
			wantErr:       entity.ErrInvalidMovieReleaseStatus,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := entity.NewMovieReleaseStatus(tt.releaseStatus)
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("NewReleaseStatus(): error is %v, want error is %v", err, tt.wantErr)
			}
		})
	}
}

func TestUpdateMovieReleaseStatus(t *testing.T) {
	tests := []struct {
		name     string
		funcName string
		want     entity.MovieReleaseStatus
	}{
		{
			name:     "to NowOpen",
			funcName: "ToNowOpen",
			want:     entity.MovieReleaseStatusNowOpen,
		},
		{
			name:     "to Released",
			funcName: "ToReleased",
			want:     entity.MovieReleaseStatusReleased,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			movie := newMovie()

			r := reflect.ValueOf(movie)
			method := r.MethodByName(tt.funcName)
			method.Call(nil)

			if movie.ReleaseStatus != tt.want {
				t.Errorf("%s(): got is %v, want is %v", tt.funcName, movie.ReleaseStatus, tt.want)
			}
		})
	}
}

func newMovie() *entity.Movie {
	title, _ := entity.NewMovieTitle("RRR")
	jst, _ := time.LoadLocation("Asia/Tokyo")
	releaseDate := time.Date(2022, 10, 21, 0, 0, 0, 0, jst)
	return entity.NewMovie(title, releaseDate, entity.MovieReleaseStatusComingSoon)
}
