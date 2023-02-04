//go:generate mockgen -source=$GOFILE -destination=../mock_$GOPACKAGE/$GOFILE
package repository

import (
	"context"

	"github.com/candy12t/cinemarch-server/domain/entity"
)

type ScreenMovieScreenType interface {
	FindByID(ctx context.Context, screenMovieScreenTypeID entity.UUID) (*entity.ScreenMovieScreenType, error)
	FindByScreenMovieID(ctx context.Context, screenMovieID entity.UUID) (*entity.ScreenMovieScreenType, error)
	FindByScreenTypeID(ctx context.Context, screenTypeID entity.UUID) (*entity.ScreenMovieScreenType, error)
	Create(ctx context.Context, screenMovieScreenType *entity.ScreenMovieScreenType) error
}
