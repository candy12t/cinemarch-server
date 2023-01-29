package database

import (
	"context"
	"errors"
	"reflect"
	"testing"

	"github.com/candy12t/cinemarch-server/domain/entity"
	"github.com/jmoiron/sqlx"
)

func TestCinemaRepository_FindByID(t *testing.T) {
	db := prepareTestCinemaRepository(t)

	tests := []struct {
		name     string
		cinemaID entity.UUID
		want     *entity.Cinema
		wantErr  error
	}{
		{
			name:     "get existing cinema",
			cinemaID: entity.UUID("existing_cinema_id"),
			want: &entity.Cinema{
				ID:      entity.UUID("existing_cinema_id"),
				Name:    entity.CinemaName("existing_cinema_name"),
				Address: entity.CinemaAddress("existing_cinema_address"),
				URL:     entity.CinemaURL("https://existing.cinema.url"),
			},
			wantErr: nil,
		},
		{
			name:     "not exist cinema",
			cinemaID: entity.UUID("not_exist_cinema_id"),
			want:     nil,
			wantErr:  entity.ErrCinemaNotFound,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewCinemaRepository(db)
			got, err := r.FindByID(context.Background(), tt.cinemaID)
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("CinemaRepository.FindByID() error is %v, want error is %v", err, tt.wantErr)
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CinemaRepository.FindByID() got is %v, want is %v", got, tt.want)
			}
		})
	}
}

func TestCinemaRepository_Create(t *testing.T) {
	db := prepareTestCinemaRepository(t)

	tests := []struct {
		name    string
		cinema  *entity.Cinema
		wantErr error
	}{
		{
			name: "create cinema",
			cinema: &entity.Cinema{
				ID:      entity.UUID("new_cinema_id"),
				Name:    entity.CinemaName("new_cinema_name"),
				Address: entity.CinemaAddress("new_cinema_address"),
				URL:     entity.CinemaURL("https://new.cinema.url"),
			},
			wantErr: nil,
		},
		{
			name: "existing cinema",
			cinema: &entity.Cinema{
				ID:      entity.UUID("existing_cinema_id"),
				Name:    entity.CinemaName("existing_cinema_name"),
				Address: entity.CinemaAddress("existing_cinema_address"),
				URL:     entity.CinemaURL("https://existing.cinema.url"),
			},
			wantErr: entity.ErrCinemaAlreadyExisted,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewCinemaRepository(db)
			if err := r.Create(context.Background(), tt.cinema); !errors.Is(err, tt.wantErr) {
				t.Errorf("CinemaRepository.Create()	error is %v, want error is %v", err, tt.wantErr)
			}
			t.Cleanup(func() {
				if _, err := db.NamedExec("DELETE FROM cinemas WHERE id = :id", r.cinemaToDTO(tt.cinema)); err != nil {
					t.Fatal(err)
				}
			})
		})
	}
}

func prepareTestCinemaRepository(t *testing.T) *sqlx.DB {
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

	cinema := &cinemaDTO{
		ID:      "existing_cinema_id",
		Name:    "existing_cinema_name",
		Address: "existing_cinema_address",
		URL:     "https://existing.cinema.url",
	}
	if _, err := db.NamedExec(`INSERT INTO cinemas (id, name, address, url) VALUES (:id, :name, :address, :url)`, cinema); err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() {
		if _, err := db.NamedExec(`DELETE FROM cinemas WHERE id = :id`, cinema); err != nil {
			t.Fatal(err)
		}
	})

	return db
}
