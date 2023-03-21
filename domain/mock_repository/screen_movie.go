// Code generated by MockGen. DO NOT EDIT.
// Source: screen_movie.go

// Package mock_repository is a generated GoMock package.
package mock_repository

import (
	context "context"
	reflect "reflect"

	entity "github.com/candy12t/cinemarch-server/domain/entity"
	gomock "github.com/golang/mock/gomock"
)

// MockScreenMovie is a mock of ScreenMovie interface.
type MockScreenMovie struct {
	ctrl     *gomock.Controller
	recorder *MockScreenMovieMockRecorder
}

// MockScreenMovieMockRecorder is the mock recorder for MockScreenMovie.
type MockScreenMovieMockRecorder struct {
	mock *MockScreenMovie
}

// NewMockScreenMovie creates a new mock instance.
func NewMockScreenMovie(ctrl *gomock.Controller) *MockScreenMovie {
	mock := &MockScreenMovie{ctrl: ctrl}
	mock.recorder = &MockScreenMovieMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockScreenMovie) EXPECT() *MockScreenMovieMockRecorder {
	return m.recorder
}

// CreateScreenMovie mocks base method.
func (m *MockScreenMovie) CreateScreenMovie(ctx context.Context, screenMovie *entity.ScreenMovie) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateScreenMovie", ctx, screenMovie)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateScreenMovie indicates an expected call of CreateScreenMovie.
func (mr *MockScreenMovieMockRecorder) CreateScreenMovie(ctx, screenMovie interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateScreenMovie", reflect.TypeOf((*MockScreenMovie)(nil).CreateScreenMovie), ctx, screenMovie)
}

// CreateScreenSchedules mocks base method.
func (m *MockScreenMovie) CreateScreenSchedules(ctx context.Context, screenSchedules entity.ScreenSchedules) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateScreenSchedules", ctx, screenSchedules)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateScreenSchedules indicates an expected call of CreateScreenSchedules.
func (mr *MockScreenMovieMockRecorder) CreateScreenSchedules(ctx, screenSchedules interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateScreenSchedules", reflect.TypeOf((*MockScreenMovie)(nil).CreateScreenSchedules), ctx, screenSchedules)
}

// FindByUniqKey mocks base method.
func (m *MockScreenMovie) FindByUniqKey(ctx context.Context, cinemaID, movieID entity.UUID, screenType entity.ScreenType, translateType entity.TranslateType, threeD bool) (*entity.ScreenMovie, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByUniqKey", ctx, cinemaID, movieID, screenType, translateType, threeD)
	ret0, _ := ret[0].(*entity.ScreenMovie)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByUniqKey indicates an expected call of FindByUniqKey.
func (mr *MockScreenMovieMockRecorder) FindByUniqKey(ctx, cinemaID, movieID, screenType, translateType, threeD interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByUniqKey", reflect.TypeOf((*MockScreenMovie)(nil).FindByUniqKey), ctx, cinemaID, movieID, screenType, translateType, threeD)
}

// Search mocks base method.
func (m *MockScreenMovie) Search(ctx context.Context) (entity.ScreenMovies, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Search", ctx)
	ret0, _ := ret[0].(entity.ScreenMovies)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Search indicates an expected call of Search.
func (mr *MockScreenMovieMockRecorder) Search(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Search", reflect.TypeOf((*MockScreenMovie)(nil).Search), ctx)
}
