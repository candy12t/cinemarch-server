package entity_test

import (
	"errors"
	"testing"

	"github.com/candy12t/cinemarch-server/domain/entity"
)

func TestNewScreenTypeName(t *testing.T) {
	tests := []struct {
		name              string
		screenTypeName string
		wantErr           error
	}{
		{
			name:              "valid length name ascii",
			screenTypeName: "IMAX",
			wantErr:           nil,
		},
		{
			name:              "valid length name mulit byte",
			screenTypeName: "IMAXレーザー",
			wantErr:           nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := entity.NewScreenTypeName(tt.screenTypeName)
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("NewScreenTypeName() error is %v, want error is %v", err, tt.wantErr)
			}
		})
	}
}
