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

type ScreenMovieRepository struct {
	db *sqlx.DB
}

var _ repository.ScreenMovie = (*ScreenMovieRepository)(nil)

func NewScreenMovieRepository(db *sqlx.DB) *ScreenMovieRepository {
	return &ScreenMovieRepository{
		db: db,
	}
}

func (r *ScreenMovieRepository) FindByID(ctx context.Context, screenMovieID entity.UUID) (*entity.ScreenMovie, error) {
	dto := new(screenMovieDTO)
	query := `SELECT id, cinema_id, movie_id, start_time, end_time FROM screen_movies WHERE id = ?`
	if err := r.db.GetContext(ctx, dto, query, screenMovieID); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, entity.ErrScreenMovieNotFound
		}
		return nil, err
	}
	return r.dtoToScreenMovie(dto), nil
}

func (r *ScreenMovieRepository) Create(ctx context.Context, screenMovie *entity.ScreenMovie) error {
	dto := r.screenMovieToDTO(screenMovie)
	query := `INSERT INTO screen_movies (id, cinema_id, movie_id, start_time, end_time) VALUES (:id, :cinema_id, :movie_id, :start_time, :end_time)`
	_, err := r.db.NamedExecContext(ctx, query, dto)
	if err != nil {
		var mysqlErr *mysql.MySQLError
		if errors.As(err, &mysqlErr); mysqlErr.Number == MySQLDuplicateEntryErrorCode {
			return entity.ErrScreenMovieAlreadyExisted
		}
		return err
	}
	return nil
}

type screenMovieDTO struct {
	ID        string    `db:"id"`
	CinemaID  string    `db:"cinema_id"`
	MovieID   string    `db:"movie_id"`
	StartTime time.Time `db:"start_time"`
	EndTime   time.Time `db:"end_time"`
}

func (r *ScreenMovieRepository) screenMovieToDTO(screenMovie *entity.ScreenMovie) *screenMovieDTO {
	return &screenMovieDTO{
		ID:        screenMovie.ID.String(),
		CinemaID:  screenMovie.CinemaID.String(),
		MovieID:   screenMovie.MovieID.String(),
		StartTime: screenMovie.StartTime,
		EndTime:   screenMovie.EndTime,
	}
}

func (r *ScreenMovieRepository) dtoToScreenMovie(dto *screenMovieDTO) *entity.ScreenMovie {
	return &entity.ScreenMovie{
		ID:        entity.UUID(dto.ID),
		CinemaID:  entity.UUID(dto.CinemaID),
		MovieID:   entity.UUID(dto.MovieID),
		StartTime: dto.StartTime,
		EndTime:   dto.EndTime,
	}
}
