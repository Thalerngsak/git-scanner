package mock_service

import (
	"github.com/stretchr/testify/mock"
	"github.com/thalerngsak/git-scanner/service"
)

type MockResultService struct {
	mock.Mock
}

func (m *MockResultService) Create(request service.ResultRequest) error {
	args := m.Called(request)
	return args.Error(0)
}

func (m *MockResultService) GetResult() ([]*service.ResultResponse, error) {
	args := m.Called()
	return args.Get(0).([]*service.ResultResponse), args.Error(1)
}

func (m *MockResultService) GetResultByID(id string) (*service.ResultResponse, error) {
	args := m.Called(id)
	return args.Get(0).(*service.ResultResponse), args.Error(1)
}

func (m *MockResultService) GetResultByRepositoryID(id string) ([]*service.ResultResponse, error) {
	args := m.Called(id)
	return args.Get(0).([]*service.ResultResponse), args.Error(1)
}

func (m *MockResultService) UpdateResult(request *service.ResultRequest) error {
	args := m.Called(request)
	return args.Error(0)
}
