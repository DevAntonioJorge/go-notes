package service

import (
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

func (s *FolderService) SaveFolder(input dto.CreateFolderRequest) (*models.Folder, error) {
	folder, err := models.NewFolder(input.Name, input.ParentID)
	if err != nil {
		return nil, err
	}
	if err := s.repo.SaveFolder(folder); err != nil {
		return nil, err
	}
	return folder, nil
}

func (s *FolderService) GetFolder(id string) (*models.Folder, error) {
	return s.repo.GetFolder(id)
}

func (s *FolderService) UpdateFolder(input dto.UpdateFolderRequest) (*models.Folder, error) {
	folder, err := s.repo.GetFolder(input.ID)
	if err != nil {
		return nil, err
	}
	newFolder, err := s.repo.UpdateFolder(folder)
	if err != nil {
		return nil, err
	}
	return newFolder, nil
}

func (s *FolderService) DeleteFolder(id string) error {
	return s.repo.DeleteFolder(id)
}

func (s *FolderService) GetFolders(userID string) ([]*models.Folder, error) {
	return s.repo.GetFolders(userID)
}

func (s *FolderService) GetFolderByPath(userID, path string) (*models.Folder, error) {
	return s.repo.GetFolderByPath(userID, path)
}

func (s *FolderService) MoveFolder(folderID, newParentID string) error {
	folder, err := s.repo.GetFolder(folderID)
	if err != nil {
		return err
	}

	return s.repo.MoveFolder(folder, newParentID)
}
