// Code generated by MockGen. DO NOT EDIT.
// Source: internal/list/pg_repository.go
//
// Generated by this command:
//
//	mockgen -source internal/list/pg_repository.go -destination internal/list/mock/pg_repository_mock.go -package mock
//

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"

	uuid "github.com/google/uuid"
	models "github.com/ithaquaKr/taskManager/internal/models"
	gomock "go.uber.org/mock/gomock"
)

// MockRepository is a mock of Repository interface.
type MockRepository struct {
	ctrl     *gomock.Controller
	recorder *MockRepositoryMockRecorder
}

// MockRepositoryMockRecorder is the mock recorder for MockRepository.
type MockRepositoryMockRecorder struct {
	mock *MockRepository
}

// NewMockRepository creates a new mock instance.
func NewMockRepository(ctrl *gomock.Controller) *MockRepository {
	mock := &MockRepository{ctrl: ctrl}
	mock.recorder = &MockRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRepository) EXPECT() *MockRepositoryMockRecorder {
	return m.recorder
}

// AllLists mocks base method.
func (m *MockRepository) AllLists(ctx context.Context, offset, limit int) ([]*models.List, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AllLists", ctx, offset, limit)
	ret0, _ := ret[0].([]*models.List)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AllLists indicates an expected call of AllLists.
func (mr *MockRepositoryMockRecorder) AllLists(ctx, offset, limit any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AllLists", reflect.TypeOf((*MockRepository)(nil).AllLists), ctx, offset, limit)
}

// CreateList mocks base method.
func (m *MockRepository) CreateList(ctx context.Context, list *models.List) (*models.List, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateList", ctx, list)
	ret0, _ := ret[0].(*models.List)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateList indicates an expected call of CreateList.
func (mr *MockRepositoryMockRecorder) CreateList(ctx, list any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateList", reflect.TypeOf((*MockRepository)(nil).CreateList), ctx, list)
}

// DeleteList mocks base method.
func (m *MockRepository) DeleteList(ctx context.Context, id uuid.UUID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteList", ctx, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteList indicates an expected call of DeleteList.
func (mr *MockRepositoryMockRecorder) DeleteList(ctx, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteList", reflect.TypeOf((*MockRepository)(nil).DeleteList), ctx, id)
}

// GetList mocks base method.
func (m *MockRepository) GetList(ctx context.Context, id uuid.UUID) (*models.List, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetList", ctx, id)
	ret0, _ := ret[0].(*models.List)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetList indicates an expected call of GetList.
func (mr *MockRepositoryMockRecorder) GetList(ctx, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetList", reflect.TypeOf((*MockRepository)(nil).GetList), ctx, id)
}

// UpdateList mocks base method.
func (m *MockRepository) UpdateList(ctx context.Context, list *models.List) (*models.List, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateList", ctx, list)
	ret0, _ := ret[0].(*models.List)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateList indicates an expected call of UpdateList.
func (mr *MockRepositoryMockRecorder) UpdateList(ctx, list any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateList", reflect.TypeOf((*MockRepository)(nil).UpdateList), ctx, list)
}
