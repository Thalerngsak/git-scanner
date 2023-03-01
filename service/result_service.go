package service

import (
	"github.com/thalerngsak/git-scanner/repository"
)

type resultService struct {
	resRepo repository.ResultStore
}

func NewResultService(resRepo repository.ResultStore) ResultService {
	return resultService{resRepo: resRepo}
}

func (s resultService) Create(request ResultRequest) error {
	res := &repository.Result{
		ID:            request.ID,
		Status:        request.Status,
		RepositoryID:  request.RepositoryID,
		RepositoryURL: request.RepositoryURL,
		EnqueuedAt:    request.EnqueuedAt,
		StartedAt:     request.StartedAt,
	}
	err := s.resRepo.Create(res)
	if err != nil {
		return err
	}
	return nil
}

func (s resultService) GetResult() ([]*ResultResponse, error) {
	repositories, _ := s.resRepo.GetAll()
	var repos []*ResultResponse
	for _, v := range repositories {
		response := &ResultResponse{
			ID:            v.ID,
			Status:        v.Status,
			RepositoryID:  v.RepositoryID,
			RepositoryURL: v.RepositoryURL,
			Findings:      v.Findings,
			EnqueuedAt:    v.EnqueuedAt,
			StartedAt:     v.StartedAt,
			FinishedAt:    v.FinishedAt,
		}
		repos = append(repos, response)
	}
	return repos, nil
}
func (s resultService) GetResultByID(id string) (*ResultResponse, error) {
	response, err := s.resRepo.GetByID(id)
	if err != nil {
		return nil, err
	}
	res := ResultResponse{
		ID:            response.ID,
		Status:        response.Status,
		RepositoryID:  response.RepositoryID,
		RepositoryURL: response.RepositoryURL,
		Findings:      response.Findings,
		EnqueuedAt:    response.EnqueuedAt,
		StartedAt:     response.StartedAt,
		FinishedAt:    response.FinishedAt,
	}

	return &res, nil
}

func (s resultService) GetResultByRepositoryID(id string) ([]*ResultResponse, error) {
	repositories, err := s.resRepo.GetByRepositoryID(id)
	if err != nil {
		return nil, err
	}
	var repos []*ResultResponse
	for _, v := range repositories {
		response := &ResultResponse{
			ID:            v.ID,
			Status:        v.Status,
			RepositoryID:  v.RepositoryID,
			RepositoryURL: v.RepositoryURL,
			Findings:      v.Findings,
			EnqueuedAt:    v.EnqueuedAt,
			StartedAt:     v.StartedAt,
			FinishedAt:    v.FinishedAt,
		}
		repos = append(repos, response)
	}
	return repos, nil
}

func (s resultService) UpdateResult(request *ResultRequest) error {
	res := &repository.Result{
		ID:            request.ID,
		Status:        request.Status,
		RepositoryID:  request.RepositoryID,
		RepositoryURL: request.RepositoryURL,
		Findings:      request.Findings,
		EnqueuedAt:    request.EnqueuedAt,
		FinishedAt:    request.FinishedAt,
	}
	err := s.resRepo.Update(res)
	if err != nil {
		return err
	}
	return nil
}
