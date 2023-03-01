package mock_repository

import "github.com/thalerngsak/git-scanner/repository"

type MockRepositoryRepository struct {
	CreateFn     func(*repository.Repository) error
	GetAllFn     func() ([]*repository.Repository, error)
	GetByIDFn    func(string) (*repository.Repository, error)
	UpdateFn     func(*repository.Repository) error
	DeleteByIDFn func(string) error
}

func (r *MockRepositoryRepository) Delete(id string) error {
	return r.DeleteByIDFn(id)
}

func (r *MockRepositoryRepository) Create(repository *repository.Repository) error {
	if r.CreateFn == nil {
		return nil
	}
	return r.CreateFn(repository)
}

func (r *MockRepositoryRepository) GetAll() ([]*repository.Repository, error) {
	if r.GetAllFn == nil {
		return nil, nil
	}
	return r.GetAllFn()
}

func (r *MockRepositoryRepository) GetByID(id string) (*repository.Repository, error) {
	if r.GetByIDFn == nil {
		return nil, nil
	}
	return r.GetByIDFn(id)
}

func (r *MockRepositoryRepository) Update(repository *repository.Repository) error {
	if r.UpdateFn == nil {
		return nil
	}
	return r.UpdateFn(repository)
}

func (r *MockRepositoryRepository) DeleteByID(id string) error {
	if r.DeleteByIDFn == nil {
		return nil
	}
	return r.DeleteByIDFn(id)
}

func (r *MockRepositoryRepository) DeleteFunc(id string) error {
	return r.DeleteFunc(id)
}

func (r *MockRepositoryRepository) DeleteCalls() {
	r.DeleteCalls()
}
