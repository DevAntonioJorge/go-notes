package dto

type CreateFolderRequest struct {
	Name     string  `json:"name" validate:"min=4,max=10"`
	ParentID *string `json:"parent_id"`
}
type UpdateFolderRequest struct {
	ID   string `param:"id"`
	Name string `json:"name" validate:"min=4,max=10"`
}

type GetFolderRequest struct {
	ID string `param:"id"`
}

type DeleteFolderRequest struct {
	ID string `param:"id"`
}

type GetFoldersRequest struct {
	UserID string `json:"user_id"`
}

type GetFolderByPathRequest struct {
	UserID string `json:"user_id"`
	Path   string `json:"path"`
}

type MoveFolderRequest struct {
	FolderID string `json:"id"`
	ParentID string `json:"id"`
}
