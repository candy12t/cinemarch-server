package database

import (
	"context"
	"errors"
	"reflect"
	"testing"
	"time"

	"github.com/candy12t/cinemarch-server/domain/entity"
	"github.com/jmoiron/sqlx"
)

func TestScreenMovieRepository_FindByID(t *testing.T) {
	db := prepareTestScreenMovieRepository(t)
	tests := []struct {
		name             string
		screenMovieID entity.UUID
		want             *entity.ScreenMovie
		wantErr          error
	}{
		{
			name:             "get screenMovie",
			screenMovieID: entity.UUID("existing_screen_movie_id"),
			want: &entity.ScreenMovie{
				ID:        entity.UUID("existing_screen_movie_id"),
				CinemaID:  entity.UUID("existing_cinema_id"),
				MovieID:   entity.UUID("existing_movie_id"),
				StartTime: time.Date(2023, 1, 1, 15, 0, 0, 0, time.UTC),
				EndTime:   time.Date(2023, 1, 1, 17, 0, 0, 0, time.UTC),
			},
			wantErr: nil,
		},
		{
			name:             "not exist screenMovie",
			screenMovieID: entity.UUID("not_exist_screen_movie_id"),
			want:             nil,
			wantErr:          entity.ErrScreenMovieNotFound,
		},
	}

	for _, tt := range tests {
		r := NewScreenMovieRepository(db)
		got, err := r.FindByID(context.Background(), tt.screenMovieID)
		if !errors.Is(err, tt.wantErr) {
			t.Errorf("ScreenMovieRepository.FindByID() error is %v, want error is %v", err, tt.wantErr)
		}
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("ScreenMovieRepository.FindByID() got is %v, want is %v", got, tt.want)
		}
	}
}

func TestScreenMovieRepository_Create(t *testing.T) {
	db := prepareTestScreenMovieRepository(t)
	tests := []struct {
		name           string
		screenMovie *entity.ScreenMovie
		wantErr        error
	}{
		{
			name: "create screenMovie",
			screenMovie: &entity.ScreenMovie{
				ID:        entity.UUID("new_screen_movie_id"),
				CinemaID:  entity.UUID("new_cinema_id"),
				MovieID:   entity.UUID("new_movie_id"),
				StartTime: time.Date(2023, 1, 1, 15, 0, 0, 0, time.UTC),
				EndTime:   time.Date(2023, 1, 1, 17, 0, 0, 0, time.UTC),
			},
			wantErr: nil,
		},
		{
			name: "existing screenType",
			screenMovie: &entity.ScreenMovie{
				ID:        entity.UUID("existing_screen_movie_id"),
				CinemaID:  entity.UUID("existing_cinema_id"),
				MovieID:   entity.UUID("existing_movie_id"),
				StartTime: time.Date(2023, 1, 1, 15, 0, 0, 0, time.UTC),
				EndTime:   time.Date(2023, 1, 1, 17, 0, 0, 0, time.UTC),
			},
			wantErr: entity.ErrScreenMovieAlreadyExisted,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewScreenMovieRepository(db)
			if err := r.Create(context.Background(), tt.screenMovie); !errors.Is(err, tt.wantErr) {
				t.Errorf("ScreenMovieRepository.Create() error is %v, want error is %v", err, tt.wantErr)
			}
			t.Cleanup(func() {
				if _, err := db.NamedExec("DELETE FROM screen_movies WHERE id = :id", r.screenMovieToDTO(tt.screenMovie)); err != nil {
					t.Fatal(err)
				}
			})
		})
	}
}

func prepareTestScreenMovieRepository(t *testing.T) *sqlx.DB {
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

	screenMovie := &screenMovieDTO{
		ID:        "existing_screen_movie_id",
		CinemaID:  "existing_cinema_id",
		MovieID:   "existing_movie_id",
		StartTime: time.Date(2023, 1, 1, 15, 0, 0, 0, time.UTC),
		EndTime:   time.Date(2023, 1, 1, 17, 0, 0, 0, time.UTC),
	}
	if _, err := db.NamedExec("INSERT INTO screen_movies (id, cinema_id, movie_id, start_time, end_time) VALUES (:id, :cinema_id, :movie_id, :start_time, :end_time)", screenMovie); err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() {
		if _, err := db.NamedExec("DELETE FROM screen_movies WHERE id = :id", screenMovie); err != nil {
			t.Fatal(err)
		}
	})

	return db
}
