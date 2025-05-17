package service

import (
	"context"

	"github.com/DevAntonioJorge/go-notes/internal/dto"
	"github.com/DevAntonioJorge/go-notes/internal/interfaces"
	"github.com/DevAntonioJorge/go-notes/internal/models"
)

type FolderService struct {
	repo interfaces.IFolderRepository
}

func NewFolderService(repository interfaces.IFolderRepository) *FolderService {
	return &FolderService{repository}
}

func (s *FolderService) SaveFolder(ctx context.Context, input dto.CreateFolderRequest) (*models.Folder, error) {
	folder, err := models.NewFolder(input.Name, input.ParentID)
	if err != nil {
		return nil, err
	}
	if err := s.repo.SaveFolder(ctx, folder); err != nil {
		return nil, err
	}
	return folder, nil
}

func (s *FolderService) GetFolder(ctx context.Context, id string) (*models.Folder, error) {
	return s.repo.GetFolder(ctx, id)
}

func (s *FolderService) UpdateFolder(ctx context.Context, input dto.UpdateFolderRequest) (*models.Folder, error) {
	folder, err := s.repo.GetFolder(ctx, input.ID)
	if err != nil {
		return nil, err
	}
	newFolder, err := s.repo.UpdateFolder(ctx, folder)
	if err != nil {
		return nil, err
	}
	return newFolder, nil
}

func (s *FolderService) DeleteFolder(ctx context.Context, id string) error {
	return s.repo.DeleteFolder(ctx, id)
}

func (s *FolderService) GetFolders(ctx context.Context, userID string) ([]*models.Folder, error) {
	return s.repo.GetFolders(ctx, userID)
}

func (s *FolderService) GetFolderByPath(ctx context.Context, userID, path string) (*models.Folder, error) {
	return s.repo.GetFolderByPath(ctx, userID, path)
}

func (s *FolderService) MoveFolder(ctx context.Context, folderID, newParentID string) (*models.Folder, error) {
	folder, err := s.repo.GetFolder(ctx, folderID)
	if err != nil {
		return nil, err
	}
	if err := s.repo.MoveFolder(ctx, folder, newParentID); err != nil {
		return nil, err
	}
	return folder, nil
}
