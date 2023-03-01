package repository

import (
	"errors"
	"time"
)

type resultRepositoryMock struct {
	results []*Result
}

func (r resultRepositoryMock) GetByID(s string) (*Result, error) {
	//TODO implement me
	panic("implement me")
}

func NewResultRepositoryMock() ResultStore {
	finding := make([]string, 0)
	finding = append(finding, "Found secret in file repos")
	results := []*Result{
		{
			ID:            "1",
			Status:        "Success",
			RepositoryID:  "af1a94c4-8882-424f-9fb4-b3976775a4eb",
			RepositoryURL: "https://github.com/Thalerngsak/apitest.git",
			Findings:      finding,
			EnqueuedAt:    time.Now(),
			StartedAt:     time.Now(),
			FinishedAt:    time.Now(),
		},
		{
			ID:            "2",
			Status:        "Success",
			RepositoryID:  "af1a94c4-8882-424f-9fb4-b3976775a4eb",
			RepositoryURL: "https://github.com/Thalerngsak/apitest.git",
			Findings:      finding,
			EnqueuedAt:    time.Now(),
			StartedAt:     time.Now(),
			FinishedAt:    time.Now(),
		},
	}
	return resultRepositoryMock{results: results}
}

func (r resultRepositoryMock) GetById(id string) (*Result, error) {
	for _, customer := range r.results {
		if customer.ID == id {
			return customer, nil
		}
	}

	return nil, errors.New("customer not found")
}

func (r resultRepositoryMock) Create(res *Result) error {
	//TODO implement me
	panic("implement me")
}

func (r resultRepositoryMock) GetAll() ([]*Result, error) {
	return r.results, nil
}

func (r resultRepositoryMock) GetByRepositoryID(id string) ([]*Result, error) {
	var results []*Result
	for _, result := range r.results {
		if result.RepositoryID == id {
			results = append(results, result)
		}
	}
	return results, nil
}

func (r resultRepositoryMock) Update(res *Result) error {
	//TODO implement me
	panic("implement me")
}
