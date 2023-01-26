package entity_test

import (
	"errors"
	"testing"

	"github.com/candy12t/cinemarch-server/domain/entity"
)

func TestNewCinemaName(t *testing.T) {
	tests := []struct {
		name       string
		cinemaName string
		wantErr    error
	}{
		{
			name:       "valid cinema name",
			cinemaName: "TOHOシネマズ 新宿",
			wantErr:    nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := entity.NewCinemaName(tt.cinemaName)
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("NewCinemaName(): error is %v, want error is %v", err, tt.wantErr)
			}
		})
	}
}

func TestNewCinemaAddress(t *testing.T) {
	tests := []struct {
		name          string
		cinemaAddress string
		wantErr       error
	}{
		{
			name:          "valid cinema address",
			cinemaAddress: "東京都新宿区歌舞伎町 1-19-1　新宿東宝ビル３階",
			wantErr:       nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := entity.NewCinemaAddress(tt.cinemaAddress)
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("NewCinemaAddress(): error is %v, want error is %v", err, tt.wantErr)
			}
		})
	}
}

func TestNewCinemaURL(t *testing.T) {
	tests := []struct {
		name      string
		cinemaURL string
		wantErr   error
	}{
		{
			name:      "invalid cinema url",
			cinemaURL: "https://hlo.tohotheater.jp/net/schedule/076/TNPI2000J01.do",
			wantErr:   nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := entity.NewCinemaURL(tt.cinemaURL)
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("NewCinemaURL(): error is %v, want error is %v", err, tt.wantErr)
			}
		})
	}
}
