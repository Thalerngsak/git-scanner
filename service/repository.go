package service

type RepositoryRequest struct {
	Name string `json:"name" binding:"required"`
	URL  string `json:"url" binding:"required"`
}

type RepositoryResponse struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	URL  string `json:"url"`
}

type RepositoryService interface {
	NewRepository(RepositoryRequest) (*RepositoryResponse, error)
	GetRepository() ([]*RepositoryResponse, error)
	GetRepositoryByID(string) (*RepositoryResponse, error)
	UpdateRepository(string, *RepositoryRequest) error
	DeleteRepository(string) error
}
