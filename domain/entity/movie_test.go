package entity_test

import (
	"errors"
	"testing"
	"time"

	"github.com/candy12t/cinemarch-server/domain/entity"
	"github.com/candy12t/cinemarch-server/lib"
)

func TestNewMovieTitle(t *testing.T) {
	tests := []struct {
		name    string
		title   string
		want    entity.MovieTitle
		wantErr error
	}{
		{
			name:    "valid movie title",
			title:   "valid movie ttile",
			want:    entity.MovieTitle("valid movie ttile"),
			wantErr: nil,
		},
		{
			name:    "invalid movie title",
			title:   "",
			want:    "",
			wantErr: entity.ErrInvalidLengthMovieTitle,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := entity.NewMovieTitle(tt.title)
			if err != nil && !errors.Is(err, tt.wantErr) {
				t.Errorf("NewMovieTitle() error is %v, want error is %v", err, tt.wantErr)
			}
			if got != tt.want {
				t.Errorf("NewMovieTitle() got is %v, want is %v", got, tt.want)
			}
		})
	}
}

func TestNewReleaseStatus(t *testing.T) {
	tests := []struct {
		name          string
		releaseStatus string
		want          entity.ReleaseStatus
		wantErr       error
	}{
		{
			name:          "coming soon",
			releaseStatus: "COMING SOON",
			want:          entity.ComingSoon,
			wantErr:       nil,
		},
		{
			name:          "now open",
			releaseStatus: "NOW OPEN",
			want:          entity.NowOpen,
			wantErr:       nil,
		},
		{
			name:          "released",
			releaseStatus: "RELEASED",
			want:          entity.Released,
			wantErr:       nil,
		},
		{
			name:          "invalid release status",
			releaseStatus: "INVALID",
			want:          "",
			wantErr:       entity.ErrInvalidReleaseStatus,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := entity.NewReleaseStatus(tt.releaseStatus)
			if err != nil && !errors.Is(err, tt.wantErr) {
				t.Errorf("NewReleaseStatus() error is %v, want error is %v", err, tt.wantErr)
			}
			if got != tt.want {
				t.Errorf("NewReleaseStatus() got is %v, want is %v", got, tt.want)
			}
		})
	}
}

func TestToNowOpen(t *testing.T) {
	lib.TimeNow = func() time.Time {
		return time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC)
	}

	tests := []struct {
		name    string
		movie   *entity.Movie
		wantErr error
	}{
		{
			name: "change release status from `COMING SOON`",
			movie: &entity.Movie{
				ReleaseDate:   time.Date(2022, 12, 31, 0, 0, 0, 0, time.UTC),
				ReleaseStatus: entity.ComingSoon,
			},
			wantErr: nil,
		},
		{
			name: "change release status from `RELEASED`",
			movie: &entity.Movie{
				ReleaseDate:   time.Date(2022, 12, 31, 0, 0, 0, 0, time.UTC),
				ReleaseStatus: entity.Released,
			},
			wantErr: nil,
		},
		{
			name: "not change release status because ReleaseDate is before",
			movie: &entity.Movie{
				ReleaseDate:   time.Date(2023, 12, 31, 0, 0, 0, 0, time.UTC),
				ReleaseStatus: entity.ComingSoon,
			},
			wantErr: entity.ErrNotAllowChangeMovieReleaseStatus,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.movie.ToNowOpen(); err != nil {
				if !errors.Is(err, tt.wantErr) {
					t.Errorf("ToNowOpen() error is %v, want error is %v", err, tt.wantErr)
				}
				return
			}

			if tt.movie.ReleaseStatus != entity.NowOpen {
				t.Errorf("TestToNowOpen() got is %v, want is %v", tt.movie.ReleaseStatus, entity.NowOpen)
			}
		})
	}
}

func TestToReleased(t *testing.T) {
	tests := []struct {
		name    string
		movie   *entity.Movie
		wantErr error
	}{
		{
			name: "now open",
			movie: &entity.Movie{
				ReleaseStatus: entity.NowOpen,
			},
			wantErr: nil,
		},
		{
			name: "coming soon",
			movie: &entity.Movie{
				ReleaseStatus: entity.ComingSoon,
			},
			wantErr: entity.ErrNotAllowChangeMovieReleaseStatus,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.movie.ToReleased(); err != nil {
				if !errors.Is(err, tt.wantErr) {
					t.Errorf("ToReleased() error is %v, want error is %v", err, tt.wantErr)
				}
				return
			}

			if tt.movie.ReleaseStatus != entity.Released {
				t.Errorf("ToReleased() is %v, want %v", tt.movie.ReleaseStatus, entity.Released)
			}
		})
	}
}
