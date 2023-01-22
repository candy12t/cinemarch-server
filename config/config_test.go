package config

import "testing"

func TestDSN(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{
			name: "be able to read environment variable to configure DSN",
			want: "cs:password@tcp(127.0.0.1:3306)/cinema-search?parseTime=true",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := DSN()
			if got != tt.want {
				t.Errorf("DNS() got is %v, want %v", got, tt.want)
			}
		})
	}
}
