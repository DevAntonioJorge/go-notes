package service

import (
	"context"

	"github.com/DevAntonioJorge/go-notes/internal/domain/models"
	"github.com/DevAntonioJorge/go-notes/internal/domain/repository"
	"github.com/DevAntonioJorge/go-notes/internal/infra/dto"
)

type NoteService struct {
	repo *repository.Repository
}

func NewNoteService(repo *repository.Repository) *NoteService {
	return &NoteService{repo}
}

func (s *NoteService) SaveNote(ctx context.Context, input dto.CreateNoteRequest) (*models.Note, error) {
	note := models.NewNote(input.UserID, input.Title, input.Content, input.FolderID, input.Tags)
	if err := s.repo.Note.SaveNote(ctx, note); err != nil {
		return nil, err
	}
	return note, nil
}

func (s *NoteService) GetNote(ctx context.Context, id string) (*models.Note, error) {
	note, err := s.repo.Note.GetNote(ctx, id)
	if err != nil {
		return nil, err
	}
	return note, nil
}
func (s *NoteService) DeleteNote(ctx context.Context, id string) error {
	_, err := s.repo.Note.GetNote(ctx, id)
	if err != nil {
		return err
	}
	return s.repo.Note.DeleteNote(ctx, id)
}

func (s *NoteService) UpdateNoteTitle(ctx context.Context, id string, input dto.UpdateNoteTitleRequest) error {
	_, err := s.repo.Note.GetNote(ctx, id)
	if err != nil {
		return err
	}
	return s.repo.Note.UpdateNoteTitle(ctx, id, input.Title)
}

func (s *NoteService) UpdateNoteContent(ctx context.Context, id string, input dto.UpdateNoteContentRequest) error {
	_, err := s.repo.Note.GetNote(ctx, id)
	if err != nil {
		return err
	}
	return s.repo.Note.UpdateNoteContent(ctx, id, input.Content)
}
func (s *NoteService) UpdateNoteTags(ctx context.Context, id string, input dto.UpdateNoteTagsRequest) error {
	_, err := s.repo.Note.GetNote(ctx, id)
	if err != nil {
		return err
	}
	return s.repo.Note.UpdateNoteTags(ctx, id, input.Tags)
}

func (s *NoteService) UpdateNoteFolder(ctx context.Context, id string, input dto.UpdateNoteFolderRequest) error {
	_, err := s.repo.Note.GetNote(ctx, id)
	if err != nil {
		return err
	}
	return s.repo.Note.UpdateNoteFolder(ctx, id, input.NewFolderID)
}

func (s *NoteService) GetNotes(ctx context.Context, userID string) ([]*models.Note, error) {
	_, err := s.repo.User.GetUserByID(userID)
	if err != nil {
		return nil, err
	}
	return s.repo.Note.GetNotes(ctx, userID)
}

func (s *NoteService) SearchNotes(ctx context.Context, userID string, query string) ([]*models.Note, error) {
	_, err := s.repo.User.GetUserByID(userID)
	if err != nil {
		return nil, err
	}
	return s.repo.Note.SearchNotes(ctx, userID, query)
}

func (s *NoteService) GetNotesByFolder(ctx context.Context, folderID string) ([]*models.Note, error) {
	_, err := s.repo.Folder.GetFolder(ctx, folderID)
	if err != nil {
		return nil, err
	}
	return s.repo.Note.GetNotesByFolder(ctx, folderID)
}

func (s *NoteService) MoveNote(ctx context.Context, noteID string, newFolderID string) error {
	_, err := s.repo.Note.GetNote(ctx, noteID)
	if err != nil {
		return err
	}
	_, err = s.repo.Folder.GetFolder(ctx, newFolderID)
	if err != nil {
		return err
	}
	return s.repo.Note.MoveNote(ctx, noteID, newFolderID)
}
func (s *NoteService) GetNotesByTag(ctx context.Context, userID string, tag string) ([]*models.Note, error) {
	_, err := s.repo.User.GetUserByID(userID)
	if err != nil {
		return nil, err
	}
	return s.repo.Note.GetNotesByTag(ctx, userID, tag)
}
func (s *NoteService) GetRecentNotes(ctx context.Context, userID string, limit int) ([]*models.Note, error) {
	_, err := s.repo.User.GetUserByID(userID)
	if err != nil {
		return nil, err
	}
	return s.repo.Note.GetRecentNotes(ctx, userID, limit)
}
