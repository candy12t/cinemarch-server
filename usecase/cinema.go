package usecase

import (
	"context"

	"github.com/candy12t/cinemarch-server/domain/entity"
	"github.com/candy12t/cinemarch-server/domain/repository"
)

type Cinema interface {
	Show(ctx context.Context, cinemaID string) (*CinemaDTO, error)
	Create(ctx context.Context, params CreateCinemaParams) (*CinemaDTO, error)
}

type CinemaUseCase struct {
	cinemaRepo repository.Cinema
}

var _ Cinema = (*CinemaUseCase)(nil)

func NewCinemaUseCase(cinemaRepo repository.Cinema) *CinemaUseCase {
	return &CinemaUseCase{
		cinemaRepo: cinemaRepo,
	}
}

func (u *CinemaUseCase) Show(ctx context.Context, cinemaID string) (*CinemaDTO, error) {
	cinema, err := u.cinemaRepo.FindByID(ctx, entity.UUID(cinemaID))
	if err != nil {
		return nil, err
	}
	return u.cinemaToDTO(cinema), nil
}

func (u *CinemaUseCase) Create(ctx context.Context, params CreateCinemaParams) (*CinemaDTO, error) {
	name, err := entity.NewCinemaName(params.Name)
	if err != nil {
		return nil, err
	}

	address, err := entity.NewCinemaAddress(params.Address)
	if err != nil {
		return nil, err
	}

	url, err := entity.NewCinemaURL(params.URL)
	if err != nil {
		return nil, err
	}

	cinema := entity.NewCinema(name, address, url)
	if err := u.cinemaRepo.Create(ctx, cinema); err != nil {
		return nil, err
	}
	return u.cinemaToDTO(cinema), nil
}

type CreateCinemaParams struct {
	Name    string
	Address string
	URL     string
}

type CinemaDTO struct {
	ID      string
	Name    string
	Address string
	URL     string
}

func (u *CinemaUseCase) cinemaToDTO(cinema *entity.Cinema) *CinemaDTO {
	return &CinemaDTO{
		ID:      cinema.ID.String(),
		Name:    cinema.Name.String(),
		Address: cinema.Address.String(),
		URL:     cinema.URL.String(),
	}
}
