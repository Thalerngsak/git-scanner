package service

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/thalerngsak/git-scanner/mock/mock_repository"
	"github.com/thalerngsak/git-scanner/repository"
	"reflect"
	"testing"
)

func TestRepositoryService_NewRepository(t *testing.T) {
	t.Parallel()

	// Create a mock repository with a predictable result
	mockRepo := &mock_repository.MockRepositoryRepository{
		CreateFn: func(repository *repository.Repository) error {
			if repository.Name == "test" {
				return nil
			} else {
				return errors.New("error creating repository")
			}
		},
	}
	repositoryService := NewRepositoryService(mockRepo)

	// Call the NewRepository method with a valid request
	request := RepositoryRequest{
		Name: "test",
		URL:  "https://github.com/thalerngsak/git-scanner",
	}
	response, err := repositoryService.NewRepository(request)

	// Assert that the error is nil
	assert.NoError(t, err)

	// Assert that the response has the expected values
	assert.NotEmpty(t, response.ID)
	assert.Equal(t, request.Name, response.Name)
	assert.Equal(t, request.URL, response.URL)

	// Call the NewRepository method with an invalid request
	request = RepositoryRequest{
		Name: "",
		URL:  "",
	}
	response, err = repositoryService.NewRepository(request)

	// Assert that the error is not nil and contains the expected message
	assert.EqualError(t, err, errors.New("error creating repository").Error())

	// Assert that the response is nil
	assert.Nil(t, response)
}

func TestRepositoryService_GetRepository(t *testing.T) {
	// Create a mock RepositoryRepository.
	repoRepo := &mock_repository.MockRepositoryRepository{}

	// Create a new RepositoryService with the mock RepositoryRepository.
	repoService := NewRepositoryService(repoRepo)

	// Set up the test case.
	repoRepo.GetAllFn = func() ([]*repository.Repository, error) {
		return []*repository.Repository{
			{
				ID:   "test-id-1",
				Name: "test-repo-1",
				URL:  "https://github.com/test/test-repo-1",
			},
			{
				ID:   "test-id-2",
				Name: "test-repo-2",
				URL:  "https://github.com/test/test-repo-2",
			},
		}, nil
	}

	// Call the GetRepository method with the test case.
	response, err := repoService.GetRepository()

	// Check that the response and error are as expected.
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
	}
	if response == nil {
		t.Errorf("Expected a non-nil response")
	} else {
		expectedResponse := []*RepositoryResponse{
			{
				ID:   "test-id-1",
				Name: "test-repo-1",
				URL:  "https://github.com/test/test-repo-1",
			},
			{
				ID:   "test-id-2",
				Name: "test-repo-2",
				URL:  "https://github.com/test/test-repo-2",
			},
		}
		if !reflect.DeepEqual(response, expectedResponse) {
			t.Errorf("Response is not as expected. Expected %v but got %v", expectedResponse, response)
		}
	}

	// Check that the GetAll method was called on the mock RepositoryRepository.
	if repoRepo.GetAllFn == nil {
		t.Errorf("GetAll method was not called")
	}
}
