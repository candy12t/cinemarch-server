package repository

import (
	"context"

	"github.com/candy12t/cinema-search-server/domain/entity"
)

type ScreeningMovieScreeningType interface {
	FindByID(ctx context.Context, screeningMovieScreeningTypeID entity.UUID) (*entity.ScreeningMovieScreeningType, error)
	FindByScreeningMovieID(ctx context.Context, screeningMovieID entity.UUID) (*entity.ScreeningMovieScreeningType, error)
	FindByScreeningTypeID(ctx context.Context, screeningTypeID entity.UUID) (*entity.ScreeningMovieScreeningType, error)
	Save(ctx context.Context, screeningMovieScreeningType *entity.ScreeningMovieScreeningType) error
}