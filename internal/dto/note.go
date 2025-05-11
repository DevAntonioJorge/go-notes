package dto

type CreateNoteRequest struct{
	UserID string `json:"user_id"`
	FolderID string `json:"folder_id"`
	Title string `json:"title"`
	Content string `json:"content"`
	Tags []string `json:"tags"`
}

type UpdateNoteRequest struct{
	ID string `json:"id"`
	Title string `json:"title"`
	Content string `json:"content"`
	Tags []string `json:"tags"`
}