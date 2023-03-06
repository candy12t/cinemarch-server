package entity

import (
	"errors"
	"testing"
)

func TestNewCinemaName(t *testing.T) {
	tests := []struct {
		name       string
		cinemaName string
		wantErr    error
	}{
		{
			name:       "valid cinema name",
			cinemaName: "cinema name",
			wantErr:    nil,
		},
		{
			name:       "valid cinema name, length is 128",
			cinemaName: randomString(128),
			wantErr:    nil,
		},
		{
			name:       "invalid cinema name, length is 0",
			cinemaName: "",
			wantErr:    ErrInvalidLengthCinemaName,
		},
		{
			name:       "invalid cinema name, length is 129",
			cinemaName: randomString(129),
			wantErr:    ErrInvalidLengthCinemaName,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := NewCinemaName(tt.cinemaName)
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("NewCinemaName() error is %v, wantErr is %v", err, tt.wantErr)
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
			cinemaAddress: "東京都新宿区新宿1-1-1",
			wantErr:       nil,
		},
		{
			name:          "valid cinema address, length is 128",
			cinemaAddress: randomString(128),
			wantErr:       nil,
		},
		{
			name:          "invalid cinema address, length is 0",
			cinemaAddress: "",
			wantErr:       ErrInvalidLengthCinemaAddress,
		},
		{
			name:          "invalid cinema address, length is 129",
			cinemaAddress: randomString(129),
			wantErr:       ErrInvalidLengthCinemaAddress,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := NewCinemaAddress(tt.cinemaAddress)
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("NewCinemaAddress() error is %v, wantErr is %v", err, tt.wantErr)
			}
		})
	}
}

func TestNewCinemaWebSite(t *testing.T) {
	tests := []struct {
		name          string
		cinemaWebSite string
		wantErr       error
	}{
		{
			name:          "valid cinema web site",
			cinemaWebSite: "https://example.com",
			wantErr:       nil,
		},
		{
			name:          "valid cinema web site, length is 128",
			cinemaWebSite: randomString(128),
			wantErr:       nil,
		},
		{
			name:          "invalid cinema web site, length is 0",
			cinemaWebSite: "",
			wantErr:       ErrInvalidLengthCinemaWebSite,
		},
		{
			name:          "invalid cinema web site, length is 129",
			cinemaWebSite: randomString(129),
			wantErr:       ErrInvalidLengthCinemaWebSite,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := NewCinemaWebSite(tt.cinemaWebSite)
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("NewCinemaWebSite() error is %v, wantErr is %v", err, tt.wantErr)
			}
		})
	}
}
