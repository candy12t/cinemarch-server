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

type ScreeningTypeRepository struct {
	db *sqlx.DB
}

var _ repository.ScreeningType = (*ScreeningTypeRepository)(nil)

func NewScreeningTypeRepository(db *sqlx.DB) *ScreeningTypeRepository {
	return &ScreeningTypeRepository{
		db: db,
	}
}

func (r *ScreeningTypeRepository) FindByID(ctx context.Context, screeningTypeID entity.UUID) (*entity.ScreeningType, error) {
	dto := new(screeningTypeDTO)
	query := `SELECT id, name FROM screening_types WHERE id = ?`
	if err := r.db.GetContext(ctx, dto, query, screeningTypeID); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, entity.ErrScreeningTypeNotFound
		}
		return nil, err
	}
	return r.dtoToScreeningType(dto), nil
}

func (r *ScreeningTypeRepository) Create(ctx context.Context, screeningType *entity.ScreeningType) error {
	dto := r.screeningTypeToDTO(screeningType)
	query := `INSERT INTO screening_types (id, name) VALUES (:id, :name)`
	if _, err := r.db.NamedQueryContext(ctx, query, dto); err != nil {
		var mysqlErr *mysql.MySQLError
		if errors.As(err, &mysqlErr); mysqlErr.Number == MySQLDuplicateEntryErrorCode {
			return entity.ErrScreeningTypeAlreadyExisted
		}
		return err
	}
	return nil
}

type screeningTypeDTO struct {
	ID   string `db:"id"`
	Name string `db:"name"`
}

func (r *ScreeningTypeRepository) screeningTypeToDTO(screeningType *entity.ScreeningType) *screeningTypeDTO {
	return &screeningTypeDTO{
		ID:   screeningType.ID.String(),
		Name: screeningType.Name.String(),
	}
}

func (r *ScreeningTypeRepository) dtoToScreeningType(dto *screeningTypeDTO) *entity.ScreeningType {
	return &entity.ScreeningType{
		ID:   entity.UUID(dto.ID),
		Name: entity.ScreeningTypeName(dto.Name),
	}
}
