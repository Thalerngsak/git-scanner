package mock_service

import (
	"github.com/stretchr/testify/mock"
	"github.com/thalerngsak/git-scanner/service"
)

type MockRepositoryService struct {
	mock.Mock
}

func (m *MockRepositoryService) NewRepository(request service.RepositoryRequest) (*service.RepositoryResponse, error) {
	args := m.Called(request)
	return args.Get(0).(*service.RepositoryResponse), args.Error(1)
}

func (m *MockRepositoryService) GetRepository() ([]*service.RepositoryResponse, error) {
	args := m.Called()
	return args.Get(0).([]*service.RepositoryResponse), args.Error(1)
}

func (m *MockRepositoryService) GetRepositoryByID(id string) (*service.RepositoryResponse, error) {
	args := m.Called(id)
	return args.Get(0).(*service.RepositoryResponse), args.Error(1)
}

func (m *MockRepositoryService) UpdateRepository(id string, request *service.RepositoryRequest) error {
	args := m.Called(id, request)
	return args.Error(0)
}

func (m *MockRepositoryService) DeleteRepository(id string) error {
	args := m.Called(id)
	return args.Error(0)
}
