package repository

import (
	"context"

	"github.com/candy12t/cinemarch-server/domain/entity"
)

type ScreeningType interface {
	FindByID(ctx context.Context, screeningTypeID entity.UUID) (*entity.ScreeningType, error)
	Create(ctx context.Context, screeningType *entity.ScreeningType) error
	Update(ctx context.Context, screeningType *entity.ScreeningType) error
}
