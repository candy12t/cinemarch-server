//go:generate mockgen -source=$GOFILE -destination=../mock_$GOPACKAGE/$GOFILE
package repository

import (
	"context"

	"github.com/candy12t/cinemarch-server/domain/entity"
)

type ScreenType interface {
	FindByID(ctx context.Context, screenTypeID entity.UUID) (*entity.ScreenType, error)
	Create(ctx context.Context, screenType *entity.ScreenType) error
}
