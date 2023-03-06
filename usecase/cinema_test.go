package usecase

import (
	"context"
	"errors"
	"reflect"
	"testing"

	"github.com/candy12t/cinemarch-server/domain/entity"
	"github.com/candy12t/cinemarch-server/domain/mock_repository"
	"github.com/golang/mock/gomock"
)

func TestCinemaUseCase_FindByID(t *testing.T) {
	tests := []struct {
		name                    string
		cinemaID                string
		prepareMockCinemaRepoFn func(m *mock_repository.MockCinema)
		want                    *CinemaDTO
		wantErr                 error
	}{
		{
			name:     "get existing cinema",
			cinemaID: "existing_cinema_id",
			prepareMockCinemaRepoFn: func(m *mock_repository.MockCinema) {
				m.EXPECT().FindByID(gomock.Any(), entity.UUID("existing_cinema_id")).Return(&entity.Cinema{
					ID:         "existing_cinema_id",
					Name:       "existing_cinema_name",
					Prefecture: "東京都",
					Address:    "東京都新宿区新宿1-1-1",
					WebSite:    "https://example.com",
				}, nil)
			},
			want: &CinemaDTO{
				ID:         "existing_cinema_id",
				Name:       "existing_cinema_name",
				Prefecture: "東京都",
				Address:    "東京都新宿区新宿1-1-1",
				WebSite:    "https://example.com",
			},
			wantErr: nil,
		},
		{
			name:     "not found cinema",
			cinemaID: "not_exist_cinema_id",
			prepareMockCinemaRepoFn: func(m *mock_repository.MockCinema) {
				m.EXPECT().FindByID(gomock.Any(), entity.UUID("not_exist_cinema_id")).Return(nil, entity.ErrCinemaNotFound)
			},
			want:    nil,
			wantErr: entity.ErrCinemaNotFound,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			mockCinemaRepo := mock_repository.NewMockCinema(ctrl)
			tt.prepareMockCinemaRepoFn(mockCinemaRepo)
			cinemaUC := NewCinemaUseCase(mockCinemaRepo)

			got, err := cinemaUC.FindByID(context.Background(), tt.cinemaID)
			if err != nil {
				if !errors.Is(err, tt.wantErr) {
					t.Errorf("CinemaUseCase.FindByID() error is %v, wantErr is %v", err, tt.wantErr)
					return
				}
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CinemaUseCase.FindByID() got is %v, want is %v", got, tt.want)
			}
		})
	}
}

func TestCinemaUseCase_Create(t *testing.T) {
	tests := []struct {
		name                    string
		params                  CreateCinemaParams
		prepareMockCinemaRepoFn func(m *mock_repository.MockCinema)
		stubUUID                string
		want                    *CinemaDTO
		wantErr                 error
	}{
		{
			name: "create new cinema",
			params: CreateCinemaParams{
				Name:       "new_cinema_name",
				Prefecture: "東京都",
				Address:    "東京都新宿区新宿1-1-1",
				WebSite:    "https://example.com",
			},
			prepareMockCinemaRepoFn: func(m *mock_repository.MockCinema) {
				m.EXPECT().Create(gomock.Any(), &entity.Cinema{
					ID:         "new_cinema_id",
					Name:       "new_cinema_name",
					Prefecture: "東京都",
					Address:    "東京都新宿区新宿1-1-1",
					WebSite:    "https://example.com",
				}).Return(nil)
			},
			stubUUID: "new_cinema_id",
			want: &CinemaDTO{
				ID:         "new_cinema_id",
				Name:       "new_cinema_name",
				Prefecture: "東京都",
				Address:    "東京都新宿区新宿1-1-1",
				WebSite:    "https://example.com",
			},
			wantErr: nil,
		},
		{
			name: "cinema has already existed",
			params: CreateCinemaParams{
				Name:       "existing_cinema_name",
				Prefecture: "東京都",
				Address:    "東京都新宿区新宿1-1-1",
				WebSite:    "https://example.com",
			},
			prepareMockCinemaRepoFn: func(m *mock_repository.MockCinema) {
				m.EXPECT().Create(gomock.Any(), &entity.Cinema{
					ID:         "new_cinema_id",
					Name:       "existing_cinema_name",
					Prefecture: "東京都",
					Address:    "東京都新宿区新宿1-1-1",
					WebSite:    "https://example.com",
				}).Return(entity.ErrCinemaAlreadyExisted)
			},
			stubUUID: "new_cinema_id",
			want:     nil,
			wantErr:  entity.ErrCinemaAlreadyExisted,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			entity.NewUUID = func() entity.UUID {
				return entity.UUID(tt.stubUUID)
			}

			ctrl := gomock.NewController(t)
			mockCinemaRepo := mock_repository.NewMockCinema(ctrl)
			tt.prepareMockCinemaRepoFn(mockCinemaRepo)
			cinemaUC := NewCinemaUseCase(mockCinemaRepo)

			got, err := cinemaUC.Create(context.Background(), tt.params)
			if err != nil {
				if !errors.Is(err, tt.wantErr) {
					t.Errorf("CinemaUseCase.Create() error is %v, wantErr is %v", err, tt.wantErr)
					return
				}
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CinemaUseCase.Create() got is %v, want is %v", got, tt.want)
			}
		})
	}
}
