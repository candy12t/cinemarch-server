package database

import (
	"context"
	"database/sql"
	"errors"

	"github.com/candy12t/cinemarch-server/domain/entity"
	"github.com/candy12t/cinemarch-server/domain/repository"
	"github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type CinemaRepository struct {
	db *sqlx.DB
}

var _ repository.Cinema = (*CinemaRepository)(nil)

func NewCinemaRepository(db *sqlx.DB) *CinemaRepository {
	return &CinemaRepository{
		db: db,
	}
}

func (r *CinemaRepository) FindByID(ctx context.Context, cinemaID entity.UUID) (*entity.Cinema, error) {
	dto := new(cinemaDTO)
	query := `SELECT id, name, address, url FROM cinemas WHERE id = ?`
	if err := r.db.GetContext(ctx, dto, query, cinemaID); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, entity.ErrCinemaNotFound
		}
		return nil, err
	}
	return r.dtoToCinema(dto), nil
}

func (r *CinemaRepository) Create(ctx context.Context, cinema *entity.Cinema) error {
	dto := r.cinemaToDTO(cinema)
	query := `INSERT INTO cinemas (id, name, address, url) VALUES (:id, :name, :address, :url)`
	if _, err := r.db.NamedExecContext(ctx, query, dto); err != nil {
		var mysqlErr *mysql.MySQLError
		if errors.As(err, &mysqlErr); mysqlErr.Number == MySQLDuplicateEntryErrorCode {
			return entity.ErrCinemaAlreadyExisted
		}
		return err
	}
	return nil
}

type cinemaDTO struct {
	ID      string `db:"id"`
	Name    string `db:"name"`
	Address string `db:"address"`
	URL     string `db:"url"`
}

func (r *CinemaRepository) cinemaToDTO(cinema *entity.Cinema) *cinemaDTO {
	return &cinemaDTO{
		ID:      cinema.ID.String(),
		Name:    cinema.Name.String(),
		Address: cinema.Address.String(),
		URL:     cinema.URL.String(),
	}
}

func (r *CinemaRepository) dtoToCinema(dto *cinemaDTO) *entity.Cinema {
	return &entity.Cinema{
		ID:      entity.UUID(dto.ID),
		Name:    entity.CinemaName(dto.Name),
		Address: entity.CinemaAddress(dto.Address),
		URL:     entity.CinemaURL(dto.URL),
	}
}
