package usecase

import (
	"context"
	"errors"

	"github.com/candy12t/cinemarch-server/domain/entity"
	"github.com/candy12t/cinemarch-server/domain/repository"
	"github.com/candy12t/cinemarch-server/lib"
)

type Movie interface {
	FindByID(ctx context.Context, movieID string) (*MovieDTO, error)
	FindAllByTitle(ctx context.Context, title string) (MovieDTOs, error)
	Upsert(ctx context.Context, params UpsertMovieParams) (*MovieDTO, error)
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

func (u *MovieUseCase) FindByID(ctx context.Context, movieID string) (*MovieDTO, error) {
	movie, err := u.movieRepo.FindByID(ctx, entity.UUID(movieID))
	if err != nil {
		return nil, err
	}
	return movieToDTO(movie), nil
}

func (u *MovieUseCase) FindAllByTitle(ctx context.Context, title string) (MovieDTOs, error) {
	conditions := entity.Conditions{{Query: "title LIKE ?", Arg: "%" + title + "%"}}
	query, args := conditions.Build()

	movies, err := u.movieRepo.Search(ctx, query, args)
	if err != nil {
		return nil, err
	}

	movieDTOs := make([]*MovieDTO, 0, len(movies))
	for _, movie := range movies {
		movieDTOs = append(movieDTOs, movieToDTO(movie))
	}
	return movieDTOs, nil
}

func (u *MovieUseCase) Upsert(ctx context.Context, params UpsertMovieParams) (*MovieDTO, error) {
	releaseDate, err := lib.ParseJSTDateInUTC(params.ReleaseDate)
	if err != nil {
		return nil, err
	}

	movieTitle, err := entity.NewMovieTitle(params.Title)
	if err != nil {
		return nil, err
	}

	releaseStatus, err := entity.NewReleaseStatus(params.ReleaseStatus)
	if err != nil {
		return nil, err
	}

	movie, err := u.movieRepo.FindByTitle(ctx, movieTitle)
	if err != nil {
		if errors.Is(err, entity.ErrMovieNotFound) {
			movie := entity.NewMovie(movieTitle, releaseDate, releaseStatus)
			if err := u.movieRepo.Create(ctx, movie); err != nil {
				return nil, err
			}
			return movieToDTO(movie), nil
		}
		return nil, err
	}

	switch releaseStatus {
	case entity.ComingSoon:
		return nil, entity.ErrNotChangeReleaseStatus
	case entity.NowOpen:
		if err := movie.ToNowOpen(); err != nil {
			return nil, err
		}
	case entity.Released:
		if err := movie.ToReleased(); err != nil {
			return nil, err
		}
	}
	movie.UpdateReleaseDate(releaseDate)
	if err := u.movieRepo.Update(ctx, movie); err != nil {
		return nil, err
	}

	return movieToDTO(movie), nil
}

type UpsertMovieParams struct {
	Title         string
	ReleaseDate   string
	ReleaseStatus string
}

type MovieDTO struct {
	ID            string
	Title         string
	ReleaseDate   string
	ReleaseStatus string
}

type MovieDTOs []*MovieDTO

func movieToDTO(movie *entity.Movie) *MovieDTO {
	return &MovieDTO{
		ID:            movie.ID.String(),
		Title:         movie.Title.String(),
		ReleaseDate:   lib.FormatDateInJST(movie.ReleaseDate),
		ReleaseStatus: movie.ReleaseStatus.String(),
	}
}
