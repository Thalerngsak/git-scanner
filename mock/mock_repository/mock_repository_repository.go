package mock_repository

import (
	repository "github.com/thalerngsak/git-scanner/repository"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockRepositoryRepository is a mock of AccountRepository interface.
type MockRepositoryRepository struct {
	ctrl     *gomock.Controller
	recorder *MockRepositoryRepositoryMockRecorder
}

// MockRepositoryRepositoryMockRecorder is the mock recorder for MockRepositoryRepository.
type MockRepositoryRepositoryMockRecorder struct {
	mock *MockRepositoryRepository
}

// NewMockAccountRepository creates a new mock instance.
func NewMockAccountRepository(ctrl *gomock.Controller) *MockRepositoryRepository {
	mock := &MockRepositoryRepository{ctrl: ctrl}
	mock.recorder = &MockRepositoryRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRepositoryRepository) EXPECT() *MockRepositoryRepositoryMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockRepositoryRepository) Create(arg0 repository.Repository) (*repository.Repository, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0)
	ret0, _ := ret[0].(*repository.Repository)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockRepositoryRepositoryMockRecorder) Create(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockRepositoryRepository)(nil).Create), arg0)
}

// GetAll mocks base method.
func (m *MockRepositoryRepository) GetAll(arg0 int) ([]repository.Repository, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAll", arg0)
	ret0, _ := ret[0].([]repository.Repository)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAll indicates an expected call of GetAll.
func (mr *MockRepositoryRepositoryMockRecorder) GetAll(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAll", reflect.TypeOf((*MockRepositoryRepository)(nil).GetAll), arg0)
}
