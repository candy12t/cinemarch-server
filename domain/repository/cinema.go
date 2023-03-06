//go:generate mockgen -source=$GOFILE -destination=../mock_$GOPACKAGE/$GOFILE
package repository

import (
	"context"

	"github.com/candy12t/cinemarch-server/domain/entity"
)

type Cinema interface {
	FindByID(ctx context.Context, cinemaID entity.UUID) (*entity.Cinema, error)
	FindAllByPrefecture(ctx context.Context, prefecture entity.Prefecture) (entity.Cinemas, error)
	Create(ctx context.Context, cinema *entity.Cinema) error
}
