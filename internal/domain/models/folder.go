package models

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type Folder struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	ParentID  string `json:"parent_id"`
	CreatedAt string `json:"created_at"`
}

func NewFolder(name string, parentID string) (*Folder, error) {
	if name == "" {
		return nil, errors.New("name is required")
	}
	return &Folder{
		ID:        uuid.NewString(),
		Name:      name,
		ParentID:  parentID,
		CreatedAt: time.Now().Format(time.DateTime),
	}, nil
}
