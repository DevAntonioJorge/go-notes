package dto

type CreateFolderRequest struct {
	Name     string `json:"name" validate:"min=4,max=10"`
	ParentID string `json:"parent_id" validate:"required"`
}
type UpdateFolderRequest struct {
	ID   string `param:"id" validate:"required"`
	Name string `json:"name" validate:"min=4,max=10"`
}

type GetFolderRequest struct {
	ID string `param:"id"`
}

type DeleteFolderRequest struct {
	ID string `param:"id"`
}

type GetFoldersRequest struct {
	UserID string `json:"user_id" validate:"required"`
}

type GetFolderByPathRequest struct {
	UserID string `json:"user_id" validate:"required"`
	Path   string `json:"path" validate:"required"`
}

type MoveFolderRequest struct {
	FolderID string `json:"id" validate:"required"`
	ParentID string `json:"parent_id" validate:"required"`
}
