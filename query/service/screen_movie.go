package service

import (
	"context"

	"github.com/candy12t/cinemarch-server/query/dto"
)

type ScreenMovie interface {
	Search(ctx context.Context, condition dto.ScreenMovieSearchCondition) (dto.ScreenMovies, error)
}
