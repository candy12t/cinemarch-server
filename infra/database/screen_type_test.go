package database

import (
	"context"
	"errors"
	"reflect"
	"testing"

	"github.com/candy12t/cinemarch-server/domain/entity"
	"github.com/jmoiron/sqlx"
)

func TestScreenTypeRepository_FindByID(t *testing.T) {
	db := prepareTestScreenTypeRepository(t)

	tests := []struct {
		name            string
		screenTypeID entity.UUID
		want            *entity.ScreenType
		wantErr         error
	}{
		{
			name:            "get existing screenType",
			screenTypeID: entity.UUID("existing_screen_type_id"),
			want: &entity.ScreenType{
				ID:   entity.UUID("existing_screen_type_id"),
				Name: entity.ScreenTypeName("existing_screen_type_name"),
			},
			wantErr: nil,
		},
		{
			name:            "not exist screenType",
			screenTypeID: entity.UUID("not_exist_screen_type_id"),
			want:            nil,
			wantErr:         entity.ErrScreenTypeNotFound,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewScreenTypeRepository(db)
			got, err := r.FindByID(context.Background(), tt.screenTypeID)
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("ScreenTypeRepository.FindByID() error is %v, want error is %v", err, tt.wantErr)
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ScreenTypeRepository.FindByID() got is %v, want is %v", got, tt.want)
			}
		})
	}
}

func TestScreenTypeRepository_Create(t *testing.T) {
	db := prepareTestScreenTypeRepository(t)

	tests := []struct {
		name          string
		screenType *entity.ScreenType
		wantErr       error
	}{
		{
			name: "create screenType",
			screenType: &entity.ScreenType{
				ID:   entity.UUID("new_screen_type_id"),
				Name: entity.ScreenTypeName("new_screen_type_name"),
			},
			wantErr: nil,
		},
		{
			name: "existing screenType",
			screenType: &entity.ScreenType{
				ID:   entity.UUID("existing_screen_type_id"),
				Name: entity.ScreenTypeName("existing_screen_type_name"),
			},
			wantErr: entity.ErrScreenTypeAlreadyExisted,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewScreenTypeRepository(db)
			if err := r.Create(context.Background(), tt.screenType); !errors.Is(err, tt.wantErr) {
				t.Errorf("ScreenTypeRepository.Create() error is %v, want error is %v", err, tt.wantErr)
			}
			t.Cleanup(func() {
				if _, err := db.NamedExec("DELETE FROM screen_types WHERE id = :id", r.screenTypeToDTO(tt.screenType)); err != nil {
					t.Fatal(err)
				}
			})
		})
	}
}

func prepareTestScreenTypeRepository(t *testing.T) *sqlx.DB {
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

	screenType := &screenTypeDTO{
		ID:   "existing_screen_type_id",
		Name: "existing_screen_type_name",
	}
	if _, err := db.NamedExec(`INSERT INTO screen_types (id, name) VALUES (:id, :name)`, screenType); err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() {
		if _, err := db.NamedExec("DELETE FROM screen_types WHERE id = :id", screenType); err != nil {
			t.Fatal(err)
		}
	})

	return db
}
