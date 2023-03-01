package mock_repository

import "github.com/thalerngsak/git-scanner/repository"

type MockResultRepository struct {
	CreateFn            func(result repository.Result) error
	GetByIDFn           func(id string) (*repository.Result, error)
	UpdateFn            func(result *repository.Result) error
	DeleteByIDFn        func(id string) error
	GetAllFn            func() ([]repository.Result, error)
	DeleteAllFn         func() error
	GetByRepositoryIDFn func(id string) ([]*repository.Result, error)
}

func (r *MockResultRepository) GetByRepositoryID(id string) ([]*repository.Result, error) {
	return r.GetByRepositoryIDFn(id)
}

func (r *MockResultRepository) Create(result repository.Result) error {
	return r.CreateFn(result)
}

func (r *MockResultRepository) GetByID(id string) (*repository.Result, error) {
	return r.GetByIDFn(id)
}

func (r *MockResultRepository) Update(result *repository.Result) error {
	return r.UpdateFn(result)
}

func (r *MockResultRepository) DeleteByID(id string) error {
	return r.DeleteByIDFn(id)
}

func (r *MockResultRepository) GetAll() ([]repository.Result, error) {
	return r.GetAllFn()
}

func (r *MockResultRepository) DeleteAll() error {
	return r.DeleteAllFn()
}
