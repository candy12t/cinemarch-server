package entity

import (
	"errors"
	"testing"
	"time"

	"github.com/candy12t/cinemarch-server/lib"
)

func TestToNowOpen(t *testing.T) {
	lib.TimeNow = func() time.Time {
		return time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC)
	}

	tests := []struct {
		name    string
		movie   *Movie
		wantErr error
	}{
		{
			name: "to NowOen from ComingSoon",
			movie: &Movie{
				ReleaseDate:   time.Date(2022, 12, 31, 0, 0, 0, 0, time.UTC),
				ReleaseStatus: ComingSoon,
			},
			wantErr: nil,
		},
		{
			name: "to NowOpen from Released",
			movie: &Movie{
				ReleaseDate:   time.Date(2022, 12, 31, 0, 0, 0, 0, time.UTC),
				ReleaseStatus: Released,
			},
			wantErr: nil,
		},
		{
			name: "not change to NowOpen, ReleaseDate is before",
			movie: &Movie{
				ReleaseDate:   time.Date(2023, 2, 1, 0, 0, 0, 0, time.UTC),
				ReleaseStatus: ComingSoon,
			},
			wantErr: ErrNotChangeReleaseStatus,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.movie.ToNowOpen(); !errors.Is(err, tt.wantErr) {
				t.Errorf("ToNowOpen() error is %v, wantErr is %v", err, tt.wantErr)
			}
		})
	}
}

func TestToReleased(t *testing.T) {
	tests := []struct {
		name    string
		movie   *Movie
		wantErr error
	}{
		{
			name: "to Released from NowOpen",
			movie: &Movie{
				ReleaseStatus: NowOpen,
			},
			wantErr: nil,
		},
		{
			name: "not chagne to Released, ReleaseStatus is ComingSoon",
			movie: &Movie{
				ReleaseStatus: ComingSoon,
			},
			wantErr: ErrNotChangeReleaseStatus,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.movie.ToReleased(); !errors.Is(err, tt.wantErr) {
				t.Errorf("ToReleased() error is %v, wantErr is %v", err, tt.wantErr)
			}
		})
	}
}

func TestMovieTitle(t *testing.T) {
	tests := []struct {
		name    string
		title   string
		wantErr error
	}{
		{
			name:    "valid movie title",
			title:   "movie title",
			wantErr: nil,
		},
		{
			name:    "valid movie title, lenght is 128",
			title:   randomString(128),
			wantErr: nil,
		},
		{
			name:    "invalid movie title, length is 0",
			title:   "",
			wantErr: ErrInvalidLengthMovieTitle,
		},
		{
			name:    "invalid movie title, length is 129",
			title:   randomString(129),
			wantErr: ErrInvalidLengthMovieTitle,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := NewMovieTitle(tt.title)
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("NewMovieTitle() error is %v, wantErr is %v", err, tt.wantErr)
			}
		})
	}
}

func TestNewReleaseStatus(t *testing.T) {
	tests := []struct {
		name          string
		releaseStatus string
		want          ReleaseStatus
		wantErr       error
	}{
		{
			name:          "coming soon",
			releaseStatus: "COMING SOON",
			want:          ComingSoon,
			wantErr:       nil,
		},
		{
			name:          "now open",
			releaseStatus: "NOW OPEN",
			want:          NowOpen,
			wantErr:       nil,
		},
		{
			name:          "released",
			releaseStatus: "RELEASED",
			want:          Released,
			wantErr:       nil,
		},
		{
			name:          "invalid release status",
			releaseStatus: "INVALID RELEASE STATUS",
			want:          "",
			wantErr:       ErrInvalidReleaseStatus,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewReleaseStatus(tt.releaseStatus)
			if err != nil {
				if !errors.Is(err, tt.wantErr) {
					t.Errorf("NewReleaseStatus() error is %v, wantErr is %v", err, tt.wantErr)
				}
				return
			}
			if got != tt.want {
				t.Errorf("NewReleaseStatus() got is %v, want %v", got, tt.want)
			}
		})
	}
}
