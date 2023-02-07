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

// Create mocks base method.
func (m *MockScreenMovie) Create(ctx context.Context, screenMovie *entity.ScreenMovie) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, screenMovie)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockScreenMovieMockRecorder) Create(ctx, screenMovie interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockScreenMovie)(nil).Create), ctx, screenMovie)
}

// FindByID mocks base method.
func (m *MockScreenMovie) FindByID(ctx context.Context, screenMovieID entity.UUID) (*entity.ScreenMovie, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByID", ctx, screenMovieID)
	ret0, _ := ret[0].(*entity.ScreenMovie)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByID indicates an expected call of FindByID.
func (mr *MockScreenMovieMockRecorder) FindByID(ctx, screenMovieID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByID", reflect.TypeOf((*MockScreenMovie)(nil).FindByID), ctx, screenMovieID)
}