package database

import (
	"context"
	"time"

	"github.com/candy12t/cinemarch-server/domain/entity"
	"github.com/candy12t/cinemarch-server/domain/repository"
	"github.com/jmoiron/sqlx"
)

type ScreenMovieRepository struct {
	db *sqlx.DB
}

var _ repository.ScreenMovie = (*ScreenMovieRepository)(nil)

func NewScreenMovieRepository(db *sqlx.DB) *ScreenMovieRepository {
	return &ScreenMovieRepository{
		db: db,
	}
}
func (r *ScreenMovieRepository) FindByUniqKey(ctx context.Context, cinemaID, movieID entity.UUID, screenType entity.ScreenType, translateType entity.TranslateType, threeD bool) (*entity.ScreenMovie, error) {
	return nil, nil
}

func (r *ScreenMovieRepository) Search(ctx context.Context) (entity.ScreenMovies, error) {
	// query := `
	// 	SELECT
	// 		screen_movies.id AS id,
	// 		screen_movies.cinema_id AS cinema_id,
	// 		screen_movies.movie_id AS movie_id,
	// 		screen_movies.screen_type AS screen_type,
	// 		screen_movies.translate_type AS translate_type,
	// 		screen_movies.three_d AS three_d,
	// 		schedules.start_time AS start_time,
	// 		schedules.end_time AS end_time
	// 	FROM
	// 		screen_movies
	// 	INNER JOIN screen_schedules AS schedules
	// 		ON screen_movies.id = schedules.screen_movie_id
	// `
	return nil, nil
}

func (r *ScreenMovieRepository) CreateScreenMovie(ctx context.Context, screenMovie *entity.ScreenMovie) error {
	return nil
}

func (r *ScreenMovieRepository) CreateScreenSchedules(ctx context.Context, screenSchedule entity.ScreenSchedules) error {
	return nil
}

type screenMovieDTO struct {
	ID            string    `db:"id"`
	CinemaID      string    `db:"cinema_id"`
	MovieID       string    `db:"movie_id"`
	TranslateType string    `db:"translate_type"`
	ScreenType    string    `db:"screen_type"`
	ThreeD        bool      `db:"three_d"`
	StartTime     time.Time `db:"start_time"`
	EndTime       time.Time `db:"end_time"`
}

func dtoToScreenMovie(dto *screenMovieDTO) *entity.ScreenMovie {
	return &entity.ScreenMovie{
		ID:            entity.UUID(dto.ID),
		CinemaID:      entity.UUID(dto.CinemaID),
		MovieID:       entity.UUID(dto.MovieID),
		TranslateType: entity.TranslateType(dto.TranslateType),
		TreeD:         dto.ThreeD,
	}
}

func screenMovieToDTO(screenMovie *entity.ScreenMovie) *screenMovieDTO {
	return &screenMovieDTO{
		ID:       screenMovie.ID.String(),
		CinemaID: screenMovie.CinemaID.String(),
		MovieID:  screenMovie.MovieID.String(),
	}
}
