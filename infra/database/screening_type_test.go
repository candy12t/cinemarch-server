package database

import (
	"context"
	"errors"
	"reflect"
	"testing"

	"github.com/candy12t/cinemarch-server/domain/entity"
	"github.com/jmoiron/sqlx"
)

func TestScreeningTypeRepository_FindByID(t *testing.T) {
	db := prepareTestScreeningTypeRepository(t)

	tests := []struct {
		name            string
		screeningTypeID entity.UUID
		want            *entity.ScreeningType
		wantErr         error
	}{
		{
			name:            "get existing screeningType",
			screeningTypeID: entity.UUID("existing_screening_type_id"),
			want: &entity.ScreeningType{
				ID:   entity.UUID("existing_screening_type_id"),
				Name: entity.ScreeningTypeName("existing_screening_type_name"),
			},
			wantErr: nil,
		},
		{
			name:            "not exist screeningType",
			screeningTypeID: entity.UUID("not_exist_screening_type_id"),
			want:            nil,
			wantErr:         entity.ErrScreeningTypeNotFound,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewScreeningTypeRepository(db)
			got, err := r.FindByID(context.Background(), tt.screeningTypeID)
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("ScreeningTypeRepository.FindByID() error is %v, want error is %v", err, tt.wantErr)
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ScreeningTypeRepository.FindByID() got is %v, want is %v", got, tt.want)
			}
		})
	}
}

func TestScreeningTypeRepository_Create(t *testing.T) {
	db := prepareTestScreeningTypeRepository(t)

	tests := []struct {
		name          string
		screeningType *entity.ScreeningType
		wantErr       error
	}{
		{
			name: "create screeningType",
			screeningType: &entity.ScreeningType{
				ID:   entity.UUID("new_screening_type_id"),
				Name: entity.ScreeningTypeName("new_screening_type_name"),
			},
			wantErr: nil,
		},
		{
			name: "existing screeningType",
			screeningType: &entity.ScreeningType{
				ID:   entity.UUID("existing_screening_type_id"),
				Name: entity.ScreeningTypeName("existing_screening_type_name"),
			},
			wantErr: entity.ErrScreeningTypeAlreadyExisted,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewScreeningTypeRepository(db)
			if err := r.Create(context.Background(), tt.screeningType); !errors.Is(err, tt.wantErr) {
				t.Errorf("ScreeningTypeRepository.Create() error is %v, want error is %v", err, tt.wantErr)
			}
			t.Cleanup(func() {
				if _, err := db.NamedExec("DELETE FROM screening_types WHERE id = :id", r.screeningTypeToDTO(tt.screeningType)); err != nil {
					t.Fatal(err)
				}
			})
		})
	}
}

func prepareTestScreeningTypeRepository(t *testing.T) *sqlx.DB {
	t.Helper()
	db, cleanup, err := NewDB()
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() {
		if err := cleanup(); err != nil {
			t.Fatal(err)
		}
	})

	screeningType := &screeningTypeDTO{
		ID:   "existing_screening_type_id",
		Name: "existing_screening_type_name",
	}
	if _, err := db.NamedExec(`INSERT INTO screening_types (id, name) VALUES (:id, :name)`, screeningType); err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() {
		if _, err := db.NamedExec("DELETE FROM screening_types WHERE id = :id", screeningType); err != nil {
			t.Fatal(err)
		}
	})

	return db
}
