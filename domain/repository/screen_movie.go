//go:generate mockgen -source=$GOFILE -destination=../mock_$GOPACKAGE/$GOFILE
package repository

import (
	"context"

	"github.com/candy12t/cinemarch-server/domain/entity"
)

type ScreenMovie interface {
	FindByID(ctx context.Context, screenMovieID entity.UUID) (*entity.ScreenMovie, error)
	Create(ctx context.Context, screenMovie *entity.ScreenMovie) error
}
