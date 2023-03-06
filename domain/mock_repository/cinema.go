// Code generated by MockGen. DO NOT EDIT.
// Source: cinema.go

// Package mock_repository is a generated GoMock package.
package mock_repository

import (
	context "context"
	reflect "reflect"

	entity "github.com/candy12t/cinemarch-server/domain/entity"
	gomock "github.com/golang/mock/gomock"
)

// MockCinema is a mock of Cinema interface.
type MockCinema struct {
	ctrl     *gomock.Controller
	recorder *MockCinemaMockRecorder
}

// MockCinemaMockRecorder is the mock recorder for MockCinema.
type MockCinemaMockRecorder struct {
	mock *MockCinema
}

// NewMockCinema creates a new mock instance.
func NewMockCinema(ctrl *gomock.Controller) *MockCinema {
	mock := &MockCinema{ctrl: ctrl}
	mock.recorder = &MockCinemaMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCinema) EXPECT() *MockCinemaMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockCinema) Create(ctx context.Context, cinema *entity.Cinema) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, cinema)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockCinemaMockRecorder) Create(ctx, cinema interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockCinema)(nil).Create), ctx, cinema)
}

// FindAllByPrefecture mocks base method.
func (m *MockCinema) FindAllByPrefecture(ctx context.Context, prefecture entity.Prefecture) (entity.Cinemas, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindAllByPrefecture", ctx, prefecture)
	ret0, _ := ret[0].(entity.Cinemas)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindAllByPrefecture indicates an expected call of FindAllByPrefecture.
func (mr *MockCinemaMockRecorder) FindAllByPrefecture(ctx, prefecture interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindAllByPrefecture", reflect.TypeOf((*MockCinema)(nil).FindAllByPrefecture), ctx, prefecture)
}

// FindByID mocks base method.
func (m *MockCinema) FindByID(ctx context.Context, cinemaID entity.UUID) (*entity.Cinema, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByID", ctx, cinemaID)
	ret0, _ := ret[0].(*entity.Cinema)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByID indicates an expected call of FindByID.
func (mr *MockCinemaMockRecorder) FindByID(ctx, cinemaID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByID", reflect.TypeOf((*MockCinema)(nil).FindByID), ctx, cinemaID)
}
