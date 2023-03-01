package repository

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// resultStore is a MongoDB-based implementation of the ResultStore interface
type resultStore struct {
	collection *mongo.Collection
}

func NewResultDB(collection *mongo.Collection) ResultStore {
	return &resultStore{collection: collection}
}

func (s *resultStore) Create(res Result) error {
	_, err := s.collection.InsertOne(context.Background(), res)
	if err != nil {
		return errors.New("failed to create result")
	}
	return nil
}

func (s *resultStore) GetAll() ([]Result, error) {
	cursor, err := s.collection.Find(context.Background(), bson.M{})
	if err != nil {
		return nil, errors.New("failed to get results")
	}
	defer cursor.Close(context.Background())
	var results []Result
	for cursor.Next(context.Background()) {
		var res Result
		if err := cursor.Decode(&res); err != nil {
			return nil, errors.New("failed to decode result")
		}
		results = append(results, res)
	}
	return results, nil
}

func (s *resultStore) GetByID(id string) (*Result, error) {
	var r Result
	err := s.collection.FindOne(context.Background(), bson.M{"_id": id}).Decode(&r)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("result ID: " + id + "not found")
		}
		return nil, errors.New("failed to get result by ID")
	}
	return &r, nil
}

func (s *resultStore) GetByRepositoryID(id string) ([]*Result, error) {
	cursor, err := s.collection.Find(context.Background(), bson.M{"repository_id": id})
	if err != nil {
		return nil, errors.New("failed to get results by repository ID")
	}
	defer cursor.Close(context.Background())
	var results []*Result
	for cursor.Next(context.Background()) {
		var res Result
		if err := cursor.Decode(&res); err != nil {
			return nil, errors.New("failed to decode result")
		}
		results = append(results, &res)
	}
	if results == nil {
		return nil, errors.New("results repository ID: " + id + " not found")
	}
	return results, nil
}

func (s *resultStore) Update(res *Result) error {
	result, err := s.collection.UpdateOne(context.Background(), bson.M{"_id": res.ID}, bson.M{"$set": bson.M{"status": res.Status, "findings": res.Findings, "started_at": res.StartedAt, "finished_at": res.FinishedAt}})
	if err != nil {
		return errors.New("failed to update the result")
	}
	if result.ModifiedCount == 0 {
		return errors.New("no documents updated")
	}
	return nil
}
