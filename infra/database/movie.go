package database

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/candy12t/cinema-search-server/domain/entity"
	"github.com/candy12t/cinema-search-server/domain/repository"
)

type MovieRepository struct {
	db *sql.DB
}

var _ repository.Movie = (*MovieRepository)(nil)

func NewMovieRepository(db *sql.DB) *MovieRepository {
	return &MovieRepository{
		db: db,
	}
}

func (r *MovieRepository) FindByID(ctx context.Context, movieID entity.UUID) (*entity.Movie, error) {
	dto := new(movieDTO)
	query := `SELECT id, title, release_date, release_status FROM movies WHERE id = ?`
	if err := r.db.QueryRowContext(ctx, query, movieID).Scan(&dto.ID, &dto.Title, &dto.ReleaseDate, &dto.ReleaseStatus); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, entity.ErrMovieNotFound
		}
		return nil, err
	}
	return r.dtoToMovie(dto)
}

func (r *MovieRepository) Save(ctx context.Context, movie *entity.Movie) error {
	dto := r.movieToDTO(movie)
	query := `INSERT INTO movies (id, title, release_status, release_date) VALUES (?, ?, ?, ?)`
	_, err := r.db.ExecContext(ctx, query, dto.ID, dto.Title, dto.ReleaseStatus, dto.ReleaseDate)
	if err != nil {
		return err
	}
	return nil
}

func (r *MovieRepository) Update(ctx context.Context, movie *entity.Movie) error {
	dto := r.movieToDTO(movie)
	query := `UPDATE movies SET title = ?, release_status = ?, release_date = ? WHERE id = ?`
	_, err := r.db.ExecContext(ctx, query, dto.Title, dto.ReleaseStatus, dto.ReleaseDate, dto.ID)
	if err != nil {
		return err
	}
	return nil
}

type movieDTO struct {
	ID            string
	Title         string
	ReleaseDate   time.Time
	ReleaseStatus string
}

func (r *MovieRepository) movieToDTO(movie *entity.Movie) *movieDTO {
	return &movieDTO{
		ID:            string(movie.ID),
		Title:         string(movie.Title),
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
