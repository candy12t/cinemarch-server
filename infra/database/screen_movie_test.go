package database

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/candy12t/cinemarch-server/domain/entity"
	"github.com/google/go-cmp/cmp"
	"github.com/jmoiron/sqlx"
)

func TestScreenMovieRepository_FindByUniqKey(t *testing.T) {
	db := prepareTestScreenMovieRepository(t)

	tests := []struct {
		name string
		args struct {
			cinemaID      entity.UUID
			movieID       entity.UUID
			screenType    entity.ScreenType
			translateType entity.TranslateType
			threeD        bool
		}
		want    *entity.ScreenMovie
		wantErr error
	}{
		{
			name: "get existing screen movie",
			args: struct {
				cinemaID      entity.UUID
				movieID       entity.UUID
				screenType    entity.ScreenType
				translateType entity.TranslateType
				threeD        bool
			}{
				cinemaID:      "existing_cinema_id",
				movieID:       "existing_movie_id",
				screenType:    entity.IMAX,
				translateType: entity.Subtitle,
				threeD:        false,
			},
			want: &entity.ScreenMovie{
				ID:            "existing_screen_movie_id",
				CinemaID:      "existing_cinema_id",
				MovieID:       "existing_movie_id",
				ScreenType:    entity.IMAX,
				TranslateType: entity.Subtitle,
				TreeD:         false,
				ScreenSchedules: entity.ScreenSchedules{
					{
						ID:            "existing_screen_schedule_id_1",
						ScreenMovieID: "existing_screen_movie_id",
						StartTime:     time.Date(2022, 12, 31, 15, 0, 0, 0, time.UTC),
						EndTime:       time.Date(2022, 12, 31, 17, 0, 0, 0, time.UTC),
					},
					{
						ID:            "existing_screen_schedule_id_2",
						ScreenMovieID: "existing_screen_movie_id",
						StartTime:     time.Date(2022, 12, 31, 18, 0, 0, 0, time.UTC),
						EndTime:       time.Date(2022, 12, 31, 20, 0, 0, 0, time.UTC),
					},
				},
			},
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := NewScreenMovieRepository(db)
			got, err := repo.FindByUniqKey(context.Background(), tt.args.cinemaID, tt.args.movieID, tt.args.screenType, tt.args.translateType, tt.args.threeD)
			if err != nil {
				if !errors.Is(err, tt.wantErr) {
					t.Errorf("ScreenMovieRepository.FindByUniqKey() error is %v, wantErr is %v", err, tt.wantErr)
					return
				}
			}
			if !cmp.Equal(got, tt.want) {
				t.Errorf("ScreenMovieRepository.FindByUniqKey() got is %v, want is %v", got, tt.want)
			}
		})
	}
}

func TestScreenMovieRepository_Search(t *testing.T) {
	t.Skip()
	tests := []struct {
		name string
	}{
		{
			name: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
		})
	}
}

func TestScreenMovieRepository_CreateScreenMovie(t *testing.T) {
	db := prepareTestScreenMovieRepository(t)
	tests := []struct {
		name        string
		screenMovie *entity.ScreenMovie
		wantErr     error
	}{
		{
			name: "create new screen movie",
			screenMovie: &entity.ScreenMovie{
				ID:            "new_screen_movie_id",
				CinemaID:      "new_cinema_id",
				MovieID:       "new_movie_id",
				ScreenType:    entity.IMAX,
				TranslateType: entity.Subtitle,
				TreeD:         false,
			},
			wantErr: nil,
		},
		{
			name: "already existed",
			screenMovie: &entity.ScreenMovie{
				ID:            "new_screen_movie_id",
				CinemaID:      "existing_cinema_id",
				MovieID:       "existing_movie_id",
				ScreenType:    entity.IMAX,
				TranslateType: entity.Subtitle,
				TreeD:         false,
			},
			wantErr: entity.ErrScreenMovieAlreadyExisted,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := NewScreenMovieRepository(db)
			if err := repo.CreateScreenMovie(context.Background(), tt.screenMovie); !errors.Is(err, tt.wantErr) {
				t.Errorf("ScreenMovieRepository.CreateScreenMovie() error is %v, wantErr is %v", err, tt.wantErr)
			}
			t.Cleanup(func() {
				if _, err := db.Exec(`delete from screen_movies where id = ?`, tt.screenMovie.ID.String()); err != nil {
					t.Fatal(err)
				}
			})
		})
	}
}

func TestScreenMovieRepository_CreateScreenSchedule(t *testing.T) {
	db := prepareTestScreenMovieRepository(t)
	tests := []struct {
		name            string
		screenSchedules entity.ScreenSchedules
		wantErr         error
	}{
		{
			name: "create new screen schedules",
			screenSchedules: entity.ScreenSchedules{
				{
					ID:            "new_screen_schedule_id_1",
					ScreenMovieID: "new_screen_movie_id",
					StartTime:     time.Date(2022, 12, 31, 15, 0, 0, 0, time.UTC),
					EndTime:       time.Date(2022, 12, 31, 17, 0, 0, 0, time.UTC),
				},
				{
					ID:            "new_screen_schedule_id_2",
					ScreenMovieID: "new_screen_movie_id",
					StartTime:     time.Date(2022, 12, 31, 18, 0, 0, 0, time.UTC),
					EndTime:       time.Date(2022, 12, 31, 20, 0, 0, 0, time.UTC),
				},
			},
			wantErr: nil,
		},
		{
			name: "already existed",
			screenSchedules: entity.ScreenSchedules{
				{
					ID:            "new_screen_schedule_id_1",
					ScreenMovieID: "existing_screen_movie_id",
					StartTime:     time.Date(2022, 12, 31, 15, 0, 0, 0, time.UTC),
					EndTime:       time.Date(2022, 12, 31, 17, 0, 0, 0, time.UTC),
				},
				{
					ID:            "new_screen_schedule_id_2",
					ScreenMovieID: "existing_screen_movie_id",
					StartTime:     time.Date(2022, 12, 31, 18, 0, 0, 0, time.UTC),
					EndTime:       time.Date(2022, 12, 31, 20, 0, 0, 0, time.UTC),
				},
			},
			wantErr: entity.ErrScreenScheduleAlreadyExisted,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := NewScreenMovieRepository(db)
			if err := repo.CreateScreenSchedules(context.Background(), tt.screenSchedules); !errors.Is(err, tt.wantErr) {
				t.Errorf("ScreenMovieRepository.CreateScreenSchedules() error is %v wantErr is %v", err, tt.wantErr)
			}
			t.Cleanup(func() {
				for _, schedule := range tt.screenSchedules {
					if _, err := db.Exec(`delete from screen_schedules where id = ?`, schedule.ID.String()); err != nil {
						t.Fatal(err)
					}
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
		ID:            "existing_screen_movie_id",
		CinemaID:      "existing_cinema_id",
		MovieID:       "existing_movie_id",
		ScreenType:    "IMAX",
		TranslateType: "Subtitle",
		ThreeD:        false,
		Schedules: schedulesDTOs{
			{
				ID:            "existing_screen_schedule_id_1",
				ScreenMovieID: "existing_screen_movie_id",
				StartTime:     "2022-12-31 15:00:00",
				EndTime:       "2022-12-31 17:00:00",
			},
			{
				ID:            "existing_screen_schedule_id_2",
				ScreenMovieID: "existing_screen_movie_id",
				StartTime:     "2022-12-31 18:00:00",
				EndTime:       "2022-12-31 20:00:00",
			},
		},
	}
	if _, err := db.NamedExec(`INSERT INTO screen_movies (id, cinema_id, movie_id, screen_type, translate_type, three_d) VALUES (:id, :cinema_id, :movie_id, :screen_type, :translate_type, :three_d)`, screenMovie); err != nil {
		t.Fatal(err)
	}
	for _, schedule := range screenMovie.Schedules {
		if _, err := db.NamedExec(`INSERT INTO screen_schedules (id, screen_movie_id, start_time, end_time) VALUES (:id, :screen_movie_id, :start_time, :end_time)`, schedule); err != nil {
			t.Fatal(err)
		}
	}
	t.Cleanup(func() {
		if _, err := db.NamedExec(`DELETE FROM screen_movies WHERE id = :id`, screenMovie); err != nil {
			t.Fatal(err)
		}
		for _, schedule := range screenMovie.Schedules {
			if _, err := db.NamedExec(`DELETE FROM screen_schedules WHERE id = :id`, schedule); err != nil {
				t.Fatal(err)
			}
		}
	})

	return db
}
