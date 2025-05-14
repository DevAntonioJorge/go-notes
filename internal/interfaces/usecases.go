package interfaces

import (
	"github.com/DevAntonioJorge/go-notes/internal/dto"
	"github.com/DevAntonioJorge/go-notes/internal/models"
)

type IUserService interface{
	SaveUser(input dto.CreateUserRequest) error
	Login(input dto.LoginRequest) (*models.User, error)
	UpdatePassword(id , password string) error
}

type INoteService interface{
	SaveNote(input dto.CreateNoteRequest) (*models.Note, error)
	GetNote(id string) (*models.Note, error)
	UpdateNote(id string, input dto.UpdateNoteRequest) error
	DeleteNote(id string) error
	GetNotes(userID string) ([]*models.Note, error)
	SearchNotes(userID string, query string) ([]*models.Note, error)
	GetNotesByFolder(folderID string) ([]*models.Note, error)
    MoveNote(noteID string, newFolderID string) error
    GetNotesByTag(userID string, tag string) ([]*models.Note, error)
    GetRecentNotes(userID string, limit int) ([]*models.Note, error)
}

type IFolderService interface {
	SaveFolder(input dto.CreateFolderRequest) (*models.Folder, error)
	GetFolder(id string) (*models.Folder, error)
    UpdateFolder(input dto.UpdateFolderRequest) (*models.Folder, error)
    DeleteFolder(id string) error
    GetFolders(userID string) ([]*models.Folder, error)
    GetFolderByPath(userID string, path string) (*models.Folder, error)
    MoveFolder(folderID string, newParentID string) (*models.Folder, error)
}