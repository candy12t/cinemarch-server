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

func TestScreeningMovieRepository_FindByID(t *testing.T) {
	db := prepareTestScreeningMovieRepository(t)
	tests := []struct {
		name             string
		screeningMovieID entity.UUID
		want             *entity.ScreeningMovie
		wantErr          error
	}{
		{
			name:             "get screeningMovie",
			screeningMovieID: entity.UUID("existing_screening_movie_id"),
			want: &entity.ScreeningMovie{
				ID:        entity.UUID("existing_screening_movie_id"),
				CinemaID:  entity.UUID("existing_cinema_id"),
				MovieID:   entity.UUID("existing_movie_id"),
				StartTime: time.Date(2023, 1, 1, 15, 0, 0, 0, time.UTC),
				EndTime:   time.Date(2023, 1, 1, 17, 0, 0, 0, time.UTC),
			},
			wantErr: nil,
		},
		{
			name:             "not exist screeningMovie",
			screeningMovieID: entity.UUID("not_exist_screening_movie_id"),
			want:             nil,
			wantErr:          entity.ErrScreeningMovieNotFound,
		},
	}

	for _, tt := range tests {
		r := NewScreeningMovieRepository(db)
		got, err := r.FindByID(context.Background(), tt.screeningMovieID)
		if !errors.Is(err, tt.wantErr) {
			t.Errorf("ScreeningMovieRepository.FindByID() error is %v, want error is %v", err, tt.wantErr)
		}
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("ScreeningMovieRepository.FindByID() got is %v, want is %v", got, tt.want)
		}
	}
}

func TestScreeningMovieRepository_Create(t *testing.T) {
	db := prepareTestScreeningMovieRepository(t)
	tests := []struct {
		name           string
		screeningMovie *entity.ScreeningMovie
		wantErr        error
	}{
		{
			name: "create screeningMovie",
			screeningMovie: &entity.ScreeningMovie{
				ID:        entity.UUID("new_screening_movie_id"),
				CinemaID:  entity.UUID("new_cinema_id"),
				MovieID:   entity.UUID("new_movie_id"),
				StartTime: time.Date(2023, 1, 1, 15, 0, 0, 0, time.UTC),
				EndTime:   time.Date(2023, 1, 1, 17, 0, 0, 0, time.UTC),
			},
			wantErr: nil,
		},
		{
			name: "existing screeningType",
			screeningMovie: &entity.ScreeningMovie{
				ID:        entity.UUID("existing_screening_movie_id"),
				CinemaID:  entity.UUID("existing_cinema_id"),
				MovieID:   entity.UUID("existing_movie_id"),
				StartTime: time.Date(2023, 1, 1, 15, 0, 0, 0, time.UTC),
				EndTime:   time.Date(2023, 1, 1, 17, 0, 0, 0, time.UTC),
			},
			wantErr: entity.ErrScreeningMovieAlreadyExisted,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewScreeningMovieRepository(db)
			if err := r.Create(context.Background(), tt.screeningMovie); !errors.Is(err, tt.wantErr) {
				t.Errorf("ScreeningMovieRepository.Create() error is %v, want error is %v", err, tt.wantErr)
			}
			t.Cleanup(func() {
				if _, err := db.NamedExec("DELETE FROM screening_movies WHERE id = :id", r.screeningMovieToDTO(tt.screeningMovie)); err != nil {
					t.Fatal(err)
				}
			})
		})
	}
}

func prepareTestScreeningMovieRepository(t *testing.T) *sqlx.DB {
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

	screeningMovie := &screeningMovieDTO{
		ID:        "existing_screening_movie_id",
		CinemaID:  "existing_cinema_id",
		MovieID:   "existing_movie_id",
		StartTime: time.Date(2023, 1, 1, 15, 0, 0, 0, time.UTC),
		EndTime:   time.Date(2023, 1, 1, 17, 0, 0, 0, time.UTC),
	}
	if _, err := db.NamedExec("INSERT INTO screening_movies (id, cinema_id, movie_id, start_time, end_time) VALUES (:id, :cinema_id, :movie_id, :start_time, :end_time)", screeningMovie); err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() {
		if _, err := db.NamedExec("DELETE FROM screening_movies WHERE id = :id", screeningMovie); err != nil {
			t.Fatal(err)
		}
	})

	return db
}
