package dbrepo

import (
	"go.mongodb.org/mongo-driver/mongo"
	"test-eth/internal/repository"
)

type mongoDBRepo struct {
	DB *mongo.Client
}

func NewMongoRepo(client *mongo.Client) repository.DatabaseRepo {
	return &mongoDBRepo{DB: client}
}
