package service

import (
	"github.com/google/uuid"
	"github.com/thalerngsak/git-scanner/repository"
)

type repositoryService struct {
	resRepo repository.RepositoryStore
}

func NewRepositoryService(resRepo repository.RepositoryStore) RepositoryService {
	return repositoryService{resRepo: resRepo}
}

func (s repositoryService) NewRepository(request RepositoryRequest) (*RepositoryResponse, error) {

	repositories := &repository.Repository{
		ID:   uuid.New().String(),
		Name: request.Name,
		URL:  request.URL,
	}

	if err := s.resRepo.Create(repositories); err != nil {
		return nil, err
	}

	response := &RepositoryResponse{
		ID:   repositories.ID,
		Name: repositories.Name,
		URL:  repositories.URL,
	}

	return response, nil
}

func (s repositoryService) GetRepository() ([]*RepositoryResponse, error) {
	repositories, err := s.resRepo.GetAll()
	if err != nil {
		return nil, err
	}
	var repos []*RepositoryResponse
	for _, v := range repositories {
		response := &RepositoryResponse{
			ID:   v.ID,
			Name: v.Name,
			URL:  v.URL,
		}
		repos = append(repos, response)
	}
	return repos, nil
}

func (s repositoryService) GetRepositoryByID(id string) (*RepositoryResponse, error) {
	repositories, err := s.resRepo.GetByID(id)
	if err != nil {
		return nil, err
	}
	response := &RepositoryResponse{
		ID:   repositories.ID,
		Name: repositories.Name,
		URL:  repositories.URL,
	}
	return response, nil
}

func (s repositoryService) UpdateRepository(id string, request *RepositoryRequest) error {
	repositories := &repository.Repository{
		ID:   id,
		Name: request.Name,
		URL:  request.URL,
	}
	err := s.resRepo.Update(repositories)
	if err != nil {
		return err
	}
	return nil
}

func (s repositoryService) DeleteRepository(id string) error {
	err := s.resRepo.Delete(id)
	if err != nil {
		return err
	}
	return nil
}
