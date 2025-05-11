package models

import (
	"time"

	"github.com/google/uuid"
)

type Note struct{
	//Metadados
	ID string `json:"id"`
	UserID string `json:"user_id"`
	FolderID string `json:"folder_id"`
	Tags []string `json:"tags"`
	//Dados
	Title string `json:"title"`
	Content string `json:"content"`
	CreatedAt string `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
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

