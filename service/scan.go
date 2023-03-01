package service

import "time"

type ScanRequest struct {
	Name string `json:"name" binding:"required"`
	URL  string `json:"url" binding:"required"`
}

type ScanResponse struct {
	ID            string    `json:"id"`
	Status        string    `json:"status"`
	RepositoryID  string    `json:"repository_id"`
	RepositoryURL string    `json:"repository_url"`
	Findings      []string  `json:"findings"`
	EnqueuedAt    time.Time `json:"enqueued_at"`
	StartedAt     time.Time `json:"started_at"`
	FinishedAt    time.Time `json:"finished_at"`
}

type ScanService interface {
	GetResult() ([]*ScanResponse, error)
	GetResultByRepositoryID(string) ([]*ScanRequest, error)
}
