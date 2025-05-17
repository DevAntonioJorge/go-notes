package repository

import (
	"context"

	"github.com/DevAntonioJorge/go-notes/internal/domain/models"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type FolderRepository struct {
	db *mongo.Database
}

func NewFolderRepository(db *mongo.Database) *FolderRepository {
	return &FolderRepository{db}
}

func (r *FolderRepository) SaveFolder(ctx context.Context, folder *models.Folder) error {
	_, err := r.db.Collection("folders").InsertOne(ctx, folder)
	if err != nil {
		return err
	}
	return nil
}

func (r *FolderRepository) GetFolder(ctx context.Context, id string) (*models.Folder, error) {
	folder := new(models.Folder)
	err := r.db.Collection("folders").FindOne(ctx, bson.M{"_id": id}).Decode(folder)
	if err != nil {
		return nil, err
	}
	return folder, nil
}

func (r *FolderRepository) UpdateFolder(ctx context.Context, folder *models.Folder) (*models.Folder, error) {
	_, err := r.db.Collection("folders").UpdateOne(ctx, bson.M{"_id": folder.ID}, bson.M{"$set": folder})
	if err != nil {
		return nil, err
	}
	return folder, nil
}

func (r *FolderRepository) DeleteFolder(ctx context.Context, id string) error {
	_, err := r.db.Collection("folders").DeleteOne(ctx, bson.M{"_id": id})
	if err != nil {
		return err
	}
	return nil
}

func (r *FolderRepository) GetFolders(ctx context.Context, userID string) ([]*models.Folder, error) {
	cursor, err := r.db.Collection("folders").Find(ctx, bson.M{"user_id": userID})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)
	var folders []*models.Folder
	for cursor.Next(ctx) {
		folder := new(models.Folder)
		if err := cursor.Decode(folder); err != nil {
			return nil, err
		}
		folders = append(folders, folder)
	}
	return folders, nil
}

func (r *FolderRepository) GetFolderByPath(ctx context.Context, userID string, path string) (*models.Folder, error) {
	folder := new(models.Folder)
	err := r.db.Collection("folders").FindOne(ctx, bson.M{"user_id": userID, "path": path}).Decode(folder)
	if err != nil {
		return nil, err
	}
	return folder, nil
}

func (r *FolderRepository) MoveFolder(ctx context.Context, folder *models.Folder, newParentID string) error {
	_, err := r.db.Collection("folders").UpdateOne(ctx, bson.M{"_id": folder.ID}, bson.M{"$set": bson.M{"parent_id": newParentID}})
	if err != nil {
		return err
	}
	return nil
}
