package repository

import (
	"context"

	"github.com/candy12t/cinemarch-server/domain/entity"
)

type Cinema interface {
	FindByID(ctx context.Context, cinemaID entity.UUID) (*entity.Cinema, error)
	Create(ctx context.Context, cinema *entity.Cinema) error
}
