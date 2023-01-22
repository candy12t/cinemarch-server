package usecase

import (
	"context"
	"time"

	"github.com/candy12t/cinema-search-server/domain/entity"
	"github.com/candy12t/cinema-search-server/domain/repository"
)

type Movie interface {
	Show(ctx context.Context, movieID string) (*MovieDTO, error)
	Save(ctx context.Context, title string, releaseDate time.Time, releaseStatus string) (*MovieDTO, error)
}

type MovieUseCase struct {
	movieRepo repository.Movie
}

var _ Movie = (*MovieUseCase)(nil)

func NewMovieUseCase(movieRepo repository.Movie) *MovieUseCase {
	return &MovieUseCase{
		movieRepo: movieRepo,
	}
}

func (u *MovieUseCase) Show(ctx context.Context, movieID string) (*MovieDTO, error) {
	movie, err := u.movieRepo.FindByID(ctx, entity.UUID(movieID))
	if err != nil {
		return nil, err
	}
	return movieToDTO(movie), nil
}

func (u *MovieUseCase) Save(ctx context.Context, title string, releaseDate time.Time, releaseStatus string) (*MovieDTO, error) {
	movieTitle, err := entity.NewMovieTitle(title)
	if err != nil {
		return nil, err
	}

	movieReleaseStatus, err := entity.NewMovieReleaseStatus(releaseStatus)
	if err != nil {
		return nil, err
	}

	movie := entity.NewMovie(movieTitle, releaseDate, movieReleaseStatus)
	if err := u.movieRepo.Save(ctx, movie); err != nil {
		return nil, err
	}

	return movieToDTO(movie), nil
}

type MovieDTO struct {
	ID            string
	Title         string
	ReleaseDate   time.Time
	ReleaseStatus string
}

func movieToDTO(movie *entity.Movie) *MovieDTO {
	return &MovieDTO{
		ID:            movie.ID.String(),
		Title:         movie.Title.String(),
		ReleaseDate:   movie.ReleaseDate,
		ReleaseStatus: movie.ReleaseStatus.String(),
	}
}
