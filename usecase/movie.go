package usecase

import (
	"context"
	"time"

	"github.com/candy12t/cinemarch-server/domain/entity"
	"github.com/candy12t/cinemarch-server/domain/repository"
)

type Movie interface {
	Show(ctx context.Context, movieID string) (*MovieDTO, error)
	Create(ctx context.Context, params CreateMovieParams) (*MovieDTO, error)
	Update(ctx context.Context, movieID string, params UpdateMovieParams) (*MovieDTO, error)
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

func (u *MovieUseCase) Create(ctx context.Context, params CreateMovieParams) (*MovieDTO, error) {
	movieTitle, err := entity.NewMovieTitle(params.Title)
	if err != nil {
		return nil, err
	}

	movieReleaseDate, err := entity.NewMovieReleaseDate(params.ReleaseDate)
	if err != nil {
		return nil, err
	}

	movieReleaseStatus, err := entity.NewReleaseStatus(params.ReleaseStatus)
	if err != nil {
		return nil, err
	}

	movie := entity.NewMovie(movieTitle, movieReleaseDate, movieReleaseStatus)
	if err := u.movieRepo.Create(ctx, movie); err != nil {
		return nil, err
	}

	return movieToDTO(movie), nil
}

func (u *MovieUseCase) Update(ctx context.Context, movieID string, params UpdateMovieParams) (*MovieDTO, error) {
	movie, err := u.movieRepo.FindByID(ctx, entity.UUID(movieID))
	if err != nil {
		return nil, err
	}

	releaseDate, err := entity.NewMovieReleaseDate(params.ReleaseDate)
	if err != nil {
		return nil, err
	}
	movie.UpdateReleaseDate(releaseDate)

	releaseStatus, err := entity.NewReleaseStatus(params.ReleaseStatus)
	if err != nil {
		return nil, err
	}
	switch releaseStatus {
	case entity.NowOpen:
		movie.ToNowOpen()
	case entity.Released:
		movie.ToReleased()
	}

	if err := u.movieRepo.Update(ctx, movie); err != nil {
		return nil, err
	}
	return movieToDTO(movie), nil
}

type CreateMovieParams struct {
	Title         string
	ReleaseDate   string
	ReleaseStatus string
}

type UpdateMovieParams struct {
	ReleaseDate   string
	ReleaseStatus string
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
