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
	query := `SELECT id, name, prefecture, address, web_site FROM cinemas WHERE id = ?`
	if err := r.db.GetContext(ctx, dto, query, cinemaID); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, entity.ErrCinemaNotFound
		}
		return nil, err
	}
	return dtoToCinema(dto), nil
}

func (r *CinemaRepository) FindAllByPrefecture(ctx context.Context, prefecture entity.Prefecture) (entity.Cinemas, error) {
	dtos := cinemaDTOs{}
	query := `SELECT id, name, prefecture, address, web_site FROM cinemas WHERE prefecture = ?`
	if err := r.db.SelectContext(ctx, &dtos, query, prefecture); err != nil {
		return nil, err
	}

	if len(dtos) == 0 {
		return nil, entity.ErrCinemaNotFound
	}

	cinemas := make(entity.Cinemas, 0, len(dtos))
	for _, dto := range dtos {
		cinemas = append(cinemas, dtoToCinema(dto))
	}
	return cinemas, nil
}

func (r *CinemaRepository) Create(ctx context.Context, cinema *entity.Cinema) error {
	dto := cinemaToDTO(cinema)
	query := `INSERT INTO cinemas (id, name, prefecture, address, web_site) VALUES (:id, :name, :prefecture, :address, :web_site)`
	if _, err := r.db.NamedExecContext(ctx, query, dto); err != nil {
		var mysqlErr *mysql.MySQLError
		if errors.As(err, &mysqlErr); mysqlErr.Number == mysqlDuplicateEntryErrorCode {
			return entity.ErrCinemaAlreadyExisted
		}
		return err
	}
	return nil
}

type cinemaDTO struct {
	ID         string `db:"id"`
	Name       string `db:"name"`
	Prefecture string `db:"prefecture"`
	Address    string `db:"address"`
	WebSite    string `db:"web_site"`
}

type cinemaDTOs []*cinemaDTO

func dtoToCinema(dto *cinemaDTO) *entity.Cinema {
	return &entity.Cinema{
		ID:         entity.UUID(dto.ID),
		Name:       entity.CinemaName(dto.Name),
		Prefecture: entity.Prefecture(dto.Prefecture),
		Address:    entity.CinemaAddress(dto.Address),
		WebSite:    entity.CinemaWebSite(dto.WebSite),
	}
}

func cinemaToDTO(cinema *entity.Cinema) *cinemaDTO {
	return &cinemaDTO{
		ID:         cinema.ID.String(),
		Name:       cinema.Name.String(),
		Prefecture: cinema.Prefecture.String(),
		Address:    cinema.Address.String(),
		WebSite:    cinema.WebSite.String(),
	}
}
