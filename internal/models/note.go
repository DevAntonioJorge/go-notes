package models

import (
	"time"

	"github.com/google/uuid"
)

type Note struct{
	//Metadados
	ID string
	UserID string
	FolderID string
	Tags []string
	//Dados
	Title string
	Content string
	CreatedAt string
	UpdatedAt time.Time
}

func NewNote(userID, folderID, title, content string, tags []string) *Note{
	return &Note{
		ID: uuid.NewString(),
		UserID: userID,
		FolderID: folderID,
		Tags: tags,
		Title: title,
		Content: content,
		CreatedAt: time.Now().Format(time.DateTime),
		UpdatedAt: time.Now(),
	}
}

