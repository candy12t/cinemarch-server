package database_test

import (
	"context"
	"testing"
	"time"

	"github.com/candy12t/cinema-search-server/domain/entity"
	"github.com/candy12t/cinema-search-server/infra/database"
)

func TestCreateMovie(t *testing.T) {
	db, cleanup, err := database.NewDB()
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() {
		if err := cleanup(); err != nil {
			t.Fatal(err)
		}
	})

	tests := []struct {
		name    string
		movie   *entity.Movie
		wantErr error
	}{
		{
			name: "success",
			movie: &entity.Movie{
				ID:            entity.NewUUID(),
				Title:         entity.MovieTitle("TENET"),
				ReleaseDate:   time.Date(2020, 9, 18, 0, 0, 0, 0, time.UTC),
				ReleaseStatus: entity.ReleaseStatusReleased,
			},
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := database.NewMovieRepository(db)
			err := r.Save(context.Background(), tt.movie)
			if err != tt.wantErr {
				t.Errorf("Create() error is %v, wantErr is %v", err, tt.wantErr)
			}
		})
	}
}
