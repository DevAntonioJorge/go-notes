package models

import (
	"time"

	"github.com/google/uuid"
)

type Folder struct {
	ID string
	Name string
	ParentID *string 
	CreatedAt string
}

func NewFolder(name string, parentID *string) *Folder {
	if name == ""{
		return nil
	}
	return &Folder{
		ID: uuid.NewString(),
		Name: name,
		ParentID: parentID,
		CreatedAt: time.Now().Format(time.DateTime),
	}
}