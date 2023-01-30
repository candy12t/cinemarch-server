//go:generate mockgen -source=$GOFILE -destination=../mock_$GOPACKAGE/$GOFILE
package repository

import (
	"context"

	"github.com/candy12t/cinemarch-server/domain/entity"
)

type ScreeningMovie interface {
	FindByID(ctx context.Context, screeningMovieID entity.UUID) (*entity.ScreeningMovie, error)
	Create(ctx context.Context, screeningMovie *entity.ScreeningMovie) error
}
