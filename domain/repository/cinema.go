package repository

import (
	"context"

	"github.com/candy12t/cinema-search-server/domain/entity"
)

type Cinema interface {
	FindByID(ctx context.Context, cinemaID entity.UUID) (*entity.Cinema, error)
	Save(ctx context.Context, cinema *entity.Cinema) error
	Update(ctx context.Context, cinema *entity.Cinema) error
}