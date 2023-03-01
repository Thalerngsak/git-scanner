package repository

// Repository represents a Git repository
type Repository struct {
	ID   string `bson:"_id" json:"id"`
	Name string `bson:"name" json:"name" binding:"required"`
	URL  string `bson:"url" json:"url" binding:"required"`
}

// RepositoryStore is an interface for storing and retrieving repositories
type RepositoryStore interface {
	Create(r *Repository) error
	GetAll() ([]*Repository, error)
	GetByID(id string) (*Repository, error)
	Update(r *Repository) error
	Delete(id string) error
}
