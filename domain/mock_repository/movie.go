// Code generated by MockGen. DO NOT EDIT.
// Source: movie.go

// Package mock_repository is a generated GoMock package.
package mock_repository

import (
	context "context"
	reflect "reflect"

	entity "github.com/candy12t/cinemarch-server/domain/entity"
	gomock "github.com/golang/mock/gomock"
)

// MockMovie is a mock of Movie interface.
type MockMovie struct {
	ctrl     *gomock.Controller
	recorder *MockMovieMockRecorder
}

// MockMovieMockRecorder is the mock recorder for MockMovie.
type MockMovieMockRecorder struct {
	mock *MockMovie
}

// NewMockMovie creates a new mock instance.
func NewMockMovie(ctrl *gomock.Controller) *MockMovie {
	mock := &MockMovie{ctrl: ctrl}
	mock.recorder = &MockMovieMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMovie) EXPECT() *MockMovieMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockMovie) Create(ctx context.Context, movie *entity.Movie) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, movie)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockMovieMockRecorder) Create(ctx, movie interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockMovie)(nil).Create), ctx, movie)
}

// FindByID mocks base method.
func (m *MockMovie) FindByID(ctx context.Context, movieID entity.UUID) (*entity.Movie, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByID", ctx, movieID)
	ret0, _ := ret[0].(*entity.Movie)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByID indicates an expected call of FindByID.
func (mr *MockMovieMockRecorder) FindByID(ctx, movieID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByID", reflect.TypeOf((*MockMovie)(nil).FindByID), ctx, movieID)
}

// FindByTitle mocks base method.
func (m *MockMovie) FindByTitle(ctx context.Context, title entity.MovieTitle) (*entity.Movie, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByTitle", ctx, title)
	ret0, _ := ret[0].(*entity.Movie)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByTitle indicates an expected call of FindByTitle.
func (mr *MockMovieMockRecorder) FindByTitle(ctx, title interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByTitle", reflect.TypeOf((*MockMovie)(nil).FindByTitle), ctx, title)
}

// Search mocks base method.
func (m *MockMovie) Search(ctx context.Context, conditionQuery string, args []any) (entity.Movies, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Search", ctx, conditionQuery, args)
	ret0, _ := ret[0].(entity.Movies)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Search indicates an expected call of Search.
func (mr *MockMovieMockRecorder) Search(ctx, conditionQuery, args interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Search", reflect.TypeOf((*MockMovie)(nil).Search), ctx, conditionQuery, args)
}

// Update mocks base method.
func (m *MockMovie) Update(ctx context.Context, movie *entity.Movie) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", ctx, movie)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockMovieMockRecorder) Update(ctx, movie interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockMovie)(nil).Update), ctx, movie)
}
