package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/thalerngsak/git-scanner/mock/mock_service"
	"github.com/thalerngsak/git-scanner/service"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestScanHandler_Scan(t *testing.T) {
	mockRepoService := new(mock_service.MockRepositoryService)
	mockResultService := new(mock_service.MockResultService)

	handler := scanHandler{
		repSrv: mockRepoService,
		resSrv: mockResultService,
	}

	// Mock repository response
	repoID := "test-repo-id"
	mockRepoService.On("GetRepositoryByID", repoID).Return(&service.RepositoryResponse{
		ID:   repoID,
		Name: "test-repo-name",
		URL:  "http://test-repo-url.com",
	}, nil)

	// Mock result creation
	mockResultService.On("Create", mock.Anything).Return(nil)

	// Mock result retrieval
	mockResultService.On("GetResultByID", mock.AnythingOfType("string")).Return(&service.ResultResponse{}, nil)

	mockResultService.On("UpdateResult", mock.AnythingOfType("*service.ResultRequest")).Return(nil)

	// Call the function
	req, _ := http.NewRequest("POST", "/scan", strings.NewReader(`{"id":"test-repo-id"}`))
	req.Header.Set("Content-Type", "application/json")
	// Create a new Gin context from the HTTP request
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = req
	handler.Scan(c)

	// Check the response status code
	assert.Equal(t, http.StatusOK, w.Code)

	// Check that the result was retrieved
	mockResultService.AssertCalled(t, "GetResultByID", mock.AnythingOfType("string"))

	// Check that the repository was retrieved
	mockRepoService.AssertCalled(t, "GetRepositoryByID", repoID)

}
