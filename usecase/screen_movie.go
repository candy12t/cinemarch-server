package usecase

import (
	"context"
	"errors"

	"github.com/candy12t/cinemarch-server/domain/entity"
	"github.com/candy12t/cinemarch-server/domain/repository"
	"github.com/candy12t/cinemarch-server/lib"
	"github.com/candy12t/cinemarch-server/query/dto"
	"github.com/candy12t/cinemarch-server/query/service"
)

type ScreenMovie interface {
	Create(ctx context.Context, params CreateScreenMovieParams) (*ScreenMovieDTO, error)
	Search(ctx context.Context, searchCondition dto.ScreenMovieSearchCondition) (ScreenMovieDTOs, error)
}

type ScreenMovieUseCase struct {
	cinemaRepo         repository.Cinema
	movieRepo          repository.Movie
	screenMovieRepo    repository.ScreenMovie
	screenMovieService service.ScreenMovie
}

var _ ScreenMovie = (*ScreenMovieUseCase)(nil)

func NewScreenMovieUseCase(cinemaRepo repository.Cinema, movieRepo repository.Movie, screenMovieRepo repository.ScreenMovie, screenMovieService service.ScreenMovie) *ScreenMovieUseCase {
	return &ScreenMovieUseCase{
		cinemaRepo:         cinemaRepo,
		movieRepo:          movieRepo,
		screenMovieRepo:    screenMovieRepo,
		screenMovieService: screenMovieService,
	}
}

func (u *ScreenMovieUseCase) Search(ctx context.Context, searchCondition dto.ScreenMovieSearchCondition) (ScreenMovieDTOs, error) {
	screenMovies, err := u.screenMovieService.Search(ctx, searchCondition)
	if err != nil {
		return nil, err
	}
	dtos := make(ScreenMovieDTOs, 0, len(screenMovies))
	for _, sm := range screenMovies {
		schedules := make(ScreenScheduleDTOs, 0, len(sm.Schedules))
		for _, schedule := range sm.Schedules {
			schedules = append(schedules, &ScreenScheduleDTO{
				StartTime: schedule.StartTime,
				EndTime:   schedule.EndTime,
			})
		}
		dtos = append(dtos, &ScreenMovieDTO{
			ID:              sm.ID,
			CinemaName:      sm.CinemaName,
			MovieTitle:      sm.MovieTitle,
			ScreenType:      sm.ScreenType,
			TranslateType:   sm.TranslateType,
			ThreeD:          sm.ThreeD,
			ScreenSchedules: schedules,
		})
	}
	return dtos, nil
}

func (u *ScreenMovieUseCase) Create(ctx context.Context, params CreateScreenMovieParams) (*ScreenMovieDTO, error) {
	cinema, err := u.cinemaRepo.FindByName(ctx, entity.CinemaName(params.CinemaName))
	if err != nil {
		return nil, err
	}
	movie, err := u.movieRepo.FindByTitle(ctx, entity.MovieTitle(params.MovieTitle))
	if err != nil {
		return nil, err
	}
	screenType, err := entity.NewScreenType(params.ScreenType)
	if err != nil {
		return nil, err
	}
	translateType, err := entity.NewTranslateType(params.TranslateType)
	if err != nil {
		return nil, err
	}

	screenMovie, err := u.screenMovieRepo.FindByUniqKey(ctx, cinema.ID, movie.ID, screenType, translateType, params.ThreeD)
	if err != nil {
		if errors.Is(err, entity.ErrScreenMovieNotFound) {
			screenMovie := entity.NewScreenMovie(cinema.ID, movie.ID, screenType, translateType, params.ThreeD)
			if err := u.screenMovieRepo.CreateScreenMovie(ctx, screenMovie); err != nil {
				return nil, err
			}
			screenSchedules, err := u.createScreenSchedules(ctx, screenMovie.ID, params.ScreenSchedules)
			if err != nil {
				return nil, err
			}
			return screenMovieToDTO(screenMovie, screenSchedules, cinema.Name, movie.Title), nil
		}
		return nil, err
	}

	screenSchedules, err := u.createScreenSchedules(ctx, screenMovie.ID, params.ScreenSchedules)
	if err != nil {
		return nil, err
	}
	return screenMovieToDTO(screenMovie, screenSchedules, cinema.Name, movie.Title), nil
}

func (u *ScreenMovieUseCase) createScreenSchedules(ctx context.Context, screenMovieID entity.UUID, screenSchedulesParams []*CreateScreenScheduleParams) (entity.ScreenSchedules, error) {
	screenSchedules := make(entity.ScreenSchedules, 0, len(screenSchedulesParams))
	for _, screenSchedule := range screenSchedulesParams {
		startTime, err := lib.ParseJSTDateTimeInUTC(screenSchedule.StartTime)
		if err != nil {
			return nil, err
		}
		endTime, err := lib.ParseJSTDateTimeInUTC(screenSchedule.EndTime)
		if err != nil {
			return nil, err
		}
		screenSchedules = append(screenSchedules, entity.NewScreenSchedule(screenMovieID, startTime, endTime))
	}
	if err := u.screenMovieRepo.CreateScreenSchedules(ctx, screenSchedules); err != nil {
		return nil, err
	}
	return screenSchedules, nil
}

type CreateScreenMovieParams struct {
	CinemaName      string
	MovieTitle      string
	ScreenType      string
	TranslateType   string
	ThreeD          bool
	ScreenSchedules []*CreateScreenScheduleParams
}

type CreateScreenScheduleParams struct {
	StartTime string
	EndTime   string
}

type ScreenMovieDTO struct {
	ID              string
	CinemaName      string
	MovieTitle      string
	ScreenType      string
	TranslateType   string
	ThreeD          bool
	ScreenSchedules ScreenScheduleDTOs
}

type ScreenMovieDTOs []*ScreenMovieDTO

type ScreenScheduleDTO struct {
	StartTime string
	EndTime   string
}

type ScreenScheduleDTOs []*ScreenScheduleDTO

func screenMovieToDTO(screenMovie *entity.ScreenMovie, screenSchedules entity.ScreenSchedules, cinemaName entity.CinemaName, movieTitle entity.MovieTitle) *ScreenMovieDTO {
	screenScheduleDTOs := make([]*ScreenScheduleDTO, 0, len(screenMovie.ScreenSchedules)+len(screenSchedules))
	for _, ss := range screenMovie.ScreenSchedules {
		screenScheduleDTOs = append(screenScheduleDTOs, &ScreenScheduleDTO{
			StartTime: lib.FormatDateTimeInJST(ss.StartTime),
			EndTime:   lib.FormatDateTimeInJST(ss.EndTime),
		})
	}
	for _, ss := range screenSchedules {
		screenScheduleDTOs = append(screenScheduleDTOs, &ScreenScheduleDTO{
			StartTime: lib.FormatDateTimeInJST(ss.StartTime),
			EndTime:   lib.FormatDateTimeInJST(ss.EndTime),
		})
	}
	return &ScreenMovieDTO{
		ID:              screenMovie.ID.String(),
		CinemaName:      cinemaName.String(),
		MovieTitle:      movieTitle.String(),
		ScreenType:      screenMovie.ScreenType.String(),
		TranslateType:   screenMovie.TranslateType.String(),
		ThreeD:          screenMovie.TreeD,
		ScreenSchedules: screenScheduleDTOs,
	}
}
