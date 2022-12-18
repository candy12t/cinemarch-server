package repository

import (
	"context"

	"github.com/candy12t/cinema-search-server/domain/entity"
)

type Movie interface {
	FindByID(ctx context.Context, movieID entity.UUID) (*entity.Movie, error)
	Save(ctx context.Context, movie *entity.Movie) error
	Update(ctx context.Context, movie *entity.Movie) error
}
