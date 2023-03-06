package usecase

import (
	"context"

	"github.com/candy12t/cinemarch-server/domain/entity"
	"github.com/candy12t/cinemarch-server/domain/repository"
)

type Cinema interface {
	FindByID(ctx context.Context, cinemaID string) (*CinemaDTO, error)
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

func (u *CinemaUseCase) FindByID(ctx context.Context, cinemaID string) (*CinemaDTO, error) {
	cinema, err := u.cinemaRepo.FindByID(ctx, entity.UUID(cinemaID))
	if err != nil {
		return nil, err
	}
	return cinemaToDTO(cinema), nil
}

func (u *CinemaUseCase) Create(ctx context.Context, params CreateCinemaParams) (*CinemaDTO, error) {
	name, err := entity.NewCinemaName(params.Name)
	if err != nil {
		return nil, err
	}
	prefecture, err := entity.NewPrefecture(params.Prefecture)
	if err != nil {
		return nil, err
	}
	address, err := entity.NewCinemaAddress(params.Address)
	if err != nil {
		return nil, err
	}
	webSite, err := entity.NewCinemaWebSite(params.WebSite)
	if err != nil {
		return nil, err
	}

	cinema := entity.NewCinema(name, prefecture, address, webSite)
	if err := u.cinemaRepo.Create(ctx, cinema); err != nil {
		return nil, err
	}
	return cinemaToDTO(cinema), nil
}

type CreateCinemaParams struct {
	Name       string
	Prefecture string
	Address    string
	WebSite    string
}

type CinemaDTO struct {
	ID         string
	Name       string
	Prefecture string
	Address    string
	WebSite    string
}

type CinemaDTOs []*CinemaDTO

func cinemaToDTO(cinema *entity.Cinema) *CinemaDTO {
	return &CinemaDTO{
		ID:         cinema.ID.String(),
		Name:       cinema.Name.String(),
		Prefecture: cinema.Prefecture.String(),
		Address:    cinema.Address.String(),
		WebSite:    cinema.WebSite.String(),
	}
}
