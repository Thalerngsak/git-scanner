package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/goccy/go-json"
	"github.com/thalerngsak/git-scanner/mock/mock_service"
	"github.com/thalerngsak/git-scanner/service"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestResultHandler_GetResult(t *testing.T) {
	mockResultService := &mock_service.MockResultService{}

	// Mock the GetResult method to return a test response
	expectedResult := []*service.ResultResponse{
		{
			ID:            "test-id",
			Status:        "test-status",
			RepositoryID:  "test-repo-id",
			RepositoryURL: "test-repo-url",
			Findings: []service.Finding{
				{
					Category: "test-category",
					Message:  "test-message",
				},
			},
			EnqueuedAt: time.Now(),
			StartedAt:  time.Now(),
			FinishedAt: time.Now(),
		},
	}
	mockResultService.On("GetResult").Return(expectedResult, nil)

	// Create the result handler with the mocked service
	resultHandler := resultHandler{resultSrv: mockResultService}

	// Create a new HTTP request and recorder for testing
	req, _ := http.NewRequest("GET", "/result", nil)
	recorder := httptest.NewRecorder()

	// Call the GetResult method
	router := gin.Default()
	router.GET("/result", resultHandler.GetResult)
	router.ServeHTTP(recorder, req)

	// Check the response status code
	if recorder.Code != http.StatusOK {
		t.Errorf("expected status code %d but got %d", http.StatusOK, recorder.Code)
	}

	// Check the response body
	var responseBody []*service.ResultResponse
	err := json.Unmarshal(recorder.Body.Bytes(), &responseBody)
	if err != nil {
		t.Errorf("unexpected error while unmarshalling response body: %s", err.Error())
	}

	// Assert that the GetResult method was called once
	mockResultService.AssertCalled(t, "GetResult")
}
