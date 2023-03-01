package service

import (
	"database/sql"
	"github.com/thalerngsak/git-scanner/errs"
	"github.com/thalerngsak/git-scanner/mock/mock_repository"

	"net/http"
	"testing"

	"github.com/golang/mock/gomock"
)

func TestGetResultByRepositoryIDNotFound(t *testing.T) {
	//Arrage
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock_repository.NewMockResultRepository(ctrl)

	id := "1"

	mockRepo.EXPECT().GetByRepositoryID(id).Return(nil, sql.ErrNoRows)
	resultService := NewResultService(mockRepo)

	//Act
	_, err := resultService.GetResultByRepositoryID(id)

	//Assert
	if err == nil {
		t.Error("should be error")
		return
	}

	appErr, ok := err.(errs.AppError)
	if !ok {
		t.Error("should return AppError")
		return
	}

	if appErr.Code != http.StatusNotFound {
		t.Error("invalid error code")
	}

	if appErr.Message != "customer not found" {
		t.Error("invalid error message")
	}
}
