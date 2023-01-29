package database

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/candy12t/cinemarch-server/domain/entity"
	"github.com/candy12t/cinemarch-server/domain/repository"
	"github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type ScreeningMovieRepository struct {
	db *sqlx.DB
}

var _ repository.ScreeningMovie = (*ScreeningMovieRepository)(nil)

func NewScreeningMovieRepository(db *sqlx.DB) *ScreeningMovieRepository {
	return &ScreeningMovieRepository{
		db: db,
	}
}

func (r *ScreeningMovieRepository) FindByID(ctx context.Context, screeningMovieID entity.UUID) (*entity.ScreeningMovie, error) {
	dto := new(screeningMovieDTO)
	query := `SELECT id, cinema_id, movie_id, start_time, end_time FROM screening_movies WHERE id = ?`
	if err := r.db.GetContext(ctx, dto, query, screeningMovieID); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, entity.ErrScreeningMovieNotFound
		}
		return nil, err
	}
	return r.dtoToScreeningMovie(dto), nil
}

func (r *ScreeningMovieRepository) Create(ctx context.Context, screeningMovie *entity.ScreeningMovie) error {
	dto := r.screeningMovieToDTO(screeningMovie)
	query := `INSERT INTO screening_movies (id, cinema_id, movie_id, start_time, end_time) VALUES (:id, :cinema_id, :movie_id, :start_time, :end_time)`
	_, err := r.db.NamedExecContext(ctx, query, dto)
	if err != nil {
		var mysqlErr *mysql.MySQLError
		if errors.As(err, &mysqlErr); mysqlErr.Number == MySQLDuplicateEntryErrorCode {
			return entity.ErrScreeningMovieAlreadyExisted
		}
		return err
	}
	return nil
}

type screeningMovieDTO struct {
	ID        string    `db:"id"`
	CinemaID  string    `db:"cinema_id"`
	MovieID   string    `db:"movie_id"`
	StartTime time.Time `db:"start_time"`
	EndTime   time.Time `db:"end_time"`
}

func (r *ScreeningMovieRepository) screeningMovieToDTO(screeningMovie *entity.ScreeningMovie) *screeningMovieDTO {
	return &screeningMovieDTO{
		ID:        screeningMovie.ID.String(),
		CinemaID:  screeningMovie.CinemaID.String(),
		MovieID:   screeningMovie.MovieID.String(),
		StartTime: screeningMovie.StartTime,
		EndTime:   screeningMovie.EndTime,
	}
}

func (r *ScreeningMovieRepository) dtoToScreeningMovie(dto *screeningMovieDTO) *entity.ScreeningMovie {
	return &entity.ScreeningMovie{
		ID:        entity.UUID(dto.ID),
		CinemaID:  entity.UUID(dto.CinemaID),
		MovieID:   entity.UUID(dto.MovieID),
		StartTime: dto.StartTime,
		EndTime:   dto.EndTime,
	}
}
