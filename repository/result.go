package repository

import "time"

// Result represents a security scan result
type Result struct {
	ID            string    `bson:"_id"`
	Status        string    `bson:"status"`
	RepositoryID  string    `bson:"repository_id"`
	RepositoryURL string    `bson:"repository_url"`
	Findings      []string  `bson:"findings"`
	EnqueuedAt    time.Time `bson:"enqueued_at"`
	StartedAt     time.Time `bson:"started_at"`
	FinishedAt    time.Time `bson:"finished_at"`
}

// ResultStore is an interface for storing and retrieving security scan results
type ResultStore interface {
	Create(*Result) error
	GetAll() ([]*Result, error)
	GetByID(string) (*Result, error)
	GetByRepositoryID(id string) ([]*Result, error)
	Update(r *Result) error
}
