package service

import "time"

type ResultRequest struct {
	ID            string    `json:"id"`
	Status        string    `json:"status"`
	RepositoryID  string    `json:"repository_id"`
	RepositoryURL string    `json:"repository_url"`
	Findings      []string  `json:"findings"`
	EnqueuedAt    time.Time `json:"enqueued_at"`
	StartedAt     time.Time `json:"started_at"`
	FinishedAt    time.Time `json:"finished_at"`
}

type ResultResponse struct {
	ID            string    `json:"id"`
	Status        string    `json:"status"`
	RepositoryID  string    `json:"repository_id"`
	RepositoryURL string    `json:"repository_url"`
	Findings      []string  `json:"findings"`
	EnqueuedAt    time.Time `json:"enqueued_at"`
	StartedAt     time.Time `json:"started_at"`
	FinishedAt    time.Time `json:"finished_at"`
}

type ResultService interface {
	Create(ResultRequest) error
	GetResult() ([]*ResultResponse, error)
	GetResultByID(string) (*ResultResponse, error)
	GetResultByRepositoryID(string) ([]*ResultResponse, error)
	UpdateResult(*ResultRequest) error
}
