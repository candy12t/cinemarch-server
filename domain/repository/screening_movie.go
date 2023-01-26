package repository

import (
	"context"

	"github.com/candy12t/cinemarch-server/domain/entity"
)

type ScreeningMovie interface {
	FindByID(ctx context.Context, screeningMovieID entity.UUID) (*entity.ScreeningMovie, error)
	Save(ctx context.Context, screeningMovie *entity.ScreeningMovie) error
	Update(ctx context.Context, screeningMovie *entity.ScreeningMovie) error
}
