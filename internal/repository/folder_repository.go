package repository

import "go.mongodb.org/mongo-driver/v2/mongo"

type FolderRepository struct {
	db *mongo.Database
}

func NewFolderRepository(db *mongo.Database) *FolderRepository {
	return &FolderRepository{db}
}
