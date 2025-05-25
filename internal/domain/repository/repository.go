package repository

import (
	"context"

	"github.com/DevAntonioJorge/go-notes/internal/domain/models"
	"github.com/jackc/pgx/v5"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type Repository struct {
	User interface {
		SaveUser(user *models.User) error
		GetUserByEmail(email string) (*models.User, error)
		GetUserByID(id string) (*models.User, error)
		GetUserByName(name string) (*models.User, error)
		UpdatePassword(user *models.User, password string) error
	}

	Note interface {
		SaveNote(note *models.Note) error
		GetNote(id string) (*models.Note, error)
		UpdateNote(id string, note *models.Note) error
		DeleteNote(id string) error
		GetNotes(userID string) ([]*models.Note, error)
		SearchNotes(userID string, query string) ([]*models.Note, error)
		GetNotesByFolder(folderID string) ([]*models.Note, error)
		MoveNote(noteID string, newFolderID string) error
		GetNotesByTag(userID string, tag string) ([]*models.Note, error)
		GetRecentNotes(userID string, limit int) ([]*models.Note, error)
	}

	Folder interface {
		SaveFolder(ctx context.Context, folder *models.Folder) error
		GetFolder(ctx context.Context, id string) (*models.Folder, error)
		UpdateFolder(ctx context.Context, folder *models.Folder) (*models.Folder, error)
		DeleteFolder(ctx context.Context, id string) error
		GetFolders(ctx context.Context, userID string) ([]*models.Folder, error)
		GetFolderByPath(ctx context.Context, userID string, path string) (*models.Folder, error)
		MoveFolder(ctx context.Context, folder *models.Folder, newParentID string) error
	}
}

func NewRepository(conn *pgx.Conn, client *mongo.Client) *Repository {
	return &Repository{
		User:   NewUserRepository(conn),
		Note:   NewNoteRepository(client.Database("notes").Collection("notes")),
		Folder: NewFolderRepository(client.Database("notes").Collection("folders")),
	}
}
