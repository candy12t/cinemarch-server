package database

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/candy12t/cinemarch-server/domain/entity"
	"github.com/candy12t/cinemarch-server/domain/repository"
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
	return r.find(ctx, "id", movieID)
}

func (r *MovieRepository) FindByTitle(ctx context.Context, title entity.MovieTitle) (*entity.Movie, error) {
	return r.find(ctx, "title", title)
}

func (r *MovieRepository) find(ctx context.Context, column string, arg any) (*entity.Movie, error) {
	dto := new(movieDTO)
	query := fmt.Sprintf(`SELECT id, title, release_date, release_status FROM movies WHERE %s = ?`, column)
	if err := r.db.GetContext(ctx, dto, query, arg); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, entity.ErrMovieNotFound
		}
		return nil, err
	}

	releaseStatus, err := entity.NewReleaseStatus(dto.ReleaseStatus)
	if err != nil {
		return nil, err
	}

	return dtoToMovie(dto, releaseStatus), nil
}

func (r *MovieRepository) Search(ctx context.Context, conditionQuery string, args []any) (entity.Movies, error) {
	dtos := movieDTOs{}
	query := fmt.Sprintf(`SELECT id, title, release_date, release_status FROM movies %s`, conditionQuery)
	if err := r.db.SelectContext(ctx, &dtos, query, args...); err != nil {
		return nil, err
	}

	if len(dtos) == 0 {
		return nil, entity.ErrMovieNotFound
	}

	movies := make(entity.Movies, 0, len(dtos))
	for _, dto := range dtos {
		releaseStatus, err := entity.NewReleaseStatus(dto.ReleaseStatus)
		if err != nil {
			return nil, err
		}
		movies = append(movies, dtoToMovie(dto, releaseStatus))
	}
	return movies, nil
}

func (r *MovieRepository) Create(ctx context.Context, movie *entity.Movie) error {
	dto := movieToDTO(movie)
	query := `INSERT INTO movies (id, title, release_date, release_status) VALUES (:id, :title, :release_date, :release_status)`
	if _, err := r.db.NamedExecContext(ctx, query, dto); err != nil {
		var mysqlErr *mysql.MySQLError
		if errors.As(err, &mysqlErr); mysqlErr.Number == mysqlDuplicateEntryErrorCode {
			return entity.ErrMovieAlreadyExisted
		}
		return err
	}
	return nil
}

func (r *MovieRepository) Update(ctx context.Context, movie *entity.Movie) error {
	dto := movieToDTO(movie)
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

type movieDTOs []*movieDTO

func dtoToMovie(dto *movieDTO, releaseStatus entity.ReleaseStatus) *entity.Movie {
	return &entity.Movie{
		ID:            entity.UUID(dto.ID),
		Title:         entity.MovieTitle(dto.Title),
		ReleaseDate:   dto.ReleaseDate,
		ReleaseStatus: releaseStatus,
	}
}

func movieToDTO(movie *entity.Movie) *movieDTO {
	return &movieDTO{
		ID:            movie.ID.String(),
		Title:         movie.Title.String(),
		ReleaseDate:   movie.ReleaseDate,
		ReleaseStatus: movie.ReleaseStatus.String(),
	}
}
