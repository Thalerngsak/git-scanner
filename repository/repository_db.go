package repository

import (
	"context"
	"github.com/pkg/errors"
	"github.com/thalerngsak/git-scanner/errs"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// repositoryStore is a MongoDB-based implementation of the RepositoryStore interface
type repositoryStore struct {
	collection *mongo.Collection
}

func NewRepositoryDB(collection *mongo.Collection) repositoryStore {
	return repositoryStore{collection: collection}
}

func (s *repositoryStore) Create(r *Repository) error {
	_, err := s.collection.InsertOne(context.Background(), r)
	if err != nil {
		return errs.NewUnexpectedError()
	}
	return nil
}

func (s *repositoryStore) GetAll() ([]*Repository, error) {
	cursor, err := s.collection.Find(context.Background(), bson.M{})
	if err != nil {
		return nil, errs.NewUnexpectedError()
	}
	defer cursor.Close(context.Background())
	var repos []*Repository
	for cursor.Next(context.Background()) {
		var r Repository
		if err := cursor.Decode(&r); err != nil {
			return nil, errs.NewUnexpectedError()
		}
		repos = append(repos, &r)
	}
	return repos, nil
}

func (s *repositoryStore) GetByID(id string) (*Repository, error) {
	var r Repository
	err := s.collection.FindOne(context.Background(), bson.M{"_id": id}).Decode(&r)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errs.NewValidationError(" repository ID: " + id + " not found")
		}
		return nil, errs.NewUnexpectedError()
	}
	return &r, nil
}

func (s *repositoryStore) Update(r *Repository) error {
	result, err := s.collection.UpdateOne(context.Background(), bson.M{"_id": r.ID}, bson.M{"$set": bson.M{"name": r.Name, "url": r.URL}})
	if err != nil {
		return errs.NewUnexpectedError()
	}
	if result.ModifiedCount == 0 {
		return errors.New("no repository updated")
	}
	return nil
}

func (s *repositoryStore) Delete(id string) error {
	result, err := s.collection.DeleteOne(context.Background(), bson.M{"_id": id})
	if err != nil {
		return errs.NewUnexpectedError()
	}
	if result.DeletedCount == 0 {
		return errors.New("no repository deleted")
	}
	return nil
}
