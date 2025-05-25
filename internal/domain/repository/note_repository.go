package repository

import (
	"context"

	"github.com/DevAntonioJorge/go-notes/internal/domain/models"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type NoteRepository struct {
	db *mongo.Collection
}

func NewNoteRepository(db *mongo.Collection) *NoteRepository {
	return &NoteRepository{db}
}

func (r *NoteRepository) SaveNote(ctx context.Context, note *models.Note) error {
	_, err := r.db.InsertOne(ctx, note)
	if err != nil {
		return err
	}
	return nil
}

func (r *NoteRepository) GetNote(ctx context.Context, id string) (*models.Note, error) {
	note := new(models.Note)

	if err := r.db.FindOne(ctx, bson.M{"_id": id}).Decode(note); err != nil {
		return nil, err
	}
	return note, nil
}
func (r *NoteRepository) UpdateNote(ctx context.Context, id string, note *models.Note) error {
	_, err := r.db.UpdateOne(ctx, bson.M{"_id": id}, bson.M{"$set": note})
	if err != nil {
		return err
	}
	return nil
}

func (r *NoteRepository) DeleteNote(ctx context.Context, id string) error {
	_, err := r.db.DeleteOne(ctx, bson.M{"_id": id})
	if err != nil {
		return err
	}
	return nil
}

func (r *NoteRepository) GetNotes(ctx context.Context, userID string) ([]*models.Note, error) {
	cursor, err := r.db.Find(ctx, bson.M{"user_id": userID})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)
	return createNoteList(ctx, cursor)
}
func (r *NoteRepository) SearchNotes(ctx context.Context, userID string, query string) ([]*models.Note, error) {
	cursor, err := r.db.Find(ctx, bson.M{"user_id": userID, "title": bson.M{"$regex": query}})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)
	return createNoteList(ctx, cursor)
}
func (r *NoteRepository) GetNotesByFolder(ctx context.Context, folderID string) ([]*models.Note, error) {
	cursor, err := r.db.Find(ctx, bson.M{"folder_id": folderID})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)
	return createNoteList(ctx, cursor)
}
func (r *NoteRepository) MoveNote(ctx context.Context, noteID string, newFolderID string) error {
	_, err := r.db.UpdateOne(ctx, bson.M{"_id": noteID}, bson.M{"$set": bson.M{"folder_id": newFolderID}})
	if err != nil {
		return err
	}
	return nil
}
func (r *NoteRepository) GetNotesByTag(ctx context.Context, userID string, tag string) ([]*models.Note, error) {
	cursor, err := r.db.Find(ctx, bson.M{"user_id": userID, "tags": tag})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)
	return createNoteList(ctx, cursor)
}
func (r *NoteRepository) GetRecentNotes(ctx context.Context, userID string, limit int) ([]*models.Note, error) {
	cursor, err := r.db.Find(ctx, bson.M{"user_id": userID}, options.Find().SetSort(bson.M{"created_at": -1}).SetLimit(int64(limit)))
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	return createNoteList(ctx, cursor)
}

func createNoteList(ctx context.Context, cursor *mongo.Cursor) ([]*models.Note, error) {
	var notes []*models.Note
	for cursor.Next(ctx) {
		note := new(models.Note)
		if err := cursor.Decode(note); err != nil {
			return nil, err
		}
		notes = append(notes, note)
	}
	if err := cursor.Err(); err != nil {
		return nil, err
	}
	return notes, nil
}
