package database

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/candy12t/cinema-search-server/domain/entity"
	"github.com/candy12t/cinema-search-server/domain/repository"
	"github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type MovieRepository struct {
	db *sqlx.DB
}

var _ repository.Movie = (*MovieRepository)(nil)

func NewMovieRepository(db *sqlx.DB) *MovieRepository {
	return &MovieRepository{
		db: db,
	}
}

func (r *MovieRepository) FindByID(ctx context.Context, movieID entity.UUID) (*entity.Movie, error) {
	dto := new(movieDTO)
	query := `SELECT id, title, release_date, release_status FROM movies WHERE id = ?`
	if err := r.db.GetContext(ctx, dto, query, movieID); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, entity.ErrMovieNotFound
		}
		return nil, err
	}
	return r.dtoToMovie(dto)
}

func (r *MovieRepository) Save(ctx context.Context, movie *entity.Movie) error {
	dto := r.movieToDTO(movie)
	query := `INSERT INTO movies (id, title, release_date, release_status) VALUE (:id, :title, :release_date, :release_status)`
	if _, err := r.db.NamedExecContext(ctx, query, dto); err != nil {
		var mysqlErr *mysql.MySQLError
		if errors.As(err, &mysqlErr); mysqlErr.Number == MySQLDuplicateEntryErrorCode {
			return entity.ErrMovieAlreadyExisted
		}
		return err
	}
	return nil
}

func (r *MovieRepository) Update(ctx context.Context, movie *entity.Movie) error {
	dto := r.movieToDTO(movie)
	query := `UPDATE movies SET title = :title, release_date = :release_date, release_status = :release_status WHERE id = :id`
	if _, err := r.db.NamedExecContext(ctx, query, dto); err != nil {
		return err
	}
	return nil
}

type movieDTO struct {
	ID            string    `db:"id"`
	Title         string    `db:"title"`
	ReleaseDate   time.Time `db:"release_date"`
	ReleaseStatus string    `db:"release_status"`
}

func (r *MovieRepository) movieToDTO(movie *entity.Movie) *movieDTO {
	return &movieDTO{
		ID:            movie.ID.String(),
		Title:         movie.Title.String(),
		ReleaseDate:   movie.ReleaseDate,
		ReleaseStatus: movie.ReleaseStatus.String(),
	}
}

func (r *MovieRepository) dtoToMovie(dto *movieDTO) (*entity.Movie, error) {
	relaseStatus, err := entity.NewReleaseStatus(dto.ReleaseStatus)
	if err != nil {
		return nil, err
	}

	return &entity.Movie{
		ID:            entity.UUID(dto.ID),
		Title:         entity.MovieTitle(dto.Title),
		ReleaseDate:   dto.ReleaseDate,
		ReleaseStatus: relaseStatus,
	}, nil
}
