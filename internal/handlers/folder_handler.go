package handlers

import (
	"net/http"

	"github.com/DevAntonioJorge/go-notes/internal/dto"
	"github.com/DevAntonioJorge/go-notes/internal/interfaces"
	"github.com/go-playground/validator/v10"
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

	validate := validator.New()
	if err := validate.Struct(req); err != nil {
		return c.String(http.StatusInternalServerError, "Invalid folder name")
	}

	folder, err := h.service.SaveFolder(req)
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

	folder, err := h.service.GetFolder(req.ID)
	if err != nil {
		return c.String(http.StatusNotFound, "Folder not found")
	}

	return c.JSON(http.StatusOK, echo.Map{
		"folder": folder,
	})
}

func (h *FolderHandler) UpdateFolder(c echo.Context) error {
	var req dto.UpdateFolderRequest

	if err := c.Bind(&req); err != nil {
		return c.String(http.StatusBadRequest, "Ïnvalid request body or invalid id")
	}

	folder, err := h.service.UpdateFolder(req)
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

	if err := h.service.DeleteFolder(req.ID); err != nil {
		return c.String(http.StatusInternalServerError, "Failed to delete folder")
	}

	return c.String(http.StatusAccepted, "Folder deleted")

}
func (h *FolderHandler) GetFoldersHandler(c echo.Context) error {
	var req dto.GetFoldersRequest

	if err := c.Bind(&req); err != nil {
		return c.String(http.StatusBadRequest, "Invalid user id")
	}

	folderList, err := h.service.GetFolders(req.UserID)

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

	folder, err := h.service.GetFolderByPath(req.UserID, req.Path)
	if err != nil {
		return c.String(http.StatusInternalServerError, "Failed to get folder")
	}

	return c.JSON(http.StatusOK, echo.Map{
		"folder": folder,
	})
}

func (h *FolderHandler) MoveFolder(c echo.Context) error {
	var req dto.MoveFolderRequest

	if err := c.Bind(&req); err != nil {
		return c.String(http.StatusBadRequest, "Invalid request body")
	}

	folder, err := h.service.MoveFolder(req.FolderID, req.ParentID)
	if err != nil {
		return c.String(http.StatusInternalServerError, "Failed to move folder")
	}

	return c.JSON(http.StatusAccepted, echo.Map{
		"folder": folder,
	})
}
