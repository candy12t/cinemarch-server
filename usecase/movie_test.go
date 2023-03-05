package usecase

import (
	"context"
	"errors"
	"reflect"
	"testing"
	"time"

	"github.com/candy12t/cinemarch-server/domain/entity"
	"github.com/candy12t/cinemarch-server/domain/mock_repository"
	"github.com/golang/mock/gomock"
)

func TestMovieUseCase_FindAllByTitle(t *testing.T) {
	tests := []struct {
		name                   string
		title                  string
		prepareMockMovieRepoFn func(m *mock_repository.MockMovie)
		wants                  MovieDTOs
		wantErr                error
	}{
		{
			name:  "get existing movies",
			title: "movie",
			prepareMockMovieRepoFn: func(m *mock_repository.MockMovie) {
				conditions := entity.Conditions{{Query: "title LIKE ?", Arg: "%movie%"}}
				query, args := conditions.Build()
				m.EXPECT().Search(gomock.Any(), query, args).Return(entity.Movies{
					{
						ID:            "existing_movie_id",
						Title:         "existing_movie_title",
						ReleaseDate:   time.Date(2022, 12, 31, 15, 0, 0, 0, time.UTC),
						ReleaseStatus: entity.ComingSoon,
					},
				}, nil)
			},
			wants: MovieDTOs{
				{
					ID:            "existing_movie_id",
					Title:         "existing_movie_title",
					ReleaseDate:   "2023-01-01",
					ReleaseStatus: "COMING SOON",
				},
			},
			wantErr: nil,
		},
		{
			name:  "not found movies",
			title: "not_exist",
			prepareMockMovieRepoFn: func(m *mock_repository.MockMovie) {
				conditions := entity.Conditions{{Query: "title LIKE ?", Arg: "%not_exist%"}}
				query, args := conditions.Build()
				m.EXPECT().Search(gomock.Any(), query, args).Return(nil, entity.ErrMovieNotFound)
			},
			wants:   nil,
			wantErr: entity.ErrMovieNotFound,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			mockMovieRepo := mock_repository.NewMockMovie(ctrl)
			tt.prepareMockMovieRepoFn(mockMovieRepo)
			movieUC := NewMovieUseCase(mockMovieRepo)

			gots, err := movieUC.FindAllByTitle(context.Background(), tt.title)
			if err != nil {
				if !errors.Is(err, tt.wantErr) {
					t.Errorf("MovieUseCase.FindAllByTitle() error is %v, wantErr is %v", err, tt.wantErr)
				}
				return
			}
			if !reflect.DeepEqual(gots, tt.wants) {
				t.Errorf("MovieUseCase.FindAllByTitle() gots is %v, wants is %v", gots, tt.wants)
			}
		})
	}
}

func TestMovieUseCase_Upsert(t *testing.T) {
	tests := []struct {
		name                   string
		params                 UpsertMovieParams
		prepareMockMovieRepoFn func(m *mock_repository.MockMovie)
		stubUUID               string
		want                   *MovieDTO
		wantErr                error
	}{
		{
			name: "create new movie",
			params: UpsertMovieParams{
				Title:         "new_movie_title",
				ReleaseDate:   "2023-01-01",
				ReleaseStatus: "COMING SOON",
			},
			prepareMockMovieRepoFn: func(m *mock_repository.MockMovie) {
				m.EXPECT().FindByTitle(gomock.Any(), entity.MovieTitle("new_movie_title")).Return(nil, entity.ErrMovieNotFound)
				m.EXPECT().Create(gomock.Any(), &entity.Movie{
					ID:            "new_movie_id",
					Title:         "new_movie_title",
					ReleaseDate:   time.Date(2022, 12, 31, 15, 0, 0, 0, time.UTC),
					ReleaseStatus: entity.ComingSoon,
				}).Return(nil)
			},
			stubUUID: "new_movie_id",
			want: &MovieDTO{
				ID:            "new_movie_id",
				Title:         "new_movie_title",
				ReleaseDate:   "2023-01-01",
				ReleaseStatus: "COMING SOON",
			},
			wantErr: nil,
		},
		{
			name: "update movie",
			params: UpsertMovieParams{
				Title:         "existing_movie_title",
				ReleaseDate:   "2023-01-01",
				ReleaseStatus: "NOW OPEN",
			},
			prepareMockMovieRepoFn: func(m *mock_repository.MockMovie) {
				m.EXPECT().FindByTitle(gomock.Any(), entity.MovieTitle("existing_movie_title")).Return(&entity.Movie{
					ID:            "existing_movie_id",
					Title:         "existing_movie_title",
					ReleaseDate:   time.Date(2022, 12, 31, 15, 0, 0, 0, time.UTC),
					ReleaseStatus: entity.ComingSoon,
				}, nil)
				m.EXPECT().Update(gomock.Any(), &entity.Movie{
					ID:            "existing_movie_id",
					Title:         "existing_movie_title",
					ReleaseDate:   time.Date(2022, 12, 31, 15, 0, 0, 0, time.UTC),
					ReleaseStatus: entity.NowOpen,
				}).Return(nil)
			},
			stubUUID: "",
			want: &MovieDTO{
				ID:            "existing_movie_id",
				Title:         "existing_movie_title",
				ReleaseDate:   "2023-01-01",
				ReleaseStatus: "NOW OPEN",
			},
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			entity.NewUUID = func() entity.UUID {
				return entity.UUID(tt.stubUUID)
			}

			ctrl := gomock.NewController(t)
			mockMovieRepo := mock_repository.NewMockMovie(ctrl)
			tt.prepareMockMovieRepoFn(mockMovieRepo)
			movieUC := NewMovieUseCase(mockMovieRepo)

			got, err := movieUC.Upsert(context.Background(), tt.params)
			if err != nil {
				if !errors.Is(err, tt.wantErr) {
					t.Errorf("MovieUseCase.Upsert() error is %v, wantErr is %v", err, tt.wantErr)
				}
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MovieUseCase.Upsert() got is %v, want is %v", got, tt.want)
			}
		})
	}
}
