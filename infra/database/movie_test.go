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

func TestMovieRepository_FindByID(t *testing.T) {
	db := prepareTestMovieRepository(t)
	tests := []struct {
		name    string
		id      entity.UUID
		want    *entity.Movie
		wantErr error
	}{
		{
			name: "get existing movie by id",
			id:   "existing_movie_id",
			want: &entity.Movie{
				ID:            "existing_movie_id",
				Title:         "existing_movie_title",
				ReleaseDate:   time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
				ReleaseStatus: entity.ComingSoon,
			},
			wantErr: nil,
		},
		{
			name:    "movie not found",
			id:      "not_exist_movie_id",
			want:    nil,
			wantErr: entity.ErrMovieNotFound,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := NewMovieRepository(db)
			got, err := repo.FindByID(context.Background(), tt.id)
			if err != nil {
				if !errors.Is(err, tt.wantErr) {
					t.Errorf("MovieRepository.FindByID() error is %v, wantErr is %v", err, tt.wantErr)
				}
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MovieRepository.FindByID() got is %v, want is %v", got, tt.want)
			}
		})
	}
}

func TestMovieRepository_FindByTitle(t *testing.T) {
	db := prepareTestMovieRepository(t)
	tests := []struct {
		name    string
		title   entity.MovieTitle
		want    *entity.Movie
		wantErr error
	}{
		{
			name:  "get existing movie by title",
			title: "existing_movie_title",
			want: &entity.Movie{
				ID:            "existing_movie_id",
				Title:         "existing_movie_title",
				ReleaseDate:   time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
				ReleaseStatus: entity.ComingSoon,
			},
			wantErr: nil,
		},
		{
			name:    "movie not found",
			title:   "not_exist_movie_title",
			want:    nil,
			wantErr: entity.ErrMovieNotFound,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := NewMovieRepository(db)
			got, err := repo.FindByTitle(context.Background(), tt.title)
			if err != nil {
				if !errors.Is(err, tt.wantErr) {
					t.Errorf("MovieRepository.FindByTitle() error is %v, wantErr is %v", err, tt.wantErr)
				}
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MovieRepository.FindByTitle() got is %v, want is %v", got, tt.want)
			}
		})
	}
}

func TestMovieRepository_Search(t *testing.T) {
	db := prepareTestMovieRepository(t)
	tests := []struct {
		name    string
		query   string
		args    []any
		want    entity.Movies
		wantErr error
	}{
		{
			name:  "get existing movies by title",
			query: "WHERE title LIKE ?",
			args:  []any{"%title%"},
			want: entity.Movies{
				{
					ID:            "existing_movie_id",
					Title:         "existing_movie_title",
					ReleaseDate:   time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
					ReleaseStatus: entity.ComingSoon,
				},
			},
			wantErr: nil,
		},
		{
			name:    "not found",
			query:   "WHERE title LIKE ?",
			args:    []any{"%not%"},
			want:    nil,
			wantErr: entity.ErrMovieNotFound,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := NewMovieRepository(db)
			got, err := repo.Search(context.Background(), tt.query, tt.args)
			if err != nil {
				if !errors.Is(err, tt.wantErr) {
					t.Errorf("MovieRepository.Search() error is %v, wantErr is %v", err, tt.wantErr)
				}
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MovieRepository.Search() got is %v, want is %v", got, tt.want)
			}
		})
	}
}

func TestMovieRepository_Create(t *testing.T) {
	db := prepareTestMovieRepository(t)
	tests := []struct {
		name    string
		movie   *entity.Movie
		wantErr error
	}{
		{
			name: "create new movie",
			movie: &entity.Movie{
				ID:            "new_movie_id",
				Title:         "new_movie_title",
				ReleaseDate:   time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
				ReleaseStatus: entity.ComingSoon,
			},
			wantErr: nil,
		},
		{
			name: "already existed",
			movie: &entity.Movie{
				ID:            "new_movie_id",
				Title:         "existing_movie_title",
				ReleaseDate:   time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
				ReleaseStatus: entity.ComingSoon,
			},
			wantErr: entity.ErrMovieAlreadyExisted,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := NewMovieRepository(db)
			if err := repo.Create(context.Background(), tt.movie); !errors.Is(err, tt.wantErr) {
				t.Errorf("MovieRepository.Create() errors is %v, wantErr is %v", err, tt.wantErr)
			}

			t.Cleanup(func() {
				if _, err := db.Exec(`DELETE FROM movies WHERE id = ?`, tt.movie.ID.String()); err != nil {
					t.Fatal(err)
				}
			})
		})
	}
}

func TestMovieRepository_Update(t *testing.T) {
	db := prepareTestMovieRepository(t)
	tests := []struct {
		name    string
		movie   *entity.Movie
		wantErr error
	}{
		{
			name: "update movie",
			movie: &entity.Movie{
				ID:            "existing_movie_id",
				Title:         "update_movie_title",
				ReleaseDate:   time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
				ReleaseStatus: entity.NowOpen,
			},
			wantErr: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := NewMovieRepository(db)
			if err := repo.Update(context.Background(), tt.movie); !errors.Is(err, tt.wantErr) {
				t.Errorf("MovieRepository.Update() errors is %v, wantErr is %v", err, tt.wantErr)
			}
		})
	}
}

func prepareTestMovieRepository(t *testing.T) *sqlx.DB {
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

	movie := &movieDTO{
		ID:            "existing_movie_id",
		Title:         "existing_movie_title",
		ReleaseDate:   time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
		ReleaseStatus: "COMING SOON",
	}
	if _, err := db.NamedExec(`INSERT INTO movies (id, title, release_date, release_status) VALUES (:id, :title, :release_date, :release_status)`, movie); err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() {
		if _, err := db.NamedExec(`DELETE FROM movies WHERE id = :id`, movie); err != nil {
			t.Fatal(err)
		}
	})

	return db
}
