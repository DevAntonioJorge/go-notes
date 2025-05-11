package interfaces

import "github.com/DevAntonioJorge/go-notes/internal/models"

type IUserRepository interface {
	SaveUser(user *models.User) error
	GetUserByEmail(email string) (*models.User, error)
	GetUserByID(id string) (*models.User, error)
	GetUserByName(name string) (*models.User, error)
	UpdatePassword(user *models.User, password string) error
}

type INoteRepository interface {
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

type IFolderRepository interface {
	SaveFolder(folder *models.Folder) error
    GetFolder(id string) (*models.Folder, error)
    UpdateFolder(id string, folder *models.Folder) error
    DeleteFolder(id string) error
    GetFolders(userID string) ([]*models.Folder, error)
    GetFolderByPath(userID string, path string) (*models.Folder, error)
    MoveFolder(folderID string, newParentID string) error
}