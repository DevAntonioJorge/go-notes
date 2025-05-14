package dto

type CreateFolderRequest struct{
	Name string `json:"name" validate:"min=4,max=10"`
	ParentID *string `json:"parent_id"`
}
type UpdateFolderRequest struct{
	Name string `json:"name" validate:"min=4,max=10"`
}

type GetFolderRequest struct{
	ID string `param:"id"`
}


