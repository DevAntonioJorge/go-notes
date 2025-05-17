package interfaces

import (
	"context"

	"github.com/DevAntonioJorge/go-notes/internal/domain/models"
	"github.com/DevAntonioJorge/go-notes/internal/infra/dto"
)

type IUserService interface {
	SaveUser(input dto.CreateUserRequest) error
	Login(input dto.LoginRequest) (*models.User, error)
	UpdatePassword(id, password string) error
}

type INoteService interface {
	SaveNote(ctx context.Context, input dto.CreateNoteRequest) (*models.Note, error)
	GetNote(ctx context.Context, id string) (*models.Note, error)
	UpdateNote(ctx context.Context, id string, input dto.UpdateNoteRequest) error
	DeleteNote(ctx context.Context, id string) error
	GetNotes(ctx context.Context, userID string) ([]*models.Note, error)
	SearchNotes(ctx context.Context, userID string, query string) ([]*models.Note, error)
	GetNotesByFolder(ctx context.Context, folderID string) ([]*models.Note, error)
	MoveNote(ctx context.Context, noteID string, newFolderID string) error
	GetNotesByTag(ctx context.Context, userID string, tag string) ([]*models.Note, error)
	GetRecentNotes(ctx context.Context, userID string, limit int) ([]*models.Note, error)
}

type IFolderService interface {
	SaveFolder(ctx context.Context, input dto.CreateFolderRequest) (*models.Folder, error)
	GetFolder(ctx context.Context, id string) (*models.Folder, error)
	UpdateFolder(ctx context.Context, input dto.UpdateFolderRequest) (*models.Folder, error)
	DeleteFolder(ctx context.Context, id string) error
	GetFolders(ctx context.Context, userID string) ([]*models.Folder, error)
	GetFolderByPath(ctx context.Context, userID string, path string) (*models.Folder, error)
	MoveFolder(ctx context.Context, folderID string, newParentID string) (*models.Folder, error)
}
