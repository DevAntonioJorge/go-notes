package main

import (
	"net/http"

	"github.com/DevAntonioJorge/go-notes/internal/dto"
	"github.com/DevAntonioJorge/go-notes/internal/interfaces"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type FolderHandler struct{
	service interfaces.IFolderService
}

func NewFolderHandler(service interfaces.IFolderService) *FolderHandler{
	return &FolderHandler{service}
}

func (h *FolderHandler) CreateFolderHandler(c echo.Context) error {
	var req dto.CreateFolderRequest

	if err := c.Bind(&req); err != nil{
		return c.String(http.StatusBadRequest, "Invalid request body")
	}

	validate := validator.New()
	if err := validate.Struct(req); err != nil{
		return c.String(http.StatusInternalServerError, "Invalid folder name")
	}

	folder, err := h.service.SaveFolder(req)
	if err != nil{
		return c.String(http.StatusInternalServerError, "Failed to save folder")
	}

	return c.JSON(http.StatusCreated, echo.Map{
		"message": "Folder created",
		"folder": folder,
	})
}

func (h *FolderHandler) GetFolderHandler(c echo.Context) error{
	var req dto.GetFolderRequest
	if err := c.Bind(&req); err != nil{
		return c.String(http.StatusBadRequest, "No id specified or invalid parameter")
	}

	folder, err := h.service.GetFolder(req.ID)
	if err != nil{
		return c.String(http.StatusNotFound, "Folder not found")
	}

	return c.JSON(http.StatusAccepted, echo.Map{
		"folder": folder,
	})
}