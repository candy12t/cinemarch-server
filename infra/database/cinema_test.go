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
			cinemaID: "existing_cinema_id",
			want: &entity.Cinema{
				ID:         "existing_cinema_id",
				Name:       "existing_cinema_name",
				Prefecture: "東京都",
				Address:    "東京都新宿区新宿1-1-1",
				WebSite:    "https://example.com",
			},
			wantErr: nil,
		},
		{
			name:     "cinema not found",
			cinemaID: "not_exist_cinema_id",
			want:     nil,
			wantErr:  entity.ErrCinemaNotFound,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := NewCinemaRepository(db)
			got, err := repo.FindByID(context.Background(), tt.cinemaID)
			if err != nil {
				if !errors.Is(err, tt.wantErr) {
					t.Errorf("CinemaRepository.FindByID() error is %v, wantErr is %v", err, tt.wantErr)
					return
				}
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CinemaRepository.FindByID() got is %v, want is %v", got, tt.want)
			}
		})
	}
}

func TestCinemaRepository_FindAllByPrefeccture(t *testing.T) {
	db := prepareTestCinemaRepository(t)
	tests := []struct {
		name       string
		prefecture entity.Prefecture
		want       entity.Cinemas
		wantErr    error
	}{
		{
			name:       "get existing cinemas by prefecture",
			prefecture: "東京都",
			want: entity.Cinemas{
				{
					ID:         "existing_cinema_id",
					Name:       "existing_cinema_name",
					Prefecture: "東京都",
					Address:    "東京都新宿区新宿1-1-1",
					WebSite:    "https://example.com",
				},
			},
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := NewCinemaRepository(db)
			got, err := repo.FindAllByPrefecture(context.Background(), tt.prefecture)
			if err != nil {
				if !errors.Is(err, tt.wantErr) {
					t.Errorf("CinemaRepository.FindAllByPrefecture() error is %v, wantErr is %v", err, tt.wantErr)
				}
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CinemaRepository.FindAllByPrefecture() got is %v, want is %v", got, tt.want)
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
			name: "create new cinema",
			cinema: &entity.Cinema{
				ID:         "new_cinema_id",
				Name:       "new_cinema_title",
				Prefecture: "東京都",
				Address:    "東京都渋谷区渋谷1-1-1",
				WebSite:    "https://example.com",
			},
			wantErr: nil,
		},
		{
			name: "cinema has already existed",
			cinema: &entity.Cinema{
				ID:         "new_cinema_id",
				Name:       "existing_cinema_name",
				Prefecture: "東京都",
				Address:    "東京都渋谷区渋谷1-1-1",
				WebSite:    "https://example.com",
			},
			wantErr: entity.ErrCinemaAlreadyExisted,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := NewCinemaRepository(db)
			if err := repo.Create(context.Background(), tt.cinema); !errors.Is(err, tt.wantErr) {
				t.Errorf("CinemaRepository.Create() error is %v, wantErr is %v", err, tt.wantErr)
			}

			t.Cleanup(func() {
				if _, err := db.Exec(`DELETE FROM cinemas WHERE id = ?`, tt.cinema.ID.String()); err != nil {
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
		ID:         "existing_cinema_id",
		Name:       "existing_cinema_name",
		Prefecture: "東京都",
		Address:    "東京都新宿区新宿1-1-1",
		WebSite:    "https://example.com",
	}
	if _, err := db.NamedExec(`INSERT INTO cinemas (id, name, prefecture, address, web_site) VALUES (:id, :name, :prefecture, :address, :web_site)`, cinema); err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() {
		if _, err := db.NamedExec(`DELETE FROM cinemas WHERE id = :id`, cinema); err != nil {
			t.Fatal(err)
		}
	})

	return db
}
