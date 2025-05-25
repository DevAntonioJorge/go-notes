package repository

import "go.mongodb.org/mongo-driver/v2/mongo"

type NoteRepository struct {
	db *mongo.Collection
}

func NewNoteRepository(db *mongo.Collection) *NoteRepository {
	return &NoteRepository{db}
}

// SaveNote saves a note to the database
