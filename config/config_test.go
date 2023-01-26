package config_test

import (
	"testing"

	"github.com/candy12t/cinemarch-server/config"
)

func TestDSN(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{
			name: "be able to read environment variable to configure DSN",
			want: "cinemarch:password@tcp(127.0.0.1:3306)/cinemarch?parseTime=true",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := config.DSN()
			if got != tt.want {
				t.Errorf("DNS() got is %v, want %v", got, tt.want)
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
			name: "be able to get HTTP_PORT from environment variable",
			want: "8080",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := config.HTTPPort()
			if got != tt.want {
				t.Errorf("HTTPPort() got is %v, want %v", got, tt.want)
			}
		})
	}
}
