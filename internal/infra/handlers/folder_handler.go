package handlers

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/DevAntonioJorge/go-notes/internal/domain/interfaces"
	"github.com/DevAntonioJorge/go-notes/internal/infra/dto"
	"github.com/DevAntonioJorge/go-notes/pkg/schema"
	"github.com/labstack/echo/v4"
)

type FolderHandler struct {
	service interfaces.IFolderService
}

func NewFolderHandler(service interfaces.IFolderService) *FolderHandler {
	return &FolderHandler{service}
}

func (h *FolderHandler) CreateFolderHandler(c echo.Context) error {
	var req dto.CreateFolderRequest

	if err := c.Bind(&req); err != nil {
		return c.String(http.StatusBadRequest, "Invalid request body")
	}

	if err := schema.NewValidator().Validate(&req); err != nil {
		log.Printf("Error validating the user: %v", err)
		return c.String(http.StatusInternalServerError, "Error invalid fields")
	}

	folder, err := h.service.SaveFolder(c.Request().Context(), req)
	if err != nil {
		return c.String(http.StatusInternalServerError, "Failed to save folder")
	}

	return c.JSON(http.StatusCreated, echo.Map{
		"message": "Folder created",
		"folder":  folder,
	})
}

func (h *FolderHandler) GetFolderHandler(c echo.Context) error {
	var req dto.GetFolderRequest
	if err := c.Bind(&req); err != nil {
		return c.String(http.StatusBadRequest, "No id specified or invalid parameter")
	}
	ctx, cancel := context.WithTimeout(c.Request().Context(), 2*time.Second)
	defer cancel()
	folder, err := h.service.GetFolder(ctx, req.ID)
	if err != nil {
		return c.String(http.StatusNotFound, "Folder not found")
	}

	return c.JSON(http.StatusOK, echo.Map{
		"folder": folder,
	})
}

func (h *FolderHandler) UpdateFolderHandler(c echo.Context) error {
	var req dto.UpdateFolderRequest

	if err := c.Bind(&req); err != nil {
		return c.String(http.StatusBadRequest, "Ïnvalid request body or invalid id")
	}

	if err := schema.NewValidator().Validate(&req); err != nil {
		log.Printf("Error validating the user: %v", err)
		return c.String(http.StatusInternalServerError, "Error invalid fields")
	}
	ctx, cancel := context.WithTimeout(c.Request().Context(), 2*time.Second)
	defer cancel()
	folder, err := h.service.UpdateFolder(ctx, req)
	if err != nil {
		return c.String(http.StatusInternalServerError, "Failed to update folder name")
	}

	return c.JSON(http.StatusAccepted, echo.Map{
		"folder_id": req.ID,
		"folder":    folder,
	})
}

func (h *FolderHandler) DeleteFolderHandler(c echo.Context) error {
	var req dto.DeleteFolderRequest

	if err := c.Bind(&req); err != nil {
		return c.String(http.StatusBadRequest, "Invalid id")
	}

	ctx, cancel := context.WithTimeout(c.Request().Context(), 2*time.Second)
	defer cancel()
	if err := h.service.DeleteFolder(ctx, req.ID); err != nil {
		return c.String(http.StatusInternalServerError, "Failed to delete folder")
	}

	return c.String(http.StatusAccepted, "Folder deleted")

}
func (h *FolderHandler) GetFoldersHandler(c echo.Context) error {
	var req dto.GetFoldersRequest

	if err := c.Bind(&req); err != nil {
		return c.String(http.StatusBadRequest, "Invalid user id")
	}

	if err := schema.NewValidator().Validate(&req); err != nil {
		log.Printf("Error validating the user: %v", err)
		return c.String(http.StatusInternalServerError, "Error invalid fields")
	}

	ctx, cancel := context.WithTimeout(c.Request().Context(), 2*time.Second)
	defer cancel()
	folderList, err := h.service.GetFolders(ctx, req.UserID)

	if err != nil {
		return c.String(http.StatusInternalServerError, "Failed to get folders")
	}

	return c.JSON(http.StatusOK, echo.Map{
		"folder list": folderList,
	})
}

func (h *FolderHandler) GetFolderByPathHandler(c echo.Context) error {
	var req dto.GetFolderByPathRequest

	if err := c.Bind(&req); err != nil {
		return c.String(http.StatusBadRequest, "Invalid request body")
	}

	if err := schema.NewValidator().Validate(&req); err != nil {
		log.Printf("Error validating the user: %v", err)
		return c.String(http.StatusInternalServerError, "Error invalid fields")
	}

	ctx, cancel := context.WithTimeout(c.Request().Context(), 2*time.Second)
	defer cancel()
	folder, err := h.service.GetFolderByPath(ctx, req.UserID, req.Path)
	if err != nil {
		return c.String(http.StatusInternalServerError, "Failed to get folder")
	}

	return c.JSON(http.StatusOK, echo.Map{
		"folder": folder,
	})
}

func (h *FolderHandler) MoveFolderHandler(c echo.Context) error {
	var req dto.MoveFolderRequest

	if err := c.Bind(&req); err != nil {
		return c.String(http.StatusBadRequest, "Invalid request body")
	}

	if err := schema.NewValidator().Validate(&req); err != nil {
		log.Printf("Error validating the user: %v", err)
		return c.String(http.StatusInternalServerError, "Error invalid fields")
	}

	ctx, cancel := context.WithTimeout(c.Request().Context(), 2*time.Second)
	defer cancel()
	folder, err := h.service.MoveFolder(ctx, req.FolderID, req.ParentID)
	if err != nil {
		return c.String(http.StatusInternalServerError, "Failed to move folder")
	}

	return c.JSON(http.StatusAccepted, echo.Map{
		"folder": folder,
	})
}
