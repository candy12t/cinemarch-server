package database

import (
	"context"
	"errors"
	"reflect"
	"testing"
	"time"

	"github.com/candy12t/cinemarch-server/domain/entity"
)

func TestMovieRepository_FindByID(t *testing.T) {
	db, cleanup, err := NewDB()
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() {
		if err := cleanup(); err != nil {
			t.Fatal(err)
		}
	})

	movie := &movieDTO{
		ID:            "existing_movie_id",
		Title:         "RRR",
		ReleaseDate:   time.Date(2022, 10, 21, 0, 0, 0, 0, time.UTC),
		ReleaseStatus: "NOW OPEN",
	}
	if _, err := db.NamedExec("INSERT movies (id, title, release_date, release_status) VALUES (:id, :title, :release_date, :release_status)", movie); err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() {
		if _, err := db.NamedExec("DELETE FROM movies where id = :id", movie); err != nil {
			t.Fatal(err)
		}
	})

	tests := []struct {
		name    string
		id      entity.UUID
		want    *entity.Movie
		wantErr error
	}{
		{
			name: "be able to get existing movie",
			id:   entity.UUID("existing_movie_id"),
			want: &entity.Movie{
				ID:            entity.UUID("existing_movie_id"),
				Title:         entity.MovieTitle("RRR"),
				ReleaseDate:   time.Date(2022, 10, 21, 0, 0, 0, 0, time.UTC),
				ReleaseStatus: entity.NowOpen,
			},
			wantErr: nil,
		},
		{
			name:    "get error is ErrMovieNotFound when get not exist movie",
			id:      entity.UUID("not_exist_movie_id"),
			want:    nil,
			wantErr: entity.ErrMovieNotFound,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewMovieRepository(db)

			got, err := r.FindByID(context.Background(), tt.id)
			if err != nil {
				if !errors.Is(err, tt.wantErr) {
					t.Fatal(err)
				}
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MovieRepository.FindByID() got is %v, want is %v", got, tt.want)
			}
		})
	}
}

func TestMovieRepository_Create(t *testing.T) {
	db, cleanup, err := NewDB()
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() {
		if err := cleanup(); err != nil {
			t.Fatal(err)
		}
	})

	movie := &movieDTO{
		ID:            "existing_movie_id",
		Title:         "RRR",
		ReleaseDate:   time.Date(2022, 10, 21, 0, 0, 0, 0, time.UTC),
		ReleaseStatus: "NOW OPEN",
	}
	if _, err := db.NamedExec("INSERT movies (id, title, release_date, release_status) VALUES (:id, :title, :release_date, :release_status)", movie); err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() {
		if _, err := db.NamedExec("DELETE FROM movies WHERE id = :id", movie); err != nil {
			t.Fatal(err)
		}
	})

	tests := []struct {
		name    string
		movie   *entity.Movie
		wantErr error
	}{
		{
			name: "be able to create new movie",
			movie: &entity.Movie{
				ID:            entity.UUID("new_movie_id"),
				Title:         entity.MovieTitle("new_movie_title"),
				ReleaseDate:   time.Date(2022, 01, 01, 0, 0, 0, 0, time.UTC),
				ReleaseStatus: entity.ComingSoon,
			},
			wantErr: nil,
		},
		{
			name: "return has already existed",
			movie: &entity.Movie{
				ID:            entity.UUID("existing_movie_id"),
				Title:         entity.MovieTitle("RRR"),
				ReleaseDate:   time.Date(2022, 10, 21, 0, 0, 0, 0, time.UTC),
				ReleaseStatus: entity.NowOpen,
			},
			wantErr: entity.ErrMovieAlreadyExisted,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewMovieRepository(db)
			if err := r.Create(context.Background(), tt.movie); !errors.Is(err, tt.wantErr) {
				t.Errorf("MovieRepository.Create() got error is %v, want error is %v", err, tt.wantErr)
			}
			t.Cleanup(func() {
				if _, err := db.NamedExec("DELETE FROM movies WHERE id = :id", r.movieToDTO(tt.movie)); err != nil {
					t.Fatal(err)
				}
			})
		})
	}
}

func TestMovieRepository_Update(t *testing.T) {
	db, cleanup, err := NewDB()
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() {
		if err := cleanup(); err != nil {
			t.Fatal(err)
		}
	})

	movie := &movieDTO{
		ID:            "existing_movie_id",
		Title:         "RRR",
		ReleaseDate:   time.Date(2022, 10, 21, 0, 0, 0, 0, time.UTC),
		ReleaseStatus: "NOW OPEN",
	}
	if _, err := db.NamedExec("INSERT movies (id, title, release_date, release_status) VALUES (:id, :title, :release_date, :release_status)", movie); err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() {
		if _, err := db.NamedExec("DELETE FROM movies WHERE id = :id", movie); err != nil {
			t.Fatal(err)
		}
	})

	tests := []struct {
		name    string
		movie   *entity.Movie
		wantErr error
	}{
		{
			name: "success update movie",
			movie: &entity.Movie{
				ID:            "existing_movie_id",
				Title:         "RRR",
				ReleaseDate:   time.Date(2022, 10, 21, 0, 0, 0, 0, time.UTC),
				ReleaseStatus: entity.Released,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewMovieRepository(db)
			if err := r.Update(context.Background(), tt.movie); err != nil {
				t.Errorf("MovieRepository.Update() got error is %v, want error is %v", err, tt.wantErr)
			}
		})
	}
}
