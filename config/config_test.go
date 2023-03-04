package config

import (
	"testing"
)

func TestDSN(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{
			name: "be able to get DSN",
			want: "cinemarch:password@tcp(127.0.0.1:3306)/cinemarch?parseTime=true",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := DSN()
			if got != tt.want {
				t.Errorf("DSN() got is %v, want is %v", got, tt.want)
			}
		})
	}
}

func TestHTTPPort(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{
			name: "be able to get HTTPPort",
			want: "8080",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := HTTPPort()
			if got != tt.want {
				t.Errorf("HTTPPort() got is %v, want is %v", got, tt.want)
			}
		})
	}
}
