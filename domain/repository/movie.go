//go:generate mockgen -source=$GOFILE -destination=../mock_$GOPACKAGE/$GOFILE
package repository

import (
	"context"

	"github.com/candy12t/cinemarch-server/domain/entity"
)

type Movie interface {
	FindByID(ctx context.Context, movieID entity.UUID) (*entity.Movie, error)
	FindByTitle(ctx context.Context, title entity.MovieTitle) (*entity.Movie, error)
	Search(ctx context.Context, conditionQuery string, args []any) (entity.Movies, error)
	Create(ctx context.Context, movie *entity.Movie) error
	Update(ctx context.Context, movie *entity.Movie) error
}
