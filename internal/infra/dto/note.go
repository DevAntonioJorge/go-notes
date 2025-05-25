package dto

type CreateNoteRequest struct {
	UserID   string   `json:"user_id" validate:"required"`
	FolderID string   `json:"folder_id" validate:"required"`
	Title    string   `json:"title" validate:"required,min=4,max=20"`
	Content  string   `json:"content" validate:"max=1000"`
	Tags     []string `json:"tags" validate:"max=10"`
}

type UpdateNoteRequest struct {
	ID      string   `json:"id" validate:"required"`
	Title   string   `json:"title" validate:"required,min=4,max=20"`
	Content string   `json:"content" validate:"max=1000"`
	Tags    []string `json:"tags" validate:"max=10"`
}

type UpdateNoteFolderRequest struct {
	ID       string `json:"id" validate:"required"`
	NewFolderID string `json:"new_folder_id" validate:"required"`
}

type UpdateNoteTagsRequest struct {
	ID   string   `json:"id" validate:"required"`
	Tags []string `json:"tags" validate:"max=10"`
}

type UpdateNoteTitleRequest struct {
	ID    string `json:"id" validate:"required"`
	Title string `json:"title" validate:"required,min=4,max=20"`
}
type UpdateNoteContentRequest struct {
	ID      string `json:"id" validate:"required"`
	Content string `json:"content" validate:"max=1000"`
}