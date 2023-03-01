package service

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/thalerngsak/git-scanner/mock/mock_repository"
	"github.com/thalerngsak/git-scanner/repository"
	"testing"
	"time"
)

func TestResultService_Create(t *testing.T) {
	// Create a mock repository with a predictable result
	mockRepo := mock_repository.MockResultRepository{
		CreateFn: func(result repository.Result) error {
			if result.ID == "" {
				return errors.New("empty ID")
			}
			return nil
		},
	}

	resultService := NewResultService(&mockRepo)

	// Create a valid request
	request := ResultRequest{
		ID:            "1",
		Status:        "pending",
		RepositoryID:  "2",
		RepositoryURL: "https://github.com/thalerngsak/git-scanner",
		EnqueuedAt:    time.Now(),
		StartedAt:     time.Now(),
	}

	// Call the Create method with the valid request
	err := resultService.Create(request)

	// Assert that the error is nil
	assert.NoError(t, err)

	// Create an invalid request with an empty ID
	invalidRequest := ResultRequest{
		ID: "",
	}

	// Call the Create method with the invalid request
	err = resultService.Create(invalidRequest)

	// Assert that the error is not nil and contains the expected message
	assert.EqualError(t, err, "empty ID")
}

func TestResultService_GetResult(t *testing.T) {
	t.Parallel()

	// Create a mock repository with a predictable result
	mockRepo := &mock_repository.MockResultRepository{
		GetAllFn: func() ([]repository.Result, error) {
			results := []repository.Result{
				{
					ID:            "1",
					Status:        "pending",
					RepositoryID:  "2",
					RepositoryURL: "https://github.com/thalerngsak/git-scanner",
					Findings:      []repository.Finding{},
					EnqueuedAt:    time.Now(),
					StartedAt:     time.Now(),
					FinishedAt:    time.Now(),
				},
				{
					ID:            "2",
					Status:        "completed",
					RepositoryID:  "3",
					RepositoryURL: "https://github.com/thalerngsak/git-scanner",
					Findings: []repository.Finding{
						{
							Category: "security",
							Message:  "potential vulnerability",
						},
					},
					EnqueuedAt: time.Now(),
					StartedAt:  time.Now(),
					FinishedAt: time.Now(),
				},
			}
			return results, nil
		},
	}
	resultService := NewResultService(mockRepo)

	// Call the GetResult method
	results, err := resultService.GetResult()

	// Assert that the error is nil
	assert.NoError(t, err)

	// Assert that the results have the expected length
	assert.Len(t, results, 2)

	// Assert that the first result has the expected values
	assert.Equal(t, "1", results[0].ID)
	assert.Equal(t, "pending", results[0].Status)
	assert.Equal(t, "2", results[0].RepositoryID)
	assert.Equal(t, "https://github.com/thalerngsak/git-scanner", results[0].RepositoryURL)
	assert.Len(t, results[0].Findings, 0)

	// Assert that the second result has the expected values
	assert.Equal(t, "2", results[1].ID)
	assert.Equal(t, "completed", results[1].Status)
	assert.Equal(t, "3", results[1].RepositoryID)
	assert.Equal(t, "https://github.com/thalerngsak/git-scanner", results[1].RepositoryURL)
	assert.Len(t, results[1].Findings, 1)
	assert.Equal(t, "security", results[1].Findings[0].Category)
	assert.Equal(t, "potential vulnerability", results[1].Findings[0].Message)

	// Create a mock repository that returns an error
	mockRepo = &mock_repository.MockResultRepository{
		GetAllFn: func() ([]repository.Result, error) {
			return nil, errors.New("repository error")
		},
	}
	resultService = NewResultService(mockRepo)

	// Call the GetResult method
	results, err = resultService.GetResult()

	// Assert that the error is not nil and contains the expected message
	assert.EqualError(t, err, "repository error")

	// Assert that the results are nil
	assert.Nil(t, results)
}

func TestResultService_GetResultByID(t *testing.T) {
	t.Parallel()

	// Create a mock repository with a predictable result
	mockRepo := &mock_repository.MockResultRepository{
		GetByIDFn: func(id string) (*repository.Result, error) {
			if id == "1" {
				return &repository.Result{
					ID:            "1",
					Status:        "pending",
					RepositoryID:  "2",
					RepositoryURL: "https://github.com/thalerngsak/git-scanner",
					Findings:      []repository.Finding{},
					EnqueuedAt:    time.Now(),
					StartedAt:     time.Now(),
					FinishedAt:    time.Now(),
				}, nil
			} else if id == "2" {
				return &repository.Result{
					ID:            "2",
					Status:        "completed",
					RepositoryID:  "3",
					RepositoryURL: "https://github.com/thalerngsak/git-scanner",
					Findings: []repository.Finding{
						{
							Category: "security",
							Message:  "potential vulnerability",
						},
					},
					EnqueuedAt: time.Now(),
					StartedAt:  time.Now(),
					FinishedAt: time.Now(),
				}, nil
			} else {
				return nil, errors.New("not found")
			}
		},
	}
	resultService := NewResultService(mockRepo)

	// Call the GetResultByID method with an existing ID
	result, err := resultService.GetResultByID("1")

	// Assert that the error is nil
	assert.NoError(t, err)

	// Assert that the result has the expected values
	assert.Equal(t, "1", result.ID)
	assert.Equal(t, "pending", result.Status)
	assert.Equal(t, "2", result.RepositoryID)
	assert.Equal(t, "https://github.com/thalerngsak/git-scanner", result.RepositoryURL)
	assert.Len(t, result.Findings, 0)

	// Call the GetResultByID method with another existing ID
	result, err = resultService.GetResultByID("2")

	// Assert that the error is nil
	assert.NoError(t, err)

	// Assert that the result has the expected values
	assert.Equal(t, "2", result.ID)
	assert.Equal(t, "completed", result.Status)
	assert.Equal(t, "3", result.RepositoryID)
	assert.Equal(t, "https://github.com/thalerngsak/git-scanner", result.RepositoryURL)
	assert.Len(t, result.Findings, 1)
	assert.Equal(t, "security", result.Findings[0].Category)
	assert.Equal(t, "potential vulnerability", result.Findings[0].Message)

	// Call the GetResultByID method with a non-existing ID
	result, err = resultService.GetResultByID("3")

	// Assert that the error is not nil and contains the expected message
	assert.EqualError(t, err, "not found")

	// Assert that the result is nil
	assert.Nil(t, result)
}

func TestResultService_GetResultByRepositoryID(t *testing.T) {
	t.Parallel()

	// Create a mock repository with a predictable result
	mockRepo := &mock_repository.MockResultRepository{
		GetByRepositoryIDFn: func(id string) ([]*repository.Result, error) {
			if id == "1" {
				return []*repository.Result{
					{
						ID:            "1",
						Status:        "pending",
						RepositoryID:  "1",
						RepositoryURL: "https://github.com/thalerngsak/git-scanner",
						Findings:      []repository.Finding{},
						EnqueuedAt:    time.Now(),
						StartedAt:     time.Now(),
						FinishedAt:    time.Now(),
					},
				}, nil
			} else if id == "2" {
				return []*repository.Result{
					{
						ID:            "2",
						Status:        "completed",
						RepositoryID:  "2",
						RepositoryURL: "https://github.com/thalerngsak/git-scanner",
						Findings: []repository.Finding{
							{
								Category: "security",
								Message:  "potential vulnerability",
							},
						},
						EnqueuedAt: time.Now(),
						StartedAt:  time.Now(),
						FinishedAt: time.Now(),
					},
				}, nil
			} else {
				return nil, errors.New("not found")
			}
		},
	}
	resultService := NewResultService(mockRepo)

	// Call the GetResultByRepositoryID method with an existing ID
	results, err := resultService.GetResultByRepositoryID("1")

	// Assert that the error is nil
	assert.NoError(t, err)

	// Assert that the result has the expected values
	assert.Len(t, results, 1)
	assert.Equal(t, "1", results[0].ID)
	assert.Equal(t, "pending", results[0].Status)
	assert.Equal(t, "1", results[0].RepositoryID)
	assert.Equal(t, "https://github.com/thalerngsak/git-scanner", results[0].RepositoryURL)
	assert.Len(t, results[0].Findings, 0)

	// Call the GetResultByRepositoryID method with another existing ID
	results, err = resultService.GetResultByRepositoryID("2")

	// Assert that the error is nil
	assert.NoError(t, err)

	// Assert that the result has the expected values
	assert.Len(t, results, 1)
	assert.Equal(t, "2", results[0].ID)
	assert.Equal(t, "completed", results[0].Status)
	assert.Equal(t, "2", results[0].RepositoryID)
	assert.Equal(t, "https://github.com/thalerngsak/git-scanner", results[0].RepositoryURL)
	assert.Len(t, results[0].Findings, 1)
	assert.Equal(t, "security", results[0].Findings[0].Category)
	assert.Equal(t, "potential vulnerability", results[0].Findings[0].Message)

	// Call the GetResultByRepositoryID method with a non-existing ID
	// Call the GetResultByRepositoryID method with a non-existing ID
	results, err = resultService.GetResultByRepositoryID("3")

	// Assert that the error is not nil and contains the expected message
	assert.EqualError(t, err, "not found")

	// Assert that the result is nil
	assert.Nil(t, results)
}

func TestResultService_UpdateResult(t *testing.T) {
	t.Parallel()

	// Create a mock repository with a predictable result
	mockRepo := mock_repository.MockResultRepository{
		UpdateFn: func(result *repository.Result) error {
			if result.ID == "1" {
				return nil
			} else {
				return errors.New("not found")
			}
		},
	}
	resultService := NewResultService(&mockRepo)

	// Call the UpdateResult method with an existing ID
	request := &ResultRequest{
		ID:            "1",
		Status:        "completed",
		RepositoryID:  "1",
		RepositoryURL: "https://github.com/thalerngsak/git-scanner",
		Findings: []Finding{
			{
				Category: "security",
				Message:  "potential vulnerability",
			},
		},
		EnqueuedAt: time.Now(),
		StartedAt:  time.Now(),
		FinishedAt: time.Now(),
	}
	err := resultService.UpdateResult(request)

	// Assert that the error is nil
	assert.NoError(t, err)

	// Call the UpdateResult method with a non-existing ID
	request = &ResultRequest{
		ID:            "2",
		Status:        "completed",
		RepositoryID:  "2",
		RepositoryURL: "https://github.com/thalerngsak/git-scanner",
		Findings: []Finding{
			{
				Category: "security",
				Message:  "potential vulnerability",
			},
		},
		EnqueuedAt: time.Now(),
		StartedAt:  time.Now(),
		FinishedAt: time.Now(),
	}
	err = resultService.UpdateResult(request)

	// Assert that the error is not nil and contains the expected message
	assert.EqualError(t, err, "not found")
}
