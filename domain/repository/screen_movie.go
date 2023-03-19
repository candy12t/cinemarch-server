//go:generate mockgen -source=$GOFILE -destination=../mock_$GOPACKAGE/$GOFILE
package repository

import (
	"context"

	"github.com/candy12t/cinemarch-server/domain/entity"
)

type ScreenMovie interface {
	FindByUniqKey(ctx context.Context, cinemaID, movieID entity.UUID, screenType entity.ScreenType, translateType entity.TranslateType, threeD bool) (*entity.ScreenMovie, error)
	Search(ctx context.Context) (entity.ScreenMovies, error)
	CreateScreenMovie(ctx context.Context, screenMovie *entity.ScreenMovie) error
	CreateScreenSchedules(ctx context.Context, screenSchedule entity.ScreenSchedules) error
}
