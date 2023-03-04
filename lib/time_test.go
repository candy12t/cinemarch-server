package lib

import (
	"errors"
	"testing"
	"time"
)

func TestParseJSTDateInUTC(t *testing.T) {
	tests := []struct {
		name    string
		value   string
		want    time.Time
		wantErr error
	}{
		{
			name:    "2023-01-01",
			value:   "2023-01-01",
			want:    time.Date(2022, 12, 31, 15, 0, 0, 0, time.UTC),
			wantErr: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseJSTDateInUTC(tt.value)
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("ParseJSTDateInUTC() error is %v, wantErr %v", err, tt.want)
			}
			if got != tt.want {
				t.Errorf("ParseJSTDateInUTC() got is %v, want is %v", got, tt.want)
			}
		})
	}
}

func TestParseJSTDateTimeInUTC(t *testing.T) {
	tests := []struct {
		name    string
		value   string
		want    time.Time
		wantErr error
	}{
		{
			name:    "2023-01-01 15:04",
			value:   "2023-01-01 15:04",
			want:    time.Date(2023, 1, 1, 6, 4, 0, 0, time.UTC),
			wantErr: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseJSTDateTimeInUTC(tt.value)
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("ParseJSTDateTimeInUTC() error is %v, wantErr %v", err, tt.want)
			}
			if got != tt.want {
				t.Errorf("ParseJSTDateTimeInUTC() got is %v, want is %v", got, tt.want)
			}
		})
	}
}

func TestFormatDateInJST(t *testing.T) {
	tests := []struct {
		name  string
		value time.Time
		want  string
	}{
		{
			name:  "2023-01-01",
			value: time.Date(2022, 12, 31, 15, 0, 0, 0, time.UTC),
			want:  "2023-01-01",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FormatDateInJST(tt.value)
			if got != tt.want {
				t.Errorf("FormatDateInJST() got is %v, want is %v", got, tt.want)
			}
		})
	}
}

func TestFormatDateTimeInJST(t *testing.T) {
	tests := []struct {
		name  string
		value time.Time
		want  string
	}{
		{
			name:  "2023-01-01 15:04",
			value: time.Date(2023, 1, 1, 6, 4, 0, 0, time.UTC),
			want:  "2023-01-01 15:04",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FormatDateTimeInJST(tt.value)
			if got != tt.want {
				t.Errorf("FormatDateTimeInJST() got is %v, want is %v", got, tt.want)
			}
		})
	}
}
