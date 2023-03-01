package mock_repository

import (
	repository "github.com/thalerngsak/git-scanner/repository"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockResultRepository is a mock of CustomerRepository interface.
type MockResultRepository struct {
	ctrl     *gomock.Controller
	recorder *MockResultRepositoryMockRecorder
}

func (m *MockResultRepository) GetByID(s string) (*repository.Result, error) {
	//TODO implement me
	panic("implement me")
}

func (m *MockResultRepository) Create(res *repository.Result) error {
	//TODO implement me
	panic("implement me")
}

// GetAll mocks base method.
func (m *MockResultRepository) GetAll() ([]*repository.Result, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAll")
	ret0, _ := ret[0].([]*repository.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (m *MockResultRepository) GetByRepositoryID(id string) ([]*repository.Result, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByRepositoryID")
	ret0, _ := ret[0].([]*repository.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (m *MockResultRepository) Update(res *repository.Result) error {
	//TODO implement me
	panic("implement me")
}

// MockResultRepositoryMockRecorder is the mock recorder for MockResultRepository.
type MockResultRepositoryMockRecorder struct {
	mock *MockResultRepository
}

// NewMockResultRepository creates a new mock instance.
func NewMockResultRepository(ctrl *gomock.Controller) *MockResultRepository {
	mock := &MockResultRepository{ctrl: ctrl}
	mock.recorder = &MockResultRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockResultRepository) EXPECT() *MockResultRepositoryMockRecorder {
	return m.recorder
}

// GetAll indicates an expected call of GetAll.
func (mr *MockResultRepositoryMockRecorder) GetAll() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAll", reflect.TypeOf((*MockResultRepository)(nil).GetAll))
}

// GetById mocks base method.
func (m *MockResultRepository) GetById(arg0 int) (*repository.Result, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetById", arg0)
	ret0, _ := ret[0].(*repository.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetById indicates an expected call of GetById.
func (mr *MockResultRepositoryMockRecorder) GetById(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetById", reflect.TypeOf((*MockResultRepository)(nil).GetById), arg0)
}

// GetByRepositoryID indicates an expected call of GetResultByRepositoryID.
func (mr *MockResultRepositoryMockRecorder) GetByRepositoryID(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByRepositoryID", reflect.TypeOf((*MockResultRepository)(nil).GetByRepositoryID), arg0)
}
