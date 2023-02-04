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

type ScreenTypeRepository struct {
	db *sqlx.DB
}

var _ repository.ScreenType = (*ScreenTypeRepository)(nil)

func NewScreenTypeRepository(db *sqlx.DB) *ScreenTypeRepository {
	return &ScreenTypeRepository{
		db: db,
	}
}

func (r *ScreenTypeRepository) FindByID(ctx context.Context, screenTypeID entity.UUID) (*entity.ScreenType, error) {
	dto := new(screenTypeDTO)
	query := `SELECT id, name FROM screen_types WHERE id = ?`
	if err := r.db.GetContext(ctx, dto, query, screenTypeID); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, entity.ErrScreenTypeNotFound
		}
		return nil, err
	}
	return r.dtoToScreenType(dto), nil
}

func (r *ScreenTypeRepository) Create(ctx context.Context, screenType *entity.ScreenType) error {
	dto := r.screenTypeToDTO(screenType)
	query := `INSERT INTO screen_types (id, name) VALUES (:id, :name)`
	if _, err := r.db.NamedQueryContext(ctx, query, dto); err != nil {
		var mysqlErr *mysql.MySQLError
		if errors.As(err, &mysqlErr); mysqlErr.Number == MySQLDuplicateEntryErrorCode {
			return entity.ErrScreenTypeAlreadyExisted
		}
		return err
	}
	return nil
}

type screenTypeDTO struct {
	ID   string `db:"id"`
	Name string `db:"name"`
}

func (r *ScreenTypeRepository) screenTypeToDTO(screenType *entity.ScreenType) *screenTypeDTO {
	return &screenTypeDTO{
		ID:   screenType.ID.String(),
		Name: screenType.Name.String(),
	}
}

func (r *ScreenTypeRepository) dtoToScreenType(dto *screenTypeDTO) *entity.ScreenType {
	return &entity.ScreenType{
		ID:   entity.UUID(dto.ID),
		Name: entity.ScreenTypeName(dto.Name),
	}
}
