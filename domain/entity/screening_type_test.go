package entity_test

import (
	"errors"
	"testing"

	"github.com/candy12t/cinemarch-server/domain/entity"
)

func TestNewScreeningTypeName(t *testing.T) {
	tests := []struct {
		name              string
		screeningTypeName string
		wantErr           error
	}{
		{
			name:              "valid length name ascii",
			screeningTypeName: "IMAX",
			wantErr:           nil,
		},
		{
			name:              "valid length name mulit byte",
			screeningTypeName: "IMAXレーザー",
			wantErr:           nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := entity.NewScreeningTypeName(tt.screeningTypeName)
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("NewScreeningTypeName(): error is %v, want error is %v", err, tt.wantErr)
			}
		})
	}
}
