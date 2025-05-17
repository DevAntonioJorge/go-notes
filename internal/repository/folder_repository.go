package repository

import (
	"context"

	"github.com/DevAntonioJorge/go-notes/internal/models"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type FolderRepository struct {
	db *mongo.Database
}

func NewFolderRepository(db *mongo.Database) *FolderRepository {
	return &FolderRepository{db}
}

func (r *FolderRepository) SaveFolder(folder *models.Folder) error {
	_, err := r.db.Collection("folders").InsertOne(context.Background(), folder)
	if err != nil {
		return err
	}
	return nil
}
