package models

import "time"

type Blog struct{
	ID string
	UserID string
	Title string
	Content string
	Description string
	Tags []string
	CreatedAt string
	UpdatedAt time.Time
	
}