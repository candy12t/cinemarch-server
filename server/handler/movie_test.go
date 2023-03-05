package handler

import (
	"context"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/candy12t/cinemarch-server/domain/entity"
	"github.com/candy12t/cinemarch-server/domain/mock_repository"
	"github.com/candy12t/cinemarch-server/usecase"
	"github.com/go-chi/chi/v5"
	"github.com/golang/mock/gomock"
)

func TestMovieHandler_Show(t *testing.T) {
	tests := []struct {
		name                   string
		movieID                string
		prepareMockMovieRepoFn func(m *mock_repository.MockMovie)
		wantResp               string
		wantCode               int
	}{
		{
			name:    "get existing movie",
			movieID: "existing_movie_id",
			prepareMockMovieRepoFn: func(m *mock_repository.MockMovie) {
				m.EXPECT().FindByID(gomock.Any(), entity.UUID("existing_movie_id")).Return(&entity.Movie{
					ID:            "existing_movie_id",
					Title:         "existing_movie_title",
					ReleaseDate:   time.Date(2022, 12, 31, 15, 0, 0, 0, time.UTC),
					ReleaseStatus: entity.ComingSoon,
				}, nil)
			},
			wantResp: `{"id":"existing_movie_id","title":"existing_movie_title","release_date":"2023-01-01","release_status":"COMING SOON"}`,
			wantCode: 200,
		},
		{
			name:    "not found movie",
			movieID: "not_exist_movie_id",
			prepareMockMovieRepoFn: func(m *mock_repository.MockMovie) {
				m.EXPECT().FindByID(gomock.Any(), entity.UUID("not_exist_movie_id")).Return(nil, entity.ErrMovieNotFound)
			},
			wantResp: `{"message":"movie not found"}`,
			wantCode: 404,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodGet, "/", nil)
			rec := httptest.NewRecorder()

			rctx := chi.NewRouteContext()
			rctx.URLParams.Add("movieID", tt.movieID)
			ctx := context.WithValue(req.Context(), chi.RouteCtxKey, rctx)
			req = req.WithContext(ctx)

			ctrl := gomock.NewController(t)
			mockMovieRepo := mock_repository.NewMockMovie(ctrl)
			tt.prepareMockMovieRepoFn(mockMovieRepo)
			movieUC := usecase.NewMovieUseCase(mockMovieRepo)
			h := NewMovieHandler(movieUC)

			h.Show(rec, req)

			if rec.Code != tt.wantCode {
				t.Errorf("MovieHandler.Show() return code is %v, wantCode is %v", rec.Code, tt.wantCode)
			}

			if string(rec.Body.Bytes()) != tt.wantResp {
				t.Errorf("MovieHandler.Show() return response is %v, wantResp is %v", string(rec.Body.Bytes()), tt.wantResp)
			}
		})
	}
}

func TestMovieHandler_Upsert(t *testing.T) {
	tests := []struct {
		name                   string
		body                   string
		prepareMockMovieRepoFn func(m *mock_repository.MockMovie)
		stubUUID               string
		wantResp               string
		wantCode               int
	}{
		{
			name: "create new movie",
			body: `{"title":"new_movie_title","release_date":"2023-01-01","release_status":"COMING SOON"}`,
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
			wantResp: `{"id":"new_movie_id","title":"new_movie_title","release_date":"2023-01-01","release_status":"COMING SOON"}`,
			wantCode: 200,
		},
		{
			name: "update existing movie",
			body: `{"title":"existing_movie_title","release_date":"2023-01-01","release_status":"NOW OPEN"}`,
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
			wantResp: `{"id":"existing_movie_id","title":"existing_movie_title","release_date":"2023-01-01","release_status":"NOW OPEN"}`,
			wantCode: 200,
		},
		{
			name: "not change release status",
			body: `{"title":"existing_movie_title","release_date":"2023-01-01","release_status":"RELEASED"}`,
			prepareMockMovieRepoFn: func(m *mock_repository.MockMovie) {
				m.EXPECT().FindByTitle(gomock.Any(), entity.MovieTitle("existing_movie_title")).Return(&entity.Movie{
					ID:            "existing_movie_id",
					Title:         "existing_movie_title",
					ReleaseDate:   time.Date(2022, 12, 31, 15, 0, 0, 0, time.UTC),
					ReleaseStatus: entity.ComingSoon,
				}, nil)
			},
			stubUUID: "",
			wantResp: `{"message":"can not change release status"}`,
			wantCode: 400,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			entity.NewUUID = func() entity.UUID {
				return entity.UUID(tt.stubUUID)
			}

			req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(tt.body))
			rec := httptest.NewRecorder()

			ctrl := gomock.NewController(t)
			mockMovieRepo := mock_repository.NewMockMovie(ctrl)
			tt.prepareMockMovieRepoFn(mockMovieRepo)
			movieUC := usecase.NewMovieUseCase(mockMovieRepo)
			h := NewMovieHandler(movieUC)

			h.Upsert(rec, req)

			if rec.Code != tt.wantCode {
				t.Errorf("MovieHandler.Upsert() return code is %v, wantCode is %v", rec.Code, tt.wantCode)
			}

			if string(rec.Body.Bytes()) != tt.wantResp {
				t.Errorf("MovieHandler.Upsert() return response is %v, wantResp is %v", string(rec.Body.Bytes()), tt.wantResp)
			}
		})
	}
}
