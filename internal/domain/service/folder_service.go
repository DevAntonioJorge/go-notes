package service

import (
	"context"

	"github.com/DevAntonioJorge/go-notes/internal/domain/models"
	"github.com/DevAntonioJorge/go-notes/internal/domain/repository"
	"github.com/DevAntonioJorge/go-notes/internal/domain/utils/path"
	"github.com/DevAntonioJorge/go-notes/internal/infra/dto"
)

type FolderService struct {
	repo *repository.Repository
}

func NewFolderService(repo *repository.Repository) *FolderService {
	return &FolderService{
		repo: repo,
	}
}
func (s *FolderService) SaveFolder(ctx context.Context, input dto.CreateFolderRequest) (*models.Folder, error) {
	normalizedPath := path.NormalizePath(input.Name)
	if !path.IsValidPath(normalizedPath) {
		return nil, models.ErrInvalidFolderPath
	}
	folder, err := models.NewFolder(normalizedPath, input.ParentID)
	if err != nil {
		return nil, err
	}
	if err := s.repo.Folder.SaveFolder(ctx, folder); err != nil {
		return nil, err
	}
	return folder, nil
}

func (s *FolderService) GetFolder(ctx context.Context, id string) (*models.Folder, error) {
	return s.repo.Folder.GetFolder(ctx, id)
}

func (s *FolderService) UpdateFolder(ctx context.Context, input dto.UpdateFolderRequest) (*models.Folder, error) {
	folder, err := s.repo.Folder.GetFolder(ctx, input.ID)
	if err != nil {
		return nil, err
	}
	normalizedPath := path.NormalizePath(folder.Name)
	if !path.IsValidPath(normalizedPath) {
		return nil, models.ErrInvalidFolderPath
	}
	folder.Name = normalizedPath
	newFolder, err := s.repo.Folder.UpdateFolder(ctx, folder)
	if err != nil {
		return nil, err
	}
	return newFolder, nil
}

func (s *FolderService) DeleteFolder(ctx context.Context, id string) error {
	folder, err := s.repo.Folder.GetFolder(ctx, id)
	if err != nil {
		return err
	}
	if folder == nil {
		return models.ErrFolderNotFound
	}

	if folder.ParentID == "" || folder.ParentID == "0" || folder.ParentID == "root" {
		return models.ErrCannotDeleteRootFolder
	}

	return s.repo.Folder.DeleteFolder(ctx, id)
}

func (s *FolderService) GetFolders(ctx context.Context, userID string) ([]*models.Folder, error) {
	if _, err := s.repo.User.GetUserByID(userID); err != nil {
		return nil, models.ErrUserNotFound
	}
	return s.repo.Folder.GetFolders(ctx, userID)
}

func (s *FolderService) GetFolderByPath(ctx context.Context, userID, path string) (*models.Folder, error) {
	if _, err := s.repo.User.GetUserByID(userID); err != nil {
		return nil, models.ErrUserNotFound
	}

	return s.repo.Folder.GetFolderByPath(ctx, userID, path)
}

func (s *FolderService) MoveFolder(ctx context.Context, folderID, newParentID string) (*models.Folder, error) {
	folder, err := s.repo.Folder.GetFolder(ctx, folderID)
	if err != nil {
		return nil, err
	}
	if err := s.repo.Folder.MoveFolder(ctx, folder, newParentID); err != nil {
		return nil, err
	}
	return folder, nil
}
